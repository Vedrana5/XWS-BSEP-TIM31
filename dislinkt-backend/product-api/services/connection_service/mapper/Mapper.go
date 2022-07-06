package mapper

import (
	"common/module/proto/connection_service"
	"connection/module/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapNewConnection(connectionPb *connection_service.Connection) *model.Connection {
	connection := &model.Connection{
		Id:             primitive.NewObjectID(),
		FirstUsername:  connectionPb.FirstUsername,
		SecondUsername: connectionPb.SecondUsername,
		IsConfirmed:    connectionPb.IsConfirmed,
		IsDeleted:      false,
	}
	return connection
}

func MapConnectionReply(connection *model.Connection) *connection_service.Connection {
	id := connection.Id.Hex()

	connectionPb := &connection_service.Connection{
		Id:             id,
		FirstUsername:  connection.FirstUsername,
		SecondUsername: connection.SecondUsername,
		IsConfirmed:    connection.IsConfirmed,
		IsDeleted:      connection.IsDeleted,
	}

	return connectionPb
}
