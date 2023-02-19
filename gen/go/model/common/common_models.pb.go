// common models

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: common_models.proto

package common

import (
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

type ServerType int32

const (
	ServerType_PROXY           ServerType = 0
	ServerType_LOBBY           ServerType = 1
	ServerType_MARATHON        ServerType = 2
	ServerType_BLOCK_SUMO      ServerType = 3
	ServerType_PARKOURTAG      ServerType = 4
	ServerType_LAZERTAG        ServerType = 5
	ServerType_HOLEY_MOLEY     ServerType = 6
	ServerType_MARATHON_RACING ServerType = 7
	ServerType_BATTLE          ServerType = 8
	ServerType_MINESWEEPER     ServerType = 9
)

// Enum value maps for ServerType.
var (
	ServerType_name = map[int32]string{
		0: "PROXY",
		1: "LOBBY",
		2: "MARATHON",
		3: "BLOCK_SUMO",
		4: "PARKOURTAG",
		5: "LAZERTAG",
		6: "HOLEY_MOLEY",
		7: "MARATHON_RACING",
		8: "BATTLE",
		9: "MINESWEEPER",
	}
	ServerType_value = map[string]int32{
		"PROXY":           0,
		"LOBBY":           1,
		"MARATHON":        2,
		"BLOCK_SUMO":      3,
		"PARKOURTAG":      4,
		"LAZERTAG":        5,
		"HOLEY_MOLEY":     6,
		"MARATHON_RACING": 7,
		"BATTLE":          8,
		"MINESWEEPER":     9,
	}
)

func (x ServerType) Enum() *ServerType {
	p := new(ServerType)
	*p = x
	return p
}

func (x ServerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServerType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_models_proto_enumTypes[0].Descriptor()
}

func (ServerType) Type() protoreflect.EnumType {
	return &file_common_models_proto_enumTypes[0]
}

func (x ServerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServerType.Descriptor instead.
func (ServerType) EnumDescriptor() ([]byte, []int) {
	return file_common_models_proto_rawDescGZIP(), []int{0}
}

type ConnectableServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kubernetes pod name, e.g lobby-wgxtd-llffk
	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Port    uint32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *ConnectableServer) Reset() {
	*x = ConnectableServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectableServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectableServer) ProtoMessage() {}

func (x *ConnectableServer) ProtoReflect() protoreflect.Message {
	mi := &file_common_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectableServer.ProtoReflect.Descriptor instead.
func (*ConnectableServer) Descriptor() ([]byte, []int) {
	return file_common_models_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectableServer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ConnectableServer) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ConnectableServer) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

// Pageable utilities
// NOTE: Page values start at 0, not 1
type Pageable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page uint64  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size *uint64 `protobuf:"varint,2,opt,name=size,proto3,oneof" json:"size,omitempty"`
}

func (x *Pageable) Reset() {
	*x = Pageable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pageable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pageable) ProtoMessage() {}

func (x *Pageable) ProtoReflect() protoreflect.Message {
	mi := &file_common_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pageable.ProtoReflect.Descriptor instead.
func (*Pageable) Descriptor() ([]byte, []int) {
	return file_common_models_proto_rawDescGZIP(), []int{1}
}

func (x *Pageable) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Pageable) GetSize() uint64 {
	if x != nil && x.Size != nil {
		return *x.Size
	}
	return 0
}

type Sortable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field     string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Ascending bool   `protobuf:"varint,2,opt,name=ascending,proto3" json:"ascending,omitempty"`
}

func (x *Sortable) Reset() {
	*x = Sortable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sortable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sortable) ProtoMessage() {}

func (x *Sortable) ProtoReflect() protoreflect.Message {
	mi := &file_common_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sortable.ProtoReflect.Descriptor instead.
func (*Sortable) Descriptor() ([]byte, []int) {
	return file_common_models_proto_rawDescGZIP(), []int{2}
}

func (x *Sortable) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *Sortable) GetAscending() bool {
	if x != nil {
		return x.Ascending
	}
	return false
}

