package service

import "post/module/repository"

type PostService struct {
	Repo repository.PostRepo
}

func NewPostService(repository repository.PostRepo) *PostService {
	return &PostService{Repo: repository}
}
