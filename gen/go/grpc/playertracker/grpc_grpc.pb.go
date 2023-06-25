// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: player_tracker/grpc.proto

package playertracker

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
//
// Deprecated: Do not use.
type PlayerTrackerClient interface {
	// Deprecated: Do not use.
	GetPlayerServer(ctx context.Context, in *GetPlayerServerRequest, opts ...grpc.CallOption) (*GetPlayerServerResponse, error)
	// Deprecated: Do not use.
	GetPlayerServers(ctx context.Context, in *GetPlayerServersRequest, opts ...grpc.CallOption) (*GetPlayerServersResponse, error)
	// Deprecated: Do not use.
	GetServerPlayers(ctx context.Context, in *GetServerPlayersRequest, opts ...grpc.CallOption) (*GetServerPlayersResponse, error)
	// Deprecated: Do not use.
	GetServerPlayerCount(ctx context.Context, in *GetServerPlayerCountRequest, opts ...grpc.CallOption) (*GetServerPlayerCountResponse, error)
	// Deprecated: Do not use.
	GetServerTypePlayerCount(ctx context.Context, in *GetServerTypePlayerCountRequest, opts ...grpc.CallOption) (*ServerTypePlayerCountResponse, error)
	// Deprecated: Do not use.
	GetServerTypesPlayerCount(ctx context.Context, in *GetServerTypesPlayerCountRequest, opts ...grpc.CallOption) (*ServerTypesPlayerCountResponse, error)
}

type playerTrackerClient struct {
	cc grpc.ClientConnInterface
}

// Deprecated: Do not use.
func NewPlayerTrackerClient(cc grpc.ClientConnInterface) PlayerTrackerClient {
	return &playerTrackerClient{cc}
}

