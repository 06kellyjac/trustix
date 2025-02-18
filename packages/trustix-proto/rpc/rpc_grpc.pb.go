// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rpc

import (
	context "context"
	api "github.com/nix-community/trustix/packages/trustix-proto/api"
	schema "github.com/nix-community/trustix/packages/trustix-proto/schema"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RPCApiClient is the client API for RPCApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCApiClient interface {
	// Get a list of all logs published/subscribed by this node
	Logs(ctx context.Context, in *api.LogsRequest, opts ...grpc.CallOption) (*api.LogsResponse, error)
	// Decide on an output for key based on the configured decision method
	Decide(ctx context.Context, in *DecideRequest, opts ...grpc.CallOption) (*DecisionResponse, error)
	// Get values by their content-address
	GetValue(ctx context.Context, in *api.ValueRequest, opts ...grpc.CallOption) (*api.ValueResponse, error)
}

type rPCApiClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCApiClient(cc grpc.ClientConnInterface) RPCApiClient {
	return &rPCApiClient{cc}
}

func (c *rPCApiClient) Logs(ctx context.Context, in *api.LogsRequest, opts ...grpc.CallOption) (*api.LogsResponse, error) {
	out := new(api.LogsResponse)
	err := c.cc.Invoke(ctx, "/trustix.RPCApi/Logs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCApiClient) Decide(ctx context.Context, in *DecideRequest, opts ...grpc.CallOption) (*DecisionResponse, error) {
	out := new(DecisionResponse)
	err := c.cc.Invoke(ctx, "/trustix.RPCApi/Decide", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCApiClient) GetValue(ctx context.Context, in *api.ValueRequest, opts ...grpc.CallOption) (*api.ValueResponse, error) {
	out := new(api.ValueResponse)
	err := c.cc.Invoke(ctx, "/trustix.RPCApi/GetValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCApiServer is the server API for RPCApi service.
// All implementations must embed UnimplementedRPCApiServer
// for forward compatibility
type RPCApiServer interface {
	// Get a list of all logs published/subscribed by this node
	Logs(context.Context, *api.LogsRequest) (*api.LogsResponse, error)
	// Decide on an output for key based on the configured decision method
	Decide(context.Context, *DecideRequest) (*DecisionResponse, error)
	// Get values by their content-address
	GetValue(context.Context, *api.ValueRequest) (*api.ValueResponse, error)
	mustEmbedUnimplementedRPCApiServer()
}

// UnimplementedRPCApiServer must be embedded to have forward compatible implementations.
type UnimplementedRPCApiServer struct {
}

func (UnimplementedRPCApiServer) Logs(context.Context, *api.LogsRequest) (*api.LogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logs not implemented")
}
func (UnimplementedRPCApiServer) Decide(context.Context, *DecideRequest) (*DecisionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decide not implemented")
}
func (UnimplementedRPCApiServer) GetValue(context.Context, *api.ValueRequest) (*api.ValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValue not implemented")
}
func (UnimplementedRPCApiServer) mustEmbedUnimplementedRPCApiServer() {}

// UnsafeRPCApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCApiServer will
// result in compilation errors.
type UnsafeRPCApiServer interface {
	mustEmbedUnimplementedRPCApiServer()
}

func RegisterRPCApiServer(s grpc.ServiceRegistrar, srv RPCApiServer) {
	s.RegisterService(&RPCApi_ServiceDesc, srv)
}

func _RPCApi_Logs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.LogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCApiServer).Logs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.RPCApi/Logs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCApiServer).Logs(ctx, req.(*api.LogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCApi_Decide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCApiServer).Decide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.RPCApi/Decide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCApiServer).Decide(ctx, req.(*DecideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCApi_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.ValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCApiServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.RPCApi/GetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCApiServer).GetValue(ctx, req.(*api.ValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPCApi_ServiceDesc is the grpc.ServiceDesc for RPCApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPCApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "trustix.RPCApi",
	HandlerType: (*RPCApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Logs",
			Handler:    _RPCApi_Logs_Handler,
		},
		{
			MethodName: "Decide",
			Handler:    _RPCApi_Decide_Handler,
		},
		{
			MethodName: "GetValue",
			Handler:    _RPCApi_GetValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

// LogRPCClient is the client API for LogRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogRPCClient interface {
	GetHead(ctx context.Context, in *api.LogHeadRequest, opts ...grpc.CallOption) (*schema.LogHead, error)
	GetLogEntries(ctx context.Context, in *api.GetLogEntriesRequest, opts ...grpc.CallOption) (*api.LogEntriesResponse, error)
	Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error)
	Flush(ctx context.Context, in *FlushRequest, opts ...grpc.CallOption) (*FlushResponse, error)
}

type logRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewLogRPCClient(cc grpc.ClientConnInterface) LogRPCClient {
	return &logRPCClient{cc}
}

func (c *logRPCClient) GetHead(ctx context.Context, in *api.LogHeadRequest, opts ...grpc.CallOption) (*schema.LogHead, error) {
	out := new(schema.LogHead)
	err := c.cc.Invoke(ctx, "/trustix.LogRPC/GetHead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logRPCClient) GetLogEntries(ctx context.Context, in *api.GetLogEntriesRequest, opts ...grpc.CallOption) (*api.LogEntriesResponse, error) {
	out := new(api.LogEntriesResponse)
	err := c.cc.Invoke(ctx, "/trustix.LogRPC/GetLogEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logRPCClient) Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error) {
	out := new(SubmitResponse)
	err := c.cc.Invoke(ctx, "/trustix.LogRPC/Submit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logRPCClient) Flush(ctx context.Context, in *FlushRequest, opts ...grpc.CallOption) (*FlushResponse, error) {
	out := new(FlushResponse)
	err := c.cc.Invoke(ctx, "/trustix.LogRPC/Flush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogRPCServer is the server API for LogRPC service.
// All implementations must embed UnimplementedLogRPCServer
// for forward compatibility
type LogRPCServer interface {
	GetHead(context.Context, *api.LogHeadRequest) (*schema.LogHead, error)
	GetLogEntries(context.Context, *api.GetLogEntriesRequest) (*api.LogEntriesResponse, error)
	Submit(context.Context, *SubmitRequest) (*SubmitResponse, error)
	Flush(context.Context, *FlushRequest) (*FlushResponse, error)
	mustEmbedUnimplementedLogRPCServer()
}

// UnimplementedLogRPCServer must be embedded to have forward compatible implementations.
type UnimplementedLogRPCServer struct {
}

func (UnimplementedLogRPCServer) GetHead(context.Context, *api.LogHeadRequest) (*schema.LogHead, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHead not implemented")
}
func (UnimplementedLogRPCServer) GetLogEntries(context.Context, *api.GetLogEntriesRequest) (*api.LogEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogEntries not implemented")
}
func (UnimplementedLogRPCServer) Submit(context.Context, *SubmitRequest) (*SubmitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}
func (UnimplementedLogRPCServer) Flush(context.Context, *FlushRequest) (*FlushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Flush not implemented")
}
func (UnimplementedLogRPCServer) mustEmbedUnimplementedLogRPCServer() {}

// UnsafeLogRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogRPCServer will
// result in compilation errors.
type UnsafeLogRPCServer interface {
	mustEmbedUnimplementedLogRPCServer()
}

func RegisterLogRPCServer(s grpc.ServiceRegistrar, srv LogRPCServer) {
	s.RegisterService(&LogRPC_ServiceDesc, srv)
}

func _LogRPC_GetHead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.LogHeadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogRPCServer).GetHead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.LogRPC/GetHead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogRPCServer).GetHead(ctx, req.(*api.LogHeadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogRPC_GetLogEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.GetLogEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogRPCServer).GetLogEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.LogRPC/GetLogEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogRPCServer).GetLogEntries(ctx, req.(*api.GetLogEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogRPC_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogRPCServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.LogRPC/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogRPCServer).Submit(ctx, req.(*SubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogRPC_Flush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogRPCServer).Flush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.LogRPC/Flush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogRPCServer).Flush(ctx, req.(*FlushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogRPC_ServiceDesc is the grpc.ServiceDesc for LogRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "trustix.LogRPC",
	HandlerType: (*LogRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHead",
			Handler:    _LogRPC_GetHead_Handler,
		},
		{
			MethodName: "GetLogEntries",
			Handler:    _LogRPC_GetLogEntries_Handler,
		},
		{
			MethodName: "Submit",
			Handler:    _LogRPC_Submit_Handler,
		},
		{
			MethodName: "Flush",
			Handler:    _LogRPC_Flush_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}
