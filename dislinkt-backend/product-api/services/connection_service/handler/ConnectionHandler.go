package handler

import (
	"connection/module/service"
)

type ConnectionHandler struct {
	ConnectionService *service.ConnectionService
}

func NewConnectionHandler(connectionService *service.ConnectionService) *ConnectionHandler {
	return &ConnectionHandler{connectionService}
}
