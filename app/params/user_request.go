package params

type UserReguest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Division string `json:"division" validate:"required"`
	Position string `json:"position" validate:"required"`
	Skill    string `json:"skill" validate:"required"`
	Quantity int    `json:"quantity" validate:"required"`
}
