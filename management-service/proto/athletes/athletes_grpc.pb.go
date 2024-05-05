// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: athletes.proto

package athletes

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
	AthleteService_CreateAthlete_FullMethodName               = "/athletes.AthleteService/CreateAthlete"
	AthleteService_GetAthletes_FullMethodName                 = "/athletes.AthleteService/GetAthletes"
	AthleteService_DeleteAthletes_FullMethodName              = "/athletes.AthleteService/DeleteAthletes"
	AthleteService_EditAthlete_FullMethodName                 = "/athletes.AthleteService/EditAthlete"
	AthleteService_AddTraining_FullMethodName                 = "/athletes.AthleteService/AddTraining"
	AthleteService_MinusTrainingOfSubscription_FullMethodName = "/athletes.AthleteService/MinusTrainingOfSubscription"
)

// AthleteServiceClient is the client API for AthleteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AthleteServiceClient interface {
	CreateAthlete(ctx context.Context, in *AthleteRequest, opts ...grpc.CallOption) (*AthleteResponse, error)
	GetAthletes(ctx context.Context, in *GetAthletesRequest, opts ...grpc.CallOption) (AthleteService_GetAthletesClient, error)
	DeleteAthletes(ctx context.Context, in *DeleteAthletesRequest, opts ...grpc.CallOption) (*AthleteResponse, error)
	EditAthlete(ctx context.Context, in *EditAthletesRequest, opts ...grpc.CallOption) (*AthleteResponse, error)
	AddTraining(ctx context.Context, in *AddTrainingRequest, opts ...grpc.CallOption) (*AthleteResponse, error)
	MinusTrainingOfSubscription(ctx context.Context, in *MinusTrainingOfSubscriptionRequest, opts ...grpc.CallOption) (*AthleteResponse, error)
}

type athleteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAthleteServiceClient(cc grpc.ClientConnInterface) AthleteServiceClient {
	return &athleteServiceClient{cc}
}

