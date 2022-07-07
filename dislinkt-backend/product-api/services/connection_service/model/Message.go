package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Message struct {
	Id             primitive.ObjectID `bson:"_id"`
	FirstUsername  string             `bson:"first_username"`
	SecondUsername string             `bson:"second_username"`
	MessageText    string             `bson:"message_text"`
	DatePosted     time.Time          `bson:"date_posted"`
	IsRead         bool               `bson:"is_read"`
}
