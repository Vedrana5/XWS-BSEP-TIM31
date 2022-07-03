// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: post_service.proto

package post_service

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetReactionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LikesNumber    int32 `protobuf:"varint,1,opt,name=LikesNumber,proto3" json:"LikesNumber,omitempty"`
	DislikesNumber int32 `protobuf:"varint,2,opt,name=DislikesNumber,proto3" json:"DislikesNumber,omitempty"`
}

func (x *GetReactionsResponse) Reset() {
	*x = GetReactionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReactionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReactionsResponse) ProtoMessage() {}

func (x *GetReactionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReactionsResponse.ProtoReflect.Descriptor instead.
func (*GetReactionsResponse) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetReactionsResponse) GetLikesNumber() int32 {
	if x != nil {
		return x.LikesNumber
	}
	return 0
}

func (x *GetReactionsResponse) GetDislikesNumber() int32 {
	if x != nil {
		return x.DislikesNumber
	}
	return 0
}

type ReactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId   string `protobuf:"bytes,1,opt,name=PostId,proto3" json:"PostId,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *ReactionRequest) Reset() {
	*x = ReactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReactionRequest) ProtoMessage() {}

func (x *ReactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReactionRequest.ProtoReflect.Descriptor instead.
func (*ReactionRequest) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{1}
}

func (x *ReactionRequest) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *ReactionRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type CreateCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId  string   `protobuf:"bytes,1,opt,name=PostId,proto3" json:"PostId,omitempty"`
	Comment *Comment `protobuf:"bytes,2,opt,name=Comment,proto3" json:"Comment,omitempty"`
}

func (x *CreateCommentRequest) Reset() {
	*x = CreateCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentRequest) ProtoMessage() {}

func (x *CreateCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentRequest.ProtoReflect.Descriptor instead.
func (*CreateCommentRequest) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCommentRequest) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *CreateCommentRequest) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type CreateCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment *Comment `protobuf:"bytes,1,opt,name=Comment,proto3" json:"Comment,omitempty"`
}

func (x *CreateCommentResponse) Reset() {
	*x = CreateCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentResponse) ProtoMessage() {}

func (x *CreateCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentResponse.ProtoReflect.Descriptor instead.
func (*CreateCommentResponse) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateCommentResponse) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type CreatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=Post,proto3" json:"Post,omitempty"`
}

func (x *CreatePostRequest) Reset() {
	*x = CreatePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostRequest) ProtoMessage() {}

func (x *CreatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostRequest.ProtoReflect.Descriptor instead.
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreatePostRequest) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username    string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	CommentText string `protobuf:"bytes,2,opt,name=CommentText,proto3" json:"CommentText,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{5}
}

func (x *Comment) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Comment) GetCommentText() string {
	if x != nil {
		return x.CommentText
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{6}
}

type GetMultipleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*Post `protobuf:"bytes,1,rep,name=Posts,proto3" json:"Posts,omitempty"`
}

func (x *GetMultipleResponse) Reset() {
	*x = GetMultipleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMultipleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMultipleResponse) ProtoMessage() {}

func (x *GetMultipleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMultipleResponse.ProtoReflect.Descriptor instead.
func (*GetMultipleResponse) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetMultipleResponse) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAllCommentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments []*Comment `protobuf:"bytes,1,rep,name=Comments,proto3" json:"Comments,omitempty"`
}

func (x *GetAllCommentsResponse) Reset() {
	*x = GetAllCommentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCommentsResponse) ProtoMessage() {}

func (x *GetAllCommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCommentsResponse.ProtoReflect.Descriptor instead.
func (*GetAllCommentsResponse) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{9}
}

func (x *GetAllCommentsResponse) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=Post,proto3" json:"Post,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{10}
}

func (x *GetResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Username       string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	PostText       string `protobuf:"bytes,3,opt,name=PostText,proto3" json:"PostText,omitempty"`
	ImagePaths     string `protobuf:"bytes,4,opt,name=ImagePaths,proto3" json:"ImagePaths,omitempty"`
	DatePosted     string `protobuf:"bytes,6,opt,name=DatePosted,proto3" json:"DatePosted,omitempty"`
	LikesNumber    int32  `protobuf:"varint,7,opt,name=LikesNumber,proto3" json:"LikesNumber,omitempty"`
	DislikesNumber int32  `protobuf:"varint,8,opt,name=DislikesNumber,proto3" json:"DislikesNumber,omitempty"`
	CommentsNumber int32  `protobuf:"varint,10,opt,name=CommentsNumber,proto3" json:"CommentsNumber,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{11}
}

