package params

type ConversationRequest struct {
	UserId   *string           `json:"user_id" validate:"required"`
	Messages []MessagesRequest `json:"messages" validate:"required"`
}

type MessagesRequest struct {
	Sender  string `bson:"sender" json:"sender"`
	Message string `bson:"message" json:"message"`
}
