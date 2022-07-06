package repository

import (
	"connection/module/model"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE             = "connections"
	CollectionConnection = "connectionsData"
)

type ConnectionRepo struct {
	connections *mongo.Collection
}

func NewConnectionRepository(client *mongo.Client) ConnectionRepo {

	connections := client.Database(DATABASE).Collection(CollectionConnection)

	return ConnectionRepo{connections: connections}

}

func (r ConnectionRepo) Create(connection *model.Connection) error {
	result, err := r.connections.InsertOne(context.TODO(), connection)
	if err != nil {
		return err
	}
	connection.Id = result.InsertedID.(primitive.ObjectID)

	return nil
}
