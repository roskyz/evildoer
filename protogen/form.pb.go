// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: form.proto

package protogen

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

type FormAndGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DefaultResponse *DefaultResponse `protobuf:"bytes,1,opt,name=default_response,json=defaultResponse,proto3" json:"default_response,omitempty"`
	Group           *Group           `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	Form            *Form            `protobuf:"bytes,3,opt,name=form,proto3" json:"form,omitempty"`
}

func (x *FormAndGroupResponse) Reset() {
	*x = FormAndGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_form_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FormAndGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FormAndGroupResponse) ProtoMessage() {}

func (x *FormAndGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_form_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FormAndGroupResponse.ProtoReflect.Descriptor instead.
func (*FormAndGroupResponse) Descriptor() ([]byte, []int) {
	return file_form_proto_rawDescGZIP(), []int{0}
}

func (x *FormAndGroupResponse) GetDefaultResponse() *DefaultResponse {
	if x != nil {
		return x.DefaultResponse
	}
	return nil
}

func (x *FormAndGroupResponse) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *FormAndGroupResponse) GetForm() *Form {
	if x != nil {
		return x.Form
	}
	return nil
}

type UpdateFormReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   *IDReqOrResponse `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Form *Form            `protobuf:"bytes,2,opt,name=form,proto3" json:"form,omitempty"`
}

func (x *UpdateFormReq) Reset() {
	*x = UpdateFormReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_form_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFormReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFormReq) ProtoMessage() {}

func (x *UpdateFormReq) ProtoReflect() protoreflect.Message {
	mi := &file_form_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFormReq.ProtoReflect.Descriptor instead.
func (*UpdateFormReq) Descriptor() ([]byte, []int) {
	return file_form_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateFormReq) GetId() *IDReqOrResponse {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *UpdateFormReq) GetForm() *Form {
	if x != nil {
		return x.Form
	}
	return nil
}

type ChangeStateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    *IDReqOrResponse `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	State string           `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *ChangeStateReq) Reset() {
	*x = ChangeStateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_form_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeStateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeStateReq) ProtoMessage() {}

func (x *ChangeStateReq) ProtoReflect() protoreflect.Message {
	mi := &file_form_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeStateReq.ProtoReflect.Descriptor instead.
func (*ChangeStateReq) Descriptor() ([]byte, []int) {
	return file_form_proto_rawDescGZIP(), []int{2}
}

func (x *ChangeStateReq) GetId() *IDReqOrResponse {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ChangeStateReq) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type ListFormResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DefaultResponse *DefaultResponse `protobuf:"bytes,1,opt,name=default_response,json=defaultResponse,proto3" json:"default_response,omitempty"`
	Forms           []*Form          `protobuf:"bytes,2,rep,name=forms,proto3" json:"forms,omitempty"`
}

func (x *ListFormResponse) Reset() {
	*x = ListFormResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_form_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFormResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFormResponse) ProtoMessage() {}

func (x *ListFormResponse) ProtoReflect() protoreflect.Message {
	mi := &file_form_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFormResponse.ProtoReflect.Descriptor instead.
func (*ListFormResponse) Descriptor() ([]byte, []int) {
	return file_form_proto_rawDescGZIP(), []int{3}
}

func (x *ListFormResponse) GetDefaultResponse() *DefaultResponse {
	if x != nil {
		return x.DefaultResponse
	}
	return nil
}

func (x *ListFormResponse) GetForms() []*Form {
	if x != nil {
		return x.Forms
	}
	return nil
}

type GetFormResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DefaultResponse *DefaultResponse `protobuf:"bytes,1,opt,name=default_response,json=defaultResponse,proto3" json:"default_response,omitempty"`
	Form            *Form            `protobuf:"bytes,2,opt,name=form,proto3" json:"form,omitempty"`
}

func (x *GetFormResponse) Reset() {
	*x = GetFormResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_form_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFormResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFormResponse) ProtoMessage() {}

