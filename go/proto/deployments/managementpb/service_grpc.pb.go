// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: deployments/management/service.proto

package managementpb

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
	Service_CreateDeployment_FullMethodName       = "/freym.deployments.management.Service/CreateDeployment"
	Service_ActivateDeployment_FullMethodName     = "/freym.deployments.management.Service/ActivateDeployment"
	Service_ConfirmDeployment_FullMethodName      = "/freym.deployments.management.Service/ConfirmDeployment"
	Service_RollbackDeployment_FullMethodName     = "/freym.deployments.management.Service/RollbackDeployment"
	Service_RollbackDeploymentById_FullMethodName = "/freym.deployments.management.Service/RollbackDeploymentById"
	Service_GetDeployment_FullMethodName          = "/freym.deployments.management.Service/GetDeployment"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	CreateDeployment(ctx context.Context, in *CreateDeploymentRequest, opts ...grpc.CallOption) (*CreateDeploymentResponse, error)
	ActivateDeployment(ctx context.Context, in *ActivateDeploymentRequest, opts ...grpc.CallOption) (*ActivateDeploymentResponse, error)
	ConfirmDeployment(ctx context.Context, in *ConfirmDeploymentRequest, opts ...grpc.CallOption) (*ConfirmDeploymentResponse, error)
	RollbackDeployment(ctx context.Context, in *RollbackDeploymentRequest, opts ...grpc.CallOption) (*RollbackDeploymentResponse, error)
	RollbackDeploymentById(ctx context.Context, in *RollbackDeploymentByIdRequest, opts ...grpc.CallOption) (*RollbackDeploymentResponse, error)
	GetDeployment(ctx context.Context, in *GetDeploymentRequest, opts ...grpc.CallOption) (*GetDeploymentResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) CreateDeployment(ctx context.Context, in *CreateDeploymentRequest, opts ...grpc.CallOption) (*CreateDeploymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDeploymentResponse)
	err := c.cc.Invoke(ctx, Service_CreateDeployment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ActivateDeployment(ctx context.Context, in *ActivateDeploymentRequest, opts ...grpc.CallOption) (*ActivateDeploymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ActivateDeploymentResponse)
	err := c.cc.Invoke(ctx, Service_ActivateDeployment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ConfirmDeployment(ctx context.Context, in *ConfirmDeploymentRequest, opts ...grpc.CallOption) (*ConfirmDeploymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConfirmDeploymentResponse)
	err := c.cc.Invoke(ctx, Service_ConfirmDeployment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) RollbackDeployment(ctx context.Context, in *RollbackDeploymentRequest, opts ...grpc.CallOption) (*RollbackDeploymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RollbackDeploymentResponse)
	err := c.cc.Invoke(ctx, Service_RollbackDeployment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) RollbackDeploymentById(ctx context.Context, in *RollbackDeploymentByIdRequest, opts ...grpc.CallOption) (*RollbackDeploymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RollbackDeploymentResponse)
	err := c.cc.Invoke(ctx, Service_RollbackDeploymentById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetDeployment(ctx context.Context, in *GetDeploymentRequest, opts ...grpc.CallOption) (*GetDeploymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDeploymentResponse)
	err := c.cc.Invoke(ctx, Service_GetDeployment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility.
type ServiceServer interface {
	CreateDeployment(context.Context, *CreateDeploymentRequest) (*CreateDeploymentResponse, error)
	ActivateDeployment(context.Context, *ActivateDeploymentRequest) (*ActivateDeploymentResponse, error)
	ConfirmDeployment(context.Context, *ConfirmDeploymentRequest) (*ConfirmDeploymentResponse, error)
	RollbackDeployment(context.Context, *RollbackDeploymentRequest) (*RollbackDeploymentResponse, error)
	RollbackDeploymentById(context.Context, *RollbackDeploymentByIdRequest) (*RollbackDeploymentResponse, error)
	GetDeployment(context.Context, *GetDeploymentRequest) (*GetDeploymentResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedServiceServer struct{}

func (UnimplementedServiceServer) CreateDeployment(context.Context, *CreateDeploymentRequest) (*CreateDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeployment not implemented")
}
func (UnimplementedServiceServer) ActivateDeployment(context.Context, *ActivateDeploymentRequest) (*ActivateDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateDeployment not implemented")
}
func (UnimplementedServiceServer) ConfirmDeployment(context.Context, *ConfirmDeploymentRequest) (*ConfirmDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmDeployment not implemented")
}
func (UnimplementedServiceServer) RollbackDeployment(context.Context, *RollbackDeploymentRequest) (*RollbackDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackDeployment not implemented")
}
func (UnimplementedServiceServer) RollbackDeploymentById(context.Context, *RollbackDeploymentByIdRequest) (*RollbackDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackDeploymentById not implemented")
}
func (UnimplementedServiceServer) GetDeployment(context.Context, *GetDeploymentRequest) (*GetDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeployment not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}
func (UnimplementedServiceServer) testEmbeddedByValue()                 {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	// If the following call pancis, it indicates UnimplementedServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_CreateDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_CreateDeployment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateDeployment(ctx, req.(*CreateDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ActivateDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ActivateDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_ActivateDeployment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ActivateDeployment(ctx, req.(*ActivateDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ConfirmDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ConfirmDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_ConfirmDeployment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ConfirmDeployment(ctx, req.(*ConfirmDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_RollbackDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).RollbackDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_RollbackDeployment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).RollbackDeployment(ctx, req.(*RollbackDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_RollbackDeploymentById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackDeploymentByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).RollbackDeploymentById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_RollbackDeploymentById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).RollbackDeploymentById(ctx, req.(*RollbackDeploymentByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetDeployment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetDeployment(ctx, req.(*GetDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "freym.deployments.management.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDeployment",
			Handler:    _Service_CreateDeployment_Handler,
		},
		{
			MethodName: "ActivateDeployment",
			Handler:    _Service_ActivateDeployment_Handler,
		},
		{
			MethodName: "ConfirmDeployment",
			Handler:    _Service_ConfirmDeployment_Handler,
		},
		{
			MethodName: "RollbackDeployment",
			Handler:    _Service_RollbackDeployment_Handler,
		},
		{
			MethodName: "RollbackDeploymentById",
			Handler:    _Service_RollbackDeploymentById_Handler,
		},
		{
			MethodName: "GetDeployment",
			Handler:    _Service_GetDeployment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deployments/management/service.proto",
}
