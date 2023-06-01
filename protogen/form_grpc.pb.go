// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: form.proto

package protogen

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

const (
	FormSrv_ListForm_FullMethodName    = "/idl.FormSrv/ListForm"
	FormSrv_GetForm_FullMethodName     = "/idl.FormSrv/GetForm"
	FormSrv_ChangeState_FullMethodName = "/idl.FormSrv/ChangeState"
	FormSrv_CreateForm_FullMethodName  = "/idl.FormSrv/CreateForm"
	FormSrv_UpdateForm_FullMethodName  = "/idl.FormSrv/UpdateForm"
)

// FormSrvClient is the client API for FormSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FormSrvClient interface {
	ListForm(ctx context.Context, in *Page, opts ...grpc.CallOption) (*ListFormResponse, error)
	GetForm(ctx context.Context, in *IDReqOrResponse, opts ...grpc.CallOption) (*GetFormResponse, error)
	ChangeState(ctx context.Context, in *ChangeStateReq, opts ...grpc.CallOption) (*DefaultResponse, error)
	CreateForm(ctx context.Context, in *Form, opts ...grpc.CallOption) (*FormAndGroupResponse, error)
	UpdateForm(ctx context.Context, in *UpdateFormReq, opts ...grpc.CallOption) (*FormAndGroupResponse, error)
}

type formSrvClient struct {
	cc grpc.ClientConnInterface
}

func NewFormSrvClient(cc grpc.ClientConnInterface) FormSrvClient {
	return &formSrvClient{cc}
}

func (c *formSrvClient) ListForm(ctx context.Context, in *Page, opts ...grpc.CallOption) (*ListFormResponse, error) {
	out := new(ListFormResponse)
	err := c.cc.Invoke(ctx, FormSrv_ListForm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *formSrvClient) GetForm(ctx context.Context, in *IDReqOrResponse, opts ...grpc.CallOption) (*GetFormResponse, error) {
	out := new(GetFormResponse)
	err := c.cc.Invoke(ctx, FormSrv_GetForm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *formSrvClient) ChangeState(ctx context.Context, in *ChangeStateReq, opts ...grpc.CallOption) (*DefaultResponse, error) {
	out := new(DefaultResponse)
	err := c.cc.Invoke(ctx, FormSrv_ChangeState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *formSrvClient) CreateForm(ctx context.Context, in *Form, opts ...grpc.CallOption) (*FormAndGroupResponse, error) {
	out := new(FormAndGroupResponse)
	err := c.cc.Invoke(ctx, FormSrv_CreateForm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *formSrvClient) UpdateForm(ctx context.Context, in *UpdateFormReq, opts ...grpc.CallOption) (*FormAndGroupResponse, error) {
	out := new(FormAndGroupResponse)
	err := c.cc.Invoke(ctx, FormSrv_UpdateForm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FormSrvServer is the server API for FormSrv service.
// All implementations must embed UnimplementedFormSrvServer
// for forward compatibility
type FormSrvServer interface {
	ListForm(context.Context, *Page) (*ListFormResponse, error)
	GetForm(context.Context, *IDReqOrResponse) (*GetFormResponse, error)
	ChangeState(context.Context, *ChangeStateReq) (*DefaultResponse, error)
	CreateForm(context.Context, *Form) (*FormAndGroupResponse, error)
	UpdateForm(context.Context, *UpdateFormReq) (*FormAndGroupResponse, error)
	mustEmbedUnimplementedFormSrvServer()
}

// UnimplementedFormSrvServer must be embedded to have forward compatible implementations.
type UnimplementedFormSrvServer struct {
}

func (UnimplementedFormSrvServer) ListForm(context.Context, *Page) (*ListFormResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListForm not implemented")
}
func (UnimplementedFormSrvServer) GetForm(context.Context, *IDReqOrResponse) (*GetFormResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetForm not implemented")
}
func (UnimplementedFormSrvServer) ChangeState(context.Context, *ChangeStateReq) (*DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeState not implemented")
}
func (UnimplementedFormSrvServer) CreateForm(context.Context, *Form) (*FormAndGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateForm not implemented")
}
func (UnimplementedFormSrvServer) UpdateForm(context.Context, *UpdateFormReq) (*FormAndGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateForm not implemented")
}
func (UnimplementedFormSrvServer) mustEmbedUnimplementedFormSrvServer() {}

// UnsafeFormSrvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FormSrvServer will
// result in compilation errors.
type UnsafeFormSrvServer interface {
	mustEmbedUnimplementedFormSrvServer()
}

func RegisterFormSrvServer(s grpc.ServiceRegistrar, srv FormSrvServer) {
	s.RegisterService(&FormSrv_ServiceDesc, srv)
}

func _FormSrv_ListForm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Page)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FormSrvServer).ListForm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FormSrv_ListForm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FormSrvServer).ListForm(ctx, req.(*Page))
	}
	return interceptor(ctx, in, info, handler)
}

func _FormSrv_GetForm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDReqOrResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FormSrvServer).GetForm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FormSrv_GetForm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FormSrvServer).GetForm(ctx, req.(*IDReqOrResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _FormSrv_ChangeState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeStateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FormSrvServer).ChangeState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FormSrv_ChangeState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FormSrvServer).ChangeState(ctx, req.(*ChangeStateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FormSrv_CreateForm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Form)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FormSrvServer).CreateForm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FormSrv_CreateForm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FormSrvServer).CreateForm(ctx, req.(*Form))
	}
	return interceptor(ctx, in, info, handler)
}

func _FormSrv_UpdateForm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFormReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FormSrvServer).UpdateForm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FormSrv_UpdateForm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FormSrvServer).UpdateForm(ctx, req.(*UpdateFormReq))
	}
	return interceptor(ctx, in, info, handler)
}

// FormSrv_ServiceDesc is the grpc.ServiceDesc for FormSrv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FormSrv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "idl.FormSrv",
	HandlerType: (*FormSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListForm",
			Handler:    _FormSrv_ListForm_Handler,
		},
		{
			MethodName: "GetForm",
			Handler:    _FormSrv_GetForm_Handler,
		},
		{
			MethodName: "ChangeState",
			Handler:    _FormSrv_ChangeState_Handler,
		},
		{
			MethodName: "CreateForm",
			Handler:    _FormSrv_CreateForm_Handler,
		},
		{
			MethodName: "UpdateForm",
			Handler:    _FormSrv_UpdateForm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "form.proto",
}