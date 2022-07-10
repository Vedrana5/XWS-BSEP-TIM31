package model

import (
	"time"

	"github.com/google/uuid"
)

type ValidationCode struct {
	ID          uuid.UUID `json:"id"`
	Code        uuid.UUID `json:"code" gorm:"not null"`
	UserId      uuid.UUID `json:"user_id" gorm:"not null"`
	CreatedTime time.Time `json:"created_time" gorm:"not null"`
	ExpiredTime time.Time `json:"expired_time" gorm:"not null"`
	IsValid     bool      `json:"is_valid" gorm:"not null"`
	IsUsed      bool      `json:"is_used" gorm:"not null"`
}
