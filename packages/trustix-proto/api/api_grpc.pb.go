// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	schema "github.com/nix-community/trustix/packages/trustix-proto/schema"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NodeAPIClient is the client API for NodeAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeAPIClient interface {
	// Get a list of all logs published by this node
	Logs(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (*LogsResponse, error)
	// Get values by their content-address
	GetValue(ctx context.Context, in *ValueRequest, opts ...grpc.CallOption) (*ValueResponse, error)
}

type nodeAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeAPIClient(cc grpc.ClientConnInterface) NodeAPIClient {
	return &nodeAPIClient{cc}
}

func (c *nodeAPIClient) Logs(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (*LogsResponse, error) {
	out := new(LogsResponse)
	err := c.cc.Invoke(ctx, "/trustix.NodeAPI/Logs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeAPIClient) GetValue(ctx context.Context, in *ValueRequest, opts ...grpc.CallOption) (*ValueResponse, error) {
	out := new(ValueResponse)
	err := c.cc.Invoke(ctx, "/trustix.NodeAPI/GetValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeAPIServer is the server API for NodeAPI service.
// All implementations must embed UnimplementedNodeAPIServer
// for forward compatibility
type NodeAPIServer interface {
	// Get a list of all logs published by this node
	Logs(context.Context, *LogsRequest) (*LogsResponse, error)
	// Get values by their content-address
	GetValue(context.Context, *ValueRequest) (*ValueResponse, error)
	mustEmbedUnimplementedNodeAPIServer()
}

// UnimplementedNodeAPIServer must be embedded to have forward compatible implementations.
type UnimplementedNodeAPIServer struct {
}

func (UnimplementedNodeAPIServer) Logs(context.Context, *LogsRequest) (*LogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logs not implemented")
}
func (UnimplementedNodeAPIServer) GetValue(context.Context, *ValueRequest) (*ValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValue not implemented")
}
func (UnimplementedNodeAPIServer) mustEmbedUnimplementedNodeAPIServer() {}

// UnsafeNodeAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeAPIServer will
// result in compilation errors.
type UnsafeNodeAPIServer interface {
	mustEmbedUnimplementedNodeAPIServer()
}

func RegisterNodeAPIServer(s grpc.ServiceRegistrar, srv NodeAPIServer) {
	s.RegisterService(&NodeAPI_ServiceDesc, srv)
}

func _NodeAPI_Logs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAPIServer).Logs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.NodeAPI/Logs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAPIServer).Logs(ctx, req.(*LogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeAPI_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAPIServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.NodeAPI/GetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAPIServer).GetValue(ctx, req.(*ValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeAPI_ServiceDesc is the grpc.ServiceDesc for NodeAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "trustix.NodeAPI",
	HandlerType: (*NodeAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Logs",
			Handler:    _NodeAPI_Logs_Handler,
		},
		{
			MethodName: "GetValue",
			Handler:    _NodeAPI_GetValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// LogAPIClient is the client API for LogAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogAPIClient interface {
	// Get signed head
	GetHead(ctx context.Context, in *LogHeadRequest, opts ...grpc.CallOption) (*schema.LogHead, error)
	GetLogConsistencyProof(ctx context.Context, in *GetLogConsistencyProofRequest, opts ...grpc.CallOption) (*ProofResponse, error)
	GetLogAuditProof(ctx context.Context, in *GetLogAuditProofRequest, opts ...grpc.CallOption) (*ProofResponse, error)
	GetLogEntries(ctx context.Context, in *GetLogEntriesRequest, opts ...grpc.CallOption) (*LogEntriesResponse, error)
	GetMapValue(ctx context.Context, in *GetMapValueRequest, opts ...grpc.CallOption) (*MapValueResponse, error)
	GetMHLogConsistencyProof(ctx context.Context, in *GetLogConsistencyProofRequest, opts ...grpc.CallOption) (*ProofResponse, error)
	GetMHLogAuditProof(ctx context.Context, in *GetLogAuditProofRequest, opts ...grpc.CallOption) (*ProofResponse, error)
	GetMHLogEntries(ctx context.Context, in *GetLogEntriesRequest, opts ...grpc.CallOption) (*LogEntriesResponse, error)
}

type logAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewLogAPIClient(cc grpc.ClientConnInterface) LogAPIClient {
	return &logAPIClient{cc}
}

func (c *logAPIClient) GetHead(ctx context.Context, in *LogHeadRequest, opts ...grpc.CallOption) (*schema.LogHead, error) {
	out := new(schema.LogHead)
	err := c.cc.Invoke(ctx, "/trustix.LogAPI/GetHead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logAPIClient) GetLogConsistencyProof(ctx context.Context, in *GetLogConsistencyProofRequest, opts ...grpc.CallOption) (*ProofResponse, error) {
	out := new(ProofResponse)
	err := c.cc.Invoke(ctx, "/trustix.LogAPI/GetLogConsistencyProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logAPIClient) GetLogAuditProof(ctx context.Context, in *GetLogAuditProofRequest, opts ...grpc.CallOption) (*ProofResponse, error) {
	out := new(ProofResponse)
	err := c.cc.Invoke(ctx, "/trustix.LogAPI/GetLogAuditProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logAPIClient) GetLogEntries(ctx context.Context, in *GetLogEntriesRequest, opts ...grpc.CallOption) (*LogEntriesResponse, error) {
	out := new(LogEntriesResponse)
	err := c.cc.Invoke(ctx, "/trustix.LogAPI/GetLogEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logAPIClient) GetMapValue(ctx context.Context, in *GetMapValueRequest, opts ...grpc.CallOption) (*MapValueResponse, error) {
	out := new(MapValueResponse)
	err := c.cc.Invoke(ctx, "/trustix.LogAPI/GetMapValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logAPIClient) GetMHLogConsistencyProof(ctx context.Context, in *GetLogConsistencyProofRequest, opts ...grpc.CallOption) (*ProofResponse, error) {
	out := new(ProofResponse)
	err := c.cc.Invoke(ctx, "/trustix.LogAPI/GetMHLogConsistencyProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logAPIClient) GetMHLogAuditProof(ctx context.Context, in *GetLogAuditProofRequest, opts ...grpc.CallOption) (*ProofResponse, error) {
	out := new(ProofResponse)
	err := c.cc.Invoke(ctx, "/trustix.LogAPI/GetMHLogAuditProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logAPIClient) GetMHLogEntries(ctx context.Context, in *GetLogEntriesRequest, opts ...grpc.CallOption) (*LogEntriesResponse, error) {
	out := new(LogEntriesResponse)
	err := c.cc.Invoke(ctx, "/trustix.LogAPI/GetMHLogEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogAPIServer is the server API for LogAPI service.
// All implementations must embed UnimplementedLogAPIServer
// for forward compatibility
type LogAPIServer interface {
	// Get signed head
	GetHead(context.Context, *LogHeadRequest) (*schema.LogHead, error)
	GetLogConsistencyProof(context.Context, *GetLogConsistencyProofRequest) (*ProofResponse, error)
	GetLogAuditProof(context.Context, *GetLogAuditProofRequest) (*ProofResponse, error)
	GetLogEntries(context.Context, *GetLogEntriesRequest) (*LogEntriesResponse, error)
	GetMapValue(context.Context, *GetMapValueRequest) (*MapValueResponse, error)
	GetMHLogConsistencyProof(context.Context, *GetLogConsistencyProofRequest) (*ProofResponse, error)
	GetMHLogAuditProof(context.Context, *GetLogAuditProofRequest) (*ProofResponse, error)
	GetMHLogEntries(context.Context, *GetLogEntriesRequest) (*LogEntriesResponse, error)
	mustEmbedUnimplementedLogAPIServer()
}

// UnimplementedLogAPIServer must be embedded to have forward compatible implementations.
type UnimplementedLogAPIServer struct {
}

func (UnimplementedLogAPIServer) GetHead(context.Context, *LogHeadRequest) (*schema.LogHead, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHead not implemented")
}
func (UnimplementedLogAPIServer) GetLogConsistencyProof(context.Context, *GetLogConsistencyProofRequest) (*ProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogConsistencyProof not implemented")
}
func (UnimplementedLogAPIServer) GetLogAuditProof(context.Context, *GetLogAuditProofRequest) (*ProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogAuditProof not implemented")
}
func (UnimplementedLogAPIServer) GetLogEntries(context.Context, *GetLogEntriesRequest) (*LogEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogEntries not implemented")
}
func (UnimplementedLogAPIServer) GetMapValue(context.Context, *GetMapValueRequest) (*MapValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMapValue not implemented")
}
func (UnimplementedLogAPIServer) GetMHLogConsistencyProof(context.Context, *GetLogConsistencyProofRequest) (*ProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMHLogConsistencyProof not implemented")
}
func (UnimplementedLogAPIServer) GetMHLogAuditProof(context.Context, *GetLogAuditProofRequest) (*ProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMHLogAuditProof not implemented")
}
func (UnimplementedLogAPIServer) GetMHLogEntries(context.Context, *GetLogEntriesRequest) (*LogEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMHLogEntries not implemented")
}
func (UnimplementedLogAPIServer) mustEmbedUnimplementedLogAPIServer() {}

// UnsafeLogAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogAPIServer will
// result in compilation errors.
type UnsafeLogAPIServer interface {
	mustEmbedUnimplementedLogAPIServer()
}

func RegisterLogAPIServer(s grpc.ServiceRegistrar, srv LogAPIServer) {
	s.RegisterService(&LogAPI_ServiceDesc, srv)
}

func _LogAPI_GetHead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogHeadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogAPIServer).GetHead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.LogAPI/GetHead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogAPIServer).GetHead(ctx, req.(*LogHeadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogAPI_GetLogConsistencyProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogConsistencyProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogAPIServer).GetLogConsistencyProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.LogAPI/GetLogConsistencyProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogAPIServer).GetLogConsistencyProof(ctx, req.(*GetLogConsistencyProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogAPI_GetLogAuditProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogAuditProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogAPIServer).GetLogAuditProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.LogAPI/GetLogAuditProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogAPIServer).GetLogAuditProof(ctx, req.(*GetLogAuditProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogAPI_GetLogEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogAPIServer).GetLogEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.LogAPI/GetLogEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogAPIServer).GetLogEntries(ctx, req.(*GetLogEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogAPI_GetMapValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMapValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogAPIServer).GetMapValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.LogAPI/GetMapValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogAPIServer).GetMapValue(ctx, req.(*GetMapValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogAPI_GetMHLogConsistencyProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogConsistencyProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogAPIServer).GetMHLogConsistencyProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.LogAPI/GetMHLogConsistencyProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogAPIServer).GetMHLogConsistencyProof(ctx, req.(*GetLogConsistencyProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogAPI_GetMHLogAuditProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogAuditProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogAPIServer).GetMHLogAuditProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.LogAPI/GetMHLogAuditProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogAPIServer).GetMHLogAuditProof(ctx, req.(*GetLogAuditProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogAPI_GetMHLogEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogAPIServer).GetMHLogEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustix.LogAPI/GetMHLogEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogAPIServer).GetMHLogEntries(ctx, req.(*GetLogEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogAPI_ServiceDesc is the grpc.ServiceDesc for LogAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "trustix.LogAPI",
	HandlerType: (*LogAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHead",
			Handler:    _LogAPI_GetHead_Handler,
		},
		{
			MethodName: "GetLogConsistencyProof",
			Handler:    _LogAPI_GetLogConsistencyProof_Handler,
		},
		{
			MethodName: "GetLogAuditProof",
			Handler:    _LogAPI_GetLogAuditProof_Handler,
		},
		{
			MethodName: "GetLogEntries",
			Handler:    _LogAPI_GetLogEntries_Handler,
		},
		{
			MethodName: "GetMapValue",
			Handler:    _LogAPI_GetMapValue_Handler,
		},
		{
			MethodName: "GetMHLogConsistencyProof",
			Handler:    _LogAPI_GetMHLogConsistencyProof_Handler,
		},
		{
			MethodName: "GetMHLogAuditProof",
			Handler:    _LogAPI_GetMHLogAuditProof_Handler,
		},
		{
			MethodName: "GetMHLogEntries",
			Handler:    _LogAPI_GetMHLogEntries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
