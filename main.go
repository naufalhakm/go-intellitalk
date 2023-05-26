package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufalhakm/go-intellitalk/app/controller"
	"github.com/naufalhakm/go-intellitalk/app/repository"
	"github.com/naufalhakm/go-intellitalk/app/service"
	"github.com/naufalhakm/go-intellitalk/config"
	"github.com/naufalhakm/go-intellitalk/database"
)

func main() {
	config.LoadConfig()
	client := database.NewMgoConnection()
	defer client.Disconnect(context.TODO())

	UserRepository := repository.NewUserRepository()
	UserService := service.NewUserService(client, UserRepository)
	UserController := controller.NewAuthContoller(UserService)
	router := gin.New()

	router.Use(gin.Logger())

	api := router.Group("/api")
	{
		v1 := api.Group("v1")
		{
			v1.GET("/users", UserController.GetAllCandidate)
			v1.GET("/users/:id", UserController.FindById)
			v1.POST("/users", UserController.Create)
		}
	}

	router.GET("/ping", Ping)

	router.Run(":" + config.ENV.PortServer)
}
func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
	})
}
