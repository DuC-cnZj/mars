// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package mars

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

// MarsClient is the client API for Mars service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarsClient interface {
	Show(ctx context.Context, in *MarsShowRequest, opts ...grpc.CallOption) (*MarsShowResponse, error)
	GlobalConfig(ctx context.Context, in *MarsGlobalConfigRequest, opts ...grpc.CallOption) (*MarsGlobalConfigResponse, error)
	ToggleEnabled(ctx context.Context, in *MarsToggleEnabledRequest, opts ...grpc.CallOption) (*MarsToggleEnabledResponse, error)
	Update(ctx context.Context, in *MarsUpdateRequest, opts ...grpc.CallOption) (*MarsUpdateResponse, error)
	GetDefaultChartValues(ctx context.Context, in *MarsDefaultChartValuesRequest, opts ...grpc.CallOption) (*MarsDefaultChartValuesResponse, error)
}

type marsClient struct {
	cc grpc.ClientConnInterface
}

func NewMarsClient(cc grpc.ClientConnInterface) MarsClient {
	return &marsClient{cc}
}

func (c *marsClient) Show(ctx context.Context, in *MarsShowRequest, opts ...grpc.CallOption) (*MarsShowResponse, error) {
	out := new(MarsShowResponse)
	err := c.cc.Invoke(ctx, "/Mars/Show", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marsClient) GlobalConfig(ctx context.Context, in *MarsGlobalConfigRequest, opts ...grpc.CallOption) (*MarsGlobalConfigResponse, error) {
	out := new(MarsGlobalConfigResponse)
	err := c.cc.Invoke(ctx, "/Mars/GlobalConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marsClient) ToggleEnabled(ctx context.Context, in *MarsToggleEnabledRequest, opts ...grpc.CallOption) (*MarsToggleEnabledResponse, error) {
	out := new(MarsToggleEnabledResponse)
	err := c.cc.Invoke(ctx, "/Mars/ToggleEnabled", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marsClient) Update(ctx context.Context, in *MarsUpdateRequest, opts ...grpc.CallOption) (*MarsUpdateResponse, error) {
	out := new(MarsUpdateResponse)
	err := c.cc.Invoke(ctx, "/Mars/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marsClient) GetDefaultChartValues(ctx context.Context, in *MarsDefaultChartValuesRequest, opts ...grpc.CallOption) (*MarsDefaultChartValuesResponse, error) {
	out := new(MarsDefaultChartValuesResponse)
	err := c.cc.Invoke(ctx, "/Mars/GetDefaultChartValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarsServer is the server API for Mars service.
// All implementations must embed UnimplementedMarsServer
// for forward compatibility
type MarsServer interface {
	Show(context.Context, *MarsShowRequest) (*MarsShowResponse, error)
	GlobalConfig(context.Context, *MarsGlobalConfigRequest) (*MarsGlobalConfigResponse, error)
	ToggleEnabled(context.Context, *MarsToggleEnabledRequest) (*MarsToggleEnabledResponse, error)
	Update(context.Context, *MarsUpdateRequest) (*MarsUpdateResponse, error)
	GetDefaultChartValues(context.Context, *MarsDefaultChartValuesRequest) (*MarsDefaultChartValuesResponse, error)
	mustEmbedUnimplementedMarsServer()
}

// UnimplementedMarsServer must be embedded to have forward compatible implementations.
type UnimplementedMarsServer struct {
}

func (UnimplementedMarsServer) Show(context.Context, *MarsShowRequest) (*MarsShowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Show not implemented")
}
func (UnimplementedMarsServer) GlobalConfig(context.Context, *MarsGlobalConfigRequest) (*MarsGlobalConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GlobalConfig not implemented")
}
func (UnimplementedMarsServer) ToggleEnabled(context.Context, *MarsToggleEnabledRequest) (*MarsToggleEnabledResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToggleEnabled not implemented")
}
func (UnimplementedMarsServer) Update(context.Context, *MarsUpdateRequest) (*MarsUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMarsServer) GetDefaultChartValues(context.Context, *MarsDefaultChartValuesRequest) (*MarsDefaultChartValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefaultChartValues not implemented")
}
func (UnimplementedMarsServer) mustEmbedUnimplementedMarsServer() {}

// UnsafeMarsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarsServer will
// result in compilation errors.
type UnsafeMarsServer interface {
	mustEmbedUnimplementedMarsServer()
}

func RegisterMarsServer(s grpc.ServiceRegistrar, srv MarsServer) {
	s.RegisterService(&Mars_ServiceDesc, srv)
}

func _Mars_Show_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarsShowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarsServer).Show(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Mars/Show",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarsServer).Show(ctx, req.(*MarsShowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mars_GlobalConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarsGlobalConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarsServer).GlobalConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Mars/GlobalConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarsServer).GlobalConfig(ctx, req.(*MarsGlobalConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mars_ToggleEnabled_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarsToggleEnabledRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarsServer).ToggleEnabled(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Mars/ToggleEnabled",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarsServer).ToggleEnabled(ctx, req.(*MarsToggleEnabledRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mars_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarsUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Mars/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarsServer).Update(ctx, req.(*MarsUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mars_GetDefaultChartValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarsDefaultChartValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarsServer).GetDefaultChartValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Mars/GetDefaultChartValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarsServer).GetDefaultChartValues(ctx, req.(*MarsDefaultChartValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Mars_ServiceDesc is the grpc.ServiceDesc for Mars service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mars_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Mars",
	HandlerType: (*MarsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Show",
			Handler:    _Mars_Show_Handler,
		},
		{
			MethodName: "GlobalConfig",
			Handler:    _Mars_GlobalConfig_Handler,
		},
		{
			MethodName: "ToggleEnabled",
			Handler:    _Mars_ToggleEnabled_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Mars_Update_Handler,
		},
		{
			MethodName: "GetDefaultChartValues",
			Handler:    _Mars_GetDefaultChartValues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mars/mars.proto",
}
