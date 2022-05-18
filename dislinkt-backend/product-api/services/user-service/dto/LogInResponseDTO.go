package dto

import (
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/model"
	"github.com/google/uuid"
)

type LogInResponseDTO struct {
	ID         uuid.UUID        `json:"id"`
	Token      string           `json:"token"`
	TypeOfUser model.TypeOfUser `json:"typeOfUser"`
}
