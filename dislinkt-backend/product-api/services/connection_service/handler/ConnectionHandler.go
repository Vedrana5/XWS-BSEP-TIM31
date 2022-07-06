package handler

import (
	"common/module/proto/connection_service"
	"connection/module/mapper"
	"connection/module/service"
	"context"
)

type ConnectionHandler struct {
	ConnectionService *service.ConnectionService
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

func (c ConnectionHandler) MustEmbedUnimplementedConnectionServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewConnectionHandler(connectionService *service.ConnectionService) *ConnectionHandler {
	return &ConnectionHandler{connectionService}
}
