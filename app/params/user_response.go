package params

type UserResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Division string `json:"division"`
	Position string `json:"position"`
	Skill    string `json:"skill"`
	Quantity int    `json:"quantity"`
	Link     string `json:"link"`
}

type UserCandidateResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Division string `json:"division"`
	Position string `json:"position"`
	Skill    string `json:"skill"`
	Quantity int    `json:"quantity"`
	Link     string `json:"link"`
}

type UserConversationResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Division  string `json:"division"`
	Position  string `json:"position"`
	Skill     string `json:"skill"`
	Quantity  int    `json:"quantity"`
	LinkVideo string `json:"link_vidio"`
}
