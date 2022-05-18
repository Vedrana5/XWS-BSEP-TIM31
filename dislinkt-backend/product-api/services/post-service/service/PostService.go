package service

import (
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/post-service/model"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/post-service/repository"
	"github.com/google/uuid"
)

type PostService struct {
	Repo *repository.PostRepositoty
}

func (service *PostService) CreatePost(post *model.Post) error {
	err := service.Repo.CreatePost(post)
	if err != nil {
		return err
	}
	return nil
}

func (service *PostService) FindByID(ID uuid.UUID) *model.Post {
	post := service.Repo.FindByID(ID)
	return post
}
