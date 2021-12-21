// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package event

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

// EventClient is the client API for Event service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventClient interface {
	EnableProject(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventList, error)
}

type eventClient struct {
	cc grpc.ClientConnInterface
}

func NewEventClient(cc grpc.ClientConnInterface) EventClient {
	return &eventClient{cc}
}

func (c *eventClient) EnableProject(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventList, error) {
	out := new(EventList)
	err := c.cc.Invoke(ctx, "/Event/EnableProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServer is the server API for Event service.
// All implementations must embed UnimplementedEventServer
// for forward compatibility
type EventServer interface {
	EnableProject(context.Context, *EventRequest) (*EventList, error)
	mustEmbedUnimplementedEventServer()
}

// UnimplementedEventServer must be embedded to have forward compatible implementations.
type UnimplementedEventServer struct {
}

func (UnimplementedEventServer) EnableProject(context.Context, *EventRequest) (*EventList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableProject not implemented")
}
func (UnimplementedEventServer) mustEmbedUnimplementedEventServer() {}

// UnsafeEventServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventServer will
// result in compilation errors.
type UnsafeEventServer interface {
	mustEmbedUnimplementedEventServer()
}

func RegisterEventServer(s grpc.ServiceRegistrar, srv EventServer) {
	s.RegisterService(&Event_ServiceDesc, srv)
}

func _Event_EnableProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).EnableProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Event/EnableProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).EnableProject(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Event_ServiceDesc is the grpc.ServiceDesc for Event service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Event_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Event",
	HandlerType: (*EventServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnableProject",
			Handler:    _Event_EnableProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "event/event.proto",
}
