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
func (service *ConnectionService) GetAllByUsername(username string) ([]*model.Connection, error) {
	return service.Repo.GetAllByUsername(username)
}

func (service *ConnectionService) GetConnection(firstUsername string, secondUsername string) *model.Connection {
	var connection, _ = service.Repo.GetConnection(firstUsername, secondUsername)
	if connection == nil {
		return nil
	} else {
		return connection
	}

}
