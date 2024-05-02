// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: party/messages.proto

package party

import (
	party "github.com/emortalmc/proto-specs/gen/go/model/party"
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

type PartyCreatedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Party *party.Party `protobuf:"bytes,1,opt,name=party,proto3" json:"party,omitempty"`
}

func (x *PartyCreatedMessage) Reset() {
	*x = PartyCreatedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyCreatedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyCreatedMessage) ProtoMessage() {}

func (x *PartyCreatedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_party_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyCreatedMessage.ProtoReflect.Descriptor instead.
func (*PartyCreatedMessage) Descriptor() ([]byte, []int) {
	return file_party_messages_proto_rawDescGZIP(), []int{0}
}

func (x *PartyCreatedMessage) GetParty() *party.Party {
	if x != nil {
		return x.Party
	}
	return nil
}

// PartyDeletedMessage is sent when a party is deleted
type PartyDeletedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Party *party.Party `protobuf:"bytes,1,opt,name=party,proto3" json:"party,omitempty"`
}

func (x *PartyDeletedMessage) Reset() {
	*x = PartyDeletedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyDeletedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyDeletedMessage) ProtoMessage() {}

func (x *PartyDeletedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_party_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyDeletedMessage.ProtoReflect.Descriptor instead.
func (*PartyDeletedMessage) Descriptor() ([]byte, []int) {
	return file_party_messages_proto_rawDescGZIP(), []int{1}
}

func (x *PartyDeletedMessage) GetParty() *party.Party {
	if x != nil {
		return x.Party
	}
	return nil
}

// PartyEmptiedMessage is sent when a party is emptied
// Note, this is different to a party being deleted. When it is emptied, the leader remains as a member but all others
// are removed and notified of a disband. The party is also set to closed.
type PartyEmptiedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Party *party.Party `protobuf:"bytes,1,opt,name=party,proto3" json:"party,omitempty"`
}

func (x *PartyEmptiedMessage) Reset() {
	*x = PartyEmptiedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyEmptiedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyEmptiedMessage) ProtoMessage() {}

func (x *PartyEmptiedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_party_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyEmptiedMessage.ProtoReflect.Descriptor instead.
func (*PartyEmptiedMessage) Descriptor() ([]byte, []int) {
	return file_party_messages_proto_rawDescGZIP(), []int{2}
}

func (x *PartyEmptiedMessage) GetParty() *party.Party {
	if x != nil {
		return x.Party
	}
	return nil
}

type PartyInviteCreatedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Invite *party.PartyInvite `protobuf:"bytes,1,opt,name=invite,proto3" json:"invite,omitempty"`
}

func (x *PartyInviteCreatedMessage) Reset() {
	*x = PartyInviteCreatedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyInviteCreatedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyInviteCreatedMessage) ProtoMessage() {}

func (x *PartyInviteCreatedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_party_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyInviteCreatedMessage.ProtoReflect.Descriptor instead.
func (*PartyInviteCreatedMessage) Descriptor() ([]byte, []int) {
	return file_party_messages_proto_rawDescGZIP(), []int{3}
}

func (x *PartyInviteCreatedMessage) GetInvite() *party.PartyInvite {
	if x != nil {
		return x.Invite
	}
	return nil
}

type PartyPlayerJoinedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// party_id of type mongo ObjectId
	PartyId string             `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	Member  *party.PartyMember `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty"`
}

func (x *PartyPlayerJoinedMessage) Reset() {
	*x = PartyPlayerJoinedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyPlayerJoinedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyPlayerJoinedMessage) ProtoMessage() {}

func (x *PartyPlayerJoinedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_party_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyPlayerJoinedMessage.ProtoReflect.Descriptor instead.
func (*PartyPlayerJoinedMessage) Descriptor() ([]byte, []int) {
	return file_party_messages_proto_rawDescGZIP(), []int{4}
}

func (x *PartyPlayerJoinedMessage) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *PartyPlayerJoinedMessage) GetMember() *party.PartyMember {
	if x != nil {
		return x.Member
	}
	return nil
}

type PartyPlayerLeftMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// party_id of type mongo ObjectId
	PartyId string             `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	Member  *party.PartyMember `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty"`
	// kicked_by only present if the player was kicked, not if they left
	KickedBy *party.PartyMember `protobuf:"bytes,3,opt,name=kicked_by,json=kickedBy,proto3,oneof" json:"kicked_by,omitempty"`
}

func (x *PartyPlayerLeftMessage) Reset() {
	*x = PartyPlayerLeftMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyPlayerLeftMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyPlayerLeftMessage) ProtoMessage() {}

