// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: message_handler/grpc.proto

package messagehandler

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MessageHandler_SendPrivateMessage_FullMethodName = "/emortal.grpc.messagehandler.MessageHandler/SendPrivateMessage"
)

// MessageHandlerClient is the client API for MessageHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageHandlerClient interface {
	SendPrivateMessage(ctx context.Context, in *PrivateMessageRequest, opts ...grpc.CallOption) (*PrivateMessageResponse, error)
}

type messageHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageHandlerClient(cc grpc.ClientConnInterface) MessageHandlerClient {
	return &messageHandlerClient{cc}
}

func (c *messageHandlerClient) SendPrivateMessage(ctx context.Context, in *PrivateMessageRequest, opts ...grpc.CallOption) (*PrivateMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PrivateMessageResponse)
	err := c.cc.Invoke(ctx, MessageHandler_SendPrivateMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageHandlerServer is the server API for MessageHandler service.
// All implementations must embed UnimplementedMessageHandlerServer
// for forward compatibility.
type MessageHandlerServer interface {
	SendPrivateMessage(context.Context, *PrivateMessageRequest) (*PrivateMessageResponse, error)
	mustEmbedUnimplementedMessageHandlerServer()
}

// UnimplementedMessageHandlerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMessageHandlerServer struct{}

func (UnimplementedMessageHandlerServer) SendPrivateMessage(context.Context, *PrivateMessageRequest) (*PrivateMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPrivateMessage not implemented")
}
func (UnimplementedMessageHandlerServer) mustEmbedUnimplementedMessageHandlerServer() {}
func (UnimplementedMessageHandlerServer) testEmbeddedByValue()                        {}

// UnsafeMessageHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageHandlerServer will
// result in compilation errors.
type UnsafeMessageHandlerServer interface {
	mustEmbedUnimplementedMessageHandlerServer()
}

func RegisterMessageHandlerServer(s grpc.ServiceRegistrar, srv MessageHandlerServer) {
	// If the following call pancis, it indicates UnimplementedMessageHandlerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MessageHandler_ServiceDesc, srv)
}

func _MessageHandler_SendPrivateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrivateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageHandlerServer).SendPrivateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageHandler_SendPrivateMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageHandlerServer).SendPrivateMessage(ctx, req.(*PrivateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageHandler_ServiceDesc is the grpc.ServiceDesc for MessageHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "emortal.grpc.messagehandler.MessageHandler",
	HandlerType: (*MessageHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendPrivateMessage",
			Handler:    _MessageHandler_SendPrivateMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message_handler/grpc.proto",
}
