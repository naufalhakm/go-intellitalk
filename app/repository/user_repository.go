package repository

import (
	"context"

	"github.com/naufalhakm/go-intellitalk/app/model"
	"github.com/naufalhakm/go-intellitalk/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(ctx context.Context, dbMgo *mongo.Client, user *model.User) (*mongo.InsertOneResult, error)
}

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, dbMgo *mongo.Client, user *model.User) (*mongo.InsertOneResult, error) {
	var table = database.MgoCollection("users", dbMgo)
	result, err := table.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return result, nil
}
