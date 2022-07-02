package handler

import (
	"common/module/proto/post_service"
	"context"

	"post/module/service"
)

type PostHandler struct {
	UserService *service.PostService
}

func NewPostHandler(postService *service.PostService) *PostHandler {
	return &PostHandler{postService}
}

func (p PostHandler) CreatePost(ctx context.Context, request *post_service.CreateRequest) (*post_service.CreateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostHandler) MustEmbedUnimplementedPostServiceServer() {
	//TODO implement me
	panic("implement me")
}