func (x *GetFormResponse) ProtoReflect() protoreflect.Message {
	mi := &file_form_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFormResponse.ProtoReflect.Descriptor instead.
func (*GetFormResponse) Descriptor() ([]byte, []int) {
	return file_form_proto_rawDescGZIP(), []int{4}
}

func (x *GetFormResponse) GetDefaultResponse() *DefaultResponse {
	if x != nil {
		return x.DefaultResponse
	}
	return nil
}

func (x *GetFormResponse) GetForm() *Form {
	if x != nil {
		return x.Form
	}
	return nil
}

type Form struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key             string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Name            string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description     string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	GroupKey        []byte `protobuf:"bytes,5,opt,name=group_key,json=groupKey,proto3" json:"group_key,omitempty"`
	Version         int32  `protobuf:"varint,6,opt,name=version,proto3" json:"version,omitempty"`
	LatestVersion   bool   `protobuf:"varint,7,opt,name=latest_version,json=latestVersion,proto3" json:"latest_version,omitempty"`
	DataScheme      []byte `protobuf:"bytes,8,opt,name=data_scheme,json=dataScheme,proto3" json:"data_scheme,omitempty"`
	EsIndexInterval string `protobuf:"bytes,9,opt,name=es_index_interval,json=esIndexInterval,proto3" json:"es_index_interval,omitempty"`
	State           string `protobuf:"bytes,10,opt,name=state,proto3" json:"state,omitempty"`
	Creator         string `protobuf:"bytes,11,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (x *Form) Reset() {
	*x = Form{}
	if protoimpl.UnsafeEnabled {
		mi := &file_form_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Form) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Form) ProtoMessage() {}

func (x *Form) ProtoReflect() protoreflect.Message {
	mi := &file_form_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Form.ProtoReflect.Descriptor instead.
func (*Form) Descriptor() ([]byte, []int) {
	return file_form_proto_rawDescGZIP(), []int{5}
}

func (x *Form) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Form) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Form) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Form) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Form) GetGroupKey() []byte {
	if x != nil {
		return x.GroupKey
	}
	return nil
}

func (x *Form) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Form) GetLatestVersion() bool {
	if x != nil {
		return x.LatestVersion
	}
	return false
}

func (x *Form) GetDataScheme() []byte {
	if x != nil {
		return x.DataScheme
	}
	return nil
}

func (x *Form) GetEsIndexInterval() string {
	if x != nil {
		return x.EsIndexInterval
	}
	return ""
}

func (x *Form) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Form) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

var File_form_proto protoreflect.FileDescriptor

var file_form_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x69, 0x64,
	0x6c, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a,
	0x14, 0x46, 0x6f, 0x72, 0x6d, 0x41, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x10, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1d, 0x0a, 0x04, 0x66, 0x6f, 0x72, 0x6d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x46, 0x6f, 0x72,
	0x6d, 0x52, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x54, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x4f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x69,
	0x64, 0x6c, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x4c, 0x0a,
	0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x24, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x69, 0x64,
	0x6c, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x4f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x74, 0x0a, 0x10, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3f, 0x0a, 0x10, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x69, 0x64, 0x6c, 0x2e,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x6d,
	0x73, 0x22, 0x71, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x10, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x04,
	0x66, 0x6f, 0x72, 0x6d, 0x22, 0xb9, 0x02, 0x0a, 0x04, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4b,
	0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x65, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x32, 0xa3, 0x02, 0x0a, 0x07, 0x46, 0x6f, 0x72, 0x6d, 0x53, 0x72, 0x76, 0x12, 0x2e, 0x0a, 0x08,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x09, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x1a, 0x15, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f,
	0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x14, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x4f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x14, 0x2e,
	0x69, 0x64, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x69, 0x64, 0x6c, 0x2e,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x34, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x12,
	0x09, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x1a, 0x19, 0x2e, 0x69, 0x64, 0x6c,
	0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x41, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x12, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x69, 0x64, 0x6c, 0x2e,
	0x46, 0x6f, 0x72, 0x6d, 0x41, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1e, 0x50, 0x01, 0x5a, 0x1a, 0x65, 0x76, 0x69, 0x6c,
	0x64, 0x6f, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_form_proto_rawDescOnce sync.Once
	file_form_proto_rawDescData = file_form_proto_rawDesc
)

func file_form_proto_rawDescGZIP() []byte {
	file_form_proto_rawDescOnce.Do(func() {
		file_form_proto_rawDescData = protoimpl.X.CompressGZIP(file_form_proto_rawDescData)
	})
	return file_form_proto_rawDescData
}

var file_form_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_form_proto_goTypes = []interface{}{
	(*FormAndGroupResponse)(nil), // 0: idl.FormAndGroupResponse
	(*UpdateFormReq)(nil),        // 1: idl.UpdateFormReq
	(*ChangeStateReq)(nil),       // 2: idl.ChangeStateReq
	(*ListFormResponse)(nil),     // 3: idl.ListFormResponse
	(*GetFormResponse)(nil),      // 4: idl.GetFormResponse
	(*Form)(nil),                 // 5: idl.Form
	(*DefaultResponse)(nil),      // 6: idl.DefaultResponse
	(*Group)(nil),                // 7: idl.Group
	(*IDReqOrResponse)(nil),      // 8: idl.IDReqOrResponse
	(*Page)(nil),                 // 9: idl.Page
}
var file_form_proto_depIdxs = []int32{
	6,  // 0: idl.FormAndGroupResponse.default_response:type_name -> idl.DefaultResponse
	7,  // 1: idl.FormAndGroupResponse.group:type_name -> idl.Group
	5,  // 2: idl.FormAndGroupResponse.form:type_name -> idl.Form
	8,  // 3: idl.UpdateFormReq.id:type_name -> idl.IDReqOrResponse
	5,  // 4: idl.UpdateFormReq.form:type_name -> idl.Form
	8,  // 5: idl.ChangeStateReq.id:type_name -> idl.IDReqOrResponse
	6,  // 6: idl.ListFormResponse.default_response:type_name -> idl.DefaultResponse
	5,  // 7: idl.ListFormResponse.forms:type_name -> idl.Form
	6,  // 8: idl.GetFormResponse.default_response:type_name -> idl.DefaultResponse
	5,  // 9: idl.GetFormResponse.form:type_name -> idl.Form
	9,  // 10: idl.FormSrv.ListForm:input_type -> idl.Page
	8,  // 11: idl.FormSrv.GetForm:input_type -> idl.IDReqOrResponse
	2,  // 12: idl.FormSrv.ChangeState:input_type -> idl.ChangeStateReq
	5,  // 13: idl.FormSrv.CreateForm:input_type -> idl.Form
	1,  // 14: idl.FormSrv.UpdateForm:input_type -> idl.UpdateFormReq
	3,  // 15: idl.FormSrv.ListForm:output_type -> idl.ListFormResponse
	4,  // 16: idl.FormSrv.GetForm:output_type -> idl.GetFormResponse
	6,  // 17: idl.FormSrv.ChangeState:output_type -> idl.DefaultResponse
	0,  // 18: idl.FormSrv.CreateForm:output_type -> idl.FormAndGroupResponse
	0,  // 19: idl.FormSrv.UpdateForm:output_type -> idl.FormAndGroupResponse
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_form_proto_init() }
func file_form_proto_init() {
	if File_form_proto != nil {
		return
	}
	file_common_proto_init()
	file_group_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_form_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FormAndGroupResponse); i {
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
		file_form_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFormReq); i {
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
		file_form_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeStateReq); i {
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
		file_form_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFormResponse); i {
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
		file_form_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFormResponse); i {
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
		file_form_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Form); i {
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
			RawDescriptor: file_form_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_form_proto_goTypes,
		DependencyIndexes: file_form_proto_depIdxs,
		MessageInfos:      file_form_proto_msgTypes,
	}.Build()
	File_form_proto = out.File
	file_form_proto_rawDesc = nil
	file_form_proto_goTypes = nil
	file_form_proto_depIdxs = nil
}