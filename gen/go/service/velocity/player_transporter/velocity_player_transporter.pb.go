// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: velocity/velocity_player_transporter.proto

package player_transporter

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	server_discovery "service/server_discovery"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TransportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server    *server_discovery.ConnectableServer `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	PlayerIds []string                            `protobuf:"bytes,2,rep,name=player_ids,json=playerIds,proto3" json:"player_ids,omitempty"`
}

func (x *TransportRequest) Reset() {
	*x = TransportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_velocity_velocity_player_transporter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransportRequest) ProtoMessage() {}

func (x *TransportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_velocity_velocity_player_transporter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransportRequest.ProtoReflect.Descriptor instead.
func (*TransportRequest) Descriptor() ([]byte, []int) {
	return file_velocity_velocity_player_transporter_proto_rawDescGZIP(), []int{0}
}

func (x *TransportRequest) GetServer() *server_discovery.ConnectableServer {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *TransportRequest) GetPlayerIds() []string {
	if x != nil {
		return x.PlayerIds
	}
	return nil
}

var File_velocity_velocity_player_transporter_proto protoreflect.FileDescriptor

var file_velocity_velocity_player_transporter_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x65, 0x6c, 0x6f, 0x63,
	0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x33, 0x74, 0x6f,
	0x77, 0x65, 0x72, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x63, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x53, 0x0a, 0x06, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x74, 0x6f,
	0x77, 0x65, 0x72, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x63, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x73, 0x32,
	0x8a, 0x01, 0x0a, 0x19, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x12, 0x6d, 0x0a,
	0x0c, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x45, 0x2e,
	0x74, 0x6f, 0x77, 0x65, 0x72, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x63, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x67, 0x0a, 0x20,
	0x64, 0x65, 0x76, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79,
	0x42, 0x1e, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x5a, 0x23, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69,
	0x74, 0x79, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_velocity_velocity_player_transporter_proto_rawDescOnce sync.Once
	file_velocity_velocity_player_transporter_proto_rawDescData = file_velocity_velocity_player_transporter_proto_rawDesc
)

func file_velocity_velocity_player_transporter_proto_rawDescGZIP() []byte {
	file_velocity_velocity_player_transporter_proto_rawDescOnce.Do(func() {
		file_velocity_velocity_player_transporter_proto_rawDescData = protoimpl.X.CompressGZIP(file_velocity_velocity_player_transporter_proto_rawDescData)
	})
	return file_velocity_velocity_player_transporter_proto_rawDescData
}

var file_velocity_velocity_player_transporter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_velocity_velocity_player_transporter_proto_goTypes = []interface{}{
	(*TransportRequest)(nil),                   // 0: towerdefence.cc.service.velocity.player_transporter.TransportRequest
	(*server_discovery.ConnectableServer)(nil), // 1: towerdefence.cc.service.server_discovery.ConnectableServer
	(*emptypb.Empty)(nil),                      // 2: google.protobuf.Empty
}
var file_velocity_velocity_player_transporter_proto_depIdxs = []int32{
	1, // 0: towerdefence.cc.service.velocity.player_transporter.TransportRequest.server:type_name -> towerdefence.cc.service.server_discovery.ConnectableServer
	0, // 1: towerdefence.cc.service.velocity.player_transporter.VelocityPlayerTransporter.SendToServer:input_type -> towerdefence.cc.service.velocity.player_transporter.TransportRequest
	2, // 2: towerdefence.cc.service.velocity.player_transporter.VelocityPlayerTransporter.SendToServer:output_type -> google.protobuf.Empty
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_velocity_velocity_player_transporter_proto_init() }
func file_velocity_velocity_player_transporter_proto_init() {
	if File_velocity_velocity_player_transporter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_velocity_velocity_player_transporter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransportRequest); i {
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
			RawDescriptor: file_velocity_velocity_player_transporter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_velocity_velocity_player_transporter_proto_goTypes,
		DependencyIndexes: file_velocity_velocity_player_transporter_proto_depIdxs,
		MessageInfos:      file_velocity_velocity_player_transporter_proto_msgTypes,
	}.Build()
	File_velocity_velocity_player_transporter_proto = out.File
	file_velocity_velocity_player_transporter_proto_rawDesc = nil
	file_velocity_velocity_player_transporter_proto_goTypes = nil
	file_velocity_velocity_player_transporter_proto_depIdxs = nil
}
