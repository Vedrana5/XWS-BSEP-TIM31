package dto

type LogInUserPasswordlessDTO struct {
	Username string `json:"Username" validate:"required,min=2,max=30"`
	Code     string `json:"Code"`
}
