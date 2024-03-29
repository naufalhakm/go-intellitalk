package repository

import (
	"context"
	"errors"

	"github.com/naufalhakm/go-intellitalk/app/model"
	"github.com/naufalhakm/go-intellitalk/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(ctx context.Context, dbMgo *mongo.Client, user *model.User) (*mongo.InsertOneResult, error)
	FindById(ctx context.Context, dbMgo *mongo.Client, user *model.User, id string) (*model.User, error)
	FindByEmail(ctx context.Context, dbMgo *mongo.Client, user *model.User, email string) (*model.User, error)
	GetAllUser(ctx context.Context, dbMgo *mongo.Client, users []*model.User) ([]*model.User, error)
	UpdateUserStatus(ctx context.Context, dbMgo *mongo.Client, id string) (*mongo.UpdateResult, error)
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

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, dbMgo *mongo.Client, user *model.User, email string) (*model.User, error) {
	var table = database.MgoCollection("users", dbMgo)

	err := table.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repository *UserRepositoryImpl) GetAllUser(ctx context.Context, dbMgo *mongo.Client, users []*model.User) ([]*model.User, error) {
	var table = database.MgoCollection("users", dbMgo)
	cursor, err := table.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var user model.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(ctx)

	if len(users) == 0 {
		return nil, errors.New("documents not found")
	}

	return users, nil

}
func (repository *UserRepositoryImpl) UpdateUserStatus(ctx context.Context, dbMgo *mongo.Client, id string) (*mongo.UpdateResult, error) {
	var table = database.MgoCollection("users", dbMgo)

	objectId, errId := primitive.ObjectIDFromHex(id)
	if errId != nil {
		return nil, errId
	}

	result, err := table.UpdateOne(
		ctx,
		bson.M{"_id": objectId},
		bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: 1}}}})
	if err != nil {
		return nil, err
	}

	return result, nil
}
