// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: connection_service.proto

package connection_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ConnectionServiceClient is the client API for ConnectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectionServiceClient interface {
	Create(ctx context.Context, in *CreateConnectionRequest, opts ...grpc.CallOption) (*Empty, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetMultipleUsernameResponse, error)
	GetConnect(ctx context.Context, in *GetUsernameRequest, opts ...grpc.CallOption) (*GetConnectionResponse, error)
	FindConnectionByUsername(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetMultipleConnectionResponse, error)
	AcceptRequest(ctx context.Context, in *EditRequest, opts ...grpc.CallOption) (*EditResponse, error)
	RejectRequest(ctx context.Context, in *EditRequest, opts ...grpc.CallOption) (*EditResponse, error)
	CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*Empty, error)
	GetMessages(ctx context.Context, in *GetUsernameRequest, opts ...grpc.CallOption) (*GetMultipleMessagesResponse, error)
	GetConnUsername(ctx context.Context, in *GetUsernamRequest, opts ...grpc.CallOption) (*GetMultipleConnectionResponse, error)
	GetUnreadMessagesByUsername(ctx context.Context, in *GetUsernamRequest, opts ...grpc.CallOption) (*GetMultipleMessagesResponse, error)
}

type connectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectionServiceClient(cc grpc.ClientConnInterface) ConnectionServiceClient {
	return &connectionServiceClient{cc}
}

func (c *connectionServiceClient) Create(ctx context.Context, in *CreateConnectionRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetMultipleUsernameResponse, error) {
	out := new(GetMultipleUsernameResponse)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) GetConnect(ctx context.Context, in *GetUsernameRequest, opts ...grpc.CallOption) (*GetConnectionResponse, error) {
	out := new(GetConnectionResponse)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/getConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) FindConnectionByUsername(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetMultipleConnectionResponse, error) {
	out := new(GetMultipleConnectionResponse)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/findConnectionByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) AcceptRequest(ctx context.Context, in *EditRequest, opts ...grpc.CallOption) (*EditResponse, error) {
	out := new(EditResponse)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/AcceptRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) RejectRequest(ctx context.Context, in *EditRequest, opts ...grpc.CallOption) (*EditResponse, error) {
	out := new(EditResponse)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/RejectRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/createMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) GetMessages(ctx context.Context, in *GetUsernameRequest, opts ...grpc.CallOption) (*GetMultipleMessagesResponse, error) {
	out := new(GetMultipleMessagesResponse)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/getMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) GetConnUsername(ctx context.Context, in *GetUsernamRequest, opts ...grpc.CallOption) (*GetMultipleConnectionResponse, error) {
	out := new(GetMultipleConnectionResponse)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/getConnUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) GetUnreadMessagesByUsername(ctx context.Context, in *GetUsernamRequest, opts ...grpc.CallOption) (*GetMultipleMessagesResponse, error) {
	out := new(GetMultipleMessagesResponse)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/getUnreadMessagesByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectionServiceServer is the server API for ConnectionService service.
// All implementations must embed UnimplementedConnectionServiceServer
// for forward compatibility
type ConnectionServiceServer interface {
	Create(context.Context, *CreateConnectionRequest) (*Empty, error)
	Get(context.Context, *GetRequest) (*GetMultipleUsernameResponse, error)
	GetConnect(context.Context, *GetUsernameRequest) (*GetConnectionResponse, error)
	FindConnectionByUsername(context.Context, *GetRequest) (*GetMultipleConnectionResponse, error)
	AcceptRequest(context.Context, *EditRequest) (*EditResponse, error)
	RejectRequest(context.Context, *EditRequest) (*EditResponse, error)
	CreateMessage(context.Context, *CreateMessageRequest) (*Empty, error)
	GetMessages(context.Context, *GetUsernameRequest) (*GetMultipleMessagesResponse, error)
	GetConnUsername(context.Context, *GetUsernamRequest) (*GetMultipleConnectionResponse, error)
	GetUnreadMessagesByUsername(context.Context, *GetUsernamRequest) (*GetMultipleMessagesResponse, error)
	MustEmbedUnimplementedConnectionServiceServer()
}

// UnimplementedConnectionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConnectionServiceServer struct {
}

