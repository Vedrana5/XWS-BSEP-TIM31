package handler

import (
	post_service "common/module/proto/post_service"
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"post/module/mapper"

	"post/module/service"
)

type PostHandler struct {
	PostService *service.PostService
}

func NewPostHandler(postService *service.PostService) *PostHandler {
	return &PostHandler{postService}
}

func (p PostHandler) Get(ctx context.Context, request *post_service.GetRequest) (*post_service.GetResponse, error) {

	objectId, err := primitive.ObjectIDFromHex(request.GetId())
	if err != nil {
		return nil, err
	}
	post, err := p.PostService.GetById(objectId)
	if err != nil {
		return nil, err
	}
	postPb := mapper.MapPostReply(post)
	response := &post_service.GetResponse{Post: postPb}
	return response, nil
}

func (p PostHandler) GetAllByUsername(_ context.Context, request *pb.GetRequest) (*pb.GetMultipleResponse, error) {
	//request = p.sanitizeGetRequest(request)

	posts, err := p.postService.GetAllByUsername(request.Id)
	if err != nil {
		p.logError.Logger.WithFields(logrus.Fields{
			"userId": request.Id,
		}).Errorf("ERR:GET ALL POSTS FOR USER FROM DB")
		return nil, err
	}
	response := &pb.GetMultipleResponse{Posts: []*pb.Post{}}

	for _, post := range posts {
		current := api.MapPostReply(post)
		response.Posts = append(response.Posts, current)
	}

	return response, nil
}

func (p PostHandler) MustEmbedUnimplementedPostServiceServer() {
	//TODO implement me
	panic("implement me")
}
