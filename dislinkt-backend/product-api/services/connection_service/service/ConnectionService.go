package service

import (
	"common/module/proto/connection_service"
	"connection/module/model"
	"connection/module/repository"
	"fmt"
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

func (service *ConnectionService) GetAllConnByUsername(username string) ([]*model.Connection, error) {
	return service.Repo.GetAllConnByUsername(username)
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

func (service *ConnectionService) GetUnreadMessages(username string) ([]*model.Message, error) {
	return service.Repo.GetUnreadMessages(username)
}

func (service *ConnectionService) ReadMessage(message []*connection_service.Message) ([]*model.Message, error, error) {
	var collections []*model.Message
	for i := 0; i < len(message); i++ {
		fmt.Println("USLA U FOR i id je" + message[i].Id)
		collections = append(collections, service.Repo.ReadMessage(message[i]))
	}
	return collections, nil, nil
}
