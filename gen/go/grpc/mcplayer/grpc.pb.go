// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: mc_player/grpc.proto

package mcplayer

import (
	common "github.com/emortalmc/proto-specs/gen/go/model/common"
	mcplayer "github.com/emortalmc/proto-specs/gen/go/model/mcplayer"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SearchPlayersByUsernameRequest_FilterMethod int32

const (
	SearchPlayersByUsernameRequest_NONE    SearchPlayersByUsernameRequest_FilterMethod = 0
	SearchPlayersByUsernameRequest_ONLINE  SearchPlayersByUsernameRequest_FilterMethod = 1
	SearchPlayersByUsernameRequest_FRIENDS SearchPlayersByUsernameRequest_FilterMethod = 2
)

// Enum value maps for SearchPlayersByUsernameRequest_FilterMethod.
var (
	SearchPlayersByUsernameRequest_FilterMethod_name = map[int32]string{
		0: "NONE",
		1: "ONLINE",
		2: "FRIENDS",
	}
	SearchPlayersByUsernameRequest_FilterMethod_value = map[string]int32{
		"NONE":    0,
		"ONLINE":  1,
		"FRIENDS": 2,
	}
)

func (x SearchPlayersByUsernameRequest_FilterMethod) Enum() *SearchPlayersByUsernameRequest_FilterMethod {
	p := new(SearchPlayersByUsernameRequest_FilterMethod)
	*p = x
	return p
}

func (x SearchPlayersByUsernameRequest_FilterMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SearchPlayersByUsernameRequest_FilterMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_mc_player_grpc_proto_enumTypes[0].Descriptor()
}

func (SearchPlayersByUsernameRequest_FilterMethod) Type() protoreflect.EnumType {
	return &file_mc_player_grpc_proto_enumTypes[0]
}

func (x SearchPlayersByUsernameRequest_FilterMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SearchPlayersByUsernameRequest_FilterMethod.Descriptor instead.
func (SearchPlayersByUsernameRequest_FilterMethod) EnumDescriptor() ([]byte, []int) {
	return file_mc_player_grpc_proto_rawDescGZIP(), []int{7, 0}
}

type GetPlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
}

func (x *GetPlayerRequest) Reset() {
	*x = GetPlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mc_player_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerRequest) ProtoMessage() {}

func (x *GetPlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mc_player_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerRequest.ProtoReflect.Descriptor instead.
func (*GetPlayerRequest) Descriptor() ([]byte, []int) {
	return file_mc_player_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *GetPlayerRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type GetPlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *mcplayer.McPlayer `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *GetPlayerResponse) Reset() {
	*x = GetPlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mc_player_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerResponse) ProtoMessage() {}

func (x *GetPlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mc_player_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerResponse.ProtoReflect.Descriptor instead.
func (*GetPlayerResponse) Descriptor() ([]byte, []int) {
	return file_mc_player_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *GetPlayerResponse) GetPlayer() *mcplayer.McPlayer {
	if x != nil {
		return x.Player
	}
	return nil
}

type GetPlayersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerIds []string `protobuf:"bytes,1,rep,name=player_ids,json=playerIds,proto3" json:"player_ids,omitempty"`
}

func (x *GetPlayersRequest) Reset() {
	*x = GetPlayersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mc_player_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayersRequest) ProtoMessage() {}

func (x *GetPlayersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mc_player_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayersRequest.ProtoReflect.Descriptor instead.
func (*GetPlayersRequest) Descriptor() ([]byte, []int) {
	return file_mc_player_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *GetPlayersRequest) GetPlayerIds() []string {
	if x != nil {
		return x.PlayerIds
	}
	return nil
}

type GetPlayersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players []*mcplayer.McPlayer `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *GetPlayersResponse) Reset() {
	*x = GetPlayersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mc_player_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayersResponse) ProtoMessage() {}

func (x *GetPlayersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mc_player_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayersResponse.ProtoReflect.Descriptor instead.
func (*GetPlayersResponse) Descriptor() ([]byte, []int) {
	return file_mc_player_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *GetPlayersResponse) GetPlayers() []*mcplayer.McPlayer {
	if x != nil {
		return x.Players
	}
	return nil
}

type PlayerUsernameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"` // ignore case when using the username
}

func (x *PlayerUsernameRequest) Reset() {
	*x = PlayerUsernameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mc_player_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerUsernameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerUsernameRequest) ProtoMessage() {}

func (x *PlayerUsernameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mc_player_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerUsernameRequest.ProtoReflect.Descriptor instead.
func (*PlayerUsernameRequest) Descriptor() ([]byte, []int) {
	return file_mc_player_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *PlayerUsernameRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type GetPlayerByUsernameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *mcplayer.McPlayer `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *GetPlayerByUsernameResponse) Reset() {
	*x = GetPlayerByUsernameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mc_player_grpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerByUsernameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerByUsernameResponse) ProtoMessage() {}

func (x *GetPlayerByUsernameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mc_player_grpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerByUsernameResponse.ProtoReflect.Descriptor instead.
func (*GetPlayerByUsernameResponse) Descriptor() ([]byte, []int) {
	return file_mc_player_grpc_proto_rawDescGZIP(), []int{5}
}

func (x *GetPlayerByUsernameResponse) GetPlayer() *mcplayer.McPlayer {
	if x != nil {
		return x.Player
	}
	return nil
}

type McPageablePlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Page     uint32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size     uint32 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *McPageablePlayerRequest) Reset() {
	*x = McPageablePlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mc_player_grpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *McPageablePlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*McPageablePlayerRequest) ProtoMessage() {}

func (x *McPageablePlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mc_player_grpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use McPageablePlayerRequest.ProtoReflect.Descriptor instead.
func (*McPageablePlayerRequest) Descriptor() ([]byte, []int) {
	return file_mc_player_grpc_proto_rawDescGZIP(), []int{6}
}

func (x *McPageablePlayerRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *McPageablePlayerRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *McPageablePlayerRequest) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type SearchPlayersByUsernameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IssuerId       string `protobuf:"bytes,1,opt,name=issuer_id,json=issuerId,proto3" json:"issuer_id,omitempty"`
	SearchUsername string `protobuf:"bytes,2,opt,name=search_username,json=searchUsername,proto3" json:"search_username,omitempty"`
	// starts with 0
	Pageable     *common.Pageable                            `protobuf:"bytes,3,opt,name=pageable,proto3" json:"pageable,omitempty"`
	FilterMethod SearchPlayersByUsernameRequest_FilterMethod `protobuf:"varint,4,opt,name=filter_method,json=filterMethod,proto3,enum=emortal.grpc.SearchPlayersByUsernameRequest_FilterMethod" json:"filter_method,omitempty"`
}

func (x *SearchPlayersByUsernameRequest) Reset() {
	*x = SearchPlayersByUsernameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mc_player_grpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchPlayersByUsernameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchPlayersByUsernameRequest) ProtoMessage() {}

func (x *SearchPlayersByUsernameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mc_player_grpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchPlayersByUsernameRequest.ProtoReflect.Descriptor instead.
func (*SearchPlayersByUsernameRequest) Descriptor() ([]byte, []int) {
	return file_mc_player_grpc_proto_rawDescGZIP(), []int{7}
}

func (x *SearchPlayersByUsernameRequest) GetIssuerId() string {
	if x != nil {
		return x.IssuerId
	}
	return ""
}

func (x *SearchPlayersByUsernameRequest) GetSearchUsername() string {
	if x != nil {
		return x.SearchUsername
	}
	return ""
}

func (x *SearchPlayersByUsernameRequest) GetPageable() *common.Pageable {
	if x != nil {
		return x.Pageable
	}
	return nil
}

func (x *SearchPlayersByUsernameRequest) GetFilterMethod() SearchPlayersByUsernameRequest_FilterMethod {
	if x != nil {
		return x.FilterMethod
	}
	return SearchPlayersByUsernameRequest_NONE
}

type SearchPlayersByUsernameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players  []*mcplayer.McPlayer `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
	PageData *common.PageData     `protobuf:"bytes,2,opt,name=page_data,json=pageData,proto3" json:"page_data,omitempty"`
}

func (x *SearchPlayersByUsernameResponse) Reset() {
	*x = SearchPlayersByUsernameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mc_player_grpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchPlayersByUsernameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchPlayersByUsernameResponse) ProtoMessage() {}

func (x *SearchPlayersByUsernameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mc_player_grpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchPlayersByUsernameResponse.ProtoReflect.Descriptor instead.
func (*SearchPlayersByUsernameResponse) Descriptor() ([]byte, []int) {
	return file_mc_player_grpc_proto_rawDescGZIP(), []int{8}
}

func (x *SearchPlayersByUsernameResponse) GetPlayers() []*mcplayer.McPlayer {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *SearchPlayersByUsernameResponse) GetPageData() *common.PageData {
	if x != nil {
		return x.PageData
	}
	return nil
}

type GetLoginSessionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string           `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Pageable *common.Pageable `protobuf:"bytes,2,opt,name=pageable,proto3" json:"pageable,omitempty"`
}

func (x *GetLoginSessionsRequest) Reset() {
	*x = GetLoginSessionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mc_player_grpc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLoginSessionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoginSessionsRequest) ProtoMessage() {}

func (x *GetLoginSessionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mc_player_grpc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLoginSessionsRequest.ProtoReflect.Descriptor instead.
func (*GetLoginSessionsRequest) Descriptor() ([]byte, []int) {
	return file_mc_player_grpc_proto_rawDescGZIP(), []int{9}
}

func (x *GetLoginSessionsRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *GetLoginSessionsRequest) GetPageable() *common.Pageable {
	if x != nil {
		return x.Pageable
	}
	return nil
}

type LoginSessionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sessions []*mcplayer.LoginSession `protobuf:"bytes,1,rep,name=sessions,proto3" json:"sessions,omitempty"`
	PageData *common.PageData         `protobuf:"bytes,2,opt,name=page_data,json=pageData,proto3" json:"page_data,omitempty"`
}

func (x *LoginSessionsResponse) Reset() {
	*x = LoginSessionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mc_player_grpc_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginSessionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginSessionsResponse) ProtoMessage() {}

func (x *LoginSessionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mc_player_grpc_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginSessionsResponse.ProtoReflect.Descriptor instead.
func (*LoginSessionsResponse) Descriptor() ([]byte, []int) {
	return file_mc_player_grpc_proto_rawDescGZIP(), []int{10}
}

func (x *LoginSessionsResponse) GetSessions() []*mcplayer.LoginSession {
	if x != nil {
		return x.Sessions
	}
	return nil
}

func (x *LoginSessionsResponse) GetPageData() *common.PageData {
	if x != nil {
		return x.PageData
	}
	return nil
}

var File_mc_player_grpc_proto protoreflect.FileDescriptor

var file_mc_player_grpc_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x1a, 0x16, 0x6d, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x2f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x44, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x63, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x32, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x47, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x63, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x22, 0x33, 0x0a, 0x15, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4e, 0x0a, 0x1b, 0x47, 0x65,
	0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6d, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x63, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x5e, 0x0a, 0x17, 0x4d, 0x63,
	0x50, 0x61, 0x67, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0xae, 0x02, 0x0a, 0x1e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x42, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x5e, 0x0a, 0x0d, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x39, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x42, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x31, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x4e, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x53, 0x10, 0x02, 0x22, 0x8a, 0x01, 0x0a, 0x1f,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x42, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x31, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x4d, 0x63, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x12, 0x34, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x6b, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x33, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x15, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x37, 0x0a, 0x08, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x34, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6d,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x32, 0xf2,
	0x03, 0x0a, 0x08, 0x4d, 0x63, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x4e, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x2e, 0x65, 0x6d, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x65, 0x6d, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x65, 0x6d, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x78, 0x0a, 0x17, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2c, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2d, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x42, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x60, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x65,
	0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x65, 0x0a, 0x1d, 0x64, 0x65, 0x76, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x63, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x42, 0x0d, 0x4d, 0x63, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x6d, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d,
	0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x6d, 0x63, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_mc_player_grpc_proto_rawDescOnce sync.Once
	file_mc_player_grpc_proto_rawDescData = file_mc_player_grpc_proto_rawDesc
)

func file_mc_player_grpc_proto_rawDescGZIP() []byte {
	file_mc_player_grpc_proto_rawDescOnce.Do(func() {
		file_mc_player_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_mc_player_grpc_proto_rawDescData)
	})
	return file_mc_player_grpc_proto_rawDescData
}

var file_mc_player_grpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mc_player_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_mc_player_grpc_proto_goTypes = []interface{}{
	(SearchPlayersByUsernameRequest_FilterMethod)(0), // 0: emortal.grpc.SearchPlayersByUsernameRequest.FilterMethod
	(*GetPlayerRequest)(nil),                         // 1: emortal.grpc.GetPlayerRequest
	(*GetPlayerResponse)(nil),                        // 2: emortal.grpc.GetPlayerResponse
	(*GetPlayersRequest)(nil),                        // 3: emortal.grpc.GetPlayersRequest
	(*GetPlayersResponse)(nil),                       // 4: emortal.grpc.GetPlayersResponse
	(*PlayerUsernameRequest)(nil),                    // 5: emortal.grpc.PlayerUsernameRequest
	(*GetPlayerByUsernameResponse)(nil),              // 6: emortal.grpc.GetPlayerByUsernameResponse
	(*McPageablePlayerRequest)(nil),                  // 7: emortal.grpc.McPageablePlayerRequest
	(*SearchPlayersByUsernameRequest)(nil),           // 8: emortal.grpc.SearchPlayersByUsernameRequest
	(*SearchPlayersByUsernameResponse)(nil),          // 9: emortal.grpc.SearchPlayersByUsernameResponse
	(*GetLoginSessionsRequest)(nil),                  // 10: emortal.grpc.GetLoginSessionsRequest
	(*LoginSessionsResponse)(nil),                    // 11: emortal.grpc.LoginSessionsResponse
	(*mcplayer.McPlayer)(nil),                        // 12: emortal.model.McPlayer
	(*common.Pageable)(nil),                          // 13: emortal.model.Pageable
	(*common.PageData)(nil),                          // 14: emortal.model.PageData
	(*mcplayer.LoginSession)(nil),                    // 15: emortal.model.LoginSession
}
var file_mc_player_grpc_proto_depIdxs = []int32{
	12, // 0: emortal.grpc.GetPlayerResponse.player:type_name -> emortal.model.McPlayer
	12, // 1: emortal.grpc.GetPlayersResponse.players:type_name -> emortal.model.McPlayer
	12, // 2: emortal.grpc.GetPlayerByUsernameResponse.player:type_name -> emortal.model.McPlayer
	13, // 3: emortal.grpc.SearchPlayersByUsernameRequest.pageable:type_name -> emortal.model.Pageable
	0,  // 4: emortal.grpc.SearchPlayersByUsernameRequest.filter_method:type_name -> emortal.grpc.SearchPlayersByUsernameRequest.FilterMethod
	12, // 5: emortal.grpc.SearchPlayersByUsernameResponse.players:type_name -> emortal.model.McPlayer
	14, // 6: emortal.grpc.SearchPlayersByUsernameResponse.page_data:type_name -> emortal.model.PageData
	13, // 7: emortal.grpc.GetLoginSessionsRequest.pageable:type_name -> emortal.model.Pageable
	15, // 8: emortal.grpc.LoginSessionsResponse.sessions:type_name -> emortal.model.LoginSession
	14, // 9: emortal.grpc.LoginSessionsResponse.page_data:type_name -> emortal.model.PageData
	1,  // 10: emortal.grpc.McPlayer.GetPlayer:input_type -> emortal.grpc.GetPlayerRequest
	3,  // 11: emortal.grpc.McPlayer.GetPlayers:input_type -> emortal.grpc.GetPlayersRequest
	5,  // 12: emortal.grpc.McPlayer.GetPlayerByUsername:input_type -> emortal.grpc.PlayerUsernameRequest
	8,  // 13: emortal.grpc.McPlayer.SearchPlayersByUsername:input_type -> emortal.grpc.SearchPlayersByUsernameRequest
	10, // 14: emortal.grpc.McPlayer.GetLoginSessions:input_type -> emortal.grpc.GetLoginSessionsRequest
	2,  // 15: emortal.grpc.McPlayer.GetPlayer:output_type -> emortal.grpc.GetPlayerResponse
	4,  // 16: emortal.grpc.McPlayer.GetPlayers:output_type -> emortal.grpc.GetPlayersResponse
	6,  // 17: emortal.grpc.McPlayer.GetPlayerByUsername:output_type -> emortal.grpc.GetPlayerByUsernameResponse
	9,  // 18: emortal.grpc.McPlayer.SearchPlayersByUsername:output_type -> emortal.grpc.SearchPlayersByUsernameResponse
	11, // 19: emortal.grpc.McPlayer.GetLoginSessions:output_type -> emortal.grpc.LoginSessionsResponse
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_mc_player_grpc_proto_init() }
func file_mc_player_grpc_proto_init() {
	if File_mc_player_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mc_player_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerRequest); i {
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
		file_mc_player_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerResponse); i {
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
		file_mc_player_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayersRequest); i {
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
		file_mc_player_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayersResponse); i {
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
		file_mc_player_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerUsernameRequest); i {
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
		file_mc_player_grpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerByUsernameResponse); i {
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
		file_mc_player_grpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*McPageablePlayerRequest); i {
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
		file_mc_player_grpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchPlayersByUsernameRequest); i {
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
		file_mc_player_grpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchPlayersByUsernameResponse); i {
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
		file_mc_player_grpc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLoginSessionsRequest); i {
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
		file_mc_player_grpc_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginSessionsResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mc_player_grpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mc_player_grpc_proto_goTypes,
		DependencyIndexes: file_mc_player_grpc_proto_depIdxs,
		EnumInfos:         file_mc_player_grpc_proto_enumTypes,
		MessageInfos:      file_mc_player_grpc_proto_msgTypes,
	}.Build()
	File_mc_player_grpc_proto = out.File
	file_mc_player_grpc_proto_rawDesc = nil
	file_mc_player_grpc_proto_goTypes = nil
	file_mc_player_grpc_proto_depIdxs = nil
}
