// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package metrics

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

// MetricsClient is the client API for Metrics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricsClient interface {
	//  Show 获取 pod 的 cpu memory 信息
	Show(ctx context.Context, in *MetricsShowRequest, opts ...grpc.CallOption) (*MetricsShowResponse, error)
	//  StreamShow stream 的方式获取 pod 的 cpu memory 信息
	StreamShow(ctx context.Context, in *MetricsShowRequest, opts ...grpc.CallOption) (Metrics_StreamShowClient, error)
}

type metricsClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsClient(cc grpc.ClientConnInterface) MetricsClient {
	return &metricsClient{cc}
}

func (c *metricsClient) Show(ctx context.Context, in *MetricsShowRequest, opts ...grpc.CallOption) (*MetricsShowResponse, error) {
	out := new(MetricsShowResponse)
	err := c.cc.Invoke(ctx, "/Metrics/Show", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsClient) StreamShow(ctx context.Context, in *MetricsShowRequest, opts ...grpc.CallOption) (Metrics_StreamShowClient, error) {
	stream, err := c.cc.NewStream(ctx, &Metrics_ServiceDesc.Streams[0], "/Metrics/StreamShow", opts...)
	if err != nil {
		return nil, err
	}
	x := &metricsStreamShowClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Metrics_StreamShowClient interface {
	Recv() (*MetricsShowResponse, error)
	grpc.ClientStream
}

type metricsStreamShowClient struct {
	grpc.ClientStream
}

func (x *metricsStreamShowClient) Recv() (*MetricsShowResponse, error) {
	m := new(MetricsShowResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MetricsServer is the server API for Metrics service.
// All implementations must embed UnimplementedMetricsServer
// for forward compatibility
type MetricsServer interface {
	//  Show 获取 pod 的 cpu memory 信息
	Show(context.Context, *MetricsShowRequest) (*MetricsShowResponse, error)
	//  StreamShow stream 的方式获取 pod 的 cpu memory 信息
	StreamShow(*MetricsShowRequest, Metrics_StreamShowServer) error
	mustEmbedUnimplementedMetricsServer()
}

// UnimplementedMetricsServer must be embedded to have forward compatible implementations.
type UnimplementedMetricsServer struct {
}

func (UnimplementedMetricsServer) Show(context.Context, *MetricsShowRequest) (*MetricsShowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Show not implemented")
}
func (UnimplementedMetricsServer) StreamShow(*MetricsShowRequest, Metrics_StreamShowServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamShow not implemented")
}
func (UnimplementedMetricsServer) mustEmbedUnimplementedMetricsServer() {}

// UnsafeMetricsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricsServer will
// result in compilation errors.
type UnsafeMetricsServer interface {
	mustEmbedUnimplementedMetricsServer()
}

func RegisterMetricsServer(s grpc.ServiceRegistrar, srv MetricsServer) {
	s.RegisterService(&Metrics_ServiceDesc, srv)
}

func _Metrics_Show_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricsShowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServer).Show(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Metrics/Show",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServer).Show(ctx, req.(*MetricsShowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metrics_StreamShow_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MetricsShowRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MetricsServer).StreamShow(m, &metricsStreamShowServer{stream})
}

type Metrics_StreamShowServer interface {
	Send(*MetricsShowResponse) error
	grpc.ServerStream
}

type metricsStreamShowServer struct {
	grpc.ServerStream
}

func (x *metricsStreamShowServer) Send(m *MetricsShowResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Metrics_ServiceDesc is the grpc.ServiceDesc for Metrics service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Metrics_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Metrics",
	HandlerType: (*MetricsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Show",
			Handler:    _Metrics_Show_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamShow",
			Handler:       _Metrics_StreamShow_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "metrics/metrics.proto",
}
