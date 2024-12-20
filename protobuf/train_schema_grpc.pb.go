// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: train_schema.proto

package train

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TrainService_PurchaseTicket_FullMethodName    = "/train.TrainService/PurchaseTicket"
	TrainService_GetReceipt_FullMethodName        = "/train.TrainService/GetReceipt"
	TrainService_GetUsersBySection_FullMethodName = "/train.TrainService/GetUsersBySection"
	TrainService_RemoveUser_FullMethodName        = "/train.TrainService/RemoveUser"
	TrainService_ModifyUserSeat_FullMethodName    = "/train.TrainService/ModifyUserSeat"
)

// TrainServiceClient is the client API for TrainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrainServiceClient interface {
	PurchaseTicket(ctx context.Context, in *PurchaseTicketRequest, opts ...grpc.CallOption) (*TicketReceipt, error)
	GetReceipt(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*TicketReceipt, error)
	GetUsersBySection(ctx context.Context, in *SectionRequest, opts ...grpc.CallOption) (*UsersResponse, error)
	RemoveUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	ModifyUserSeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*TicketReceipt, error)
}

type trainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrainServiceClient(cc grpc.ClientConnInterface) TrainServiceClient {
	return &trainServiceClient{cc}
}

func (c *trainServiceClient) PurchaseTicket(ctx context.Context, in *PurchaseTicketRequest, opts ...grpc.CallOption) (*TicketReceipt, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TicketReceipt)
	err := c.cc.Invoke(ctx, TrainService_PurchaseTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainServiceClient) GetReceipt(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*TicketReceipt, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TicketReceipt)
	err := c.cc.Invoke(ctx, TrainService_GetReceipt_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainServiceClient) GetUsersBySection(ctx context.Context, in *SectionRequest, opts ...grpc.CallOption) (*UsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UsersResponse)
	err := c.cc.Invoke(ctx, TrainService_GetUsersBySection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainServiceClient) RemoveUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, TrainService_RemoveUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainServiceClient) ModifyUserSeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*TicketReceipt, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TicketReceipt)
	err := c.cc.Invoke(ctx, TrainService_ModifyUserSeat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrainServiceServer is the server API for TrainService service.
// All implementations must embed UnimplementedTrainServiceServer
// for forward compatibility.
type TrainServiceServer interface {
	PurchaseTicket(context.Context, *PurchaseTicketRequest) (*TicketReceipt, error)
	GetReceipt(context.Context, *UserRequest) (*TicketReceipt, error)
	GetUsersBySection(context.Context, *SectionRequest) (*UsersResponse, error)
	RemoveUser(context.Context, *UserRequest) (*EmptyResponse, error)
	ModifyUserSeat(context.Context, *ModifySeatRequest) (*TicketReceipt, error)
	mustEmbedUnimplementedTrainServiceServer()
}

// UnimplementedTrainServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTrainServiceServer struct{}

func (UnimplementedTrainServiceServer) PurchaseTicket(context.Context, *PurchaseTicketRequest) (*TicketReceipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurchaseTicket not implemented")
}
func (UnimplementedTrainServiceServer) GetReceipt(context.Context, *UserRequest) (*TicketReceipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReceipt not implemented")
}
func (UnimplementedTrainServiceServer) GetUsersBySection(context.Context, *SectionRequest) (*UsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersBySection not implemented")
}
func (UnimplementedTrainServiceServer) RemoveUser(context.Context, *UserRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedTrainServiceServer) ModifyUserSeat(context.Context, *ModifySeatRequest) (*TicketReceipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyUserSeat not implemented")
}
func (UnimplementedTrainServiceServer) mustEmbedUnimplementedTrainServiceServer() {}
func (UnimplementedTrainServiceServer) testEmbeddedByValue()                      {}

// UnsafeTrainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrainServiceServer will
// result in compilation errors.
type UnsafeTrainServiceServer interface {
	mustEmbedUnimplementedTrainServiceServer()
}

func RegisterTrainServiceServer(s grpc.ServiceRegistrar, srv TrainServiceServer) {
	// If the following call pancis, it indicates UnimplementedTrainServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TrainService_ServiceDesc, srv)
}

func _TrainService_PurchaseTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).PurchaseTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainService_PurchaseTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).PurchaseTicket(ctx, req.(*PurchaseTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainService_GetReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).GetReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainService_GetReceipt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).GetReceipt(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainService_GetUsersBySection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).GetUsersBySection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainService_GetUsersBySection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).GetUsersBySection(ctx, req.(*SectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainService_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainService_RemoveUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).RemoveUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainService_ModifyUserSeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifySeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).ModifyUserSeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainService_ModifyUserSeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).ModifyUserSeat(ctx, req.(*ModifySeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TrainService_ServiceDesc is the grpc.ServiceDesc for TrainService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "train.TrainService",
	HandlerType: (*TrainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PurchaseTicket",
			Handler:    _TrainService_PurchaseTicket_Handler,
		},
		{
			MethodName: "GetReceipt",
			Handler:    _TrainService_GetReceipt_Handler,
		},
		{
			MethodName: "GetUsersBySection",
			Handler:    _TrainService_GetUsersBySection_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _TrainService_RemoveUser_Handler,
		},
		{
			MethodName: "ModifyUserSeat",
			Handler:    _TrainService_ModifyUserSeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "train_schema.proto",
}
