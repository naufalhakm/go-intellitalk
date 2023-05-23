package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufalhakm/go-intellitalk/app/commons/response"
	"github.com/naufalhakm/go-intellitalk/app/params"
	"github.com/naufalhakm/go-intellitalk/app/service"
)

type AuthController interface {
	Create(ctx *gin.Context)
	FindById(ctx *gin.Context)
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

	result, custErr := controller.UserService.Create(ctx, &user)
	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.CreatedSuccessWithPayload(result)
	ctx.JSON(resp.StatusCode, resp)
}

func (controller *AuthControllerImpl) FindById(ctx *gin.Context) {
	var email string = ctx.Param("id")

	result, err := controller.UserService.FindById(ctx, email)
	if err != nil {
		ctx.AbortWithStatusJSON(err.StatusCode, err)
		return
	}
	resp := response.GeneralSuccessCustomMessageAndPayload("Success get data user!", result)
	ctx.JSON(resp.StatusCode, resp)
}
