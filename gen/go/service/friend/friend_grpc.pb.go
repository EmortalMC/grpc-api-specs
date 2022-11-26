// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: friend.proto

package friend

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

// FriendClient is the client API for Friend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FriendClient interface {
	AddFriend(ctx context.Context, in *AddFriendRequest, opts ...grpc.CallOption) (*AddFriendResponse, error)
	RemoveFriend(ctx context.Context, in *RemoveFriendRequest, opts ...grpc.CallOption) (*RemoveFriendResponse, error)
	DenyFriendRequest(ctx context.Context, in *DenyFriendRequestRequest, opts ...grpc.CallOption) (*DenyFriendRequestResponse, error)
	MassDenyFriendRequest(ctx context.Context, in *MassDenyFriendRequestRequest, opts ...grpc.CallOption) (*MassDenyFriendRequestResponse, error)
	GetFriendList(ctx context.Context, in *PlayerRequest, opts ...grpc.CallOption) (*FriendListResponse, error)
	GetPendingFriendRequestList(ctx context.Context, in *GetPendingFriendRequestListRequest, opts ...grpc.CallOption) (*PendingFriendListResponse, error)
}

type friendClient struct {
	cc grpc.ClientConnInterface
}

func NewFriendClient(cc grpc.ClientConnInterface) FriendClient {
	return &friendClient{cc}
}

func (c *friendClient) AddFriend(ctx context.Context, in *AddFriendRequest, opts ...grpc.CallOption) (*AddFriendResponse, error) {
	out := new(AddFriendResponse)
	err := c.cc.Invoke(ctx, "/towerdefence.cc.service.friend.Friend/AddFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) RemoveFriend(ctx context.Context, in *RemoveFriendRequest, opts ...grpc.CallOption) (*RemoveFriendResponse, error) {
	out := new(RemoveFriendResponse)
	err := c.cc.Invoke(ctx, "/towerdefence.cc.service.friend.Friend/RemoveFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) DenyFriendRequest(ctx context.Context, in *DenyFriendRequestRequest, opts ...grpc.CallOption) (*DenyFriendRequestResponse, error) {
	out := new(DenyFriendRequestResponse)
	err := c.cc.Invoke(ctx, "/towerdefence.cc.service.friend.Friend/DenyFriendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) MassDenyFriendRequest(ctx context.Context, in *MassDenyFriendRequestRequest, opts ...grpc.CallOption) (*MassDenyFriendRequestResponse, error) {
	out := new(MassDenyFriendRequestResponse)
	err := c.cc.Invoke(ctx, "/towerdefence.cc.service.friend.Friend/MassDenyFriendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) GetFriendList(ctx context.Context, in *PlayerRequest, opts ...grpc.CallOption) (*FriendListResponse, error) {
	out := new(FriendListResponse)
	err := c.cc.Invoke(ctx, "/towerdefence.cc.service.friend.Friend/GetFriendList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) GetPendingFriendRequestList(ctx context.Context, in *GetPendingFriendRequestListRequest, opts ...grpc.CallOption) (*PendingFriendListResponse, error) {
	out := new(PendingFriendListResponse)
	err := c.cc.Invoke(ctx, "/towerdefence.cc.service.friend.Friend/GetPendingFriendRequestList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FriendServer is the server API for Friend service.
// All implementations must embed UnimplementedFriendServer
// for forward compatibility
type FriendServer interface {
	AddFriend(context.Context, *AddFriendRequest) (*AddFriendResponse, error)
	RemoveFriend(context.Context, *RemoveFriendRequest) (*RemoveFriendResponse, error)
	DenyFriendRequest(context.Context, *DenyFriendRequestRequest) (*DenyFriendRequestResponse, error)
	MassDenyFriendRequest(context.Context, *MassDenyFriendRequestRequest) (*MassDenyFriendRequestResponse, error)
	GetFriendList(context.Context, *PlayerRequest) (*FriendListResponse, error)
	GetPendingFriendRequestList(context.Context, *GetPendingFriendRequestListRequest) (*PendingFriendListResponse, error)
	mustEmbedUnimplementedFriendServer()
}

// UnimplementedFriendServer must be embedded to have forward compatible implementations.
type UnimplementedFriendServer struct {
}

func (UnimplementedFriendServer) AddFriend(context.Context, *AddFriendRequest) (*AddFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}
func (UnimplementedFriendServer) RemoveFriend(context.Context, *RemoveFriendRequest) (*RemoveFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFriend not implemented")
}
func (UnimplementedFriendServer) DenyFriendRequest(context.Context, *DenyFriendRequestRequest) (*DenyFriendRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenyFriendRequest not implemented")
}
func (UnimplementedFriendServer) MassDenyFriendRequest(context.Context, *MassDenyFriendRequestRequest) (*MassDenyFriendRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MassDenyFriendRequest not implemented")
}
func (UnimplementedFriendServer) GetFriendList(context.Context, *PlayerRequest) (*FriendListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendList not implemented")
}
func (UnimplementedFriendServer) GetPendingFriendRequestList(context.Context, *GetPendingFriendRequestListRequest) (*PendingFriendListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPendingFriendRequestList not implemented")
}
func (UnimplementedFriendServer) mustEmbedUnimplementedFriendServer() {}

// UnsafeFriendServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FriendServer will
// result in compilation errors.
type UnsafeFriendServer interface {
	mustEmbedUnimplementedFriendServer()
}

func RegisterFriendServer(s grpc.ServiceRegistrar, srv FriendServer) {
	s.RegisterService(&Friend_ServiceDesc, srv)
}

func _Friend_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/towerdefence.cc.service.friend.Friend/AddFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).AddFriend(ctx, req.(*AddFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_RemoveFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).RemoveFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/towerdefence.cc.service.friend.Friend/RemoveFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).RemoveFriend(ctx, req.(*RemoveFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_DenyFriendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DenyFriendRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).DenyFriendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/towerdefence.cc.service.friend.Friend/DenyFriendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).DenyFriendRequest(ctx, req.(*DenyFriendRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_MassDenyFriendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MassDenyFriendRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).MassDenyFriendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/towerdefence.cc.service.friend.Friend/MassDenyFriendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).MassDenyFriendRequest(ctx, req.(*MassDenyFriendRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_GetFriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).GetFriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/towerdefence.cc.service.friend.Friend/GetFriendList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).GetFriendList(ctx, req.(*PlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_GetPendingFriendRequestList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPendingFriendRequestListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).GetPendingFriendRequestList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/towerdefence.cc.service.friend.Friend/GetPendingFriendRequestList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).GetPendingFriendRequestList(ctx, req.(*GetPendingFriendRequestListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Friend_ServiceDesc is the grpc.ServiceDesc for Friend service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Friend_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "towerdefence.cc.service.friend.Friend",
	HandlerType: (*FriendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFriend",
			Handler:    _Friend_AddFriend_Handler,
		},
		{
			MethodName: "RemoveFriend",
			Handler:    _Friend_RemoveFriend_Handler,
		},
		{
			MethodName: "DenyFriendRequest",
			Handler:    _Friend_DenyFriendRequest_Handler,
		},
		{
			MethodName: "MassDenyFriendRequest",
			Handler:    _Friend_MassDenyFriendRequest_Handler,
		},
		{
			MethodName: "GetFriendList",
			Handler:    _Friend_GetFriendList_Handler,
		},
		{
			MethodName: "GetPendingFriendRequestList",
			Handler:    _Friend_GetPendingFriendRequestList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "friend.proto",
}
