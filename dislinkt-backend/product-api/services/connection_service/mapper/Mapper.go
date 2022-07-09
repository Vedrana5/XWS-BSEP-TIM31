package mapper

import (
	"common/module/proto/connection_service"
	"connection/module/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
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

func MapNewMessage(messagePb *connection_service.Message) *model.Message {
	message := &model.Message{
		Id:             primitive.NewObjectID(),
		FirstUsername:  messagePb.FirstUsername,
		SecondUsername: messagePb.SecondUsername,
		MessageText:    messagePb.MessageText,
		DatePosted:     time.Now(),
		IsRead:         false,
	}
	return message
}
func MapNewBlock(blockPb *connection_service.Block) *model.Block {
	block := &model.Block{
		Id:             primitive.NewObjectID(),
		FirstUsername:  blockPb.FirstUsername,
		SecondUsername: blockPb.SecondUsername,
	}
	return block
}

func MapMessagesReply(message *model.Message) *connection_service.Message {
	id := message.Id.Hex()

	connectionPb := &connection_service.Message{
		Id:             id,
		FirstUsername:  message.FirstUsername,
		SecondUsername: message.SecondUsername,
		MessageText:    message.MessageText,
		DateCreated:    message.DatePosted.String(),
		IsRead:         message.IsRead,
	}

	return connectionPb
}

func MapBlockReply(block *model.Block) *connection_service.Block {
	id := block.Id.Hex()

	blockPb := &connection_service.Block{
		Id:             id,
		FirstUsername:  block.FirstUsername,
		SecondUsername: block.SecondUsername,
	}

	return blockPb
}

func MapBlockReply1(block model.Block) *connection_service.Block {
	id := block.Id.Hex()

	blockPb := &connection_service.Block{
		Id:             id,
		FirstUsername:  block.FirstUsername,
		SecondUsername: block.SecondUsername,
	}

	return blockPb
}

func MapConnectionReply1(connection model.Connection) *connection_service.Connection {
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
