package service

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"post/module/model"
	"post/module/repository"
)

type PostService struct {
	Repo *repository.PostRepo
}

func NewPostService(repository repository.PostRepo) *PostService {
	return &PostService{Repo: &repository}
}

func (service *PostService) GetById(id primitive.ObjectID) *model.Post {
	return service.Repo.GetById(id)
}

func (service *PostService) GetAll() []*model.Post {
	return service.Repo.GetAll()
}

func (service *PostService) GetAllByUsername(username string) []*model.Post {
	return service.Repo.GetAllByUsername(username)
}

func (service *PostService) Create(post *model.Post) error {
	return service.Repo.Create(post)
}

func (service *PostService) CreateComment(post *model.Post, comment *model.Comment) error {
	return service.Repo.CreateComment(post, comment)
}

func (service *PostService) LikePost(post *model.Post, userId uuid.UUID) error {
	return service.Repo.LikePost(post, userId)
}

func (service *PostService) DislikePost(post *model.Post, userId uuid.UUID) error {
	return service.Repo.DislikePost(post, userId)
}
