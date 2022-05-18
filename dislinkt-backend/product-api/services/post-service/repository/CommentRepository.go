package repository

import (
	"fmt"

	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/post-service/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentRepository struct {
	Database *gorm.DB
}

func (repo *CommentRepository) CreateComment(comment *model.Comment) error {
	result := repo.Database.Create(comment)
	fmt.Print(result)
	return nil
}

func (repo *CommentRepository) FindAllCommentsForPost(postId uuid.UUID) []model.Comment {
	var comments []model.Comment
	repo.Database.Select("*").Where("post_id = ?", postId).Find(&comments)
	return comments
}

func (repo *CommentRepository) FindAllUserComments(userId uuid.UUID) []model.Comment {
	var comments []model.Comment
	repo.Database.Select("*").Where("user_id = ?", userId).Find(&comments)
	return comments
}
