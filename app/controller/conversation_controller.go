package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufalhakm/go-intellitalk/app/commons/response"
	"github.com/naufalhakm/go-intellitalk/app/params"
	"github.com/naufalhakm/go-intellitalk/app/service"
)

type ConversationController interface {
	Create(ctx *gin.Context)
	FindById(ctx *gin.Context)
	GetAllConversation(ctx *gin.Context)
}

type ConversationControllerImpl struct {
	ConversationService service.ConversationService
}

func NewoConversationContoller(conversationService service.ConversationService) ConversationController {
	return &ConversationControllerImpl{
		ConversationService: conversationService,
	}
}

func (controller *ConversationControllerImpl) Create(ctx *gin.Context) {
	var conversation params.ConversationRequest

	err := ctx.ShouldBindJSON(&conversation)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	result, custErr := controller.ConversationService.Create(ctx, &conversation)
	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.CreatedSuccessWithPayload(result)
	ctx.JSON(resp.StatusCode, resp)
}

func (controller *ConversationControllerImpl) FindById(ctx *gin.Context) {
	var id string = ctx.Param("id")

	result, err := controller.ConversationService.FindById(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(err.StatusCode, err)
		return
	}
	resp := response.GeneralSuccessCustomMessageAndPayload("Success get data conversation!", result)
	ctx.JSON(resp.StatusCode, resp)
}

func (controller *ConversationControllerImpl) GetAllConversation(ctx *gin.Context) {
	result, err := controller.ConversationService.GetAllConversation(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(err.StatusCode, err)
		return
	}
	resp := response.GeneralSuccessCustomMessageAndPayload("Success get all data conversation!", result)
	ctx.JSON(resp.StatusCode, resp)
}
