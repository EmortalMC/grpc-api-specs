// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: permission/grpc.proto

package permission

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

// PermissionServiceClient is the client API for PermissionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PermissionServiceClient interface {
	GetAllRoles(ctx context.Context, in *GetAllRolesRequest, opts ...grpc.CallOption) (*GetAllRolesResponse, error)
	GetPlayerRoles(ctx context.Context, in *GetPlayerRolesRequest, opts ...grpc.CallOption) (*PlayerRolesResponse, error)
	CreateRole(ctx context.Context, in *RoleCreateRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error)
	UpdateRole(ctx context.Context, in *RoleUpdateRequest, opts ...grpc.CallOption) (*UpdateRoleResponse, error)
	AddRoleToPlayer(ctx context.Context, in *AddRoleToPlayerRequest, opts ...grpc.CallOption) (*AddRoleToPlayerResponse, error)
	RemoveRoleFromPlayer(ctx context.Context, in *RemoveRoleFromPlayerRequest, opts ...grpc.CallOption) (*RemoveRoleFromPlayerResponse, error)
}

type permissionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPermissionServiceClient(cc grpc.ClientConnInterface) PermissionServiceClient {
	return &permissionServiceClient{cc}
}

func (c *permissionServiceClient) GetAllRoles(ctx context.Context, in *GetAllRolesRequest, opts ...grpc.CallOption) (*GetAllRolesResponse, error) {
	out := new(GetAllRolesResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.permission.PermissionService/GetAllRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) GetPlayerRoles(ctx context.Context, in *GetPlayerRolesRequest, opts ...grpc.CallOption) (*PlayerRolesResponse, error) {
	out := new(PlayerRolesResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.permission.PermissionService/GetPlayerRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) CreateRole(ctx context.Context, in *RoleCreateRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error) {
	out := new(CreateRoleResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.permission.PermissionService/CreateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) UpdateRole(ctx context.Context, in *RoleUpdateRequest, opts ...grpc.CallOption) (*UpdateRoleResponse, error) {
	out := new(UpdateRoleResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.permission.PermissionService/UpdateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) AddRoleToPlayer(ctx context.Context, in *AddRoleToPlayerRequest, opts ...grpc.CallOption) (*AddRoleToPlayerResponse, error) {
	out := new(AddRoleToPlayerResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.permission.PermissionService/AddRoleToPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) RemoveRoleFromPlayer(ctx context.Context, in *RemoveRoleFromPlayerRequest, opts ...grpc.CallOption) (*RemoveRoleFromPlayerResponse, error) {
	out := new(RemoveRoleFromPlayerResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.permission.PermissionService/RemoveRoleFromPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionServiceServer is the server API for PermissionService service.
// All implementations must embed UnimplementedPermissionServiceServer
// for forward compatibility
type PermissionServiceServer interface {
	GetAllRoles(context.Context, *GetAllRolesRequest) (*GetAllRolesResponse, error)
	GetPlayerRoles(context.Context, *GetPlayerRolesRequest) (*PlayerRolesResponse, error)
	CreateRole(context.Context, *RoleCreateRequest) (*CreateRoleResponse, error)
	UpdateRole(context.Context, *RoleUpdateRequest) (*UpdateRoleResponse, error)
	AddRoleToPlayer(context.Context, *AddRoleToPlayerRequest) (*AddRoleToPlayerResponse, error)
	RemoveRoleFromPlayer(context.Context, *RemoveRoleFromPlayerRequest) (*RemoveRoleFromPlayerResponse, error)
	mustEmbedUnimplementedPermissionServiceServer()
}

// UnimplementedPermissionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPermissionServiceServer struct {
}

func (UnimplementedPermissionServiceServer) GetAllRoles(context.Context, *GetAllRolesRequest) (*GetAllRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRoles not implemented")
}
func (UnimplementedPermissionServiceServer) GetPlayerRoles(context.Context, *GetPlayerRolesRequest) (*PlayerRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerRoles not implemented")
}
func (UnimplementedPermissionServiceServer) CreateRole(context.Context, *RoleCreateRequest) (*CreateRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedPermissionServiceServer) UpdateRole(context.Context, *RoleUpdateRequest) (*UpdateRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedPermissionServiceServer) AddRoleToPlayer(context.Context, *AddRoleToPlayerRequest) (*AddRoleToPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRoleToPlayer not implemented")
}
func (UnimplementedPermissionServiceServer) RemoveRoleFromPlayer(context.Context, *RemoveRoleFromPlayerRequest) (*RemoveRoleFromPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRoleFromPlayer not implemented")
}
func (UnimplementedPermissionServiceServer) mustEmbedUnimplementedPermissionServiceServer() {}

// UnsafePermissionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PermissionServiceServer will
// result in compilation errors.
type UnsafePermissionServiceServer interface {
	mustEmbedUnimplementedPermissionServiceServer()
}

func RegisterPermissionServiceServer(s grpc.ServiceRegistrar, srv PermissionServiceServer) {
	s.RegisterService(&PermissionService_ServiceDesc, srv)
}

func _PermissionService_GetAllRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).GetAllRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.permission.PermissionService/GetAllRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).GetAllRoles(ctx, req.(*GetAllRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_GetPlayerRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).GetPlayerRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.permission.PermissionService/GetPlayerRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).GetPlayerRoles(ctx, req.(*GetPlayerRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.permission.PermissionService/CreateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).CreateRole(ctx, req.(*RoleCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.permission.PermissionService/UpdateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).UpdateRole(ctx, req.(*RoleUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_AddRoleToPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRoleToPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).AddRoleToPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.permission.PermissionService/AddRoleToPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).AddRoleToPlayer(ctx, req.(*AddRoleToPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_RemoveRoleFromPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRoleFromPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).RemoveRoleFromPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.permission.PermissionService/RemoveRoleFromPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).RemoveRoleFromPlayer(ctx, req.(*RemoveRoleFromPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PermissionService_ServiceDesc is the grpc.ServiceDesc for PermissionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PermissionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "emortal.grpc.permission.PermissionService",
	HandlerType: (*PermissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllRoles",
			Handler:    _PermissionService_GetAllRoles_Handler,
		},
		{
			MethodName: "GetPlayerRoles",
			Handler:    _PermissionService_GetPlayerRoles_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _PermissionService_CreateRole_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _PermissionService_UpdateRole_Handler,
		},
		{
			MethodName: "AddRoleToPlayer",
			Handler:    _PermissionService_AddRoleToPlayer_Handler,
		},
		{
			MethodName: "RemoveRoleFromPlayer",
			Handler:    _PermissionService_RemoveRoleFromPlayer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "permission/grpc.proto",
}
