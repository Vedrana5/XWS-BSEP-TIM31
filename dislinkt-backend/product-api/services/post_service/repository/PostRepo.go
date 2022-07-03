package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
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

func (r PostRepo) DislikePost(post *model.Post, userName string) error {
	var reactions []model.Reaction

	reactionExists := false
	for _, reaction := range post.Reactions {
		if reaction.UserName != userName {
			reactions = append(reactions, reaction)
		} else {
			if reaction.Reaction != model.DISLIKED {
				reaction.Reaction = model.DISLIKED
				reactions = append(reactions, reaction)
			}
			reactionExists = true
		}
	}
	if !reactionExists {
		reaction := model.Reaction{
			UserName: userName,
			Reaction: model.DISLIKED,
		}
		reactions = append(reactions, reaction)
	}

	_, err := r.posts.UpdateOne(context.TODO(), bson.M{"_id": post.Id}, bson.D{
		{"$set", bson.D{{"reactions", reactions}}},
	},
	)
	if err != nil {
		return err
	}

	return nil
}

func (r PostRepo) LikePost(post *model.Post, userName string) error {
	var reactions []model.Reaction

	reactionExists := false
	for _, reaction := range post.Reactions {
		if reaction.UserName != userName {
			reactions = append(reactions, reaction)
		} else {
			if reaction.Reaction != model.LIKED {
				reaction.Reaction = model.LIKED
				reactions = append(reactions, reaction)
			}
			reactionExists = true
		}

	}
	if !reactionExists {
		reaction := model.Reaction{
			UserName: userName,
			Reaction: model.LIKED,
		}
		reactions = append(reactions, reaction)
	}

	_, err := r.posts.UpdateOne(context.TODO(), bson.M{"_id": post.Id}, bson.D{
		{"$set", bson.D{{"reactions", reactions}}},
	},
	)
	if err != nil {
		return err
	}

	return nil

}

func (r PostRepo) CreateComment(post *model.Post, comment *model.Comment) error {
	comments := append(post.Comments, *comment)

	_, err := r.posts.UpdateOne(context.TODO(), bson.M{"_id": "507f1f77bcf86cd799439010"}, bson.D{
		{"$set", bson.D{{"comments", comments}}},
	},
	)
	if err != nil {
		return err
	}

	return nil
}

func (r PostRepo) Create(post *model.Post) error {
	result, err := r.posts.InsertOne(context.TODO(), post)
	if err != nil {
		return err
	}
	post.Id = result.InsertedID.(primitive.ObjectID)

	return nil
}

func (r PostRepo) GetAllByUsername(username string) ([]*model.Post, error) {
	filter := bson.M{"username": username}
	return r.filter(filter)
}

func (r PostRepo) GetAll() ([]*model.Post, error) {
	filter := bson.D{}
	return r.filter(filter)
}

func (r PostRepo) GetById(id primitive.ObjectID) (*model.Post, error) {
	filter := bson.M{"_id": id}
	return r.filterOne(filter)
}

func (r PostRepo) filterOne(filter bson.M) (post *model.Post, err error) {
	result := r.posts.FindOne(context.TODO(), filter)
	err = result.Decode(&post)
	return
}

func (r PostRepo) filter(filter interface{}) ([]*model.Post, error) {
	cursor, err := r.posts.Find(context.TODO(), filter)
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

func decode(cursor *mongo.Cursor) (posts []*model.Post, err error) {
	for cursor.Next(context.TODO()) {
		var post model.Post
		err = cursor.Decode(&post)
		if err != nil {
			return
		}
		posts = append(posts, &post)
	}
	err = cursor.Err()
	return
}
