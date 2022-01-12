// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package container

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

// ContainerSvcClient is the client API for ContainerSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContainerSvcClient interface {
	CopyToPod(ctx context.Context, in *CopyToPodRequest, opts ...grpc.CallOption) (*CopyToPodResponse, error)
	Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (ContainerSvc_ExecClient, error)
}

type containerSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewContainerSvcClient(cc grpc.ClientConnInterface) ContainerSvcClient {
	return &containerSvcClient{cc}
}

func (c *containerSvcClient) CopyToPod(ctx context.Context, in *CopyToPodRequest, opts ...grpc.CallOption) (*CopyToPodResponse, error) {
	out := new(CopyToPodResponse)
	err := c.cc.Invoke(ctx, "/ContainerSvc/CopyToPod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerSvcClient) Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (ContainerSvc_ExecClient, error) {
	stream, err := c.cc.NewStream(ctx, &ContainerSvc_ServiceDesc.Streams[0], "/ContainerSvc/Exec", opts...)
	if err != nil {
		return nil, err
	}
	x := &containerSvcExecClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ContainerSvc_ExecClient interface {
	Recv() (*ExecResponse, error)
	grpc.ClientStream
}

type containerSvcExecClient struct {
	grpc.ClientStream
}

func (x *containerSvcExecClient) Recv() (*ExecResponse, error) {
	m := new(ExecResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ContainerSvcServer is the server API for ContainerSvc service.
// All implementations must embed UnimplementedContainerSvcServer
// for forward compatibility
type ContainerSvcServer interface {
	CopyToPod(context.Context, *CopyToPodRequest) (*CopyToPodResponse, error)
	Exec(*ExecRequest, ContainerSvc_ExecServer) error
	mustEmbedUnimplementedContainerSvcServer()
}

// UnimplementedContainerSvcServer must be embedded to have forward compatible implementations.
type UnimplementedContainerSvcServer struct {
}

func (UnimplementedContainerSvcServer) CopyToPod(context.Context, *CopyToPodRequest) (*CopyToPodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CopyToPod not implemented")
}
func (UnimplementedContainerSvcServer) Exec(*ExecRequest, ContainerSvc_ExecServer) error {
	return status.Errorf(codes.Unimplemented, "method Exec not implemented")
}
func (UnimplementedContainerSvcServer) mustEmbedUnimplementedContainerSvcServer() {}

// UnsafeContainerSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContainerSvcServer will
// result in compilation errors.
type UnsafeContainerSvcServer interface {
	mustEmbedUnimplementedContainerSvcServer()
}

func RegisterContainerSvcServer(s grpc.ServiceRegistrar, srv ContainerSvcServer) {
	s.RegisterService(&ContainerSvc_ServiceDesc, srv)
}

func _ContainerSvc_CopyToPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CopyToPodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerSvcServer).CopyToPod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContainerSvc/CopyToPod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerSvcServer).CopyToPod(ctx, req.(*CopyToPodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerSvc_Exec_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExecRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ContainerSvcServer).Exec(m, &containerSvcExecServer{stream})
}

type ContainerSvc_ExecServer interface {
	Send(*ExecResponse) error
	grpc.ServerStream
}

type containerSvcExecServer struct {
	grpc.ServerStream
}

func (x *containerSvcExecServer) Send(m *ExecResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ContainerSvc_ServiceDesc is the grpc.ServiceDesc for ContainerSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContainerSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ContainerSvc",
	HandlerType: (*ContainerSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CopyToPod",
			Handler:    _ContainerSvc_CopyToPod_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Exec",
			Handler:       _ContainerSvc_Exec_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "container/container.proto",
}
