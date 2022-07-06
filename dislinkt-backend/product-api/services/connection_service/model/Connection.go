package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Connection struct {
	Id             primitive.ObjectID `bson:"_id"`
	FirstUsername  string             `bson:"first_username"`
	SecondUsername string             `bson:"second_username"`
	IsConfirmed    bool               `bson:"is_confirmed"`
}
