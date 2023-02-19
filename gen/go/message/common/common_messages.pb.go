// common messages

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: common_messages.proto

package common

import (
	common "github.com/emortalmc/proto-specs/gen/go/model/common"
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

// SwitchPlayersServerMessage tells a proxy to switch player(s) to a different server
type SwitchPlayersServerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server    *common.ConnectableServer `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	PlayerIds []string                  `protobuf:"bytes,2,rep,name=player_ids,json=playerIds,proto3" json:"player_ids,omitempty"`
}

func (x *SwitchPlayersServerMessage) Reset() {
	*x = SwitchPlayersServerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwitchPlayersServerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwitchPlayersServerMessage) ProtoMessage() {}

func (x *SwitchPlayersServerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_common_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwitchPlayersServerMessage.ProtoReflect.Descriptor instead.
func (*SwitchPlayersServerMessage) Descriptor() ([]byte, []int) {
	return file_common_messages_proto_rawDescGZIP(), []int{0}
}

func (x *SwitchPlayersServerMessage) GetServer() *common.ConnectableServer {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *SwitchPlayersServerMessage) GetPlayerIds() []string {
	if x != nil {
		return x.PlayerIds
	}
	return nil
}

type PlayerConnectMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId       string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	PlayerUsername string `protobuf:"bytes,2,opt,name=player_username,json=playerUsername,proto3" json:"player_username,omitempty"`
	ServerId       string `protobuf:"bytes,3,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
}

func (x *PlayerConnectMessage) Reset() {
	*x = PlayerConnectMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerConnectMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerConnectMessage) ProtoMessage() {}

func (x *PlayerConnectMessage) ProtoReflect() protoreflect.Message {
	mi := &file_common_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerConnectMessage.ProtoReflect.Descriptor instead.
func (*PlayerConnectMessage) Descriptor() ([]byte, []int) {
	return file_common_messages_proto_rawDescGZIP(), []int{1}
}

func (x *PlayerConnectMessage) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerConnectMessage) GetPlayerUsername() string {
	if x != nil {
		return x.PlayerUsername
	}
	return ""
}

func (x *PlayerConnectMessage) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

// PlayerSwitchServerMessage is sent when a player is confirmed to have switched servers.
// it is therefore only sent by the proxy the player is connected to.
type PlayerSwitchServerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
}

func (x *PlayerSwitchServerMessage) Reset() {
	*x = PlayerSwitchServerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerSwitchServerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerSwitchServerMessage) ProtoMessage() {}

func (x *PlayerSwitchServerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_common_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerSwitchServerMessage.ProtoReflect.Descriptor instead.
func (*PlayerSwitchServerMessage) Descriptor() ([]byte, []int) {
	return file_common_messages_proto_rawDescGZIP(), []int{2}
}

func (x *PlayerSwitchServerMessage) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerSwitchServerMessage) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

type PlayerDisconnectMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
}

func (x *PlayerDisconnectMessage) Reset() {
	*x = PlayerDisconnectMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerDisconnectMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerDisconnectMessage) ProtoMessage() {}

func (x *PlayerDisconnectMessage) ProtoReflect() protoreflect.Message {
	mi := &file_common_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerDisconnectMessage.ProtoReflect.Descriptor instead.
func (*PlayerDisconnectMessage) Descriptor() ([]byte, []int) {
	return file_common_messages_proto_rawDescGZIP(), []int{3}
}

func (x *PlayerDisconnectMessage) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

var File_common_messages_proto protoreflect.FileDescriptor

var file_common_messages_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a,
	0x1a, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6d,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x64, 0x73, 0x22, 0x79, 0x0a, 0x14, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x55, 0x0a, 0x19, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x36, 0x0a, 0x17, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x42, 0x5a,
	0x0a, 0x1e, 0x64, 0x65, 0x76, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x6d, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73,
	0x70, 0x65, 0x63, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_common_messages_proto_rawDescOnce sync.Once
	file_common_messages_proto_rawDescData = file_common_messages_proto_rawDesc
)

func file_common_messages_proto_rawDescGZIP() []byte {
	file_common_messages_proto_rawDescOnce.Do(func() {
		file_common_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_messages_proto_rawDescData)
	})
	return file_common_messages_proto_rawDescData
}

var file_common_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_common_messages_proto_goTypes = []interface{}{
	(*SwitchPlayersServerMessage)(nil), // 0: emortal.message.SwitchPlayersServerMessage
	(*PlayerConnectMessage)(nil),       // 1: emortal.message.PlayerConnectMessage
	(*PlayerSwitchServerMessage)(nil),  // 2: emortal.message.PlayerSwitchServerMessage
	(*PlayerDisconnectMessage)(nil),    // 3: emortal.message.PlayerDisconnectMessage
	(*common.ConnectableServer)(nil),   // 4: emortal.model.ConnectableServer
}
var file_common_messages_proto_depIdxs = []int32{
	4, // 0: emortal.message.SwitchPlayersServerMessage.server:type_name -> emortal.model.ConnectableServer
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_messages_proto_init() }
func file_common_messages_proto_init() {
	if File_common_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwitchPlayersServerMessage); i {
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
		file_common_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerConnectMessage); i {
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
		file_common_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerSwitchServerMessage); i {
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
		file_common_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerDisconnectMessage); i {
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
			RawDescriptor: file_common_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_messages_proto_goTypes,
		DependencyIndexes: file_common_messages_proto_depIdxs,
		MessageInfos:      file_common_messages_proto_msgTypes,
	}.Build()
	File_common_messages_proto = out.File
	file_common_messages_proto_rawDesc = nil
	file_common_messages_proto_goTypes = nil
	file_common_messages_proto_depIdxs = nil
}