type PageData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page          uint64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size          uint64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	TotalElements uint64 `protobuf:"varint,3,opt,name=total_elements,json=totalElements,proto3" json:"total_elements,omitempty"`
	TotalPages    uint64 `protobuf:"varint,4,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
}

func (x *PageData) Reset() {
	*x = PageData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageData) ProtoMessage() {}

func (x *PageData) ProtoReflect() protoreflect.Message {
	mi := &file_common_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageData.ProtoReflect.Descriptor instead.
func (*PageData) Descriptor() ([]byte, []int) {
	return file_common_models_proto_rawDescGZIP(), []int{3}
}

func (x *PageData) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PageData) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *PageData) GetTotalElements() uint64 {
	if x != nil {
		return x.TotalElements
	}
	return 0
}

func (x *PageData) GetTotalPages() uint64 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

var File_common_models_proto protoreflect.FileDescriptor

var file_common_models_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x22, 0x51, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x40, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x3e, 0x0a, 0x08, 0x53, 0x6f, 0x72,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x61, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x7a, 0x0a, 0x08, 0x50, 0x61, 0x67,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x50, 0x61, 0x67, 0x65, 0x73, 0x2a, 0xa1, 0x01, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x52, 0x4f, 0x58, 0x59, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x42, 0x42, 0x59, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x41,
	0x52, 0x41, 0x54, 0x48, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x42, 0x4c, 0x4f, 0x43,
	0x4b, 0x5f, 0x53, 0x55, 0x4d, 0x4f, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x41, 0x52, 0x4b,
	0x4f, 0x55, 0x52, 0x54, 0x41, 0x47, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x41, 0x5a, 0x45,
	0x52, 0x54, 0x41, 0x47, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x48, 0x4f, 0x4c, 0x45, 0x59, 0x5f,
	0x4d, 0x4f, 0x4c, 0x45, 0x59, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x41, 0x52, 0x41, 0x54,
	0x48, 0x4f, 0x4e, 0x5f, 0x52, 0x41, 0x43, 0x49, 0x4e, 0x47, 0x10, 0x07, 0x12, 0x0a, 0x0a, 0x06,
	0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x10, 0x08, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x49, 0x4e, 0x45,
	0x53, 0x57, 0x45, 0x45, 0x50, 0x45, 0x52, 0x10, 0x09, 0x42, 0x56, 0x0a, 0x1c, 0x64, 0x65, 0x76,
	0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x6d,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_models_proto_rawDescOnce sync.Once
	file_common_models_proto_rawDescData = file_common_models_proto_rawDesc
)

func file_common_models_proto_rawDescGZIP() []byte {
	file_common_models_proto_rawDescOnce.Do(func() {
		file_common_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_models_proto_rawDescData)
	})
	return file_common_models_proto_rawDescData
}

var file_common_models_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_models_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_common_models_proto_goTypes = []interface{}{
	(ServerType)(0),           // 0: emortal.model.ServerType
	(*ConnectableServer)(nil), // 1: emortal.model.ConnectableServer
	(*Pageable)(nil),          // 2: emortal.model.Pageable
	(*Sortable)(nil),          // 3: emortal.model.Sortable
	(*PageData)(nil),          // 4: emortal.model.PageData
}
var file_common_models_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_models_proto_init() }
func file_common_models_proto_init() {
	if File_common_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectableServer); i {
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
		file_common_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pageable); i {
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
		file_common_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sortable); i {
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
		file_common_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageData); i {
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
	file_common_models_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_models_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_models_proto_goTypes,
		DependencyIndexes: file_common_models_proto_depIdxs,
		EnumInfos:         file_common_models_proto_enumTypes,
		MessageInfos:      file_common_models_proto_msgTypes,
	}.Build()
	File_common_models_proto = out.File
	file_common_models_proto_rawDesc = nil
	file_common_models_proto_goTypes = nil
	file_common_models_proto_depIdxs = nil
}
