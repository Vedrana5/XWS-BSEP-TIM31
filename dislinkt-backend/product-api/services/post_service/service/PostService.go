package service

import "post/module/repository"

type PostService struct {
	Repo *repository.PostRepo
}
