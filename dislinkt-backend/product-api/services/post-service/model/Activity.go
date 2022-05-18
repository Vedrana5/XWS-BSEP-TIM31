package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Activity struct {
	ID          uuid.UUID   `json:"id"`
	PostID      uuid.UUID   `json:"post_id" gorm:"not null"`
	UserID      uuid.UUID   `json:"user_id" gorm:"not null"`
	LikedStatus LikedStatus `json:"liked_status" gorm:"not null"`
}
