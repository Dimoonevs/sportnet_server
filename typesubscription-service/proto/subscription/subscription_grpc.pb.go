// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: subscription.proto

package subscription

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
	TypeSubscriptionService_CreateSubscription_FullMethodName  = "/subscription.TypeSubscriptionService/CreateSubscription"
	TypeSubscriptionService_GetAllSubscriptions_FullMethodName = "/subscription.TypeSubscriptionService/GetAllSubscriptions"
	TypeSubscriptionService_EditSubscription_FullMethodName    = "/subscription.TypeSubscriptionService/EditSubscription"
)

// TypeSubscriptionServiceClient is the client API for TypeSubscriptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TypeSubscriptionServiceClient interface {
	CreateSubscription(ctx context.Context, in *SubscriptionRequest, opts ...grpc.CallOption) (*SubscriptionResponse, error)
	GetAllSubscriptions(ctx context.Context, in *GetSubscriptionRequest, opts ...grpc.CallOption) (*GetSubscriptionResponse, error)
	EditSubscription(ctx context.Context, in *SubscriptionEditRequest, opts ...grpc.CallOption) (*SubscriptionResponse, error)
}

type typeSubscriptionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTypeSubscriptionServiceClient(cc grpc.ClientConnInterface) TypeSubscriptionServiceClient {
	return &typeSubscriptionServiceClient{cc}
}

func (c *typeSubscriptionServiceClient) CreateSubscription(ctx context.Context, in *SubscriptionRequest, opts ...grpc.CallOption) (*SubscriptionResponse, error) {
	out := new(SubscriptionResponse)
	err := c.cc.Invoke(ctx, TypeSubscriptionService_CreateSubscription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *typeSubscriptionServiceClient) GetAllSubscriptions(ctx context.Context, in *GetSubscriptionRequest, opts ...grpc.CallOption) (*GetSubscriptionResponse, error) {
	out := new(GetSubscriptionResponse)
	err := c.cc.Invoke(ctx, TypeSubscriptionService_GetAllSubscriptions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *typeSubscriptionServiceClient) EditSubscription(ctx context.Context, in *SubscriptionEditRequest, opts ...grpc.CallOption) (*SubscriptionResponse, error) {
	out := new(SubscriptionResponse)
	err := c.cc.Invoke(ctx, TypeSubscriptionService_EditSubscription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TypeSubscriptionServiceServer is the server API for TypeSubscriptionService service.
// All implementations must embed UnimplementedTypeSubscriptionServiceServer
// for forward compatibility
type TypeSubscriptionServiceServer interface {
	CreateSubscription(context.Context, *SubscriptionRequest) (*SubscriptionResponse, error)
	GetAllSubscriptions(context.Context, *GetSubscriptionRequest) (*GetSubscriptionResponse, error)
	EditSubscription(context.Context, *SubscriptionEditRequest) (*SubscriptionResponse, error)
	mustEmbedUnimplementedTypeSubscriptionServiceServer()
}

// UnimplementedTypeSubscriptionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTypeSubscriptionServiceServer struct {
}

func (UnimplementedTypeSubscriptionServiceServer) CreateSubscription(context.Context, *SubscriptionRequest) (*SubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubscription not implemented")
}
func (UnimplementedTypeSubscriptionServiceServer) GetAllSubscriptions(context.Context, *GetSubscriptionRequest) (*GetSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllSubscriptions not implemented")
}
func (UnimplementedTypeSubscriptionServiceServer) EditSubscription(context.Context, *SubscriptionEditRequest) (*SubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditSubscription not implemented")
}
func (UnimplementedTypeSubscriptionServiceServer) mustEmbedUnimplementedTypeSubscriptionServiceServer() {
}

// UnsafeTypeSubscriptionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TypeSubscriptionServiceServer will
// result in compilation errors.
type UnsafeTypeSubscriptionServiceServer interface {
	mustEmbedUnimplementedTypeSubscriptionServiceServer()
}

func RegisterTypeSubscriptionServiceServer(s grpc.ServiceRegistrar, srv TypeSubscriptionServiceServer) {
	s.RegisterService(&TypeSubscriptionService_ServiceDesc, srv)
}

func _TypeSubscriptionService_CreateSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypeSubscriptionServiceServer).CreateSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TypeSubscriptionService_CreateSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypeSubscriptionServiceServer).CreateSubscription(ctx, req.(*SubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TypeSubscriptionService_GetAllSubscriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypeSubscriptionServiceServer).GetAllSubscriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TypeSubscriptionService_GetAllSubscriptions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypeSubscriptionServiceServer).GetAllSubscriptions(ctx, req.(*GetSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TypeSubscriptionService_EditSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionEditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypeSubscriptionServiceServer).EditSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TypeSubscriptionService_EditSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypeSubscriptionServiceServer).EditSubscription(ctx, req.(*SubscriptionEditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TypeSubscriptionService_ServiceDesc is the grpc.ServiceDesc for TypeSubscriptionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TypeSubscriptionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "subscription.TypeSubscriptionService",
	HandlerType: (*TypeSubscriptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSubscription",
			Handler:    _TypeSubscriptionService_CreateSubscription_Handler,
		},
		{
			MethodName: "GetAllSubscriptions",
			Handler:    _TypeSubscriptionService_GetAllSubscriptions_Handler,
		},
		{
			MethodName: "EditSubscription",
			Handler:    _TypeSubscriptionService_EditSubscription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "subscription.proto",
}
