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
	GetAllCandidate(ctx *gin.Context)
	GetAllUserConversation(ctx *gin.Context)
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
	var id string = ctx.Param("id")

	result, err := controller.UserService.FindById(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(err.StatusCode, err)
		return
	}
	resp := response.GeneralSuccessCustomMessageAndPayload("Success get data user!", result)
	ctx.JSON(resp.StatusCode, resp)
}

func (controller *AuthControllerImpl) GetAllCandidate(ctx *gin.Context) {
	result, err := controller.UserService.GetAllCandidate(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(err.StatusCode, err)
		return
	}
	resp := response.GeneralSuccessCustomMessageAndPayload("Success get all data user candidate!", result)
	ctx.JSON(resp.StatusCode, resp)
}

func (controller *AuthControllerImpl) GetAllUserConversation(ctx *gin.Context) {
	result, err := controller.UserService.GetAllUserConversation(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(err.StatusCode, err)
		return
	}
	resp := response.GeneralSuccessCustomMessageAndPayload("Success get all data user candidate conversation!", result)
	ctx.JSON(resp.StatusCode, resp)
}
