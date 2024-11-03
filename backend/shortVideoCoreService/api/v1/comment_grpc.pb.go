// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: v1/comment.proto

package v1

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

// CommentServiceClient is the client API for CommentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentServiceClient interface {
	CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error)
	RemoveComment(ctx context.Context, in *RemoveCommentRequest, opts ...grpc.CallOption) (*RemoveCommentResponse, error)
	ListComment4Video(ctx context.Context, in *ListComment4VideoRequest, opts ...grpc.CallOption) (*ListComment4VideoResponse, error)
	ListChildComment4Comment(ctx context.Context, in *ListChildComment4CommentRequest, opts ...grpc.CallOption) (*ListChildComment4CommentResponse, error)
	GetCommentById(ctx context.Context, in *GetCommentByIdRequest, opts ...grpc.CallOption) (*GetCommentByIdResponse, error)
	CountComment4Video(ctx context.Context, in *CountComment4VideoRequest, opts ...grpc.CallOption) (*CountComment4VideoResponse, error)
	CountComment4User(ctx context.Context, in *CountComment4UserRequest, opts ...grpc.CallOption) (*CountComment4UserResponse, error)
}

type commentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentServiceClient(cc grpc.ClientConnInterface) CommentServiceClient {
	return &commentServiceClient{cc}
}

