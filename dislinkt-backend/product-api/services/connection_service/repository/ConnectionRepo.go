package repository

import "go.mongodb.org/mongo-driver/mongo"

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
