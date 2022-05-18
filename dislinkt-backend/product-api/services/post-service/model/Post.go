package model

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID           uuid.UUID `json:"id"`
	Description  string    `json:"description"`
	CreationDate time.Time `json:"creationDate" gorm:"not null"`
	UserID       uuid.UUID `json:"userID" gorm:"not null"`
	LocationId   uuid.UUID `json:"locationID"`
}
