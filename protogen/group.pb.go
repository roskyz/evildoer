// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: group.proto

package protogen

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

type AccessMode int32

const (
	AccessMode_ACCESS_MODE_API AccessMode = 0
	AccessMode_ACCESS_MODE_APP AccessMode = 1
	AccessMode_ACCESS_MODE_MQ  AccessMode = 2
)

// Enum value maps for AccessMode.
var (
	AccessMode_name = map[int32]string{
		0: "ACCESS_MODE_API",
		1: "ACCESS_MODE_APP",
		2: "ACCESS_MODE_MQ",
	}
	AccessMode_value = map[string]int32{
		"ACCESS_MODE_API": 0,
		"ACCESS_MODE_APP": 1,
		"ACCESS_MODE_MQ":  2,
	}
)

func (x AccessMode) Enum() *AccessMode {
	p := new(AccessMode)
	*p = x
	return p
}

func (x AccessMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccessMode) Descriptor() protoreflect.EnumDescriptor {
	return file_group_proto_enumTypes[0].Descriptor()
}

func (AccessMode) Type() protoreflect.EnumType {
	return &file_group_proto_enumTypes[0]
}

func (x AccessMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccessMode.Descriptor instead.
func (AccessMode) EnumDescriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{0}
}

type GroupState int32

const (
	GroupState_GROUP_STATE_NORMAL   GroupState = 0
	GroupState_GROUP_STATE_DISABLED GroupState = 1
	GroupState_GROUP_STATE_DELETED  GroupState = 2
)

// Enum value maps for GroupState.
var (
	GroupState_name = map[int32]string{
		0: "GROUP_STATE_NORMAL",
		1: "GROUP_STATE_DISABLED",
		2: "GROUP_STATE_DELETED",
	}
	GroupState_value = map[string]int32{
		"GROUP_STATE_NORMAL":   0,
		"GROUP_STATE_DISABLED": 1,
		"GROUP_STATE_DELETED":  2,
	}
)

func (x GroupState) Enum() *GroupState {
	p := new(GroupState)
	*p = x
	return p
}

func (x GroupState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GroupState) Descriptor() protoreflect.EnumDescriptor {
	return file_group_proto_enumTypes[1].Descriptor()
}

func (GroupState) Type() protoreflect.EnumType {
	return &file_group_proto_enumTypes[1]
}

func (x GroupState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GroupState.Descriptor instead.
func (GroupState) EnumDescriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{1}
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          []byte       `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Key         string       `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"` // todo: mongo uniq index
	Name        string       `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string       `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	State       GroupState   `protobuf:"varint,5,opt,name=state,proto3,enum=idl.GroupState" json:"state,omitempty"`
	Creator     string       `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"` // todo: creator should not be passed
	AccessMode  []AccessMode `protobuf:"varint,7,rep,packed,name=access_mode,json=accessMode,proto3,enum=idl.AccessMode" json:"access_mode,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{0}
}

func (x *Group) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Group) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Group) GetState() GroupState {
	if x != nil {
		return x.State
	}
	return GroupState_GROUP_STATE_NORMAL
}

func (x *Group) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *Group) GetAccessMode() []AccessMode {
	if x != nil {
		return x.AccessMode
	}
	return nil
}

type ListGroupReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page         *Page   `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	SearchOption *Search `protobuf:"bytes,2,opt,name=search_option,json=searchOption,proto3,oneof" json:"search_option,omitempty"`
}

func (x *ListGroupReq) Reset() {
	*x = ListGroupReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGroupReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroupReq) ProtoMessage() {}

func (x *ListGroupReq) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroupReq.ProtoReflect.Descriptor instead.
func (*ListGroupReq) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{1}
}

func (x *ListGroupReq) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListGroupReq) GetSearchOption() *Search {
	if x != nil {
		return x.SearchOption
	}
	return nil
}

type ListGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DefaultResponse *DefaultResponse `protobuf:"bytes,1,opt,name=default_response,json=defaultResponse,proto3" json:"default_response,omitempty"`
	Count           int64            `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Groups          []*Group         `protobuf:"bytes,3,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *ListGroupResponse) Reset() {
	*x = ListGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroupResponse) ProtoMessage() {}

func (x *ListGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroupResponse.ProtoReflect.Descriptor instead.
func (*ListGroupResponse) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{2}
}

