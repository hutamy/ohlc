// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: ohlc.proto

package __

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
	OHLCService_GetOHLC_FullMethodName = "/ohlc.OHLCService/GetOHLC"
)

// OHLCServiceClient is the client API for OHLCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OHLCServiceClient interface {
	GetOHLC(ctx context.Context, in *StockRequest, opts ...grpc.CallOption) (*Summary, error)
}

type oHLCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOHLCServiceClient(cc grpc.ClientConnInterface) OHLCServiceClient {
	return &oHLCServiceClient{cc}
}

func (c *oHLCServiceClient) GetOHLC(ctx context.Context, in *StockRequest, opts ...grpc.CallOption) (*Summary, error) {
	out := new(Summary)
	err := c.cc.Invoke(ctx, OHLCService_GetOHLC_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OHLCServiceServer is the server API for OHLCService service.
// All implementations must embed UnimplementedOHLCServiceServer
// for forward compatibility
type OHLCServiceServer interface {
	GetOHLC(context.Context, *StockRequest) (*Summary, error)
	mustEmbedUnimplementedOHLCServiceServer()
}

// UnimplementedOHLCServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOHLCServiceServer struct {
}

func (UnimplementedOHLCServiceServer) GetOHLC(context.Context, *StockRequest) (*Summary, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOHLC not implemented")
}
func (UnimplementedOHLCServiceServer) mustEmbedUnimplementedOHLCServiceServer() {}

// UnsafeOHLCServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OHLCServiceServer will
// result in compilation errors.
type UnsafeOHLCServiceServer interface {
	mustEmbedUnimplementedOHLCServiceServer()
}

func RegisterOHLCServiceServer(s grpc.ServiceRegistrar, srv OHLCServiceServer) {
	s.RegisterService(&OHLCService_ServiceDesc, srv)
}

func _OHLCService_GetOHLC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OHLCServiceServer).GetOHLC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OHLCService_GetOHLC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OHLCServiceServer).GetOHLC(ctx, req.(*StockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OHLCService_ServiceDesc is the grpc.ServiceDesc for OHLCService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OHLCService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ohlc.OHLCService",
	HandlerType: (*OHLCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOHLC",
			Handler:    _OHLCService_GetOHLC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ohlc.proto",
}
