package repository

import (
	"common/module/proto/connection_service"
	"connection/module/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE             = "connections"
	CollectionConnection = "connectionsData"
	CollectionMessage    = "messageData"
	CollectionBlock      = "blockData"
)

type ConnectionRepo struct {
	connections *mongo.Collection
	messages    *mongo.Collection
	block       *mongo.Collection
}

func NewConnectionRepository(client *mongo.Client) ConnectionRepo {

	connections := client.Database(DATABASE).Collection(CollectionConnection)
	messages := client.Database(DATABASE).Collection(CollectionMessage)
	block := client.Database(DATABASE).Collection(CollectionBlock)

	return ConnectionRepo{connections: connections, messages: messages, block: block}

}

func (r ConnectionRepo) GetById(id primitive.ObjectID) (*model.Connection, error) {
	filter := bson.M{"_id": id}
	return r.filterOne(filter)
}

func (r ConnectionRepo) Create(connection *model.Connection) error {
	result, err := r.connections.InsertOne(context.TODO(), connection)
	if err != nil {
		return err
	}
	connection.Id = result.InsertedID.(primitive.ObjectID)

	return nil
}

func (r ConnectionRepo) CreateMessage(message *model.Message) error {
	result, err := r.messages.InsertOne(context.TODO(), message)
	if err != nil {
		return err
	}
	message.Id = result.InsertedID.(primitive.ObjectID)

	return nil
}

func (r ConnectionRepo) GetAllByUsername(username string) ([]*model.Connection, error) {

	filter := bson.M{"second_username": username, "is_confirmed": false, "is_deleted": false}
	return r.filter(filter)
}

func (r ConnectionRepo) GetAllConnByUsername(username string) ([]*model.Connection, error) {
	filter := bson.M{"first_username": username, "is_confirmed": true, "is_deleted": false}

	return r.filter(filter)
}

func (r ConnectionRepo) filter(filter interface{}) ([]*model.Connection, error) {
	cursor, err := r.connections.Find(context.TODO(), filter)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, context.TODO())

	if err != nil {
		return nil, err
	}

	return decode(cursor)
}

func (r ConnectionRepo) GetConnection(username string, username2 string) ([]*model.Connection, error) {
	filter := bson.M{"first_username": username, "second_username": username2}
	return r.filter(filter)

}

func (r ConnectionRepo) filterOne(filter bson.M) (connection *model.Connection, err error) {
	result := r.connections.FindOne(context.TODO(), filter)
	err = result.Decode(&connection)
	return
}

func (r ConnectionRepo) AcceptRequest(id primitive.ObjectID) (*model.Connection, error) {
	_, err := r.connections.UpdateOne(context.TODO(),
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"is_confirmed", true}}},
		})
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": id}
	return r.filterOne(filter)

}

func (r ConnectionRepo) RejectRequest(id primitive.ObjectID) (*model.Connection, error) {
	_, err := r.connections.UpdateOne(context.TODO(),
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"is_deleted", true}}},
		})
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": id}
	return r.filterOne(filter)
}

func (r ConnectionRepo) GetAllMessagesByUsernames(firstUsername string, secondUsername string) ([]*model.Message, error) {
	filter := bson.M{"first_username": firstUsername, "second_username": secondUsername}
	return r.filter2(filter)
}

func (r ConnectionRepo) filter2(filter interface{}) ([]*model.Message, error) {
	cursor, err := r.messages.Find(context.TODO(), filter)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, context.TODO())

	if err != nil {
		return nil, err
	}

	return decode2(cursor)
}

func (r ConnectionRepo) GetUnreadMessages(username string) ([]*model.Message, error) {
	filter := bson.M{"second_username": username, "is_read": false}
	return r.filter2(filter)
}

func (r ConnectionRepo) ReadMessage(message *connection_service.Message) {
	fmt.Println("ID U REPOZITORIJUMMU JE" + message.Id)
	_, err := r.connections.UpdateOne(context.TODO(),
		bson.M{"_id": message.Id},
		bson.D{
			{"$set", bson.D{{"is_read", true}}},
		})
	if err != nil {
		fmt.Println("U ERORU SAM")
		return
	}
	return
}

func (r ConnectionRepo) CreateBlock(block *model.Block) error {
	result, err := r.block.InsertOne(context.TODO(), block)
	if err != nil {
		return err
	}
	block.Id = result.InsertedID.(primitive.ObjectID)

	return nil
}

func decode2(cursor *mongo.Cursor) (messages []*model.Message, err error) {
	for cursor.Next(context.TODO()) {
		var conn model.Message
		err = cursor.Decode(&conn)
		if err != nil {
			return
		}
		messages = append(messages, &conn)
	}
	err = cursor.Err()
	return
}

func decode(cursor *mongo.Cursor) (connections []*model.Connection, err error) {
	for cursor.Next(context.TODO()) {
		var conn model.Connection
		err = cursor.Decode(&conn)
		if err != nil {
			return
		}
		connections = append(connections, &conn)
	}
	err = cursor.Err()
	return
}
