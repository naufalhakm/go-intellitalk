package params

type UserReguest struct {
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Division  string `json:"division" validate:"required"`
	Position  string `json:"position" validate:"required"`
	Parameter string `json:"parameter" validate:"required"`
}