func (x *Post) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Post) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Post) GetPostText() string {
	if x != nil {
		return x.PostText
	}
	return ""
}

func (x *Post) GetImagePaths() string {
	if x != nil {
		return x.ImagePaths
	}
	return ""
}

func (x *Post) GetDatePosted() string {
	if x != nil {
		return x.DatePosted
	}
	return ""
}

func (x *Post) GetLikesNumber() int32 {
	if x != nil {
		return x.LikesNumber
	}
	return 0
}

func (x *Post) GetDislikesNumber() int32 {
	if x != nil {
		return x.DislikesNumber
	}
	return 0
}

func (x *Post) GetCommentsNumber() int32 {
	if x != nil {
		return x.CommentsNumber
	}
	return 0
}

var File_post_service_proto protoreflect.FileDescriptor

var file_post_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x60, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x69, 0x6b, 0x65,
	0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x4c,
	0x69, 0x6b, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x44, 0x69,
	0x73, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0e, 0x44, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x45, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5f, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x07, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x48, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0x3b, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x50, 0x6f, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x50, 0x6f, 0x73,
	0x74, 0x22, 0x47, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x78, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x3f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x50, 0x6f,
	0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x50,
	0x6f, 0x73, 0x74, 0x73, 0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x64, 0x22, 0x4b, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x08,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0x35, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x22, 0x80, 0x02, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50,
	0x6f, 0x73, 0x74, 0x54, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50,
	0x6f, 0x73, 0x74, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x50, 0x61, 0x74, 0x68, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x69, 0x6b, 0x65, 0x73,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x4c, 0x69,
	0x6b, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x44, 0x69, 0x73,
	0x6c, 0x69, 0x6b, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0e, 0x44, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x26, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0xa3, 0x07, 0x0a, 0x0b, 0x50, 0x6f,
	0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x03, 0x67, 0x65, 0x74,
	0x12, 0x18, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f,
	0x70, 0x6f, 0x73, 0x74, 0x2f, 0x7b, 0x49, 0x64, 0x7d, 0x12, 0x68, 0x0a, 0x10, 0x67, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x11, 0x12, 0x0f, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b,
	0x49, 0x64, 0x7d, 0x12, 0x4f, 0x0a, 0x06, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x13, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x21, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x2f,
	0x70, 0x6f, 0x73, 0x74, 0x12, 0x53, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1f,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x05, 0x2f, 0x70,
	0x6f, 0x73, 0x74, 0x3a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x74, 0x0a, 0x15, 0x67, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x6f, 0x72, 0x50, 0x6f,
	0x73, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x70, 0x6f, 0x73,
	0x74, 0x2f, 0x7b, 0x49, 0x64, 0x7d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x81, 0x01, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x22, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x21, 0x22, 0x16, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x7b, 0x50, 0x6f, 0x73, 0x74, 0x49,
	0x64, 0x7d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x07, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x5e, 0x0a, 0x08, 0x6c, 0x69, 0x6b, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x1d, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x70, 0x6f,
	0x73, 0x74, 0x2f, 0x7b, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x7d, 0x2f, 0x6c, 0x69, 0x6b, 0x65,
	0x3a, 0x01, 0x2a, 0x12, 0x64, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16,
	0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x7b, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x7d, 0x2f, 0x64,
	0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x74, 0x0a, 0x16, 0x67, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x6f, 0x72, 0x50,
	0x6f, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x70, 0x6f, 0x73, 0x74,
	0x2f, 0x7b, 0x49, 0x64, 0x7d, 0x2f, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0x22, 0x5a, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_post_service_proto_rawDescOnce sync.Once
	file_post_service_proto_rawDescData = file_post_service_proto_rawDesc
)

