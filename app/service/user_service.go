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

type UserService interface {
	Create(ctx context.Context, req *params.UserReguest) (*params.UserResponse, *response.CustomError)
	FindById(ctx context.Context, id string) (*params.UserResponse, *response.CustomError)
	GetAllCandidate(ctx context.Context) ([]*params.UserCandidateResponse, *response.CustomError)
	GetAllUserConversation(ctx context.Context) ([]*params.UserConversationResponse, *response.CustomError)
}

type UserServiceImpl struct {
	DBMgo                  *mongo.Client
	UserRepository         repository.UserRepository
	ConversationRepository repository.ConversationRepository
}

func NewUserService(dbMgo *mongo.Client, userRepository repository.UserRepository, conversationRepository repository.ConversationRepository) UserService {
	return &UserServiceImpl{
		DBMgo:                  dbMgo,
		UserRepository:         userRepository,
		ConversationRepository: conversationRepository,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, req *params.UserReguest) (*params.UserResponse, *response.CustomError) {
	val := validator.New()
	if err := val.Struct(req); err != nil {
		return nil, response.BadRequestError()
	}

	var user = model.User{
		Name:      req.Name,
		Email:     req.Email,
		Division:  req.Division,
		Position:  req.Position,
		Skill:     req.Skill,
		Quantity:  req.Quantity,
		Status:    0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, errEmail := service.UserRepository.FindByEmail(ctx, service.DBMgo, &user, user.Email)
	if errEmail == nil {
		return nil, response.RepositoryErrorWithAdditionalInfo("Email has been used, replace with another email.")
	}

	// session, errSess := service.DBMgo.StartSession()
	// if errSess != nil {
	// 	return nil, response.GeneralError()
	// }

	// if errSess = session.StartTransaction(); errSess != nil {
	// 	return nil, response.GeneralError()
	// }

	result, err := service.UserRepository.Create(ctx, service.DBMgo, &user)
	if err != nil {
		return nil, response.BadRequestError()
	}

	var IdHex string
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		IdHex = oid.Hex()
	}

	return &params.UserResponse{
		ID:       IdHex,
		Name:     user.Name,
		Email:    user.Email,
		Division: user.Division,
		Position: user.Position,
		Skill:    user.Skill,
		Quantity: user.Quantity,
		Status:   user.Status,
		Link:     "https://arkademi-intellitalk.vercel.app/#/prepatation/" + IdHex,
	}, nil
}

func (service *UserServiceImpl) FindById(ctx context.Context, id string) (*params.UserResponse, *response.CustomError) {
	var user *model.User

	result, err := service.UserRepository.FindById(ctx, service.DBMgo, user, id)

	if err != nil {
		return nil, response.NotFoundError()
	}

	// json, errJson := json.Marshal(result)
	// if errJson != nil {
	// 	return nil, response.GeneralError()
	// }
	// Enkripsi data
	// encryptedData, errEncrypt := encryption.EncryptData(json)
	// if errEncrypt != nil {
	// 	return nil, response.GeneralError()
	// }

	return &params.UserResponse{
		ID:       result.ID.Hex(),
		Name:     result.Name,
		Email:    result.Email,
		Division: result.Division,
		Position: result.Position,
		Skill:    result.Skill,
		Quantity: result.Quantity,
		Link:     "https://arkademi-intellitalk.vercel.app/#/prepatation/" + result.ID.Hex(),
		Status:   result.Status,
	}, nil

}

func (service *UserServiceImpl) GetAllCandidate(ctx context.Context) ([]*params.UserCandidateResponse, *response.CustomError) {
	var users []*model.User

	results, err := service.UserRepository.GetAllUser(ctx, service.DBMgo, users)
	if err != nil {
		return nil, response.NotFoundError()
	}

	var responses []*params.UserCandidateResponse
	for _, result := range results {
		response := params.UserCandidateResponse{
			ID:       result.ID.Hex(),
			Name:     result.Name,
			Email:    result.Email,
			Division: result.Division,
			Position: result.Position,
			Skill:    result.Skill,
			Quantity: result.Quantity,
			Link:     "https://arkademi-intellitalk.vercel.app/#/preparation/" + result.ID.Hex(),
		}
		responses = append(responses, &response)
	}

	return responses, nil
}

func (service *UserServiceImpl) GetAllUserConversation(ctx context.Context) ([]*params.UserConversationResponse, *response.CustomError) {
	var users []*model.User

	results, err := service.UserRepository.GetAllUser(ctx, service.DBMgo, users)
	if err != nil {
		return nil, response.NotFoundError()
	}

	var responses []*params.UserConversationResponse
	for _, result := range results {
		var conversation *model.Conversation

		resultCon, errCon := service.ConversationRepository.FindByUserId(ctx, service.DBMgo, conversation, result.ID.Hex())

		if errCon == nil {

			response := params.UserConversationResponse{
				ID:       resultCon.ID.Hex(),
				Name:     result.Name,
				Email:    result.Email,
				Division: result.Division,
				Position: result.Position,
				Skill:    result.Skill,
				Quantity: result.Quantity,
				// LinkVideo: "https://drive.google.com/drive/intellitalk/guest/" + result.ID.Hex(),
			}
			responses = append(responses, &response)

		}

	}

	return responses, nil
}
