package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Post struct {
	Id         primitive.ObjectID `bson:"_id"`
	Username   string             `bson:"username"`
	PostText   string             `bson:"post_text"`
	ImagePaths []byte             `bson:"image_paths"`
	DatePosted time.Time          `bson:"date_posted"`
	Reactions  []Reaction         `bson:"reactions"`
	Comments   []Comment          `bson:"comments"`
	IsDeleted  bool               `bson:"is_deleted"`
	Links      []string           `bson:"description"`
}
