package service

import (
	"context"
	"time"

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

	var user = model.User{
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := service.UserRepository.Create(ctx, service.DBMgo, &user)
	if err != nil {
		return nil, response.BadRequestError()
	}

	var IdHex string
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		IdHex = oid.Hex()
	}

	return &params.UserResponse{
		ID:    IdHex,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (service *UserServiceImpl) FindById(ctx context.Context, id string) (*params.UserResponse, *response.CustomError) {
	var user *model.User

	result, err := service.UserRepository.FindById(ctx, service.DBMgo, user, id)

	if err != nil {
		return nil, response.NotFoundError()
	}

	return &params.UserResponse{
		ID:    result.ID.Hex(),
		Name:  result.Name,
		Email: result.Email,
	}, nil

}
