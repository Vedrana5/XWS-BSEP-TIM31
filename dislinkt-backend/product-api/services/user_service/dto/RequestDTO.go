package dto

import "github.com/google/uuid"

type RequestDTO struct {
	ID       uuid.UUID `json:"ID"`
	Password string    `json:"Password"`
	Token    string    `json:"Token"`
}
