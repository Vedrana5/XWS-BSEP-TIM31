package service

import (
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/post-service/model"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/post-service/repository"
	"github.com/google/uuid"
)

type CommentService struct {
	Repo *repository.CommentRepository
}

func (service *CommentService) CreateComment(comment *model.Comment) error {
	err := service.Repo.CreateComment(comment)
	if err != nil {
		return err
	}
	return nil
}

func (service *CommentService) FindAllCommentsForPost(postId uuid.UUID) []model.Comment {
	comments := service.Repo.FindAllCommentsForPost(postId)
	return comments
}

func (service *CommentService) FindAllUserComments(userId uuid.UUID) []model.Comment {
	comments := service.Repo.FindAllUserComments(userId)
	return comments
}
