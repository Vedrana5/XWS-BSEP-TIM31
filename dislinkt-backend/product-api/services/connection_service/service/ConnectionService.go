package service

import (
	"connection/module/model"
	"connection/module/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (service *ConnectionService) AcceptRequest(id primitive.ObjectID) (*model.Connection, error) {
	return service.Repo.AcceptRequest(id)
}

func (service *ConnectionService) RejectRequest(id primitive.ObjectID) (*model.Connection, error) {
	return service.Repo.RejectRequest(id)
}