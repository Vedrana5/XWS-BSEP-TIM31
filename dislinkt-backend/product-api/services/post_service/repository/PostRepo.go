package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE       = "posts"
	CollectionPost = "postsData"
)

type PostRepo struct {
	posts *mongo.Collection
}

func NewPostRepository(client *mongo.Client) PostRepo {
	posts := client.Database(DATABASE).Collection(CollectionPost)
	return PostRepo{posts: posts}
}