func (c *athleteServiceClient) CreateAthlete(ctx context.Context, in *AthleteRequest, opts ...grpc.CallOption) (*AthleteResponse, error) {
	out := new(AthleteResponse)
	err := c.cc.Invoke(ctx, AthleteService_CreateAthlete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *athleteServiceClient) GetAthletes(ctx context.Context, in *GetAthletesRequest, opts ...grpc.CallOption) (AthleteService_GetAthletesClient, error) {
	stream, err := c.cc.NewStream(ctx, &AthleteService_ServiceDesc.Streams[0], AthleteService_GetAthletes_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &athleteServiceGetAthletesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AthleteService_GetAthletesClient interface {
	Recv() (*AthleteRequest, error)
	grpc.ClientStream
}

type athleteServiceGetAthletesClient struct {
	grpc.ClientStream
}

func (x *athleteServiceGetAthletesClient) Recv() (*AthleteRequest, error) {
	m := new(AthleteRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *athleteServiceClient) DeleteAthletes(ctx context.Context, in *DeleteAthletesRequest, opts ...grpc.CallOption) (*AthleteResponse, error) {
	out := new(AthleteResponse)
	err := c.cc.Invoke(ctx, AthleteService_DeleteAthletes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *athleteServiceClient) EditAthlete(ctx context.Context, in *EditAthletesRequest, opts ...grpc.CallOption) (*AthleteResponse, error) {
	out := new(AthleteResponse)
	err := c.cc.Invoke(ctx, AthleteService_EditAthlete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *athleteServiceClient) AddTraining(ctx context.Context, in *AddTrainingRequest, opts ...grpc.CallOption) (*AthleteResponse, error) {
	out := new(AthleteResponse)
	err := c.cc.Invoke(ctx, AthleteService_AddTraining_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *athleteServiceClient) MinusTrainingOfSubscription(ctx context.Context, in *MinusTrainingOfSubscriptionRequest, opts ...grpc.CallOption) (*AthleteResponse, error) {
	out := new(AthleteResponse)
	err := c.cc.Invoke(ctx, AthleteService_MinusTrainingOfSubscription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AthleteServiceServer is the server API for AthleteService service.
// All implementations must embed UnimplementedAthleteServiceServer
// for forward compatibility
type AthleteServiceServer interface {
	CreateAthlete(context.Context, *AthleteRequest) (*AthleteResponse, error)
	GetAthletes(*GetAthletesRequest, AthleteService_GetAthletesServer) error
	DeleteAthletes(context.Context, *DeleteAthletesRequest) (*AthleteResponse, error)
	EditAthlete(context.Context, *EditAthletesRequest) (*AthleteResponse, error)
	AddTraining(context.Context, *AddTrainingRequest) (*AthleteResponse, error)
	MinusTrainingOfSubscription(context.Context, *MinusTrainingOfSubscriptionRequest) (*AthleteResponse, error)
	mustEmbedUnimplementedAthleteServiceServer()
}

// UnimplementedAthleteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAthleteServiceServer struct {
}

func (UnimplementedAthleteServiceServer) CreateAthlete(context.Context, *AthleteRequest) (*AthleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAthlete not implemented")
}
func (UnimplementedAthleteServiceServer) GetAthletes(*GetAthletesRequest, AthleteService_GetAthletesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAthletes not implemented")
}
func (UnimplementedAthleteServiceServer) DeleteAthletes(context.Context, *DeleteAthletesRequest) (*AthleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAthletes not implemented")
}
func (UnimplementedAthleteServiceServer) EditAthlete(context.Context, *EditAthletesRequest) (*AthleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditAthlete not implemented")
}
func (UnimplementedAthleteServiceServer) AddTraining(context.Context, *AddTrainingRequest) (*AthleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTraining not implemented")
}
func (UnimplementedAthleteServiceServer) MinusTrainingOfSubscription(context.Context, *MinusTrainingOfSubscriptionRequest) (*AthleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MinusTrainingOfSubscription not implemented")
}
func (UnimplementedAthleteServiceServer) mustEmbedUnimplementedAthleteServiceServer() {}

// UnsafeAthleteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AthleteServiceServer will
// result in compilation errors.
type UnsafeAthleteServiceServer interface {
	mustEmbedUnimplementedAthleteServiceServer()
}

func RegisterAthleteServiceServer(s grpc.ServiceRegistrar, srv AthleteServiceServer) {
	s.RegisterService(&AthleteService_ServiceDesc, srv)
}

func _AthleteService_CreateAthlete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AthleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AthleteServiceServer).CreateAthlete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AthleteService_CreateAthlete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AthleteServiceServer).CreateAthlete(ctx, req.(*AthleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AthleteService_GetAthletes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAthletesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AthleteServiceServer).GetAthletes(m, &athleteServiceGetAthletesServer{stream})
}

type AthleteService_GetAthletesServer interface {
	Send(*AthleteRequest) error
	grpc.ServerStream
}

type athleteServiceGetAthletesServer struct {
	grpc.ServerStream
}

func (x *athleteServiceGetAthletesServer) Send(m *AthleteRequest) error {
	return x.ServerStream.SendMsg(m)
}

func _AthleteService_DeleteAthletes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAthletesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AthleteServiceServer).DeleteAthletes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AthleteService_DeleteAthletes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AthleteServiceServer).DeleteAthletes(ctx, req.(*DeleteAthletesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AthleteService_EditAthlete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditAthletesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AthleteServiceServer).EditAthlete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AthleteService_EditAthlete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AthleteServiceServer).EditAthlete(ctx, req.(*EditAthletesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AthleteService_AddTraining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTrainingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AthleteServiceServer).AddTraining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AthleteService_AddTraining_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AthleteServiceServer).AddTraining(ctx, req.(*AddTrainingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AthleteService_MinusTrainingOfSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinusTrainingOfSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AthleteServiceServer).MinusTrainingOfSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AthleteService_MinusTrainingOfSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AthleteServiceServer).MinusTrainingOfSubscription(ctx, req.(*MinusTrainingOfSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AthleteService_ServiceDesc is the grpc.ServiceDesc for AthleteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AthleteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "athletes.AthleteService",
	HandlerType: (*AthleteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAthlete",
			Handler:    _AthleteService_CreateAthlete_Handler,
		},
		{
			MethodName: "DeleteAthletes",
			Handler:    _AthleteService_DeleteAthletes_Handler,
		},
		{
			MethodName: "EditAthlete",
			Handler:    _AthleteService_EditAthlete_Handler,
		},
		{
			MethodName: "AddTraining",
			Handler:    _AthleteService_AddTraining_Handler,
		},
		{
			MethodName: "MinusTrainingOfSubscription",
			Handler:    _AthleteService_MinusTrainingOfSubscription_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAthletes",
			Handler:       _AthleteService_GetAthletes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "athletes.proto",
}
