// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: permission.proto

package permission

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PermissionNode_PermissionState int32

const (
	PermissionNode_ALLOW PermissionNode_PermissionState = 0
	PermissionNode_DENY  PermissionNode_PermissionState = 1
)

// Enum value maps for PermissionNode_PermissionState.
var (
	PermissionNode_PermissionState_name = map[int32]string{
		0: "ALLOW",
		1: "DENY",
	}
	PermissionNode_PermissionState_value = map[string]int32{
		"ALLOW": 0,
		"DENY":  1,
	}
)

func (x PermissionNode_PermissionState) Enum() *PermissionNode_PermissionState {
	p := new(PermissionNode_PermissionState)
	*p = x
	return p
}

func (x PermissionNode_PermissionState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PermissionNode_PermissionState) Descriptor() protoreflect.EnumDescriptor {
	return file_permission_proto_enumTypes[0].Descriptor()
}

func (PermissionNode_PermissionState) Type() protoreflect.EnumType {
	return &file_permission_proto_enumTypes[0]
}

func (x PermissionNode_PermissionState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PermissionNode_PermissionState.Descriptor instead.
func (PermissionNode_PermissionState) EnumDescriptor() ([]byte, []int) {
	return file_permission_proto_rawDescGZIP(), []int{1, 0}
}

type PlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
}

func (x *PlayerRequest) Reset() {
	*x = PlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerRequest) ProtoMessage() {}

func (x *PlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_permission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerRequest.ProtoReflect.Descriptor instead.
func (*PlayerRequest) Descriptor() ([]byte, []int) {
	return file_permission_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type PermissionNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node  string                         `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	State PermissionNode_PermissionState `protobuf:"varint,2,opt,name=state,proto3,enum=towerdefence.cc.service.permission.PermissionNode_PermissionState" json:"state,omitempty"`
}

func (x *PermissionNode) Reset() {
	*x = PermissionNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionNode) ProtoMessage() {}

func (x *PermissionNode) ProtoReflect() protoreflect.Message {
	mi := &file_permission_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionNode.ProtoReflect.Descriptor instead.
func (*PermissionNode) Descriptor() ([]byte, []int) {
	return file_permission_proto_rawDescGZIP(), []int{1}
}

func (x *PermissionNode) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *PermissionNode) GetState() PermissionNode_PermissionState {
	if x != nil {
		return x.State
	}
	return PermissionNode_ALLOW
}

type RoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Priority      uint32            `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	DisplayPrefix *string           `protobuf:"bytes,3,opt,name=display_prefix,json=displayPrefix,proto3,oneof" json:"display_prefix,omitempty"`
	DisplayName   *string           `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3,oneof" json:"display_name,omitempty"`
	Permissions   []*PermissionNode `protobuf:"bytes,5,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *RoleResponse) Reset() {
	*x = RoleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleResponse) ProtoMessage() {}

func (x *RoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_permission_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleResponse.ProtoReflect.Descriptor instead.
func (*RoleResponse) Descriptor() ([]byte, []int) {
	return file_permission_proto_rawDescGZIP(), []int{2}
}

func (x *RoleResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoleResponse) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *RoleResponse) GetDisplayPrefix() string {
	if x != nil && x.DisplayPrefix != nil {
		return *x.DisplayPrefix
	}
	return ""
}

func (x *RoleResponse) GetDisplayName() string {
	if x != nil && x.DisplayName != nil {
		return *x.DisplayName
	}
	return ""
}

func (x *RoleResponse) GetPermissions() []*PermissionNode {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type RolesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roles []*RoleResponse `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *RolesResponse) Reset() {
	*x = RolesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RolesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RolesResponse) ProtoMessage() {}

func (x *RolesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_permission_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RolesResponse.ProtoReflect.Descriptor instead.
func (*RolesResponse) Descriptor() ([]byte, []int) {
	return file_permission_proto_rawDescGZIP(), []int{3}
}

func (x *RolesResponse) GetRoles() []*RoleResponse {
	if x != nil {
		return x.Roles
	}
	return nil
}

type PlayerRolesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleIds []string `protobuf:"bytes,1,rep,name=role_ids,json=roleIds,proto3" json:"role_ids,omitempty"`
}

func (x *PlayerRolesResponse) Reset() {
	*x = PlayerRolesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerRolesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerRolesResponse) ProtoMessage() {}

func (x *PlayerRolesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_permission_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerRolesResponse.ProtoReflect.Descriptor instead.
func (*PlayerRolesResponse) Descriptor() ([]byte, []int) {
	return file_permission_proto_rawDescGZIP(), []int{4}
}

func (x *PlayerRolesResponse) GetRoleIds() []string {
	if x != nil {
		return x.RoleIds
	}
	return nil
}

type RoleCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Priority      uint32  `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	DisplayPrefix *string `protobuf:"bytes,3,opt,name=display_prefix,json=displayPrefix,proto3,oneof" json:"display_prefix,omitempty"`
	DisplayName   *string `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3,oneof" json:"display_name,omitempty"`
}

func (x *RoleCreateRequest) Reset() {
	*x = RoleCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleCreateRequest) ProtoMessage() {}

func (x *RoleCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_permission_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleCreateRequest.ProtoReflect.Descriptor instead.
func (*RoleCreateRequest) Descriptor() ([]byte, []int) {
	return file_permission_proto_rawDescGZIP(), []int{5}
}

func (x *RoleCreateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoleCreateRequest) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *RoleCreateRequest) GetDisplayPrefix() string {
	if x != nil && x.DisplayPrefix != nil {
		return *x.DisplayPrefix
	}
	return ""
}

func (x *RoleCreateRequest) GetDisplayName() string {
	if x != nil && x.DisplayName != nil {
		return *x.DisplayName
	}
	return ""
}

type RoleUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Priority         *uint32           `protobuf:"varint,2,opt,name=priority,proto3,oneof" json:"priority,omitempty"`
	DisplayPrefix    *string           `protobuf:"bytes,3,opt,name=display_prefix,json=displayPrefix,proto3,oneof" json:"display_prefix,omitempty"`
	DisplayName      *string           `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3,oneof" json:"display_name,omitempty"`
	SetPermissions   []*PermissionNode `protobuf:"bytes,5,rep,name=set_permissions,json=setPermissions,proto3" json:"set_permissions,omitempty"`
	UnsetPermissions []string          `protobuf:"bytes,6,rep,name=unset_permissions,json=unsetPermissions,proto3" json:"unset_permissions,omitempty"`
}

