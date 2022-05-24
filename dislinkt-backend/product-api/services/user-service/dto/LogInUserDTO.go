package dto

type LogInUserDTO struct {
	Username string `json:"Username" validate:"required,min=2,max=30"`
	Password string `json:"Password" validate:"required,min=10,max=30"`
	Question string `json:"Question"`
	Answer   string `json:"Answer"`
}
