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
