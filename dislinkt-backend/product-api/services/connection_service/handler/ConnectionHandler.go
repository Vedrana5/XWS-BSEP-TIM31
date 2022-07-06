package handler

import (
	cproto "common/module/proto/connection_service"
	"connection/module/service"
	"context"
)

type ConnectionHandler struct {
	ConnectionService *service.ConnectionService
}

func (c ConnectionHandler) Create(ctx context.Context, request *cproto.CreateConnectionRequest) (*cproto.Empty, error) {
	return nil, nil
}

func (c ConnectionHandler) MustEmbedUnimplementedConnectionServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewConnectionHandler(connectionService *service.ConnectionService) *ConnectionHandler {
	return &ConnectionHandler{connectionService}
}
