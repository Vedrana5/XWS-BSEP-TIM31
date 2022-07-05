package handler

import (
	post_service "common/module/proto/post_service"
	"context"
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"post/module/mapper"
	"post/module/service"
	"strings"
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

func (p PostHandler) Create(ctx context.Context, request *post_service.CreatePostRequest) (*post_service.Empty, error) {
	post := mapper.MapNewPost(request.Post)
	err := p.PostService.Create(post)
	if err != nil {
		return nil, err

	}
	return &post_service.Empty{}, nil
}

func (p PostHandler) CreateComment(ctx context.Context, request *post_service.CreateCommentRequest) (*post_service.CreateCommentResponse, error) {

	objectId, err := primitive.ObjectIDFromHex(request.PostId)

	post, err := p.PostService.GetById(objectId)
	if err != nil {
		return nil, err
	}
	comment := mapper.MapNewComment(request.Comment)
	err = p.PostService.CreateComment(post, comment)
	if err != nil {
		return nil, err
	}

	return &post_service.CreateCommentResponse{
		Comment: request.Comment,
	}, nil
}

func (p PostHandler) GetAllCommentsForPost(_ context.Context, request *post_service.GetRequest) (*post_service.GetAllCommentsResponse, error) {
	//request = p.sanitizeGetRequest(request)
	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	post, err := p.PostService.GetById(objectId)
	if err != nil {
		return nil, err
	}

	response := &post_service.GetAllCommentsResponse{Comments: []*post_service.Comment{}}
	for _, comment := range post.Comments {
		if err != nil {
			return nil, err
		}
		current := mapper.MapUserCommentsForPost(comment.Username, comment.CommentText)
		response.Comments = append(response.Comments, current)
	}

	return response, nil
}

func (p PostHandler) LikePost(ctx context.Context, request *post_service.ReactionRequest) (*post_service.Empty, error) {

	objectId, err := primitive.ObjectIDFromHex(request.PostId)

	post, err := p.PostService.GetById(objectId)
	if err != nil {
		return nil, err
	}

	err = p.PostService.LikePost(post, request.Username)
	if err != nil {
		return nil, err
	}

	return &post_service.Empty{}, nil
}

func (p PostHandler) DislikePost(ctx context.Context, request *post_service.ReactionRequest) (*post_service.Empty, error) {

	objectId, err := primitive.ObjectIDFromHex(request.PostId)

	post, err := p.PostService.GetById(objectId)
	if err != nil {

		return nil, err
	}

	err = p.PostService.DislikePost(post, request.Username)
	if err != nil {

		return nil, err
	}

	return &post_service.Empty{}, nil
}

func (p PostHandler) GetAllReactionsForPost(_ context.Context, request *post_service.GetRequest) (*post_service.GetReactionsResponse, error) {
	policy := bluemonday.UGCPolicy()
	request.Id = strings.TrimSpace(policy.Sanitize(request.Id))

	objectId, err := primitive.ObjectIDFromHex(request.Id)

	post, err := p.PostService.GetById(objectId)
	if err != nil {
		return nil, err
	}
	likesNum, dislikesNum := mapper.FindNumberOfReactions(post)
	response := &post_service.GetReactionsResponse{}
	response.DislikesNumber = int32(dislikesNum)
	response.LikesNumber = int32(likesNum)

	return response, nil
}

func (p PostHandler) CreateJobOffer(_ context.Context, request *post_service.CreateJobOfferRequest) (*post_service.Empty, error) {

	offer := mapper.MapNewJobOffer(request.JobOffer)
	err := p.PostService.CreateJobOffer(offer)
	if err != nil {
		return nil, err
	}
	return &post_service.Empty{}, nil
}

func (p PostHandler) GetAllJobOffers(_ context.Context, _ *post_service.Empty) (*post_service.GetAllJobOffers, error) {

	offers, err := p.PostService.GetAllJobOffers()
	if err != nil {

		return nil, err
	}
	response := &post_service.GetAllJobOffers{JobOffers: []*post_service.JobOffer{}}
	for _, offer := range offers {
		current := mapper.MapJobOfferReply(offer)
		response.JobOffers = append(response.JobOffers, current)
	}
	return response, nil
}

func (p PostHandler) getOffersByPosition(_ context.Context, request *post_service.GetOfferRequest) (*post_service.GetMultipleOfferResponse, error) {

	var offers = p.PostService.GetOffersByPosition(request.Position)
	response := &post_service.GetMultipleOfferResponse{
		JobOffers: []*post_service.JobOffer{},
	}
	fmt.Printf("Pre kreiranja response-a")
	for _, JobOffer := range offers {
		current := mapper.MapFindOffers(JobOffer)
		response.Users = append(response.Users, current)
	}
	return response, nil

}

func (p PostHandler) MustEmbedUnimplementedPostServiceServer() {
	//TODO implement me
	panic("implement me")
}
