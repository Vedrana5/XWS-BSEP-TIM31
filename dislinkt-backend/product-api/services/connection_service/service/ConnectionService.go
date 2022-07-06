package service

import (
	"connection/module/model"
	"connection/module/repository"
)

type ConnectionService struct {
	Repo *repository.ConnectionRepo
}

func (service *ConnectionService) Create(post *model.Connection) error {
	return service.Repo.Create(post)
}

func NewConnectionService(repository repository.ConnectionRepo) *ConnectionService {
	return &ConnectionService{Repo: &repository}
}
