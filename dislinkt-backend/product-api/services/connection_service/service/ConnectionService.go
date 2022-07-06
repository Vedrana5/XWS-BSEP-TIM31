package service

import "connection/module/repository"

type ConnectionService struct {
	Repo *repository.ConnectionRepo
}

func NewConnectionService(repository repository.ConnectionRepo) *ConnectionService {
	return &ConnectionService{Repo: &repository}
}
