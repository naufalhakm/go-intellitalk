package service

import (
	"context"
	"time"

	"github.com/naufalhakm/go-intellitalk/app/model"
	"github.com/naufalhakm/go-intellitalk/app/params"
	"github.com/naufalhakm/go-intellitalk/app/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	Create(ctx context.Context, req *params.UserReguest) (string, error)
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

func (service *UserServiceImpl) Create(ctx context.Context, req *params.UserReguest) (string, error) {
	defer service.DBMgo.Disconnect(context.TODO())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var user = model.User{
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := service.UserRepository.Create(ctx, service.DBMgo, &user)
	if err != nil {
		return "", err
	}

	var IdHex string
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		IdHex = oid.Hex()
	}

	return IdHex, nil
}
