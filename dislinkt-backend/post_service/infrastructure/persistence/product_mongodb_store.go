package persistence

import (
	"context"
	pb "github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/common/proto/post_service"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/post_service/application"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "post_db"
	COLLECTION = "post"
)

type PostMongoDBStore struct {
	products *mongo.Collection
}

func NewProductMongoDBStore(client *mongo.Client) domain.PostStore {
	products := client.Database(DATABASE).Collection(COLLECTION)
	return &PostMongoDBStore{
		products: products,
	}
}

func (store *PostMongoDBStore) Get(id primitive.ObjectID) (*domain.Post, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *PostMongoDBStore) GetAll() ([]*domain.Post, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *PostMongoDBStore) Insert(product *domain.Post) error {
	result, err := store.products.InsertOne(context.TODO(), product)
	if err != nil {
		return err
	}
	product.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *PostMongoDBStore) DeleteAll() {
	store.products.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *PostMongoDBStore) filter(filter interface{}) ([]*domain.Post, error) {
	cursor, err := store.products.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *PostMongoDBStore) filterOne(filter interface{}) (product *domain.Post, err error) {
	result := store.products.FindOne(context.TODO(), filter)
	err = result.Decode(&product)
	return
}

func decode(cursor *mongo.Cursor) (products []*domain.Post, err error) {
	for cursor.Next(context.TODO()) {
		var product domain.Post
		err = cursor.Decode(&product)
		if err != nil {
			return
		}
		products = append(products, &product)
	}
	err = cursor.Err()
	return
}