func (x *PartyPlayerLeftMessage) ProtoReflect() protoreflect.Message {
	mi := &file_party_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyPlayerLeftMessage.ProtoReflect.Descriptor instead.
func (*PartyPlayerLeftMessage) Descriptor() ([]byte, []int) {
	return file_party_messages_proto_rawDescGZIP(), []int{5}
}

func (x *PartyPlayerLeftMessage) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *PartyPlayerLeftMessage) GetMember() *party.PartyMember {
	if x != nil {
		return x.Member
	}
	return nil
}

func (x *PartyPlayerLeftMessage) GetKickedBy() *party.PartyMember {
	if x != nil {
		return x.KickedBy
	}
	return nil
}

type PartyOpenChangedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// party_id of type mongo ObjectId
	PartyId string `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	Open    bool   `protobuf:"varint,2,opt,name=open,proto3" json:"open,omitempty"`
}

func (x *PartyOpenChangedMessage) Reset() {
	*x = PartyOpenChangedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyOpenChangedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyOpenChangedMessage) ProtoMessage() {}

func (x *PartyOpenChangedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_party_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyOpenChangedMessage.ProtoReflect.Descriptor instead.
func (*PartyOpenChangedMessage) Descriptor() ([]byte, []int) {
	return file_party_messages_proto_rawDescGZIP(), []int{6}
}

func (x *PartyOpenChangedMessage) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *PartyOpenChangedMessage) GetOpen() bool {
	if x != nil {
		return x.Open
	}
	return false
}

type PartyClosedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// party_id of type mongo ObjectId
	PartyId string `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
}

func (x *PartyClosedMessage) Reset() {
	*x = PartyClosedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyClosedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyClosedMessage) ProtoMessage() {}

func (x *PartyClosedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_party_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyClosedMessage.ProtoReflect.Descriptor instead.
func (*PartyClosedMessage) Descriptor() ([]byte, []int) {
	return file_party_messages_proto_rawDescGZIP(), []int{7}
}

func (x *PartyClosedMessage) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

type PartyLeaderChangedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// party_id of type mongo ObjectId
	PartyId   string             `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	NewLeader *party.PartyMember `protobuf:"bytes,2,opt,name=new_leader,json=newLeader,proto3" json:"new_leader,omitempty"`
}

func (x *PartyLeaderChangedMessage) Reset() {
	*x = PartyLeaderChangedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_messages_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyLeaderChangedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyLeaderChangedMessage) ProtoMessage() {}

func (x *PartyLeaderChangedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_party_messages_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyLeaderChangedMessage.ProtoReflect.Descriptor instead.
func (*PartyLeaderChangedMessage) Descriptor() ([]byte, []int) {
	return file_party_messages_proto_rawDescGZIP(), []int{8}
}

func (x *PartyLeaderChangedMessage) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *PartyLeaderChangedMessage) GetNewLeader() *party.PartyMember {
	if x != nil {
		return x.NewLeader
	}
	return nil
}

type PartySettingsChangedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// player_id of type UUID
	PlayerId string               `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Settings *party.PartySettings `protobuf:"bytes,2,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *PartySettingsChangedMessage) Reset() {
	*x = PartySettingsChangedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_messages_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartySettingsChangedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartySettingsChangedMessage) ProtoMessage() {}

func (x *PartySettingsChangedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_party_messages_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartySettingsChangedMessage.ProtoReflect.Descriptor instead.
func (*PartySettingsChangedMessage) Descriptor() ([]byte, []int) {
	return file_party_messages_proto_rawDescGZIP(), []int{9}
}

func (x *PartySettingsChangedMessage) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PartySettingsChangedMessage) GetSettings() *party.PartySettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

type PartyBroadcastMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Party *party.Party `protobuf:"bytes,1,opt,name=party,proto3" json:"party,omitempty"`
}

func (x *PartyBroadcastMessage) Reset() {
	*x = PartyBroadcastMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_messages_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyBroadcastMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyBroadcastMessage) ProtoMessage() {}

func (x *PartyBroadcastMessage) ProtoReflect() protoreflect.Message {
	mi := &file_party_messages_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyBroadcastMessage.ProtoReflect.Descriptor instead.
func (*PartyBroadcastMessage) Descriptor() ([]byte, []int) {
	return file_party_messages_proto_rawDescGZIP(), []int{10}
}

func (x *PartyBroadcastMessage) GetParty() *party.Party {
	if x != nil {
		return x.Party
	}
	return nil
}

// EventDisplayMessage is sent when an event has reached its display_time timestamp.
type EventDisplayMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *party.EventData `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *EventDisplayMessage) Reset() {
	*x = EventDisplayMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_messages_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventDisplayMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventDisplayMessage) ProtoMessage() {}

