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

func (service *ConnectionService) CreateMessage(message *model.Message) error {
	return service.Repo.CreateMessage(message)
}

func NewConnectionService(repository repository.ConnectionRepo) *ConnectionService {
	return &ConnectionService{Repo: &repository}
}
func (service *ConnectionService) GetAllByUsername(username string) ([]*model.Connection, error) {
	return service.Repo.GetAllByUsername(username)
}

func (service *ConnectionService) GetConnection(firstUsername string, secondUsername string) *model.Connection {
	var connection, _ = service.Repo.GetConnection(firstUsername, secondUsername)
	var conn = connection[len(connection)-1]
	if connection == nil {
		return nil
	} else {
		return conn
	}

}

func (service *ConnectionService) AcceptRequest(id primitive.ObjectID) (*model.Connection, error) {
	return service.Repo.AcceptRequest(id)
}

func (service *ConnectionService) RejectRequest(id primitive.ObjectID) (*model.Connection, error) {
	return service.Repo.RejectRequest(id)
}

func (service *ConnectionService) GetAllMessagesByUsernames(firstUsername string, secondUsername string) ([]*model.Message, error) {
	var collections []*model.Message
	var messages1, _ = service.Repo.GetAllMessagesByUsernames(firstUsername, secondUsername)
	var messages2, _ = service.Repo.GetAllMessagesByUsernames(secondUsername, firstUsername)

	collections = append(collections, messages1...)
	collections = append(collections, messages2...)
	return collections, nil
}
