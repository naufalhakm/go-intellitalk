package model

import (
	"time"

	"github.com/naufalhakm/go-intellitalk/app/params"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Conversation struct {
	ID        primitive.ObjectID       `bson:"_id,omitempty" json:"_id,omitempty"`
	UserId    *string                  `bson:"user_id" json:"user_id"`
	Messages  []params.MessagesRequest `bson:"messages" json:"messages"`
	CreatedAt time.Time                `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time                `bson:"updated_at" json:"updated_at"`
}
