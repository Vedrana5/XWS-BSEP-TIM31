package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"post/module/model"
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

func (r PostRepo) GetById(id primitive.ObjectID) *model.Post {
	return nil
}

func (r PostRepo) GetAll() []*model.Post {
	return nil
}

func (r PostRepo) GetAllByUsername(username string) []*model.Post {
	return nil
}

func (r PostRepo) Create(post *model.Post) error {
	return nil
}

func (r PostRepo) CreateComment(post *model.Post, comment *model.Comment) error {
	return nil
}

func (r PostRepo) LikePost(post *model.Post, id interface{}) error {
	return nil
}

func (r PostRepo) DislikePost(post *model.Post, id interface{}) error {
	return nil
}
