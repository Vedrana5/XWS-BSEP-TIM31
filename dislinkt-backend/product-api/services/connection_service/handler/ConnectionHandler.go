package handler

import (
	cproto "common/module/proto/connection_service"
	"connection/module/mapper"
	"connection/module/service"
	"context"
)

type ConnectionHandler struct {
	ConnectionService *service.ConnectionService
}

func (c ConnectionHandler) Create(ctx context.Context, request *cproto.CreateConnectionRequest) (*cproto.Empty, error) {
	connection := mapper.MapNewConnection(request.Connection)
	err := c.ConnectionService.Create(connection)
	if err != nil {
		return nil, err

	}
	return &cproto.Empty{}, nil
}

func (c ConnectionHandler) MustEmbedUnimplementedConnectionServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewConnectionHandler(connectionService *service.ConnectionService) *ConnectionHandler {
	return &ConnectionHandler{connectionService}
}
