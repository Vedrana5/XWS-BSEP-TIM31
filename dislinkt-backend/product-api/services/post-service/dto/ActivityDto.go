package dto

import (
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/model"
	"github.com/google/uuid"
)

type ActivityDTO struct {
	ID         uuid.UUID        `json:"id"`
	PostID     uuid.UUID        `json:"postID"`
	UserID     uuid.UUID        `json:"userID"`
	PostStatus model.PostStatus `json:"likedStatus"`
}