func (x *EventDisplayMessage) ProtoReflect() protoreflect.Message {
	mi := &file_party_messages_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventDisplayMessage.ProtoReflect.Descriptor instead.
func (*EventDisplayMessage) Descriptor() ([]byte, []int) {
	return file_party_messages_proto_rawDescGZIP(), []int{11}
}

func (x *EventDisplayMessage) GetEvent() *party.EventData {
	if x != nil {
		return x.Event
	}
	return nil
}

type EventStartMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *party.EventData `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *EventStartMessage) Reset() {
	*x = EventStartMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_messages_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventStartMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventStartMessage) ProtoMessage() {}

func (x *EventStartMessage) ProtoReflect() protoreflect.Message {
	mi := &file_party_messages_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventStartMessage.ProtoReflect.Descriptor instead.
func (*EventStartMessage) Descriptor() ([]byte, []int) {
	return file_party_messages_proto_rawDescGZIP(), []int{12}
}

func (x *EventStartMessage) GetEvent() *party.EventData {
	if x != nil {
		return x.Event
	}
	return nil
}

// EventDeleteMessage this may be an event ending or being deleted
// this should trigger the event stopping and also to opposite of the EventDisplayMessage
// to determine if the event was 'live', check if party_id is empty
type EventDeleteMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *party.EventData `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *EventDeleteMessage) Reset() {
	*x = EventDeleteMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_messages_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventDeleteMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventDeleteMessage) ProtoMessage() {}

func (x *EventDeleteMessage) ProtoReflect() protoreflect.Message {
	mi := &file_party_messages_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventDeleteMessage.ProtoReflect.Descriptor instead.
func (*EventDeleteMessage) Descriptor() ([]byte, []int) {
	return file_party_messages_proto_rawDescGZIP(), []int{13}
}

func (x *EventDeleteMessage) GetEvent() *party.EventData {
	if x != nil {
		return x.Event
	}
	return nil
}

var File_party_messages_proto protoreflect.FileDescriptor

var file_party_messages_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x1a, 0x12, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x47, 0x0a, 0x13, 0x50, 0x61, 0x72, 0x74, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x61,
	0x72, 0x74, 0x79, 0x52, 0x05, 0x70, 0x61, 0x72, 0x74, 0x79, 0x22, 0x47, 0x0a, 0x13, 0x50, 0x61,
	0x72, 0x74, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x30, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x05, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x22, 0x47, 0x0a, 0x13, 0x50, 0x61, 0x72, 0x74, 0x79, 0x45, 0x6d, 0x70, 0x74,
	0x69, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6d, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x05, 0x70, 0x61, 0x72, 0x74, 0x79, 0x22, 0x55, 0x0a, 0x19,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6d, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x06, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x22, 0x6f, 0x0a, 0x18, 0x50, 0x61, 0x72, 0x74, 0x79, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x4a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x06, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6d, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x06, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0xbf, 0x01, 0x0a, 0x16, 0x50, 0x61, 0x72, 0x74, 0x79, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x66, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x06, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6d, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x06, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x09, 0x6b, 0x69, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x61,
	0x72, 0x74, 0x79, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x48, 0x00, 0x52, 0x08, 0x6b, 0x69, 0x63,
	0x6b, 0x65, 0x64, 0x42, 0x79, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6b, 0x69, 0x63,
	0x6b, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x22, 0x48, 0x0a, 0x17, 0x50, 0x61, 0x72, 0x74, 0x79, 0x4f,
	0x70, 0x65, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6f, 0x70, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6f, 0x70, 0x65, 0x6e,
	0x22, 0x2f, 0x0a, 0x12, 0x50, 0x61, 0x72, 0x74, 0x79, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49,
	0x64, 0x22, 0x77, 0x0a, 0x19, 0x50, 0x61, 0x72, 0x74, 0x79, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x0a, 0x6e, 0x65, 0x77,
	0x5f, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x09, 0x6e, 0x65, 0x77, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x7a, 0x0a, 0x1b, 0x50, 0x61,
	0x72, 0x74, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x50,
	0x61, 0x72, 0x74, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x49, 0x0a, 0x15, 0x50, 0x61, 0x72, 0x74, 0x79, 0x42,
	0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x30, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x05, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x22, 0x4b, 0x0a, 0x13, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x49,
	0x0a, 0x11, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x4a, 0x0a, 0x12, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x34, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x58, 0x0a, 0x1d, 0x64, 0x65, 0x76, 0x2e, 0x65, 0x6d, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x6d, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_party_messages_proto_rawDescOnce sync.Once
	file_party_messages_proto_rawDescData = file_party_messages_proto_rawDesc
)

func file_party_messages_proto_rawDescGZIP() []byte {
	file_party_messages_proto_rawDescOnce.Do(func() {
		file_party_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_party_messages_proto_rawDescData)
	})
	return file_party_messages_proto_rawDescData
}

