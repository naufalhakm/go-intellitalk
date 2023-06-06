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

type ConversationRepository interface {
	CreateConversation(ctx context.Context, dbMgo *mongo.Client, conversation *model.Conversation) (*mongo.InsertOneResult, error)
	FindById(ctx context.Context, dbMgo *mongo.Client, conversation *model.Conversation, id string) (*model.Conversation, error)
	GetAllConversation(ctx context.Context, dbMgo *mongo.Client, conversations []*model.Conversation) ([]*model.Conversation, error)
}

type ConversationRepositoryImpl struct {
}

func NewConversationRepository() ConversationRepository {
	return &ConversationRepositoryImpl{}
}

func (repository *ConversationRepositoryImpl) CreateConversation(ctx context.Context, dbMgo *mongo.Client, conversation *model.Conversation) (*mongo.InsertOneResult, error) {
	var table = database.MgoCollection("conversations", dbMgo)
	result, err := table.InsertOne(ctx, conversation)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repository *ConversationRepositoryImpl) FindById(ctx context.Context, dbMgo *mongo.Client, conversation *model.Conversation, id string) (*model.Conversation, error) {
	var table = database.MgoCollection("conversations", dbMgo)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = table.FindOne(ctx, bson.M{"_id": objectId}).Decode(&conversation)
	if err != nil {
		return nil, err
	}
	return conversation, nil
}

func (repository *ConversationRepositoryImpl) GetAllConversation(ctx context.Context, dbMgo *mongo.Client, conversations []*model.Conversation) ([]*model.Conversation, error) {
	var table = database.MgoCollection("conversations", dbMgo)
	cursor, err := table.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var conversation model.Conversation
		err := cursor.Decode(&conversation)
		if err != nil {
			return nil, err
		}

		conversations = append(conversations, &conversation)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(ctx)

	if len(conversations) == 0 {
		return nil, errors.New("documents not found")
	}

	return conversations, nil

}
