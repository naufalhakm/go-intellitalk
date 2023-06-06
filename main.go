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
	ConversationRepository := repository.NewConversationRepository()

	UserService := service.NewUserService(client, UserRepository, ConversationRepository)
	UserController := controller.NewAuthContoller(UserService)

	ConversationService := service.NewConversationService(client, ConversationRepository)
	ConversationController := controller.NewoConversationContoller(ConversationService)

	router := gin.New()

	router.Use(gin.Logger(), CORS())

	api := router.Group("/api")
	{
		v1 := api.Group("v1")
		{
			v1.GET("/users", UserController.GetAllCandidate)
			v1.GET("/users/:id", UserController.FindById)
			v1.POST("/users", UserController.Create)
			v1.GET("/conversations", ConversationController.GetAllConversation)
			v1.GET("/conversations/:id", ConversationController.FindById)
			v1.POST("/conversations", ConversationController.Create)
			v1.GET("/users/conversations", UserController.GetAllUserConversation)
		}
	}

	router.GET("/ping", Ping)

	router.Run(":" + config.ENV.PortServer)
}

func CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Request-Methods", "GET, OPTIONS, POST, PUT, DELETE")
		ctx.Header("Access-Control-Request-Headers", "Authorization, Content-Type")
		ctx.Next()
	}
}

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
	})
}
