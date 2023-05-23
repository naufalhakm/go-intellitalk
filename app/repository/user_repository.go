package repository

import (
	"context"

	"github.com/naufalhakm/go-intellitalk/app/model"
	"github.com/naufalhakm/go-intellitalk/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(ctx context.Context, dbMgo *mongo.Client, user *model.User) (*mongo.InsertOneResult, error)
	FindById(ctx context.Context, dbMgo *mongo.Client, user *model.User, id string) (*model.User, error)
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

func (repository *UserRepositoryImpl) FindById(ctx context.Context, dbMgo *mongo.Client, user *model.User, id string) (*model.User, error) {
	var table = database.MgoCollection("users", dbMgo)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = table.FindOne(ctx, bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