var file_party_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_party_messages_proto_goTypes = []interface{}{
	(*PartyCreatedMessage)(nil),         // 0: emortal.message.party.PartyCreatedMessage
	(*PartyDeletedMessage)(nil),         // 1: emortal.message.party.PartyDeletedMessage
	(*PartyEmptiedMessage)(nil),         // 2: emortal.message.party.PartyEmptiedMessage
	(*PartyInviteCreatedMessage)(nil),   // 3: emortal.message.party.PartyInviteCreatedMessage
	(*PartyPlayerJoinedMessage)(nil),    // 4: emortal.message.party.PartyPlayerJoinedMessage
	(*PartyPlayerLeftMessage)(nil),      // 5: emortal.message.party.PartyPlayerLeftMessage
	(*PartyOpenChangedMessage)(nil),     // 6: emortal.message.party.PartyOpenChangedMessage
	(*PartyClosedMessage)(nil),          // 7: emortal.message.party.PartyClosedMessage
	(*PartyLeaderChangedMessage)(nil),   // 8: emortal.message.party.PartyLeaderChangedMessage
	(*PartySettingsChangedMessage)(nil), // 9: emortal.message.party.PartySettingsChangedMessage
	(*PartyBroadcastMessage)(nil),       // 10: emortal.message.party.PartyBroadcastMessage
	(*EventDisplayMessage)(nil),         // 11: emortal.message.party.EventDisplayMessage
	(*EventStartMessage)(nil),           // 12: emortal.message.party.EventStartMessage
	(*EventDeleteMessage)(nil),          // 13: emortal.message.party.EventDeleteMessage
	(*party.Party)(nil),                 // 14: emortal.model.party.Party
	(*party.PartyInvite)(nil),           // 15: emortal.model.party.PartyInvite
	(*party.PartyMember)(nil),           // 16: emortal.model.party.PartyMember
	(*party.PartySettings)(nil),         // 17: emortal.model.party.PartySettings
	(*party.EventData)(nil),             // 18: emortal.model.party.EventData
}
var file_party_messages_proto_depIdxs = []int32{
	14, // 0: emortal.message.party.PartyCreatedMessage.party:type_name -> emortal.model.party.Party
	14, // 1: emortal.message.party.PartyDeletedMessage.party:type_name -> emortal.model.party.Party
	14, // 2: emortal.message.party.PartyEmptiedMessage.party:type_name -> emortal.model.party.Party
	15, // 3: emortal.message.party.PartyInviteCreatedMessage.invite:type_name -> emortal.model.party.PartyInvite
	16, // 4: emortal.message.party.PartyPlayerJoinedMessage.member:type_name -> emortal.model.party.PartyMember
	16, // 5: emortal.message.party.PartyPlayerLeftMessage.member:type_name -> emortal.model.party.PartyMember
	16, // 6: emortal.message.party.PartyPlayerLeftMessage.kicked_by:type_name -> emortal.model.party.PartyMember
	16, // 7: emortal.message.party.PartyLeaderChangedMessage.new_leader:type_name -> emortal.model.party.PartyMember
	17, // 8: emortal.message.party.PartySettingsChangedMessage.settings:type_name -> emortal.model.party.PartySettings
	14, // 9: emortal.message.party.PartyBroadcastMessage.party:type_name -> emortal.model.party.Party
	18, // 10: emortal.message.party.EventDisplayMessage.event:type_name -> emortal.model.party.EventData
	18, // 11: emortal.message.party.EventStartMessage.event:type_name -> emortal.model.party.EventData
	18, // 12: emortal.message.party.EventDeleteMessage.event:type_name -> emortal.model.party.EventData
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_party_messages_proto_init() }
func file_party_messages_proto_init() {
	if File_party_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_party_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyCreatedMessage); i {
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
		file_party_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyDeletedMessage); i {
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
		file_party_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyEmptiedMessage); i {
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
		file_party_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyInviteCreatedMessage); i {
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
		file_party_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyPlayerJoinedMessage); i {
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
		file_party_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyPlayerLeftMessage); i {
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
		file_party_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyOpenChangedMessage); i {
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
		file_party_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyClosedMessage); i {
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
		file_party_messages_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyLeaderChangedMessage); i {
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
		file_party_messages_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartySettingsChangedMessage); i {
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
		file_party_messages_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyBroadcastMessage); i {
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
		file_party_messages_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventDisplayMessage); i {
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
		file_party_messages_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventStartMessage); i {
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
		file_party_messages_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventDeleteMessage); i {
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
	file_party_messages_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_party_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_party_messages_proto_goTypes,
		DependencyIndexes: file_party_messages_proto_depIdxs,
		MessageInfos:      file_party_messages_proto_msgTypes,
	}.Build()
	File_party_messages_proto = out.File
	file_party_messages_proto_rawDesc = nil
	file_party_messages_proto_goTypes = nil
	file_party_messages_proto_depIdxs = nil
}
