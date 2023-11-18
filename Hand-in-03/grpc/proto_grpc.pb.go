// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: grpc/proto.proto

package proto

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

const (
	ChitChat_Join_FullMethodName  = "/someName.ChitChat/join"
	ChitChat_Chat_FullMethodName  = "/someName.ChitChat/chat"
	ChitChat_Trans_FullMethodName = "/someName.ChitChat/trans"
)

// ChitChatClient is the client API for ChitChat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChitChatClient interface {
	Join(ctx context.Context, in *Publish, opts ...grpc.CallOption) (*Acknowledge, error)
	Chat(ctx context.Context, in *Publish, opts ...grpc.CallOption) (*Broadcast, error)
	Trans(ctx context.Context, in *Broadcast, opts ...grpc.CallOption) (*Acknowledge, error)
}

type chitChatClient struct {
	cc grpc.ClientConnInterface
}

func NewChitChatClient(cc grpc.ClientConnInterface) ChitChatClient {
	return &chitChatClient{cc}
}

func (c *chitChatClient) Join(ctx context.Context, in *Publish, opts ...grpc.CallOption) (*Acknowledge, error) {
	out := new(Acknowledge)
	err := c.cc.Invoke(ctx, ChitChat_Join_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chitChatClient) Chat(ctx context.Context, in *Publish, opts ...grpc.CallOption) (*Broadcast, error) {
	out := new(Broadcast)
	err := c.cc.Invoke(ctx, ChitChat_Chat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chitChatClient) Trans(ctx context.Context, in *Broadcast, opts ...grpc.CallOption) (*Acknowledge, error) {
	out := new(Acknowledge)
	err := c.cc.Invoke(ctx, ChitChat_Trans_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChitChatServer is the server API for ChitChat service.
// All implementations must embed UnimplementedChitChatServer
// for forward compatibility
type ChitChatServer interface {
	Join(context.Context, *Publish) (*Acknowledge, error)
	Chat(context.Context, *Publish) (*Broadcast, error)
	Trans(context.Context, *Broadcast) (*Acknowledge, error)
	mustEmbedUnimplementedChitChatServer()
}

// UnimplementedChitChatServer must be embedded to have forward compatible implementations.
type UnimplementedChitChatServer struct {
}

func (UnimplementedChitChatServer) Join(context.Context, *Publish) (*Acknowledge, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedChitChatServer) Chat(context.Context, *Publish) (*Broadcast, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (UnimplementedChitChatServer) Trans(context.Context, *Broadcast) (*Acknowledge, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Trans not implemented")
}
func (UnimplementedChitChatServer) mustEmbedUnimplementedChitChatServer() {}

// UnsafeChitChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChitChatServer will
// result in compilation errors.
type UnsafeChitChatServer interface {
	mustEmbedUnimplementedChitChatServer()
}

func RegisterChitChatServer(s grpc.ServiceRegistrar, srv ChitChatServer) {
	s.RegisterService(&ChitChat_ServiceDesc, srv)
}

func _ChitChat_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Publish)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChitChatServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChitChat_Join_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChitChatServer).Join(ctx, req.(*Publish))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChitChat_Chat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Publish)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChitChatServer).Chat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChitChat_Chat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChitChatServer).Chat(ctx, req.(*Publish))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChitChat_Trans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Broadcast)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChitChatServer).Trans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChitChat_Trans_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChitChatServer).Trans(ctx, req.(*Broadcast))
	}
	return interceptor(ctx, in, info, handler)
}

// ChitChat_ServiceDesc is the grpc.ServiceDesc for ChitChat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChitChat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "someName.ChitChat",
	HandlerType: (*ChitChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "join",
			Handler:    _ChitChat_Join_Handler,
		},
		{
			MethodName: "chat",
			Handler:    _ChitChat_Chat_Handler,
		},
		{
			MethodName: "trans",
			Handler:    _ChitChat_Trans_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto.proto",
}
