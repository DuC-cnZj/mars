// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: endpoint/endpoint.proto

package endpoint

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

// EndpointClient is the client API for Endpoint service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EndpointClient interface {
	// InNamespace 名称空间下所有的 endpoints
	InNamespace(ctx context.Context, in *InNamespaceRequest, opts ...grpc.CallOption) (*InNamespaceResponse, error)
	// InProject 项目下所有的 endpoints
	InProject(ctx context.Context, in *InProjectRequest, opts ...grpc.CallOption) (*InProjectResponse, error)
}

type endpointClient struct {
	cc grpc.ClientConnInterface
}

func NewEndpointClient(cc grpc.ClientConnInterface) EndpointClient {
	return &endpointClient{cc}
}

func (c *endpointClient) InNamespace(ctx context.Context, in *InNamespaceRequest, opts ...grpc.CallOption) (*InNamespaceResponse, error) {
	out := new(InNamespaceResponse)
	err := c.cc.Invoke(ctx, "/endpoint.Endpoint/InNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endpointClient) InProject(ctx context.Context, in *InProjectRequest, opts ...grpc.CallOption) (*InProjectResponse, error) {
	out := new(InProjectResponse)
	err := c.cc.Invoke(ctx, "/endpoint.Endpoint/InProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndpointServer is the server API for Endpoint service.
// All implementations must embed UnimplementedEndpointServer
// for forward compatibility
type EndpointServer interface {
	// InNamespace 名称空间下所有的 endpoints
	InNamespace(context.Context, *InNamespaceRequest) (*InNamespaceResponse, error)
	// InProject 项目下所有的 endpoints
	InProject(context.Context, *InProjectRequest) (*InProjectResponse, error)
	mustEmbedUnimplementedEndpointServer()
}

// UnimplementedEndpointServer must be embedded to have forward compatible implementations.
type UnimplementedEndpointServer struct {
}

func (UnimplementedEndpointServer) InNamespace(context.Context, *InNamespaceRequest) (*InNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InNamespace not implemented")
}
func (UnimplementedEndpointServer) InProject(context.Context, *InProjectRequest) (*InProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InProject not implemented")
}
func (UnimplementedEndpointServer) mustEmbedUnimplementedEndpointServer() {}

// UnsafeEndpointServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EndpointServer will
// result in compilation errors.
type UnsafeEndpointServer interface {
	mustEmbedUnimplementedEndpointServer()
}

func RegisterEndpointServer(s grpc.ServiceRegistrar, srv EndpointServer) {
	s.RegisterService(&Endpoint_ServiceDesc, srv)
}

func _Endpoint_InNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointServer).InNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoint.Endpoint/InNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointServer).InNamespace(ctx, req.(*InNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Endpoint_InProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointServer).InProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoint.Endpoint/InProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointServer).InProject(ctx, req.(*InProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Endpoint_ServiceDesc is the grpc.ServiceDesc for Endpoint service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Endpoint_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "endpoint.Endpoint",
	HandlerType: (*EndpointServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InNamespace",
			Handler:    _Endpoint_InNamespace_Handler,
		},
		{
			MethodName: "InProject",
			Handler:    _Endpoint_InProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "endpoint/endpoint.proto",
}
