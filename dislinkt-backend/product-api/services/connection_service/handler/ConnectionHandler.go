package handler

import (
	"common/module/proto/connection_service"
	"connection/module/mapper"
	"connection/module/model"
	"connection/module/service"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConnectionHandler struct {
	ConnectionService *service.ConnectionService
}

func (c ConnectionHandler) GetUnreadMessagesByUsername(ctx context.Context, request *connection_service.GetUsernamRequest) (*connection_service.GetMultipleMessagesResponse, error) {
	messages, err := c.ConnectionService.GetUnreadMessages(request.Username)
	if err != nil {
		return nil, err
	}
	response := &connection_service.GetMultipleMessagesResponse{Message: []*connection_service.Message{}}

	for _, message := range messages {
		current := mapper.MapMessagesReply(message)
		response.Message = append(response.Message, current)
	}

	return response, nil
}

func (c ConnectionHandler) GetConnUsername(ctx context.Context, request *connection_service.GetUsernamRequest) (*connection_service.GetMultipleConnectionResponse, error) {
	connections, err := c.ConnectionService.GetAllConnByUsername(request.Username)
	if err != nil {
		return nil, err
	}
	response := &connection_service.GetMultipleConnectionResponse{Connection: []*connection_service.Connection{}}

	for _, connection := range connections {
		current := mapper.MapConnectionReply(connection)
		response.Connection = append(response.Connection, current)
	}

	return response, nil
}

func (c ConnectionHandler) GetMessages(ctx context.Context, request *connection_service.GetUsernameRequest) (*connection_service.GetMultipleMessagesResponse, error) {
	messages, err := c.ConnectionService.GetAllMessagesByUsernames(request.FirstUsername, request.SecondUsername)
	if err != nil {
		return nil, err
	}
	response := &connection_service.GetMultipleMessagesResponse{Message: []*connection_service.Message{}}

	for _, message := range messages {
		current := mapper.MapMessagesReply(message)
		response.Message = append(response.Message, current)
	}

	return response, nil

}

func (c ConnectionHandler) RejectRequest(ctx context.Context, request *connection_service.EditRequest) (*connection_service.EditResponse, error) {

	objectId, err := primitive.ObjectIDFromHex(request.GetId())
	if err != nil {
		return nil, err
	}
	connection, err := c.ConnectionService.RejectRequest(objectId)
	connectionPb := mapper.MapConnectionReply(connection)
	response := &connection_service.EditResponse{Connecton: connectionPb}
	return response, nil
}

func (c ConnectionHandler) AcceptRequest(ctx context.Context, request *connection_service.EditRequest) (*connection_service.EditResponse, error) {

	objectId, err := primitive.ObjectIDFromHex(request.GetId())
	if err != nil {
		return nil, err
	}
	connection, err := c.ConnectionService.AcceptRequest(objectId)
	connectionPb := mapper.MapConnectionReply(connection)
	response := &connection_service.EditResponse{Connecton: connectionPb}
	return response, nil

}

func (c ConnectionHandler) FindConnectionByUsername(ctx context.Context, request *connection_service.GetRequest) (*connection_service.GetMultipleConnectionResponse, error) {
	connections, err := c.ConnectionService.GetAllByUsername(request.Username)
	if err != nil {
		return nil, err
	}
	response := &connection_service.GetMultipleConnectionResponse{Connection: []*connection_service.Connection{}}

	for _, connection := range connections {
		current := mapper.MapConnectionReply(connection)
		response.Connection = append(response.Connection, current)
	}

	return response, nil
}

func (c ConnectionHandler) GetConnect(ctx context.Context, request *connection_service.GetUsernameRequest) (*connection_service.GetConnectionResponse, error) {
	connection := c.ConnectionService.GetConnection(request.FirstUsername, request.SecondUsername)
	if connection == nil {

		connection1 := model.Connection{
			FirstUsername:  "",
			SecondUsername: "",
			IsConfirmed:    false,
			IsDeleted:      false,
		}
		connectionPb1 := mapper.MapConnectionReply1(connection1)
		response := &connection_service.GetConnectionResponse{Connection: connectionPb1}
		return response, nil
	} else {
		connectionPb := mapper.MapConnectionReply(connection)
		response := &connection_service.GetConnectionResponse{Connection: connectionPb}
		return response, nil
	}

}

func (c ConnectionHandler) Get(ctx context.Context, request *connection_service.GetRequest) (*connection_service.GetMultipleUsernameResponse, error) {
	return nil, nil
}

func (c ConnectionHandler) Create(ctx context.Context, request *connection_service.CreateConnectionRequest) (*connection_service.Empty, error) {
	connection := mapper.MapNewConnection(request.Connection)
	err := c.ConnectionService.Create(connection)
	if err != nil {
		return nil, err

	}
	return &connection_service.Empty{}, nil
}

func (c ConnectionHandler) CreateMessage(ctx context.Context, request *connection_service.CreateMessageRequest) (*connection_service.Empty, error) {
	message := mapper.MapNewMessage(request.Message)
	err := c.ConnectionService.CreateMessage(message)
	if err != nil {
		return nil, err

	}
	return &connection_service.Empty{}, nil
}

func (c ConnectionHandler) MustEmbedUnimplementedConnectionServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewConnectionHandler(connectionService *service.ConnectionService) *ConnectionHandler {
	return &ConnectionHandler{connectionService}
}
