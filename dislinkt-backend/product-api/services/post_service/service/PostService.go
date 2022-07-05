package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"post/module/model"
	"post/module/repository"
	"strings"
)

type PostService struct {
	Repo *repository.PostRepo
}

func NewPostService(repository repository.PostRepo) *PostService {
	return &PostService{Repo: &repository}
}

func (service *PostService) GetById(id primitive.ObjectID) (*model.Post, error) {
	return service.Repo.GetById(id)
}

func (service *PostService) GetAll() ([]*model.Post, error) {
	return service.Repo.GetAll()
}

func (service *PostService) GetAllByUsername(username string) ([]*model.Post, error) {
	return service.Repo.GetAllByUsername(username)
}

func (service *PostService) Create(post *model.Post) error {
	return service.Repo.Create(post)
}

func (service *PostService) CreateComment(post *model.Post, comment *model.Comment) error {
	return service.Repo.CreateComment(post, comment)
}

func (service *PostService) LikePost(post *model.Post, userId string) error {
	return service.Repo.LikePost(post, userId)
}

func (service *PostService) DislikePost(post *model.Post, userId string) error {
	return service.Repo.DislikePost(post, userId)
}

func (service *PostService) CreateJobOffer(offer *model.JobOffer) error {
	return service.Repo.CreateJobOffer(offer)
}

func (service *PostService) GetAllJobOffers() ([]*model.JobOffer, error) {
	return service.Repo.GetAllJobOffers()
}

func (service *PostService) GetOffersByPosition(position string) []*model.JobOffer {
	offers, _ := service.Repo.GetAllJobOffers()
	var s []*model.JobOffer
	for i := 0; i < len(offers); i++ {
		if strings.Contains(offers[i].Position, position) {
			var offer *model.JobOffer = offers[i]

			s = append(s, offer)

		}
	}
	if s != nil {
		return s
	}
	return nil
}
