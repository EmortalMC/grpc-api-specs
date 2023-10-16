// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: mc_player/grpc.proto

package mcplayer

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

// McPlayerClient is the client API for McPlayer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type McPlayerClient interface {
	GetPlayer(ctx context.Context, in *GetPlayerRequest, opts ...grpc.CallOption) (*GetPlayerResponse, error)
	GetPlayers(ctx context.Context, in *GetPlayersRequest, opts ...grpc.CallOption) (*GetPlayersResponse, error)
	GetPlayerByUsername(ctx context.Context, in *PlayerUsernameRequest, opts ...grpc.CallOption) (*GetPlayerByUsernameResponse, error)
	SearchPlayersByUsername(ctx context.Context, in *SearchPlayersByUsernameRequest, opts ...grpc.CallOption) (*SearchPlayersByUsernameResponse, error)
	GetLoginSessions(ctx context.Context, in *GetLoginSessionsRequest, opts ...grpc.CallOption) (*LoginSessionsResponse, error)
	GetStatTotalUniquePlayers(ctx context.Context, in *GetStatTotalUniquePlayersRequest, opts ...grpc.CallOption) (*GetStatTotalUniquePlayersResponse, error)
	GetStatTotalPlaytime(ctx context.Context, in *GetStatTotalPlaytimeRequest, opts ...grpc.CallOption) (*GetStatTotalPlaytimeResponse, error)
}

type mcPlayerClient struct {
	cc grpc.ClientConnInterface
}

func NewMcPlayerClient(cc grpc.ClientConnInterface) McPlayerClient {
	return &mcPlayerClient{cc}
}