func (c *commentServiceClient) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error) {
	out := new(CreateCommentResponse)
	err := c.cc.Invoke(ctx, "/shortVideoCoreService.api.v1.CommentService/CreateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) RemoveComment(ctx context.Context, in *RemoveCommentRequest, opts ...grpc.CallOption) (*RemoveCommentResponse, error) {
	out := new(RemoveCommentResponse)
	err := c.cc.Invoke(ctx, "/shortVideoCoreService.api.v1.CommentService/RemoveComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) ListComment4Video(ctx context.Context, in *ListComment4VideoRequest, opts ...grpc.CallOption) (*ListComment4VideoResponse, error) {
	out := new(ListComment4VideoResponse)
	err := c.cc.Invoke(ctx, "/shortVideoCoreService.api.v1.CommentService/ListComment4Video", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) ListChildComment4Comment(ctx context.Context, in *ListChildComment4CommentRequest, opts ...grpc.CallOption) (*ListChildComment4CommentResponse, error) {
	out := new(ListChildComment4CommentResponse)
	err := c.cc.Invoke(ctx, "/shortVideoCoreService.api.v1.CommentService/ListChildComment4Comment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetCommentById(ctx context.Context, in *GetCommentByIdRequest, opts ...grpc.CallOption) (*GetCommentByIdResponse, error) {
	out := new(GetCommentByIdResponse)
	err := c.cc.Invoke(ctx, "/shortVideoCoreService.api.v1.CommentService/GetCommentById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) CountComment4Video(ctx context.Context, in *CountComment4VideoRequest, opts ...grpc.CallOption) (*CountComment4VideoResponse, error) {
	out := new(CountComment4VideoResponse)
	err := c.cc.Invoke(ctx, "/shortVideoCoreService.api.v1.CommentService/CountComment4Video", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) CountComment4User(ctx context.Context, in *CountComment4UserRequest, opts ...grpc.CallOption) (*CountComment4UserResponse, error) {
	out := new(CountComment4UserResponse)
	err := c.cc.Invoke(ctx, "/shortVideoCoreService.api.v1.CommentService/CountComment4User", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServiceServer is the server API for CommentService service.
// All implementations must embed UnimplementedCommentServiceServer
// for forward compatibility
type CommentServiceServer interface {
	CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error)
	RemoveComment(context.Context, *RemoveCommentRequest) (*RemoveCommentResponse, error)
	ListComment4Video(context.Context, *ListComment4VideoRequest) (*ListComment4VideoResponse, error)
	ListChildComment4Comment(context.Context, *ListChildComment4CommentRequest) (*ListChildComment4CommentResponse, error)
	GetCommentById(context.Context, *GetCommentByIdRequest) (*GetCommentByIdResponse, error)
	CountComment4Video(context.Context, *CountComment4VideoRequest) (*CountComment4VideoResponse, error)
	CountComment4User(context.Context, *CountComment4UserRequest) (*CountComment4UserResponse, error)
	mustEmbedUnimplementedCommentServiceServer()
}

// UnimplementedCommentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommentServiceServer struct {
}

func (UnimplementedCommentServiceServer) CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedCommentServiceServer) RemoveComment(context.Context, *RemoveCommentRequest) (*RemoveCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveComment not implemented")
}
func (UnimplementedCommentServiceServer) ListComment4Video(context.Context, *ListComment4VideoRequest) (*ListComment4VideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListComment4Video not implemented")
}
func (UnimplementedCommentServiceServer) ListChildComment4Comment(context.Context, *ListChildComment4CommentRequest) (*ListChildComment4CommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChildComment4Comment not implemented")
}
func (UnimplementedCommentServiceServer) GetCommentById(context.Context, *GetCommentByIdRequest) (*GetCommentByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentById not implemented")
}
func (UnimplementedCommentServiceServer) CountComment4Video(context.Context, *CountComment4VideoRequest) (*CountComment4VideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountComment4Video not implemented")
}
func (UnimplementedCommentServiceServer) CountComment4User(context.Context, *CountComment4UserRequest) (*CountComment4UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountComment4User not implemented")
}
func (UnimplementedCommentServiceServer) mustEmbedUnimplementedCommentServiceServer() {}

// UnsafeCommentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServiceServer will
// result in compilation errors.
type UnsafeCommentServiceServer interface {
	mustEmbedUnimplementedCommentServiceServer()
}

func RegisterCommentServiceServer(s grpc.ServiceRegistrar, srv CommentServiceServer) {
	s.RegisterService(&CommentService_ServiceDesc, srv)
}

func _CommentService_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortVideoCoreService.api.v1.CommentService/CreateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).CreateComment(ctx, req.(*CreateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_RemoveComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).RemoveComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortVideoCoreService.api.v1.CommentService/RemoveComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).RemoveComment(ctx, req.(*RemoveCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_ListComment4Video_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListComment4VideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).ListComment4Video(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortVideoCoreService.api.v1.CommentService/ListComment4Video",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).ListComment4Video(ctx, req.(*ListComment4VideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_ListChildComment4Comment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListChildComment4CommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).ListChildComment4Comment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortVideoCoreService.api.v1.CommentService/ListChildComment4Comment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).ListChildComment4Comment(ctx, req.(*ListChildComment4CommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetCommentById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetCommentById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortVideoCoreService.api.v1.CommentService/GetCommentById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetCommentById(ctx, req.(*GetCommentByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_CountComment4Video_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountComment4VideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).CountComment4Video(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortVideoCoreService.api.v1.CommentService/CountComment4Video",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).CountComment4Video(ctx, req.(*CountComment4VideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_CountComment4User_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountComment4UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).CountComment4User(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortVideoCoreService.api.v1.CommentService/CountComment4User",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).CountComment4User(ctx, req.(*CountComment4UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentService_ServiceDesc is the grpc.ServiceDesc for CommentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shortVideoCoreService.api.v1.CommentService",
	HandlerType: (*CommentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateComment",
			Handler:    _CommentService_CreateComment_Handler,
		},
		{
			MethodName: "RemoveComment",
			Handler:    _CommentService_RemoveComment_Handler,
		},
		{
			MethodName: "ListComment4Video",
			Handler:    _CommentService_ListComment4Video_Handler,
		},
		{
			MethodName: "ListChildComment4Comment",
			Handler:    _CommentService_ListChildComment4Comment_Handler,
		},
		{
			MethodName: "GetCommentById",
			Handler:    _CommentService_GetCommentById_Handler,
		},
		{
			MethodName: "CountComment4Video",
			Handler:    _CommentService_CountComment4Video_Handler,
		},
		{
			MethodName: "CountComment4User",
			Handler:    _CommentService_CountComment4User_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/comment.proto",
}
