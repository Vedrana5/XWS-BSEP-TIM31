package handler

import (
	post_service "common/module/proto/post_service"
	"context"
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

func (p PostHandler) GetAllByUsername(_ context.Context, request *post_service.GetRequest) (*post_service.GetMultipleResponse, error) {

	posts, err := p.PostService.GetAllByUsername(request.Id)
	if err != nil {
		return nil, err
	}
	response := &post_service.GetMultipleResponse{Posts: []*post_service.Post{}}

	for _, post := range posts {
		current := mapper.MapPostReply(post)
		response.Posts = append(response.Posts, current)
	}

	return response, nil
}

func (p PostHandler) GetAll(_ context.Context, _ *post_service.Empty) (*post_service.GetMultipleResponse, error) {

	posts, err := p.PostService.GetAll()
	if err != nil {
		return nil, err
	}
	response := &post_service.GetMultipleResponse{Posts: []*post_service.Post{}}
	for _, post := range posts {
		current := mapper.MapPostReply(post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}

func (p PostHandler) MustEmbedUnimplementedPostServiceServer() {
	//TODO implement me
	panic("implement me")
}
