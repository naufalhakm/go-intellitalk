package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufalhakm/go-intellitalk/app/params"
	"github.com/naufalhakm/go-intellitalk/app/service"
)

type AuthController interface {
	Create(ctx *gin.Context)
}

type AuthControllerImpl struct {
	UserService service.UserService
}

func NewAuthContoller(userService service.UserService) AuthController {
	return &AuthControllerImpl{
		UserService: userService,
	}
}

func (controller *AuthControllerImpl) Create(ctx *gin.Context) {
	var user params.UserReguest

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	id, errCust := controller.UserService.Create(ctx, &user)
	if errCust != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": errCust,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "Success created new user",
		"data":    id,
	})
}
