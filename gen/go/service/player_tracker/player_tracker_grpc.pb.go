// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: player_tracker.proto

package player_tracker

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

// PlayerTrackerClient is the client API for PlayerTracker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlayerTrackerClient interface {
	GetPlayerServer(ctx context.Context, in *PlayerRequest, opts ...grpc.CallOption) (*GetPlayerServerResponse, error)
	GetPlayerServers(ctx context.Context, in *PlayersRequest, opts ...grpc.CallOption) (*GetPlayerServersResponse, error)
	GetServerPlayers(ctx context.Context, in *ServerIdRequest, opts ...grpc.CallOption) (*GetServerPlayersResponse, error)
	GetServerPlayerCount(ctx context.Context, in *ServerIdRequest, opts ...grpc.CallOption) (*GetServerPlayerCountResponse, error)
	GetServerTypePlayerCount(ctx context.Context, in *ServerTypeRequest, opts ...grpc.CallOption) (*ServerTypePlayerCountResponse, error)
	GetServerTypeSPlayerCount(ctx context.Context, in *ServerTypesRequest, opts ...grpc.CallOption) (*ServerTypesPlayerCountResponse, error)
}

type playerTrackerClient struct {
	cc grpc.ClientConnInterface
}

func NewPlayerTrackerClient(cc grpc.ClientConnInterface) PlayerTrackerClient {
	return &playerTrackerClient{cc}
}

