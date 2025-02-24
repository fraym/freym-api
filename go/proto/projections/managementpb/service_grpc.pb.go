// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.3
// source: projections/management/service.proto

package managementpb

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
	Service_DeploySchema_FullMethodName               = "/freym.projections.management.Service/DeploySchema"
	Service_ActivateSchema_FullMethodName             = "/freym.projections.management.Service/ActivateSchema"
	Service_ConfirmSchema_FullMethodName              = "/freym.projections.management.Service/ConfirmSchema"
	Service_RollbackSchema_FullMethodName             = "/freym.projections.management.Service/RollbackSchema"
	Service_RollbackSchemaByDeployment_FullMethodName = "/freym.projections.management.Service/RollbackSchemaByDeployment"
	Service_GetSchemaDeployment_FullMethodName        = "/freym.projections.management.Service/GetSchemaDeployment"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	DeploySchema(ctx context.Context, in *DeploySchemaRequest, opts ...grpc.CallOption) (*DeploySchemaResponse, error)
	ActivateSchema(ctx context.Context, in *ActivateSchemaRequest, opts ...grpc.CallOption) (*ActivateSchemaResponse, error)
	ConfirmSchema(ctx context.Context, in *ConfirmSchemaRequest, opts ...grpc.CallOption) (*ConfirmSchemaResponse, error)
	RollbackSchema(ctx context.Context, in *RollbackSchemaRequest, opts ...grpc.CallOption) (*RollbackSchemaResponse, error)
	RollbackSchemaByDeployment(ctx context.Context, in *RollbackSchemaByDeploymentRequest, opts ...grpc.CallOption) (*RollbackSchemaResponse, error)
	GetSchemaDeployment(ctx context.Context, in *GetSchemaDeploymentRequest, opts ...grpc.CallOption) (*GetSchemaDeploymentResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) DeploySchema(ctx context.Context, in *DeploySchemaRequest, opts ...grpc.CallOption) (*DeploySchemaResponse, error) {
	out := new(DeploySchemaResponse)
	err := c.cc.Invoke(ctx, Service_DeploySchema_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ActivateSchema(ctx context.Context, in *ActivateSchemaRequest, opts ...grpc.CallOption) (*ActivateSchemaResponse, error) {
	out := new(ActivateSchemaResponse)
	err := c.cc.Invoke(ctx, Service_ActivateSchema_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ConfirmSchema(ctx context.Context, in *ConfirmSchemaRequest, opts ...grpc.CallOption) (*ConfirmSchemaResponse, error) {
	out := new(ConfirmSchemaResponse)
	err := c.cc.Invoke(ctx, Service_ConfirmSchema_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) RollbackSchema(ctx context.Context, in *RollbackSchemaRequest, opts ...grpc.CallOption) (*RollbackSchemaResponse, error) {
	out := new(RollbackSchemaResponse)
	err := c.cc.Invoke(ctx, Service_RollbackSchema_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) RollbackSchemaByDeployment(ctx context.Context, in *RollbackSchemaByDeploymentRequest, opts ...grpc.CallOption) (*RollbackSchemaResponse, error) {
	out := new(RollbackSchemaResponse)
	err := c.cc.Invoke(ctx, Service_RollbackSchemaByDeployment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetSchemaDeployment(ctx context.Context, in *GetSchemaDeploymentRequest, opts ...grpc.CallOption) (*GetSchemaDeploymentResponse, error) {
	out := new(GetSchemaDeploymentResponse)
	err := c.cc.Invoke(ctx, Service_GetSchemaDeployment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	DeploySchema(context.Context, *DeploySchemaRequest) (*DeploySchemaResponse, error)
	ActivateSchema(context.Context, *ActivateSchemaRequest) (*ActivateSchemaResponse, error)
	ConfirmSchema(context.Context, *ConfirmSchemaRequest) (*ConfirmSchemaResponse, error)
	RollbackSchema(context.Context, *RollbackSchemaRequest) (*RollbackSchemaResponse, error)
	RollbackSchemaByDeployment(context.Context, *RollbackSchemaByDeploymentRequest) (*RollbackSchemaResponse, error)
	GetSchemaDeployment(context.Context, *GetSchemaDeploymentRequest) (*GetSchemaDeploymentResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) DeploySchema(context.Context, *DeploySchemaRequest) (*DeploySchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeploySchema not implemented")
}
func (UnimplementedServiceServer) ActivateSchema(context.Context, *ActivateSchemaRequest) (*ActivateSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateSchema not implemented")
}
func (UnimplementedServiceServer) ConfirmSchema(context.Context, *ConfirmSchemaRequest) (*ConfirmSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmSchema not implemented")
}
func (UnimplementedServiceServer) RollbackSchema(context.Context, *RollbackSchemaRequest) (*RollbackSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackSchema not implemented")
}
func (UnimplementedServiceServer) RollbackSchemaByDeployment(context.Context, *RollbackSchemaByDeploymentRequest) (*RollbackSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackSchemaByDeployment not implemented")
}
func (UnimplementedServiceServer) GetSchemaDeployment(context.Context, *GetSchemaDeploymentRequest) (*GetSchemaDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchemaDeployment not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_DeploySchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeploySchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeploySchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_DeploySchema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeploySchema(ctx, req.(*DeploySchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ActivateSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ActivateSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_ActivateSchema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ActivateSchema(ctx, req.(*ActivateSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ConfirmSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ConfirmSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_ConfirmSchema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ConfirmSchema(ctx, req.(*ConfirmSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_RollbackSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).RollbackSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_RollbackSchema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).RollbackSchema(ctx, req.(*RollbackSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_RollbackSchemaByDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackSchemaByDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).RollbackSchemaByDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_RollbackSchemaByDeployment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).RollbackSchemaByDeployment(ctx, req.(*RollbackSchemaByDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetSchemaDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSchemaDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetSchemaDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetSchemaDeployment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetSchemaDeployment(ctx, req.(*GetSchemaDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "freym.projections.management.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeploySchema",
			Handler:    _Service_DeploySchema_Handler,
		},
		{
			MethodName: "ActivateSchema",
			Handler:    _Service_ActivateSchema_Handler,
		},
		{
			MethodName: "ConfirmSchema",
			Handler:    _Service_ConfirmSchema_Handler,
		},
		{
			MethodName: "RollbackSchema",
			Handler:    _Service_RollbackSchema_Handler,
		},
		{
			MethodName: "RollbackSchemaByDeployment",
			Handler:    _Service_RollbackSchemaByDeployment_Handler,
		},
		{
			MethodName: "GetSchemaDeployment",
			Handler:    _Service_GetSchemaDeployment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "projections/management/service.proto",
}