func (x *RoleUpdateRequest) Reset() {
	*x = RoleUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleUpdateRequest) ProtoMessage() {}

func (x *RoleUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_permission_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleUpdateRequest.ProtoReflect.Descriptor instead.
func (*RoleUpdateRequest) Descriptor() ([]byte, []int) {
	return file_permission_proto_rawDescGZIP(), []int{6}
}

func (x *RoleUpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoleUpdateRequest) GetPriority() uint32 {
	if x != nil && x.Priority != nil {
		return *x.Priority
	}
	return 0
}

func (x *RoleUpdateRequest) GetDisplayPrefix() string {
	if x != nil && x.DisplayPrefix != nil {
		return *x.DisplayPrefix
	}
	return ""
}

func (x *RoleUpdateRequest) GetDisplayName() string {
	if x != nil && x.DisplayName != nil {
		return *x.DisplayName
	}
	return ""
}

func (x *RoleUpdateRequest) GetSetPermissions() []*PermissionNode {
	if x != nil {
		return x.SetPermissions
	}
	return nil
}

func (x *RoleUpdateRequest) GetUnsetPermissions() []string {
	if x != nil {
		return x.UnsetPermissions
	}
	return nil
}

type AddRoleToPlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	RoleId   string `protobuf:"bytes,2,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
}

func (x *AddRoleToPlayerRequest) Reset() {
	*x = AddRoleToPlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRoleToPlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRoleToPlayerRequest) ProtoMessage() {}

func (x *AddRoleToPlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_permission_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRoleToPlayerRequest.ProtoReflect.Descriptor instead.
func (*AddRoleToPlayerRequest) Descriptor() ([]byte, []int) {
	return file_permission_proto_rawDescGZIP(), []int{7}
}

func (x *AddRoleToPlayerRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *AddRoleToPlayerRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

type RemoveRoleFromPlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	RoleId   string `protobuf:"bytes,2,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
}

func (x *RemoveRoleFromPlayerRequest) Reset() {
	*x = RemoveRoleFromPlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRoleFromPlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRoleFromPlayerRequest) ProtoMessage() {}

