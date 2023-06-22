// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: proto/courses.proto

package courses

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CoursesService_CreateTechnology_FullMethodName = "/CoursesService/CreateTechnology"
	CoursesService_GetTechnologies_FullMethodName  = "/CoursesService/GetTechnologies"
)

// CoursesServiceClient is the client API for CoursesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoursesServiceClient interface {
	CreateTechnology(ctx context.Context, in *CreateTechnologyRequest, opts ...grpc.CallOption) (*CreateTechnologyResponse, error)
	GetTechnologies(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTechnologiesResponse, error)
}

type coursesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCoursesServiceClient(cc grpc.ClientConnInterface) CoursesServiceClient {
	return &coursesServiceClient{cc}
}

func (c *coursesServiceClient) CreateTechnology(ctx context.Context, in *CreateTechnologyRequest, opts ...grpc.CallOption) (*CreateTechnologyResponse, error) {
	out := new(CreateTechnologyResponse)
	err := c.cc.Invoke(ctx, CoursesService_CreateTechnology_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursesServiceClient) GetTechnologies(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTechnologiesResponse, error) {
	out := new(GetTechnologiesResponse)
	err := c.cc.Invoke(ctx, CoursesService_GetTechnologies_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoursesServiceServer is the server API for CoursesService service.
// All implementations must embed UnimplementedCoursesServiceServer
// for forward compatibility
type CoursesServiceServer interface {
	CreateTechnology(context.Context, *CreateTechnologyRequest) (*CreateTechnologyResponse, error)
	GetTechnologies(context.Context, *emptypb.Empty) (*GetTechnologiesResponse, error)
	mustEmbedUnimplementedCoursesServiceServer()
}

// UnimplementedCoursesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCoursesServiceServer struct {
}

func (UnimplementedCoursesServiceServer) CreateTechnology(context.Context, *CreateTechnologyRequest) (*CreateTechnologyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTechnology not implemented")
}
func (UnimplementedCoursesServiceServer) GetTechnologies(context.Context, *emptypb.Empty) (*GetTechnologiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTechnologies not implemented")
}
func (UnimplementedCoursesServiceServer) mustEmbedUnimplementedCoursesServiceServer() {}

// UnsafeCoursesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoursesServiceServer will
// result in compilation errors.
type UnsafeCoursesServiceServer interface {
	mustEmbedUnimplementedCoursesServiceServer()
}

func RegisterCoursesServiceServer(s grpc.ServiceRegistrar, srv CoursesServiceServer) {
	s.RegisterService(&CoursesService_ServiceDesc, srv)
}

func _CoursesService_CreateTechnology_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTechnologyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursesServiceServer).CreateTechnology(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoursesService_CreateTechnology_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursesServiceServer).CreateTechnology(ctx, req.(*CreateTechnologyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoursesService_GetTechnologies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursesServiceServer).GetTechnologies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoursesService_GetTechnologies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursesServiceServer).GetTechnologies(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CoursesService_ServiceDesc is the grpc.ServiceDesc for CoursesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoursesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CoursesService",
	HandlerType: (*CoursesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTechnology",
			Handler:    _CoursesService_CreateTechnology_Handler,
		},
		{
			MethodName: "GetTechnologies",
			Handler:    _CoursesService_GetTechnologies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/courses.proto",
}