// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: gitconfig/gitconfig.proto

package gitconfig

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
	GitConfig_Show_FullMethodName                  = "/gitconfig.GitConfig/Show"
	GitConfig_GlobalConfig_FullMethodName          = "/gitconfig.GitConfig/GlobalConfig"
	GitConfig_ToggleGlobalStatus_FullMethodName    = "/gitconfig.GitConfig/ToggleGlobalStatus"
	GitConfig_Update_FullMethodName                = "/gitconfig.GitConfig/Update"
	GitConfig_GetDefaultChartValues_FullMethodName = "/gitconfig.GitConfig/GetDefaultChartValues"
)

// GitConfigClient is the client API for GitConfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GitConfigClient interface {
	Show(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowResponse, error)
	GlobalConfig(ctx context.Context, in *GlobalConfigRequest, opts ...grpc.CallOption) (*GlobalConfigResponse, error)
	ToggleGlobalStatus(ctx context.Context, in *ToggleGlobalStatusRequest, opts ...grpc.CallOption) (*ToggleGlobalStatusResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	GetDefaultChartValues(ctx context.Context, in *DefaultChartValuesRequest, opts ...grpc.CallOption) (*DefaultChartValuesResponse, error)
}

type gitConfigClient struct {
	cc grpc.ClientConnInterface
}

func NewGitConfigClient(cc grpc.ClientConnInterface) GitConfigClient {
	return &gitConfigClient{cc}
}

func (c *gitConfigClient) Show(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowResponse, error) {
	out := new(ShowResponse)
	err := c.cc.Invoke(ctx, GitConfig_Show_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitConfigClient) GlobalConfig(ctx context.Context, in *GlobalConfigRequest, opts ...grpc.CallOption) (*GlobalConfigResponse, error) {
	out := new(GlobalConfigResponse)
	err := c.cc.Invoke(ctx, GitConfig_GlobalConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitConfigClient) ToggleGlobalStatus(ctx context.Context, in *ToggleGlobalStatusRequest, opts ...grpc.CallOption) (*ToggleGlobalStatusResponse, error) {
	out := new(ToggleGlobalStatusResponse)
	err := c.cc.Invoke(ctx, GitConfig_ToggleGlobalStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitConfigClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, GitConfig_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitConfigClient) GetDefaultChartValues(ctx context.Context, in *DefaultChartValuesRequest, opts ...grpc.CallOption) (*DefaultChartValuesResponse, error) {
	out := new(DefaultChartValuesResponse)
	err := c.cc.Invoke(ctx, GitConfig_GetDefaultChartValues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GitConfigServer is the server API for GitConfig service.
// All implementations must embed UnimplementedGitConfigServer
// for forward compatibility
type GitConfigServer interface {
	Show(context.Context, *ShowRequest) (*ShowResponse, error)
	GlobalConfig(context.Context, *GlobalConfigRequest) (*GlobalConfigResponse, error)
	ToggleGlobalStatus(context.Context, *ToggleGlobalStatusRequest) (*ToggleGlobalStatusResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	GetDefaultChartValues(context.Context, *DefaultChartValuesRequest) (*DefaultChartValuesResponse, error)
	mustEmbedUnimplementedGitConfigServer()
}

// UnimplementedGitConfigServer must be embedded to have forward compatible implementations.
type UnimplementedGitConfigServer struct {
}

func (UnimplementedGitConfigServer) Show(context.Context, *ShowRequest) (*ShowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Show not implemented")
}
func (UnimplementedGitConfigServer) GlobalConfig(context.Context, *GlobalConfigRequest) (*GlobalConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GlobalConfig not implemented")
}
func (UnimplementedGitConfigServer) ToggleGlobalStatus(context.Context, *ToggleGlobalStatusRequest) (*ToggleGlobalStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToggleGlobalStatus not implemented")
}
func (UnimplementedGitConfigServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedGitConfigServer) GetDefaultChartValues(context.Context, *DefaultChartValuesRequest) (*DefaultChartValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefaultChartValues not implemented")
}
func (UnimplementedGitConfigServer) mustEmbedUnimplementedGitConfigServer() {}

// UnsafeGitConfigServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GitConfigServer will
// result in compilation errors.
type UnsafeGitConfigServer interface {
	mustEmbedUnimplementedGitConfigServer()
}

func RegisterGitConfigServer(s grpc.ServiceRegistrar, srv GitConfigServer) {
	s.RegisterService(&GitConfig_ServiceDesc, srv)
}

func _GitConfig_Show_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitConfigServer).Show(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitConfig_Show_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitConfigServer).Show(ctx, req.(*ShowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitConfig_GlobalConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GlobalConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitConfigServer).GlobalConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitConfig_GlobalConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitConfigServer).GlobalConfig(ctx, req.(*GlobalConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitConfig_ToggleGlobalStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToggleGlobalStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitConfigServer).ToggleGlobalStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitConfig_ToggleGlobalStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitConfigServer).ToggleGlobalStatus(ctx, req.(*ToggleGlobalStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitConfig_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitConfigServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitConfig_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitConfigServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitConfig_GetDefaultChartValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DefaultChartValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitConfigServer).GetDefaultChartValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitConfig_GetDefaultChartValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitConfigServer).GetDefaultChartValues(ctx, req.(*DefaultChartValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GitConfig_ServiceDesc is the grpc.ServiceDesc for GitConfig service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GitConfig_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gitconfig.GitConfig",
	HandlerType: (*GitConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Show",
			Handler:    _GitConfig_Show_Handler,
		},
		{
			MethodName: "GlobalConfig",
			Handler:    _GitConfig_GlobalConfig_Handler,
		},
		{
			MethodName: "ToggleGlobalStatus",
			Handler:    _GitConfig_ToggleGlobalStatus_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _GitConfig_Update_Handler,
		},
		{
			MethodName: "GetDefaultChartValues",
			Handler:    _GitConfig_GetDefaultChartValues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gitconfig/gitconfig.proto",
}
