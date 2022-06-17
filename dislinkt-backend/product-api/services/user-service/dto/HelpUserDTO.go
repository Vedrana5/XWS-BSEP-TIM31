package dto

type HelpUserDTO struct {
	Username string `json:"Username" validate:"required,min=2,max=30"`
}
