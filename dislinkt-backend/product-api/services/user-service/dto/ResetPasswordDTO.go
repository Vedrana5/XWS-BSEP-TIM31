package dto

type ResetPasswordDTO struct {
	ID       string `json:"ID"`
	Password string `json:"Password"`
	Code     string `json:"Code"`
}