func (x *RemoveRoleFromPlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_permission_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRoleFromPlayerRequest.ProtoReflect.Descriptor instead.
func (*RemoveRoleFromPlayerRequest) Descriptor() ([]byte, []int) {
	return file_permission_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveRoleFromPlayerRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *RemoveRoleFromPlayerRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

var File_permission_proto protoreflect.FileDescriptor

var file_permission_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x22, 0x74, 0x6f, 0x77, 0x65, 0x72, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65,
	0x2e, 0x63, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x22, 0xa6, 0x01, 0x0a, 0x0e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x58, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x42, 0x2e, 0x74, 0x6f, 0x77, 0x65, 0x72, 0x64,
	0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x22, 0x26, 0x0a, 0x0f, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x4e, 0x59, 0x10, 0x01, 0x22, 0x88, 0x02, 0x0a, 0x0c, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x2a, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x54, 0x0a, 0x0b, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x32, 0x2e, 0x74, 0x6f, 0x77, 0x65, 0x72, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e,
	0x63, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x57, 0x0a, 0x0d, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x74, 0x6f, 0x77, 0x65, 0x72, 0x64, 0x65, 0x66,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x30,
	0x0a, 0x13, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x73,
	0x22, 0xb7, 0x01, 0x0a, 0x11, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x2a, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x88, 0x01, 0x01, 0x12, 0x26,
	0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xd3, 0x02, 0x0a, 0x11, 0x52,
	0x6f, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1f, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x48, 0x00, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x88, 0x01,
	0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0d, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x5b, 0x0a, 0x0f, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32,
	0x2e, 0x74, 0x6f, 0x77, 0x65, 0x72, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x63,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x0e, 0x73, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x75, 0x6e, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x75,
	0x6e, 0x73, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x42, 0x11, 0x0a, 0x0f,
	0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x42,
	0x0f, 0x0a, 0x0d, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x4e, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x6f, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64,
	0x22, 0x53, 0x0a, 0x1b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x46, 0x72,
	0x6f, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x6f, 0x6c, 0x65, 0x49, 0x64, 0x32, 0xae, 0x05, 0x0a, 0x11, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x31, 0x2e, 0x74, 0x6f, 0x77, 0x65, 0x72, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x63,
	0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x7c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x6f, 0x6c, 0x65, 0x73, 0x12, 0x31, 0x2e, 0x74, 0x6f, 0x77, 0x65, 0x72, 0x64, 0x65, 0x66, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x63, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x74, 0x6f, 0x77, 0x65, 0x72, 0x64,
	0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x75, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x35,
	0x2e, 0x74, 0x6f, 0x77, 0x65, 0x72, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x63,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x74, 0x6f, 0x77, 0x65, 0x72, 0x64, 0x65, 0x66,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x75, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x35, 0x2e, 0x74, 0x6f, 0x77, 0x65, 0x72, 0x64, 0x65, 0x66,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x74,
	0x6f, 0x77, 0x65, 0x72, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x63, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65,
	0x0a, 0x0f, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x12, 0x3a, 0x2e, 0x74, 0x6f, 0x77, 0x65, 0x72, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65,
	0x2e, 0x63, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x6f,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x6f, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x3f, 0x2e,
	0x74, 0x6f, 0x77, 0x65, 0x72, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x63, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x46, 0x72, 0x6f,
	0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x3e, 0x0a, 0x17, 0x64, 0x65, 0x76, 0x2e, 0x65, 0x6d,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x42, 0x0f, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x5a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_permission_proto_rawDescOnce sync.Once
	file_permission_proto_rawDescData = file_permission_proto_rawDesc
)

func file_permission_proto_rawDescGZIP() []byte {
	file_permission_proto_rawDescOnce.Do(func() {
		file_permission_proto_rawDescData = protoimpl.X.CompressGZIP(file_permission_proto_rawDescData)
	})
	return file_permission_proto_rawDescData
}

var file_permission_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_permission_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_permission_proto_goTypes = []interface{}{
	(PermissionNode_PermissionState)(0), // 0: towerdefence.cc.service.permission.PermissionNode.PermissionState
	(*PlayerRequest)(nil),               // 1: towerdefence.cc.service.permission.PlayerRequest
	(*PermissionNode)(nil),              // 2: towerdefence.cc.service.permission.PermissionNode
	(*RoleResponse)(nil),                // 3: towerdefence.cc.service.permission.RoleResponse
	(*RolesResponse)(nil),               // 4: towerdefence.cc.service.permission.RolesResponse
	(*PlayerRolesResponse)(nil),         // 5: towerdefence.cc.service.permission.PlayerRolesResponse
	(*RoleCreateRequest)(nil),           // 6: towerdefence.cc.service.permission.RoleCreateRequest
	(*RoleUpdateRequest)(nil),           // 7: towerdefence.cc.service.permission.RoleUpdateRequest
	(*AddRoleToPlayerRequest)(nil),      // 8: towerdefence.cc.service.permission.AddRoleToPlayerRequest
	(*RemoveRoleFromPlayerRequest)(nil), // 9: towerdefence.cc.service.permission.RemoveRoleFromPlayerRequest
	(*emptypb.Empty)(nil),               // 10: google.protobuf.Empty
}
var file_permission_proto_depIdxs = []int32{
	0,  // 0: towerdefence.cc.service.permission.PermissionNode.state:type_name -> towerdefence.cc.service.permission.PermissionNode.PermissionState
	2,  // 1: towerdefence.cc.service.permission.RoleResponse.permissions:type_name -> towerdefence.cc.service.permission.PermissionNode
	3,  // 2: towerdefence.cc.service.permission.RolesResponse.roles:type_name -> towerdefence.cc.service.permission.RoleResponse
	2,  // 3: towerdefence.cc.service.permission.RoleUpdateRequest.set_permissions:type_name -> towerdefence.cc.service.permission.PermissionNode
	10, // 4: towerdefence.cc.service.permission.PermissionService.GetRoles:input_type -> google.protobuf.Empty
	1,  // 5: towerdefence.cc.service.permission.PermissionService.GetPlayerRoles:input_type -> towerdefence.cc.service.permission.PlayerRequest
	6,  // 6: towerdefence.cc.service.permission.PermissionService.CreateRole:input_type -> towerdefence.cc.service.permission.RoleCreateRequest
	7,  // 7: towerdefence.cc.service.permission.PermissionService.UpdateRole:input_type -> towerdefence.cc.service.permission.RoleUpdateRequest
	8,  // 8: towerdefence.cc.service.permission.PermissionService.AddRoleToPlayer:input_type -> towerdefence.cc.service.permission.AddRoleToPlayerRequest
	9,  // 9: towerdefence.cc.service.permission.PermissionService.RemoveRoleFromPlayer:input_type -> towerdefence.cc.service.permission.RemoveRoleFromPlayerRequest
	4,  // 10: towerdefence.cc.service.permission.PermissionService.GetRoles:output_type -> towerdefence.cc.service.permission.RolesResponse
	5,  // 11: towerdefence.cc.service.permission.PermissionService.GetPlayerRoles:output_type -> towerdefence.cc.service.permission.PlayerRolesResponse
	3,  // 12: towerdefence.cc.service.permission.PermissionService.CreateRole:output_type -> towerdefence.cc.service.permission.RoleResponse
	3,  // 13: towerdefence.cc.service.permission.PermissionService.UpdateRole:output_type -> towerdefence.cc.service.permission.RoleResponse
	10, // 14: towerdefence.cc.service.permission.PermissionService.AddRoleToPlayer:output_type -> google.protobuf.Empty
	10, // 15: towerdefence.cc.service.permission.PermissionService.RemoveRoleFromPlayer:output_type -> google.protobuf.Empty
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_permission_proto_init() }
func file_permission_proto_init() {
	if File_permission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_permission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerRequest); i {
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
		file_permission_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionNode); i {
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
		file_permission_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleResponse); i {
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
		file_permission_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RolesResponse); i {
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
		file_permission_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerRolesResponse); i {
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
		file_permission_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleCreateRequest); i {
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
		file_permission_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleUpdateRequest); i {
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
		file_permission_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRoleToPlayerRequest); i {
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
		file_permission_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRoleFromPlayerRequest); i {
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
	file_permission_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_permission_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_permission_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_permission_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_permission_proto_goTypes,
		DependencyIndexes: file_permission_proto_depIdxs,
		EnumInfos:         file_permission_proto_enumTypes,
		MessageInfos:      file_permission_proto_msgTypes,
	}.Build()
	File_permission_proto = out.File
	file_permission_proto_rawDesc = nil
	file_permission_proto_goTypes = nil
	file_permission_proto_depIdxs = nil
}