func (UnimplementedConnectionServiceServer) Create(context.Context, *CreateConnectionRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedConnectionServiceServer) Get(context.Context, *GetRequest) (*GetMultipleUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedConnectionServiceServer) GetConnect(context.Context, *GetUsernameRequest) (*GetConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnect not implemented")
}
func (UnimplementedConnectionServiceServer) FindConnectionByUsername(context.Context, *GetRequest) (*GetMultipleConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindConnectionByUsername not implemented")
}
func (UnimplementedConnectionServiceServer) AcceptRequest(context.Context, *EditRequest) (*EditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptRequest not implemented")
}
func (UnimplementedConnectionServiceServer) RejectRequest(context.Context, *EditRequest) (*EditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectRequest not implemented")
}
func (UnimplementedConnectionServiceServer) CreateMessage(context.Context, *CreateMessageRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}
func (UnimplementedConnectionServiceServer) GetMessages(context.Context, *GetUsernameRequest) (*GetMultipleMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedConnectionServiceServer) GetConnUsername(context.Context, *GetUsernamRequest) (*GetMultipleConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnUsername not implemented")
}
func (UnimplementedConnectionServiceServer) GetUnreadMessagesByUsername(context.Context, *GetUsernamRequest) (*GetMultipleMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnreadMessagesByUsername not implemented")
}
func (UnimplementedConnectionServiceServer) MustEmbedUnimplementedConnectionServiceServer() {}

// UnsafeConnectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectionServiceServer will
// result in compilation errors.
type UnsafeConnectionServiceServer interface {
	MustEmbedUnimplementedConnectionServiceServer()
}

func RegisterConnectionServiceServer(s grpc.ServiceRegistrar, srv ConnectionServiceServer) {
	s.RegisterService(&ConnectionService_ServiceDesc, srv)
}

func _ConnectionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).Create(ctx, req.(*CreateConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_GetConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).GetConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/getConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).GetConnect(ctx, req.(*GetUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_FindConnectionByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).FindConnectionByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/findConnectionByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).FindConnectionByUsername(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_AcceptRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).AcceptRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/AcceptRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).AcceptRequest(ctx, req.(*EditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_RejectRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).RejectRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/RejectRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).RejectRequest(ctx, req.(*EditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/createMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).CreateMessage(ctx, req.(*CreateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_GetMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).GetMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/getMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).GetMessages(ctx, req.(*GetUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_GetConnUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsernamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).GetConnUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/getConnUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).GetConnUsername(ctx, req.(*GetUsernamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_GetUnreadMessagesByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsernamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).GetUnreadMessagesByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/getUnreadMessagesByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).GetUnreadMessagesByUsername(ctx, req.(*GetUsernamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConnectionService_ServiceDesc is the grpc.ServiceDesc for ConnectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConnectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "connection_service.ConnectionService",
	HandlerType: (*ConnectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _ConnectionService_Create_Handler,
		},
		{
			MethodName: "get",
			Handler:    _ConnectionService_Get_Handler,
		},
		{
			MethodName: "getConnect",
			Handler:    _ConnectionService_GetConnect_Handler,
		},
		{
			MethodName: "findConnectionByUsername",
			Handler:    _ConnectionService_FindConnectionByUsername_Handler,
		},
		{
			MethodName: "AcceptRequest",
			Handler:    _ConnectionService_AcceptRequest_Handler,
		},
		{
			MethodName: "RejectRequest",
			Handler:    _ConnectionService_RejectRequest_Handler,
		},
		{
			MethodName: "createMessage",
			Handler:    _ConnectionService_CreateMessage_Handler,
		},
		{
			MethodName: "getMessages",
			Handler:    _ConnectionService_GetMessages_Handler,
		},
		{
			MethodName: "getConnUsername",
			Handler:    _ConnectionService_GetConnUsername_Handler,
		},
		{
			MethodName: "getUnreadMessagesByUsername",
			Handler:    _ConnectionService_GetUnreadMessagesByUsername_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "connection_service.proto",
}