func (c *mcPlayerClient) GetPlayer(ctx context.Context, in *GetPlayerRequest, opts ...grpc.CallOption) (*GetPlayerResponse, error) {
	out := new(GetPlayerResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.McPlayer/GetPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mcPlayerClient) GetPlayers(ctx context.Context, in *GetPlayersRequest, opts ...grpc.CallOption) (*GetPlayersResponse, error) {
	out := new(GetPlayersResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.McPlayer/GetPlayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mcPlayerClient) GetPlayerByUsername(ctx context.Context, in *PlayerUsernameRequest, opts ...grpc.CallOption) (*GetPlayerByUsernameResponse, error) {
	out := new(GetPlayerByUsernameResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.McPlayer/GetPlayerByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mcPlayerClient) SearchPlayersByUsername(ctx context.Context, in *SearchPlayersByUsernameRequest, opts ...grpc.CallOption) (*SearchPlayersByUsernameResponse, error) {
	out := new(SearchPlayersByUsernameResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.McPlayer/SearchPlayersByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mcPlayerClient) GetLoginSessions(ctx context.Context, in *GetLoginSessionsRequest, opts ...grpc.CallOption) (*LoginSessionsResponse, error) {
	out := new(LoginSessionsResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.McPlayer/GetLoginSessions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mcPlayerClient) GetStatTotalUniquePlayers(ctx context.Context, in *GetStatTotalUniquePlayersRequest, opts ...grpc.CallOption) (*GetStatTotalUniquePlayersResponse, error) {
	out := new(GetStatTotalUniquePlayersResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.McPlayer/GetStatTotalUniquePlayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mcPlayerClient) GetStatTotalPlaytime(ctx context.Context, in *GetStatTotalPlaytimeRequest, opts ...grpc.CallOption) (*GetStatTotalPlaytimeResponse, error) {
	out := new(GetStatTotalPlaytimeResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.McPlayer/GetStatTotalPlaytime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// McPlayerServer is the server API for McPlayer service.
// All implementations must embed UnimplementedMcPlayerServer
// for forward compatibility
type McPlayerServer interface {
	GetPlayer(context.Context, *GetPlayerRequest) (*GetPlayerResponse, error)
	GetPlayers(context.Context, *GetPlayersRequest) (*GetPlayersResponse, error)
	GetPlayerByUsername(context.Context, *PlayerUsernameRequest) (*GetPlayerByUsernameResponse, error)
	SearchPlayersByUsername(context.Context, *SearchPlayersByUsernameRequest) (*SearchPlayersByUsernameResponse, error)
	GetLoginSessions(context.Context, *GetLoginSessionsRequest) (*LoginSessionsResponse, error)
	GetStatTotalUniquePlayers(context.Context, *GetStatTotalUniquePlayersRequest) (*GetStatTotalUniquePlayersResponse, error)
	GetStatTotalPlaytime(context.Context, *GetStatTotalPlaytimeRequest) (*GetStatTotalPlaytimeResponse, error)
	mustEmbedUnimplementedMcPlayerServer()
}

// UnimplementedMcPlayerServer must be embedded to have forward compatible implementations.
type UnimplementedMcPlayerServer struct {
}

func (UnimplementedMcPlayerServer) GetPlayer(context.Context, *GetPlayerRequest) (*GetPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayer not implemented")
}
func (UnimplementedMcPlayerServer) GetPlayers(context.Context, *GetPlayersRequest) (*GetPlayersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayers not implemented")
}
func (UnimplementedMcPlayerServer) GetPlayerByUsername(context.Context, *PlayerUsernameRequest) (*GetPlayerByUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerByUsername not implemented")
}
func (UnimplementedMcPlayerServer) SearchPlayersByUsername(context.Context, *SearchPlayersByUsernameRequest) (*SearchPlayersByUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPlayersByUsername not implemented")
}
func (UnimplementedMcPlayerServer) GetLoginSessions(context.Context, *GetLoginSessionsRequest) (*LoginSessionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLoginSessions not implemented")
}
func (UnimplementedMcPlayerServer) GetStatTotalUniquePlayers(context.Context, *GetStatTotalUniquePlayersRequest) (*GetStatTotalUniquePlayersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatTotalUniquePlayers not implemented")
}
func (UnimplementedMcPlayerServer) GetStatTotalPlaytime(context.Context, *GetStatTotalPlaytimeRequest) (*GetStatTotalPlaytimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatTotalPlaytime not implemented")
}
func (UnimplementedMcPlayerServer) mustEmbedUnimplementedMcPlayerServer() {}

// UnsafeMcPlayerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to McPlayerServer will
// result in compilation errors.
type UnsafeMcPlayerServer interface {
	mustEmbedUnimplementedMcPlayerServer()
}

func RegisterMcPlayerServer(s grpc.ServiceRegistrar, srv McPlayerServer) {
	s.RegisterService(&McPlayer_ServiceDesc, srv)
}

func _McPlayer_GetPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(McPlayerServer).GetPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.McPlayer/GetPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(McPlayerServer).GetPlayer(ctx, req.(*GetPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _McPlayer_GetPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(McPlayerServer).GetPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.McPlayer/GetPlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(McPlayerServer).GetPlayers(ctx, req.(*GetPlayersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _McPlayer_GetPlayerByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(McPlayerServer).GetPlayerByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.McPlayer/GetPlayerByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(McPlayerServer).GetPlayerByUsername(ctx, req.(*PlayerUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _McPlayer_SearchPlayersByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPlayersByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(McPlayerServer).SearchPlayersByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.McPlayer/SearchPlayersByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(McPlayerServer).SearchPlayersByUsername(ctx, req.(*SearchPlayersByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _McPlayer_GetLoginSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLoginSessionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(McPlayerServer).GetLoginSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.McPlayer/GetLoginSessions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(McPlayerServer).GetLoginSessions(ctx, req.(*GetLoginSessionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _McPlayer_GetStatTotalUniquePlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatTotalUniquePlayersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(McPlayerServer).GetStatTotalUniquePlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.McPlayer/GetStatTotalUniquePlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(McPlayerServer).GetStatTotalUniquePlayers(ctx, req.(*GetStatTotalUniquePlayersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _McPlayer_GetStatTotalPlaytime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatTotalPlaytimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(McPlayerServer).GetStatTotalPlaytime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.McPlayer/GetStatTotalPlaytime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(McPlayerServer).GetStatTotalPlaytime(ctx, req.(*GetStatTotalPlaytimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// McPlayer_ServiceDesc is the grpc.ServiceDesc for McPlayer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var McPlayer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "emortal.grpc.McPlayer",
	HandlerType: (*McPlayerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlayer",
			Handler:    _McPlayer_GetPlayer_Handler,
		},
		{
			MethodName: "GetPlayers",
			Handler:    _McPlayer_GetPlayers_Handler,
		},
		{
			MethodName: "GetPlayerByUsername",
			Handler:    _McPlayer_GetPlayerByUsername_Handler,
		},
		{
			MethodName: "SearchPlayersByUsername",
			Handler:    _McPlayer_SearchPlayersByUsername_Handler,
		},
		{
			MethodName: "GetLoginSessions",
			Handler:    _McPlayer_GetLoginSessions_Handler,
		},
		{
			MethodName: "GetStatTotalUniquePlayers",
			Handler:    _McPlayer_GetStatTotalUniquePlayers_Handler,
		},
		{
			MethodName: "GetStatTotalPlaytime",
			Handler:    _McPlayer_GetStatTotalPlaytime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mc_player/grpc.proto",
}

// PlayerTrackerClient is the client API for PlayerTracker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlayerTrackerClient interface {
	GetPlayerServers(ctx context.Context, in *GetPlayerServersRequest, opts ...grpc.CallOption) (*GetPlayerServersResponse, error)
	GetServerPlayers(ctx context.Context, in *GetServerPlayersRequest, opts ...grpc.CallOption) (*GetServerPlayersResponse, error)
	GetPlayerCount(ctx context.Context, in *GetPlayerCountRequest, opts ...grpc.CallOption) (*GetPlayerCountResponse, error)
	// GetFleetPlayerCounts returns the player count for each fleet
	GetFleetPlayerCounts(ctx context.Context, in *GetFleetsPlayerCountRequest, opts ...grpc.CallOption) (*GetFleetsPlayerCountResponse, error)
	// GetGlobalSummary returns every player online that can be used to create a summary.
	// You can determine the fleet by trimming the player's server_id
	GetGlobalSummary(ctx context.Context, in *GetGlobalSummaryRequest, opts ...grpc.CallOption) (*GetGlobalSummaryResponse, error)
	// GetFleetSummary returns every player online in a fleet that can be used to create a summary.
	// This is just GetGlobalSummary but restricted to a single fleet
	GetFleetSummary(ctx context.Context, in *GetFleetSummaryRequest, opts ...grpc.CallOption) (*GetFleetSummaryResponse, error)
	// GetServerSummary returns every player online in a server that can be used to create a summary.
	// This is just GetGlobalSummary but restricted to a single server
	GetServerSummary(ctx context.Context, in *GetServerSummaryRequest, opts ...grpc.CallOption) (*GetServerSummaryResponse, error)
}

type playerTrackerClient struct {
	cc grpc.ClientConnInterface
}

func NewPlayerTrackerClient(cc grpc.ClientConnInterface) PlayerTrackerClient {
	return &playerTrackerClient{cc}
}

func (c *playerTrackerClient) GetPlayerServers(ctx context.Context, in *GetPlayerServersRequest, opts ...grpc.CallOption) (*GetPlayerServersResponse, error) {
	out := new(GetPlayerServersResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.PlayerTracker/GetPlayerServers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerTrackerClient) GetServerPlayers(ctx context.Context, in *GetServerPlayersRequest, opts ...grpc.CallOption) (*GetServerPlayersResponse, error) {
	out := new(GetServerPlayersResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.PlayerTracker/GetServerPlayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerTrackerClient) GetPlayerCount(ctx context.Context, in *GetPlayerCountRequest, opts ...grpc.CallOption) (*GetPlayerCountResponse, error) {
	out := new(GetPlayerCountResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.PlayerTracker/GetPlayerCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerTrackerClient) GetFleetPlayerCounts(ctx context.Context, in *GetFleetsPlayerCountRequest, opts ...grpc.CallOption) (*GetFleetsPlayerCountResponse, error) {
	out := new(GetFleetsPlayerCountResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.PlayerTracker/GetFleetPlayerCounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerTrackerClient) GetGlobalSummary(ctx context.Context, in *GetGlobalSummaryRequest, opts ...grpc.CallOption) (*GetGlobalSummaryResponse, error) {
	out := new(GetGlobalSummaryResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.PlayerTracker/GetGlobalSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerTrackerClient) GetFleetSummary(ctx context.Context, in *GetFleetSummaryRequest, opts ...grpc.CallOption) (*GetFleetSummaryResponse, error) {
	out := new(GetFleetSummaryResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.PlayerTracker/GetFleetSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerTrackerClient) GetServerSummary(ctx context.Context, in *GetServerSummaryRequest, opts ...grpc.CallOption) (*GetServerSummaryResponse, error) {
	out := new(GetServerSummaryResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.PlayerTracker/GetServerSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlayerTrackerServer is the server API for PlayerTracker service.
// All implementations must embed UnimplementedPlayerTrackerServer
// for forward compatibility
type PlayerTrackerServer interface {
	GetPlayerServers(context.Context, *GetPlayerServersRequest) (*GetPlayerServersResponse, error)
	GetServerPlayers(context.Context, *GetServerPlayersRequest) (*GetServerPlayersResponse, error)
	GetPlayerCount(context.Context, *GetPlayerCountRequest) (*GetPlayerCountResponse, error)
	// GetFleetPlayerCounts returns the player count for each fleet
	GetFleetPlayerCounts(context.Context, *GetFleetsPlayerCountRequest) (*GetFleetsPlayerCountResponse, error)
	// GetGlobalSummary returns every player online that can be used to create a summary.
	// You can determine the fleet by trimming the player's server_id
	GetGlobalSummary(context.Context, *GetGlobalSummaryRequest) (*GetGlobalSummaryResponse, error)
	// GetFleetSummary returns every player online in a fleet that can be used to create a summary.
	// This is just GetGlobalSummary but restricted to a single fleet
	GetFleetSummary(context.Context, *GetFleetSummaryRequest) (*GetFleetSummaryResponse, error)
	// GetServerSummary returns every player online in a server that can be used to create a summary.
	// This is just GetGlobalSummary but restricted to a single server
	GetServerSummary(context.Context, *GetServerSummaryRequest) (*GetServerSummaryResponse, error)
	mustEmbedUnimplementedPlayerTrackerServer()
}

// UnimplementedPlayerTrackerServer must be embedded to have forward compatible implementations.
type UnimplementedPlayerTrackerServer struct {
}

func (UnimplementedPlayerTrackerServer) GetPlayerServers(context.Context, *GetPlayerServersRequest) (*GetPlayerServersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerServers not implemented")
}
func (UnimplementedPlayerTrackerServer) GetServerPlayers(context.Context, *GetServerPlayersRequest) (*GetServerPlayersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerPlayers not implemented")
}
func (UnimplementedPlayerTrackerServer) GetPlayerCount(context.Context, *GetPlayerCountRequest) (*GetPlayerCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerCount not implemented")
}
func (UnimplementedPlayerTrackerServer) GetFleetPlayerCounts(context.Context, *GetFleetsPlayerCountRequest) (*GetFleetsPlayerCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFleetPlayerCounts not implemented")
}
func (UnimplementedPlayerTrackerServer) GetGlobalSummary(context.Context, *GetGlobalSummaryRequest) (*GetGlobalSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGlobalSummary not implemented")
}
func (UnimplementedPlayerTrackerServer) GetFleetSummary(context.Context, *GetFleetSummaryRequest) (*GetFleetSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFleetSummary not implemented")
}
func (UnimplementedPlayerTrackerServer) GetServerSummary(context.Context, *GetServerSummaryRequest) (*GetServerSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerSummary not implemented")
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

func _PlayerTracker_GetPlayerCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackerServer).GetPlayerCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.PlayerTracker/GetPlayerCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackerServer).GetPlayerCount(ctx, req.(*GetPlayerCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerTracker_GetFleetPlayerCounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFleetsPlayerCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackerServer).GetFleetPlayerCounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.PlayerTracker/GetFleetPlayerCounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackerServer).GetFleetPlayerCounts(ctx, req.(*GetFleetsPlayerCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerTracker_GetGlobalSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGlobalSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackerServer).GetGlobalSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.PlayerTracker/GetGlobalSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackerServer).GetGlobalSummary(ctx, req.(*GetGlobalSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerTracker_GetFleetSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFleetSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackerServer).GetFleetSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.PlayerTracker/GetFleetSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackerServer).GetFleetSummary(ctx, req.(*GetFleetSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerTracker_GetServerSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackerServer).GetServerSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.PlayerTracker/GetServerSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackerServer).GetServerSummary(ctx, req.(*GetServerSummaryRequest))
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
			MethodName: "GetPlayerServers",
			Handler:    _PlayerTracker_GetPlayerServers_Handler,
		},
		{
			MethodName: "GetServerPlayers",
			Handler:    _PlayerTracker_GetServerPlayers_Handler,
		},
		{
			MethodName: "GetPlayerCount",
			Handler:    _PlayerTracker_GetPlayerCount_Handler,
		},
		{
			MethodName: "GetFleetPlayerCounts",
			Handler:    _PlayerTracker_GetFleetPlayerCounts_Handler,
		},
		{
			MethodName: "GetGlobalSummary",
			Handler:    _PlayerTracker_GetGlobalSummary_Handler,
		},
		{
			MethodName: "GetFleetSummary",
			Handler:    _PlayerTracker_GetFleetSummary_Handler,
		},
		{
			MethodName: "GetServerSummary",
			Handler:    _PlayerTracker_GetServerSummary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mc_player/grpc.proto",
}
