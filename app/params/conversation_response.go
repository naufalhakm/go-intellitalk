package params

type ConversationResponse struct {
	ID       string      `json:"id"`
	UserId   *string     `json:"user_id"`
	Messages interface{} `json:"messages"`
}
