package dto

import (
	"github.com/google/uuid"
	"user/module/model"
)

type LogInResponseDTO struct {
	ID         uuid.UUID        `json:"ID"`
	Token      string           `json:"Token"`
	TypeOfUser model.TypeOfUser `json:"TypeOfUser"`
}
