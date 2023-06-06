package service

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/naufalhakm/go-intellitalk/app/commons/response"
	"github.com/naufalhakm/go-intellitalk/app/model"
	"github.com/naufalhakm/go-intellitalk/app/params"
	"github.com/naufalhakm/go-intellitalk/app/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ConversationService interface {
	Create(ctx context.Context, req *params.ConversationRequest) (*params.ConversationResponse, *response.CustomError)
	FindById(ctx context.Context, id string) (*params.ConversationResponse, *response.CustomError)
	GetAllConversation(ctx context.Context) ([]*params.ConversationResponse, *response.CustomError)
}

type ConversationServiceImpl struct {
	DBMgo                  *mongo.Client
	ConversationRepository repository.ConversationRepository
}

func NewConversationService(dbMgo *mongo.Client, conversationRepository repository.ConversationRepository) ConversationService {
	return &ConversationServiceImpl{
		DBMgo:                  dbMgo,
		ConversationRepository: conversationRepository,
	}
}

func (service *ConversationServiceImpl) Create(ctx context.Context, req *params.ConversationRequest) (*params.ConversationResponse, *response.CustomError) {
	val := validator.New()
	if err := val.Struct(req); err != nil {
		return nil, response.BadRequestError()
	}

	var conversation = model.Conversation{
		UserId:    req.UserId,
		Messages:  req.Messages,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := service.ConversationRepository.CreateConversation(ctx, service.DBMgo, &conversation)
	if err != nil {
		return nil, response.BadRequestError()
	}

	var IdHex string
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		IdHex = oid.Hex()
	}

	return &params.ConversationResponse{
		ID:       IdHex,
		UserId:   conversation.UserId,
		Messages: conversation.Messages,
	}, nil
}

func (service *ConversationServiceImpl) FindById(ctx context.Context, id string) (*params.ConversationResponse, *response.CustomError) {
	var conversation *model.Conversation

	result, err := service.ConversationRepository.FindById(ctx, service.DBMgo, conversation, id)

	if err != nil {
		return nil, response.NotFoundError()
	}

	return &params.ConversationResponse{
		ID:       result.ID.Hex(),
		UserId:   result.UserId,
		Messages: result.Messages,
	}, nil

}

func (service *ConversationServiceImpl) GetAllConversation(ctx context.Context) ([]*params.ConversationResponse, *response.CustomError) {
	var conversations []*model.Conversation

	results, err := service.ConversationRepository.GetAllConversation(ctx, service.DBMgo, conversations)
	if err != nil {
		return nil, response.NotFoundError()
	}

	var responses []*params.ConversationResponse
	for _, result := range results {
		response := params.ConversationResponse{
			ID:       result.ID.Hex(),
			UserId:   result.UserId,
			Messages: result.Messages,
		}
		responses = append(responses, &response)
	}

	return responses, nil
}