func file_post_service_proto_rawDescGZIP() []byte {
	file_post_service_proto_rawDescOnce.Do(func() {
		file_post_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_post_service_proto_rawDescData)
	})
	return file_post_service_proto_rawDescData
}

var file_post_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_post_service_proto_goTypes = []interface{}{
	(*GetReactionsResponse)(nil),   // 0: post_service.GetReactionsResponse
	(*ReactionRequest)(nil),        // 1: post_service.ReactionRequest
	(*CreateCommentRequest)(nil),   // 2: post_service.CreateCommentRequest
	(*CreateCommentResponse)(nil),  // 3: post_service.CreateCommentResponse
	(*CreatePostRequest)(nil),      // 4: post_service.CreatePostRequest
	(*Comment)(nil),                // 5: post_service.Comment
	(*Empty)(nil),                  // 6: post_service.Empty
	(*GetMultipleResponse)(nil),    // 7: post_service.GetMultipleResponse
	(*GetRequest)(nil),             // 8: post_service.GetRequest
	(*GetAllCommentsResponse)(nil), // 9: post_service.GetAllCommentsResponse
	(*GetResponse)(nil),            // 10: post_service.GetResponse
	(*Post)(nil),                   // 11: post_service.Post
}
var file_post_service_proto_depIdxs = []int32{
	5,  // 0: post_service.CreateCommentRequest.Comment:type_name -> post_service.Comment
	5,  // 1: post_service.CreateCommentResponse.Comment:type_name -> post_service.Comment
	11, // 2: post_service.CreatePostRequest.Post:type_name -> post_service.Post
	11, // 3: post_service.GetMultipleResponse.Posts:type_name -> post_service.Post
	5,  // 4: post_service.GetAllCommentsResponse.Comments:type_name -> post_service.Comment
	11, // 5: post_service.GetResponse.Post:type_name -> post_service.Post
	8,  // 6: post_service.PostService.get:input_type -> post_service.GetRequest
	8,  // 7: post_service.PostService.getAllByUsername:input_type -> post_service.GetRequest
	6,  // 8: post_service.PostService.getAll:input_type -> post_service.Empty
	4,  // 9: post_service.PostService.create:input_type -> post_service.CreatePostRequest
	8,  // 10: post_service.PostService.getAllCommentsForPost:input_type -> post_service.GetRequest
	2,  // 11: post_service.PostService.createComment:input_type -> post_service.CreateCommentRequest
	1,  // 12: post_service.PostService.likePost:input_type -> post_service.ReactionRequest
	1,  // 13: post_service.PostService.dislikePost:input_type -> post_service.ReactionRequest
	8,  // 14: post_service.PostService.getAllReactionsForPost:input_type -> post_service.GetRequest
	10, // 15: post_service.PostService.get:output_type -> post_service.GetResponse
	7,  // 16: post_service.PostService.getAllByUsername:output_type -> post_service.GetMultipleResponse
	7,  // 17: post_service.PostService.getAll:output_type -> post_service.GetMultipleResponse
	6,  // 18: post_service.PostService.create:output_type -> post_service.Empty
	9,  // 19: post_service.PostService.getAllCommentsForPost:output_type -> post_service.GetAllCommentsResponse
	3,  // 20: post_service.PostService.createComment:output_type -> post_service.CreateCommentResponse
	6,  // 21: post_service.PostService.likePost:output_type -> post_service.Empty
	6,  // 22: post_service.PostService.dislikePost:output_type -> post_service.Empty
	0,  // 23: post_service.PostService.getAllReactionsForPost:output_type -> post_service.GetReactionsResponse
	15, // [15:24] is the sub-list for method output_type
	6,  // [6:15] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_post_service_proto_init() }
func file_post_service_proto_init() {
	if File_post_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_post_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReactionsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReactionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMultipleResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCommentsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_post_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_post_service_proto_goTypes,
		DependencyIndexes: file_post_service_proto_depIdxs,
		MessageInfos:      file_post_service_proto_msgTypes,
	}.Build()
	File_post_service_proto = out.File
	file_post_service_proto_rawDesc = nil
	file_post_service_proto_goTypes = nil
	file_post_service_proto_depIdxs = nil
}