func (x *ListGroupResponse) GetDefaultResponse() *DefaultResponse {
	if x != nil {
		return x.DefaultResponse
	}
	return nil
}

func (x *ListGroupResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListGroupResponse) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

type GetGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DefaultResponse *DefaultResponse `protobuf:"bytes,1,opt,name=default_response,json=defaultResponse,proto3" json:"default_response,omitempty"`
	Group           *Group           `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *GetGroupResponse) Reset() {
	*x = GetGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupResponse) ProtoMessage() {}

func (x *GetGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupResponse.ProtoReflect.Descriptor instead.
func (*GetGroupResponse) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{3}
}

func (x *GetGroupResponse) GetDefaultResponse() *DefaultResponse {
	if x != nil {
		return x.DefaultResponse
	}
	return nil
}

func (x *GetGroupResponse) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type CreateGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DefaultResponse *DefaultResponse `protobuf:"bytes,1,opt,name=default_response,json=defaultResponse,proto3" json:"default_response,omitempty"`
	IdResp          *IDReqOrResponse `protobuf:"bytes,2,opt,name=id_resp,json=idResp,proto3" json:"id_resp,omitempty"`
}

func (x *CreateGroupResponse) Reset() {
	*x = CreateGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupResponse) ProtoMessage() {}

func (x *CreateGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateGroupResponse) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{4}
}

func (x *CreateGroupResponse) GetDefaultResponse() *DefaultResponse {
	if x != nil {
		return x.DefaultResponse
	}
	return nil
}

func (x *CreateGroupResponse) GetIdResp() *IDReqOrResponse {
	if x != nil {
		return x.IdResp
	}
	return nil
}

type ListModeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DefaultResponse *DefaultResponse `protobuf:"bytes,1,opt,name=default_response,json=defaultResponse,proto3" json:"default_response,omitempty"`
	GroupState      []GroupState     `protobuf:"varint,2,rep,packed,name=group_state,json=groupState,proto3,enum=idl.GroupState" json:"group_state,omitempty"`
}

func (x *ListModeResponse) Reset() {
	*x = ListModeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListModeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListModeResponse) ProtoMessage() {}

func (x *ListModeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListModeResponse.ProtoReflect.Descriptor instead.
func (*ListModeResponse) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{5}
}

func (x *ListModeResponse) GetDefaultResponse() *DefaultResponse {
	if x != nil {
		return x.DefaultResponse
	}
	return nil
}

func (x *ListModeResponse) GetGroupState() []GroupState {
	if x != nil {
		return x.GroupState
	}
	return nil
}

var File_group_proto protoreflect.FileDescriptor

var file_group_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x69,
	0x64, 0x6c, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde, 0x01,
	0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x30, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x69, 0x64, 0x6c,
	0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x22, 0x76,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x12, 0x1d,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x69,
	0x64, 0x6c, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x35, 0x0a,
	0x0d, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8e, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x10,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0f, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x75, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x10, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0f, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x69, 0x64,
	0x6c, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x85,
	0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x10, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x4f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06,
	0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x85, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x10, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0f, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0b,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0e, 0x32, 0x0f, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2a, 0x4a,
	0x0a, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x13, 0x0a, 0x0f,
	0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x50, 0x49, 0x10,
	0x00, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4d, 0x4f, 0x44, 0x45,
	0x5f, 0x41, 0x50, 0x50, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x4d, 0x51, 0x10, 0x02, 0x2a, 0x57, 0x0a, 0x0a, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x47, 0x52, 0x4f, 0x55,
	0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00,
	0x12, 0x18, 0x0a, 0x14, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x47, 0x52,
	0x4f, 0x55, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45,
	0x44, 0x10, 0x02, 0x32, 0xa6, 0x02, 0x0a, 0x08, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x72, 0x76,
	0x12, 0x38, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x11, 0x2e,
	0x69, 0x64, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x1a, 0x16, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x14, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x4f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x15, 0x2e, 0x69,
	0x64, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x0a, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x1a, 0x18, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0a, 0x2e, 0x69, 0x64,
	0x6c, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x14, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x44, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3b, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1c, 0x5a, 0x1a,
	0x65, 0x76, 0x69, 0x6c, 0x64, 0x6f, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65,
	0x6e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_group_proto_rawDescOnce sync.Once
	file_group_proto_rawDescData = file_group_proto_rawDesc
)

func file_group_proto_rawDescGZIP() []byte {
	file_group_proto_rawDescOnce.Do(func() {
		file_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_group_proto_rawDescData)
	})
	return file_group_proto_rawDescData
}

var file_group_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_group_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_group_proto_goTypes = []interface{}{
	(AccessMode)(0),             // 0: idl.AccessMode
	(GroupState)(0),             // 1: idl.GroupState
	(*Group)(nil),               // 2: idl.Group
	(*ListGroupReq)(nil),        // 3: idl.ListGroupReq
	(*ListGroupResponse)(nil),   // 4: idl.ListGroupResponse
	(*GetGroupResponse)(nil),    // 5: idl.GetGroupResponse
	(*CreateGroupResponse)(nil), // 6: idl.CreateGroupResponse
	(*ListModeResponse)(nil),    // 7: idl.ListModeResponse
	(*Page)(nil),                // 8: idl.Page
	(*Search)(nil),              // 9: idl.Search
	(*DefaultResponse)(nil),     // 10: idl.DefaultResponse
	(*IDReqOrResponse)(nil),     // 11: idl.IDReqOrResponse
	(*emptypb.Empty)(nil),       // 12: google.protobuf.Empty
}
var file_group_proto_depIdxs = []int32{
	1,  // 0: idl.Group.state:type_name -> idl.GroupState
	0,  // 1: idl.Group.access_mode:type_name -> idl.AccessMode
	8,  // 2: idl.ListGroupReq.page:type_name -> idl.Page
	9,  // 3: idl.ListGroupReq.search_option:type_name -> idl.Search
	10, // 4: idl.ListGroupResponse.default_response:type_name -> idl.DefaultResponse
	2,  // 5: idl.ListGroupResponse.groups:type_name -> idl.Group
	10, // 6: idl.GetGroupResponse.default_response:type_name -> idl.DefaultResponse
	2,  // 7: idl.GetGroupResponse.group:type_name -> idl.Group
	10, // 8: idl.CreateGroupResponse.default_response:type_name -> idl.DefaultResponse
	11, // 9: idl.CreateGroupResponse.id_resp:type_name -> idl.IDReqOrResponse
	10, // 10: idl.ListModeResponse.default_response:type_name -> idl.DefaultResponse
	1,  // 11: idl.ListModeResponse.group_state:type_name -> idl.GroupState
	3,  // 12: idl.GroupSrv.ListGroup:input_type -> idl.ListGroupReq
	11, // 13: idl.GroupSrv.GetGroup:input_type -> idl.IDReqOrResponse
	2,  // 14: idl.GroupSrv.CreateGroup:input_type -> idl.Group
	2,  // 15: idl.GroupSrv.UpdateGroup:input_type -> idl.Group
	12, // 16: idl.GroupSrv.ListMode:input_type -> google.protobuf.Empty
	4,  // 17: idl.GroupSrv.ListGroup:output_type -> idl.ListGroupResponse
	5,  // 18: idl.GroupSrv.GetGroup:output_type -> idl.GetGroupResponse
	6,  // 19: idl.GroupSrv.CreateGroup:output_type -> idl.CreateGroupResponse
	10, // 20: idl.GroupSrv.UpdateGroup:output_type -> idl.DefaultResponse
	7,  // 21: idl.GroupSrv.ListMode:output_type -> idl.ListModeResponse
	17, // [17:22] is the sub-list for method output_type
	12, // [12:17] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_group_proto_init() }
func file_group_proto_init() {
	if File_group_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGroupReq); i {
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
		file_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGroupResponse); i {
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
		file_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupResponse); i {
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
		file_group_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupResponse); i {
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
		file_group_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListModeResponse); i {
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
	file_group_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_group_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_group_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_group_proto_goTypes,
		DependencyIndexes: file_group_proto_depIdxs,
		EnumInfos:         file_group_proto_enumTypes,
		MessageInfos:      file_group_proto_msgTypes,
	}.Build()
	File_group_proto = out.File
	file_group_proto_rawDesc = nil
	file_group_proto_goTypes = nil
	file_group_proto_depIdxs = nil
}
