// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: proto/audit.proto

package audit

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

// AuditLogClient is the client API for AuditLog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuditLogClient interface {
	Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*Empty, error)
}

type auditLogClient struct {
	cc grpc.ClientConnInterface
}

func NewAuditLogClient(cc grpc.ClientConnInterface) AuditLogClient {
	return &auditLogClient{cc}
}

func (c *auditLogClient) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/audit.AuditLog/Log", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuditLogServer is the server API for AuditLog service.
// All implementations must embed UnimplementedAuditLogServer
// for forward compatibility
type AuditLogServer interface {
	Log(context.Context, *LogRequest) (*Empty, error)
	//mustEmbedUnimplementedAuditLogServer()
}

// UnimplementedAuditLogServer must be embedded to have forward compatible implementations.
type UnimplementedAuditLogServer struct {
}

func (UnimplementedAuditLogServer) Log(context.Context, *LogRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Log not implemented")
}
func (UnimplementedAuditLogServer) mustEmbedUnimplementedAuditLogServer() {}

// UnsafeAuditLogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuditLogServer will
// result in compilation errors.
type UnsafeAuditLogServer interface {
	mustEmbedUnimplementedAuditLogServer()
}

func RegisterAuditLogServer(s grpc.ServiceRegistrar, srv AuditLogServer) {
	s.RegisterService(&AuditLog_ServiceDesc, srv)
}

func _AuditLog_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditLogServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/audit.AuditLog/Log",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditLogServer).Log(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuditLog_ServiceDesc is the grpc.ServiceDesc for AuditLog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuditLog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "audit.AuditLog",
	HandlerType: (*AuditLogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Log",
			Handler:    _AuditLog_Log_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/audit.proto",
}