// Deprecated: Do not use.
func (c *playerTrackerClient) GetPlayerServer(ctx context.Context, in *GetPlayerServerRequest, opts ...grpc.CallOption) (*GetPlayerServerResponse, error) {
	out := new(GetPlayerServerResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.PlayerTracker/GetPlayerServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *playerTrackerClient) GetPlayerServers(ctx context.Context, in *GetPlayerServersRequest, opts ...grpc.CallOption) (*GetPlayerServersResponse, error) {
	out := new(GetPlayerServersResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.PlayerTracker/GetPlayerServers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *playerTrackerClient) GetServerPlayers(ctx context.Context, in *GetServerPlayersRequest, opts ...grpc.CallOption) (*GetServerPlayersResponse, error) {
	out := new(GetServerPlayersResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.PlayerTracker/GetServerPlayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *playerTrackerClient) GetServerPlayerCount(ctx context.Context, in *GetServerPlayerCountRequest, opts ...grpc.CallOption) (*GetServerPlayerCountResponse, error) {
	out := new(GetServerPlayerCountResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.PlayerTracker/GetServerPlayerCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *playerTrackerClient) GetServerTypePlayerCount(ctx context.Context, in *GetServerTypePlayerCountRequest, opts ...grpc.CallOption) (*ServerTypePlayerCountResponse, error) {
	out := new(ServerTypePlayerCountResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.PlayerTracker/GetServerTypePlayerCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *playerTrackerClient) GetServerTypesPlayerCount(ctx context.Context, in *GetServerTypesPlayerCountRequest, opts ...grpc.CallOption) (*ServerTypesPlayerCountResponse, error) {
	out := new(ServerTypesPlayerCountResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.PlayerTracker/GetServerTypesPlayerCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlayerTrackerServer is the server API for PlayerTracker service.
// All implementations must embed UnimplementedPlayerTrackerServer
// for forward compatibility
//
// Deprecated: Do not use.
type PlayerTrackerServer interface {
	// Deprecated: Do not use.
	GetPlayerServer(context.Context, *GetPlayerServerRequest) (*GetPlayerServerResponse, error)
	// Deprecated: Do not use.
	GetPlayerServers(context.Context, *GetPlayerServersRequest) (*GetPlayerServersResponse, error)
	// Deprecated: Do not use.
	GetServerPlayers(context.Context, *GetServerPlayersRequest) (*GetServerPlayersResponse, error)
	// Deprecated: Do not use.
	GetServerPlayerCount(context.Context, *GetServerPlayerCountRequest) (*GetServerPlayerCountResponse, error)
	// Deprecated: Do not use.
	GetServerTypePlayerCount(context.Context, *GetServerTypePlayerCountRequest) (*ServerTypePlayerCountResponse, error)
	// Deprecated: Do not use.
	GetServerTypesPlayerCount(context.Context, *GetServerTypesPlayerCountRequest) (*ServerTypesPlayerCountResponse, error)
	mustEmbedUnimplementedPlayerTrackerServer()
}

// UnimplementedPlayerTrackerServer must be embedded to have forward compatible implementations.
type UnimplementedPlayerTrackerServer struct {
}

func (UnimplementedPlayerTrackerServer) GetPlayerServer(context.Context, *GetPlayerServerRequest) (*GetPlayerServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerServer not implemented")
}
func (UnimplementedPlayerTrackerServer) GetPlayerServers(context.Context, *GetPlayerServersRequest) (*GetPlayerServersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerServers not implemented")
}
func (UnimplementedPlayerTrackerServer) GetServerPlayers(context.Context, *GetServerPlayersRequest) (*GetServerPlayersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerPlayers not implemented")
}
func (UnimplementedPlayerTrackerServer) GetServerPlayerCount(context.Context, *GetServerPlayerCountRequest) (*GetServerPlayerCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerPlayerCount not implemented")
}
func (UnimplementedPlayerTrackerServer) GetServerTypePlayerCount(context.Context, *GetServerTypePlayerCountRequest) (*ServerTypePlayerCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerTypePlayerCount not implemented")
}
func (UnimplementedPlayerTrackerServer) GetServerTypesPlayerCount(context.Context, *GetServerTypesPlayerCountRequest) (*ServerTypesPlayerCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerTypesPlayerCount not implemented")
}
func (UnimplementedPlayerTrackerServer) mustEmbedUnimplementedPlayerTrackerServer() {}

// UnsafePlayerTrackerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlayerTrackerServer will
// result in compilation errors.
type UnsafePlayerTrackerServer interface {
	mustEmbedUnimplementedPlayerTrackerServer()
}

// Deprecated: Do not use.
func RegisterPlayerTrackerServer(s grpc.ServiceRegistrar, srv PlayerTrackerServer) {
	s.RegisterService(&PlayerTracker_ServiceDesc, srv)
}

func _PlayerTracker_GetPlayerServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackerServer).GetPlayerServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.PlayerTracker/GetPlayerServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackerServer).GetPlayerServer(ctx, req.(*GetPlayerServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerTracker_GetPlayerServers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerServersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackerServer).GetPlayerServers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.PlayerTracker/GetPlayerServers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackerServer).GetPlayerServers(ctx, req.(*GetPlayerServersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerTracker_GetServerPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerPlayersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackerServer).GetServerPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.PlayerTracker/GetServerPlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackerServer).GetServerPlayers(ctx, req.(*GetServerPlayersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerTracker_GetServerPlayerCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerPlayerCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackerServer).GetServerPlayerCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.PlayerTracker/GetServerPlayerCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackerServer).GetServerPlayerCount(ctx, req.(*GetServerPlayerCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerTracker_GetServerTypePlayerCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerTypePlayerCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackerServer).GetServerTypePlayerCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.PlayerTracker/GetServerTypePlayerCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackerServer).GetServerTypePlayerCount(ctx, req.(*GetServerTypePlayerCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerTracker_GetServerTypesPlayerCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerTypesPlayerCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackerServer).GetServerTypesPlayerCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.PlayerTracker/GetServerTypesPlayerCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackerServer).GetServerTypesPlayerCount(ctx, req.(*GetServerTypesPlayerCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PlayerTracker_ServiceDesc is the grpc.ServiceDesc for PlayerTracker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlayerTracker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "emortal.grpc.PlayerTracker",
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
			MethodName: "GetServerTypesPlayerCount",
			Handler:    _PlayerTracker_GetServerTypesPlayerCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "player_tracker/grpc.proto",
}
