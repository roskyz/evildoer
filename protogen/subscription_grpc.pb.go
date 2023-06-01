// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: subscription.proto

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
	SubscriptionSrv_ListSubscription_FullMethodName   = "/idl.SubscriptionSrv/ListSubscription"
	SubscriptionSrv_GetSubscription_FullMethodName    = "/idl.SubscriptionSrv/GetSubscription"
	SubscriptionSrv_CreateSubscription_FullMethodName = "/idl.SubscriptionSrv/CreateSubscription"
)

// SubscriptionSrvClient is the client API for SubscriptionSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubscriptionSrvClient interface {
	ListSubscription(ctx context.Context, in *Page, opts ...grpc.CallOption) (*ListSubscriptionResponse, error)
	GetSubscription(ctx context.Context, in *IDReqOrResponse, opts ...grpc.CallOption) (*SubscriptionResponse, error)
	CreateSubscription(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*SubscriptionResponse, error)
}

type subscriptionSrvClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscriptionSrvClient(cc grpc.ClientConnInterface) SubscriptionSrvClient {
	return &subscriptionSrvClient{cc}
}

func (c *subscriptionSrvClient) ListSubscription(ctx context.Context, in *Page, opts ...grpc.CallOption) (*ListSubscriptionResponse, error) {
	out := new(ListSubscriptionResponse)
	err := c.cc.Invoke(ctx, SubscriptionSrv_ListSubscription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionSrvClient) GetSubscription(ctx context.Context, in *IDReqOrResponse, opts ...grpc.CallOption) (*SubscriptionResponse, error) {
	out := new(SubscriptionResponse)
	err := c.cc.Invoke(ctx, SubscriptionSrv_GetSubscription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionSrvClient) CreateSubscription(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*SubscriptionResponse, error) {
	out := new(SubscriptionResponse)
	err := c.cc.Invoke(ctx, SubscriptionSrv_CreateSubscription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriptionSrvServer is the server API for SubscriptionSrv service.
// All implementations must embed UnimplementedSubscriptionSrvServer
// for forward compatibility
type SubscriptionSrvServer interface {
	ListSubscription(context.Context, *Page) (*ListSubscriptionResponse, error)
	GetSubscription(context.Context, *IDReqOrResponse) (*SubscriptionResponse, error)
	CreateSubscription(context.Context, *Subscription) (*SubscriptionResponse, error)
	mustEmbedUnimplementedSubscriptionSrvServer()
}

// UnimplementedSubscriptionSrvServer must be embedded to have forward compatible implementations.
type UnimplementedSubscriptionSrvServer struct {
}

func (UnimplementedSubscriptionSrvServer) ListSubscription(context.Context, *Page) (*ListSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSubscription not implemented")
}
func (UnimplementedSubscriptionSrvServer) GetSubscription(context.Context, *IDReqOrResponse) (*SubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscription not implemented")
}
func (UnimplementedSubscriptionSrvServer) CreateSubscription(context.Context, *Subscription) (*SubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubscription not implemented")
}
func (UnimplementedSubscriptionSrvServer) mustEmbedUnimplementedSubscriptionSrvServer() {}

// UnsafeSubscriptionSrvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscriptionSrvServer will
// result in compilation errors.
type UnsafeSubscriptionSrvServer interface {
	mustEmbedUnimplementedSubscriptionSrvServer()
}

func RegisterSubscriptionSrvServer(s grpc.ServiceRegistrar, srv SubscriptionSrvServer) {
	s.RegisterService(&SubscriptionSrv_ServiceDesc, srv)
}

func _SubscriptionSrv_ListSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Page)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionSrvServer).ListSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionSrv_ListSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionSrvServer).ListSubscription(ctx, req.(*Page))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionSrv_GetSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDReqOrResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionSrvServer).GetSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionSrv_GetSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionSrvServer).GetSubscription(ctx, req.(*IDReqOrResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionSrv_CreateSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Subscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionSrvServer).CreateSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionSrv_CreateSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionSrvServer).CreateSubscription(ctx, req.(*Subscription))
	}
	return interceptor(ctx, in, info, handler)
}

// SubscriptionSrv_ServiceDesc is the grpc.ServiceDesc for SubscriptionSrv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubscriptionSrv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "idl.SubscriptionSrv",
	HandlerType: (*SubscriptionSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSubscription",
			Handler:    _SubscriptionSrv_ListSubscription_Handler,
		},
		{
			MethodName: "GetSubscription",
			Handler:    _SubscriptionSrv_GetSubscription_Handler,
		},
		{
			MethodName: "CreateSubscription",
			Handler:    _SubscriptionSrv_CreateSubscription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "subscription.proto",
}
