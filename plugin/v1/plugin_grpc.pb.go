// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: plugin/v1/plugin.proto

package pluginv1

import (
	context "context"
	messages "github.com/tomvodi/limepipes-plugin-api/plugin/v1/messages"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PluginService_PluginInfo_FullMethodName = "/plugin.v1.PluginService/PluginInfo"
	PluginService_ImportFile_FullMethodName = "/plugin.v1.PluginService/ImportFile"
)

// PluginServiceClient is the client API for PluginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PluginServiceClient interface {
	PluginInfo(ctx context.Context, in *messages.PluginInfoRequest, opts ...grpc.CallOption) (*messages.PluginInfoResponse, error)
	// ImportFile imports a file into the plugin
	// When the filetype is not supported by the plugin it returns an UNIMPLEMENTED error
	ImportFile(ctx context.Context, in *messages.ImportFileRequest, opts ...grpc.CallOption) (*messages.ImportFileResponse, error)
}

type pluginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginServiceClient(cc grpc.ClientConnInterface) PluginServiceClient {
	return &pluginServiceClient{cc}
}

func (c *pluginServiceClient) PluginInfo(ctx context.Context, in *messages.PluginInfoRequest, opts ...grpc.CallOption) (*messages.PluginInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(messages.PluginInfoResponse)
	err := c.cc.Invoke(ctx, PluginService_PluginInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginServiceClient) ImportFile(ctx context.Context, in *messages.ImportFileRequest, opts ...grpc.CallOption) (*messages.ImportFileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(messages.ImportFileResponse)
	err := c.cc.Invoke(ctx, PluginService_ImportFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginServiceServer is the server API for PluginService service.
// All implementations must embed UnimplementedPluginServiceServer
// for forward compatibility.
type PluginServiceServer interface {
	PluginInfo(context.Context, *messages.PluginInfoRequest) (*messages.PluginInfoResponse, error)
	// ImportFile imports a file into the plugin
	// When the filetype is not supported by the plugin it returns an UNIMPLEMENTED error
	ImportFile(context.Context, *messages.ImportFileRequest) (*messages.ImportFileResponse, error)
	mustEmbedUnimplementedPluginServiceServer()
}

// UnimplementedPluginServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPluginServiceServer struct{}

func (UnimplementedPluginServiceServer) PluginInfo(context.Context, *messages.PluginInfoRequest) (*messages.PluginInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PluginInfo not implemented")
}
func (UnimplementedPluginServiceServer) ImportFile(context.Context, *messages.ImportFileRequest) (*messages.ImportFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportFile not implemented")
}
func (UnimplementedPluginServiceServer) mustEmbedUnimplementedPluginServiceServer() {}
func (UnimplementedPluginServiceServer) testEmbeddedByValue()                       {}

// UnsafePluginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PluginServiceServer will
// result in compilation errors.
type UnsafePluginServiceServer interface {
	mustEmbedUnimplementedPluginServiceServer()
}

func RegisterPluginServiceServer(s grpc.ServiceRegistrar, srv PluginServiceServer) {
	// If the following call pancis, it indicates UnimplementedPluginServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PluginService_ServiceDesc, srv)
}

func _PluginService_PluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(messages.PluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServiceServer).PluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PluginService_PluginInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServiceServer).PluginInfo(ctx, req.(*messages.PluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginService_ImportFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(messages.ImportFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServiceServer).ImportFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PluginService_ImportFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServiceServer).ImportFile(ctx, req.(*messages.ImportFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PluginService_ServiceDesc is the grpc.ServiceDesc for PluginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PluginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "plugin.v1.PluginService",
	HandlerType: (*PluginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PluginInfo",
			Handler:    _PluginService_PluginInfo_Handler,
		},
		{
			MethodName: "ImportFile",
			Handler:    _PluginService_ImportFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugin/v1/plugin.proto",
}
