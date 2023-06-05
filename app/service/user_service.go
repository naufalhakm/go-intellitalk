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
}

type UserServiceImpl struct {
	DBMgo          *mongo.Client
	UserRepository repository.UserRepository
}

func NewUserService(dbMgo *mongo.Client, userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		DBMgo:          dbMgo,
		UserRepository: userRepository,
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
		Link:     "http://localhost:3000/intellitalk/guest/" + IdHex,
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
		Link:     "http://localhost:3000/intellitalk/guest/" + result.ID.Hex(),
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
			Link:     "http://localhost:3000/intellitalk/guest/" + result.ID.Hex(),
		}
		responses = append(responses, &response)
	}

	return responses, nil
}
