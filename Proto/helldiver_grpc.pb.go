// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.0--rc1
// source: Proto/helldiver.proto

package Proto

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

// ServicioRecursosClient is the client API for ServicioRecursos service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServicioRecursosClient interface {
	PedirRecursos(ctx context.Context, in *ResourceRequest, opts ...grpc.CallOption) (*ResourceResponse, error)
}

type servicioRecursosClient struct {
	cc grpc.ClientConnInterface
}

func NewServicioRecursosClient(cc grpc.ClientConnInterface) ServicioRecursosClient {
	return &servicioRecursosClient{cc}
}

func (c *servicioRecursosClient) PedirRecursos(ctx context.Context, in *ResourceRequest, opts ...grpc.CallOption) (*ResourceResponse, error) {
	out := new(ResourceResponse)
	err := c.cc.Invoke(ctx, "/Proto.ServicioRecursos/PedirRecursos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServicioRecursosServer is the server API for ServicioRecursos service.
// All implementations must embed UnimplementedServicioRecursosServer
// for forward compatibility
type ServicioRecursosServer interface {
	PedirRecursos(context.Context, *ResourceRequest) (*ResourceResponse, error)
	mustEmbedUnimplementedServicioRecursosServer()
}

// UnimplementedServicioRecursosServer must be embedded to have forward compatible implementations.
type UnimplementedServicioRecursosServer struct {
}

func (UnimplementedServicioRecursosServer) PedirRecursos(context.Context, *ResourceRequest) (*ResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PedirRecursos not implemented")
}
func (UnimplementedServicioRecursosServer) mustEmbedUnimplementedServicioRecursosServer() {}

// UnsafeServicioRecursosServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServicioRecursosServer will
// result in compilation errors.
type UnsafeServicioRecursosServer interface {
	mustEmbedUnimplementedServicioRecursosServer()
}

func RegisterServicioRecursosServer(s grpc.ServiceRegistrar, srv ServicioRecursosServer) {
	s.RegisterService(&ServicioRecursos_ServiceDesc, srv)
}

func _ServicioRecursos_PedirRecursos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicioRecursosServer).PedirRecursos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Proto.ServicioRecursos/PedirRecursos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicioRecursosServer).PedirRecursos(ctx, req.(*ResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServicioRecursos_ServiceDesc is the grpc.ServiceDesc for ServicioRecursos service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServicioRecursos_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Proto.ServicioRecursos",
	HandlerType: (*ServicioRecursosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PedirRecursos",
			Handler:    _ServicioRecursos_PedirRecursos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Proto/helldiver.proto",
}