func (c *playerTrackerClient) GetPlayerServer(ctx context.Context, in *PlayerRequest, opts ...grpc.CallOption) (*GetPlayerServerResponse, error) {
	out := new(GetPlayerServerResponse)
	err := c.cc.Invoke(ctx, "/towerdefence.cc.service.player_tracker.PlayerTracker/GetPlayerServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerTrackerClient) GetPlayerServers(ctx context.Context, in *PlayersRequest, opts ...grpc.CallOption) (*GetPlayerServersResponse, error) {
	out := new(GetPlayerServersResponse)
	err := c.cc.Invoke(ctx, "/towerdefence.cc.service.player_tracker.PlayerTracker/GetPlayerServers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerTrackerClient) GetServerPlayers(ctx context.Context, in *ServerIdRequest, opts ...grpc.CallOption) (*GetServerPlayersResponse, error) {
	out := new(GetServerPlayersResponse)
	err := c.cc.Invoke(ctx, "/towerdefence.cc.service.player_tracker.PlayerTracker/GetServerPlayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerTrackerClient) GetServerPlayerCount(ctx context.Context, in *ServerIdRequest, opts ...grpc.CallOption) (*GetServerPlayerCountResponse, error) {
	out := new(GetServerPlayerCountResponse)
	err := c.cc.Invoke(ctx, "/towerdefence.cc.service.player_tracker.PlayerTracker/GetServerPlayerCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerTrackerClient) GetServerTypePlayerCount(ctx context.Context, in *ServerTypeRequest, opts ...grpc.CallOption) (*ServerTypePlayerCountResponse, error) {
	out := new(ServerTypePlayerCountResponse)
	err := c.cc.Invoke(ctx, "/towerdefence.cc.service.player_tracker.PlayerTracker/GetServerTypePlayerCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerTrackerClient) GetServerTypeSPlayerCount(ctx context.Context, in *ServerTypesRequest, opts ...grpc.CallOption) (*ServerTypesPlayerCountResponse, error) {
	out := new(ServerTypesPlayerCountResponse)
	err := c.cc.Invoke(ctx, "/towerdefence.cc.service.player_tracker.PlayerTracker/GetServerTypeSPlayerCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlayerTrackerServer is the server API for PlayerTracker service.
// All implementations must embed UnimplementedPlayerTrackerServer
// for forward compatibility
type PlayerTrackerServer interface {
	GetPlayerServer(context.Context, *PlayerRequest) (*GetPlayerServerResponse, error)
	GetPlayerServers(context.Context, *PlayersRequest) (*GetPlayerServersResponse, error)
	GetServerPlayers(context.Context, *ServerIdRequest) (*GetServerPlayersResponse, error)
	GetServerPlayerCount(context.Context, *ServerIdRequest) (*GetServerPlayerCountResponse, error)
	GetServerTypePlayerCount(context.Context, *ServerTypeRequest) (*ServerTypePlayerCountResponse, error)
	GetServerTypeSPlayerCount(context.Context, *ServerTypesRequest) (*ServerTypesPlayerCountResponse, error)
	mustEmbedUnimplementedPlayerTrackerServer()
}

// UnimplementedPlayerTrackerServer must be embedded to have forward compatible implementations.
type UnimplementedPlayerTrackerServer struct {
}

func (UnimplementedPlayerTrackerServer) GetPlayerServer(context.Context, *PlayerRequest) (*GetPlayerServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerServer not implemented")
}
func (UnimplementedPlayerTrackerServer) GetPlayerServers(context.Context, *PlayersRequest) (*GetPlayerServersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerServers not implemented")
}
func (UnimplementedPlayerTrackerServer) GetServerPlayers(context.Context, *ServerIdRequest) (*GetServerPlayersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerPlayers not implemented")
}
func (UnimplementedPlayerTrackerServer) GetServerPlayerCount(context.Context, *ServerIdRequest) (*GetServerPlayerCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerPlayerCount not implemented")
}
func (UnimplementedPlayerTrackerServer) GetServerTypePlayerCount(context.Context, *ServerTypeRequest) (*ServerTypePlayerCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerTypePlayerCount not implemented")
}
func (UnimplementedPlayerTrackerServer) GetServerTypeSPlayerCount(context.Context, *ServerTypesRequest) (*ServerTypesPlayerCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerTypeSPlayerCount not implemented")
}
func (UnimplementedPlayerTrackerServer) mustEmbedUnimplementedPlayerTrackerServer() {}

// UnsafePlayerTrackerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlayerTrackerServer will
// result in compilation errors.
type UnsafePlayerTrackerServer interface {
	mustEmbedUnimplementedPlayerTrackerServer()
}

func RegisterPlayerTrackerServer(s grpc.ServiceRegistrar, srv PlayerTrackerServer) {
	s.RegisterService(&PlayerTracker_ServiceDesc, srv)
}

func _PlayerTracker_GetPlayerServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackerServer).GetPlayerServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/towerdefence.cc.service.player_tracker.PlayerTracker/GetPlayerServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackerServer).GetPlayerServer(ctx, req.(*PlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerTracker_GetPlayerServers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackerServer).GetPlayerServers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/towerdefence.cc.service.player_tracker.PlayerTracker/GetPlayerServers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackerServer).GetPlayerServers(ctx, req.(*PlayersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerTracker_GetServerPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackerServer).GetServerPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/towerdefence.cc.service.player_tracker.PlayerTracker/GetServerPlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackerServer).GetServerPlayers(ctx, req.(*ServerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerTracker_GetServerPlayerCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackerServer).GetServerPlayerCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/towerdefence.cc.service.player_tracker.PlayerTracker/GetServerPlayerCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackerServer).GetServerPlayerCount(ctx, req.(*ServerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerTracker_GetServerTypePlayerCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackerServer).GetServerTypePlayerCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/towerdefence.cc.service.player_tracker.PlayerTracker/GetServerTypePlayerCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackerServer).GetServerTypePlayerCount(ctx, req.(*ServerTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerTracker_GetServerTypeSPlayerCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackerServer).GetServerTypeSPlayerCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/towerdefence.cc.service.player_tracker.PlayerTracker/GetServerTypeSPlayerCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackerServer).GetServerTypeSPlayerCount(ctx, req.(*ServerTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PlayerTracker_ServiceDesc is the grpc.ServiceDesc for PlayerTracker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlayerTracker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "towerdefence.cc.service.player_tracker.PlayerTracker",
	HandlerType: (*PlayerTrackerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlayerServer",
			Handler:    _PlayerTracker_GetPlayerServer_Handler,
		},
		{
			MethodName: "GetPlayerServers",
			Handler:    _PlayerTracker_GetPlayerServers_Handler,
		},
		{
			MethodName: "GetServerPlayers",
			Handler:    _PlayerTracker_GetServerPlayers_Handler,
		},
		{
			MethodName: "GetServerPlayerCount",
			Handler:    _PlayerTracker_GetServerPlayerCount_Handler,
		},
		{
			MethodName: "GetServerTypePlayerCount",
			Handler:    _PlayerTracker_GetServerTypePlayerCount_Handler,
		},
		{
			MethodName: "GetServerTypeSPlayerCount",
			Handler:    _PlayerTracker_GetServerTypeSPlayerCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "player_tracker.proto",
}
