// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: minio.proto

package minio_pb

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
	Minio_PutObject_FullMethodName    = "/Minio/PutObject"
	Minio_RemoveObject_FullMethodName = "/Minio/RemoveObject"
	Minio_ListBucket_FullMethodName   = "/Minio/ListBucket"
	Minio_ListObject_FullMethodName   = "/Minio/ListObject"
	Minio_GetObject_FullMethodName    = "/Minio/GetObject"
)

// MinioClient is the client API for Minio service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MinioClient interface {
	PutObject(ctx context.Context, in *PutRest, opts ...grpc.CallOption) (*PutResp, error)
	RemoveObject(ctx context.Context, in *RemoveObjectRest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListBucket(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*BucketListResp, error)
	ListObject(ctx context.Context, in *ListObjectRest, opts ...grpc.CallOption) (*ListObjectResp, error)
	GetObject(ctx context.Context, in *GetObjectInfo, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type minioClient struct {
	cc grpc.ClientConnInterface
}

func NewMinioClient(cc grpc.ClientConnInterface) MinioClient {
	return &minioClient{cc}
}

func (c *minioClient) PutObject(ctx context.Context, in *PutRest, opts ...grpc.CallOption) (*PutResp, error) {
	out := new(PutResp)
	err := c.cc.Invoke(ctx, Minio_PutObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minioClient) RemoveObject(ctx context.Context, in *RemoveObjectRest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Minio_RemoveObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minioClient) ListBucket(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*BucketListResp, error) {
	out := new(BucketListResp)
	err := c.cc.Invoke(ctx, Minio_ListBucket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minioClient) ListObject(ctx context.Context, in *ListObjectRest, opts ...grpc.CallOption) (*ListObjectResp, error) {
	out := new(ListObjectResp)
	err := c.cc.Invoke(ctx, Minio_ListObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minioClient) GetObject(ctx context.Context, in *GetObjectInfo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Minio_GetObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MinioServer is the server API for Minio service.
// All implementations must embed UnimplementedMinioServer
// for forward compatibility
type MinioServer interface {
	PutObject(context.Context, *PutRest) (*PutResp, error)
	RemoveObject(context.Context, *RemoveObjectRest) (*emptypb.Empty, error)
	ListBucket(context.Context, *emptypb.Empty) (*BucketListResp, error)
	ListObject(context.Context, *ListObjectRest) (*ListObjectResp, error)
	GetObject(context.Context, *GetObjectInfo) (*emptypb.Empty, error)
	mustEmbedUnimplementedMinioServer()
}

// UnimplementedMinioServer must be embedded to have forward compatible implementations.
type UnimplementedMinioServer struct {
}

func (UnimplementedMinioServer) PutObject(context.Context, *PutRest) (*PutResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutObject not implemented")
}
func (UnimplementedMinioServer) RemoveObject(context.Context, *RemoveObjectRest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveObject not implemented")
}
func (UnimplementedMinioServer) ListBucket(context.Context, *emptypb.Empty) (*BucketListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBucket not implemented")
}
func (UnimplementedMinioServer) ListObject(context.Context, *ListObjectRest) (*ListObjectResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListObject not implemented")
}
func (UnimplementedMinioServer) GetObject(context.Context, *GetObjectInfo) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObject not implemented")
}
func (UnimplementedMinioServer) mustEmbedUnimplementedMinioServer() {}

// UnsafeMinioServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MinioServer will
// result in compilation errors.
type UnsafeMinioServer interface {
	mustEmbedUnimplementedMinioServer()
}

func RegisterMinioServer(s grpc.ServiceRegistrar, srv MinioServer) {
	s.RegisterService(&Minio_ServiceDesc, srv)
}

func _Minio_PutObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinioServer).PutObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Minio_PutObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinioServer).PutObject(ctx, req.(*PutRest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Minio_RemoveObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveObjectRest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinioServer).RemoveObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Minio_RemoveObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinioServer).RemoveObject(ctx, req.(*RemoveObjectRest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Minio_ListBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinioServer).ListBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Minio_ListBucket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinioServer).ListBucket(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Minio_ListObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListObjectRest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinioServer).ListObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Minio_ListObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinioServer).ListObject(ctx, req.(*ListObjectRest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Minio_GetObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinioServer).GetObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Minio_GetObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinioServer).GetObject(ctx, req.(*GetObjectInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// Minio_ServiceDesc is the grpc.ServiceDesc for Minio service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Minio_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Minio",
	HandlerType: (*MinioServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutObject",
			Handler:    _Minio_PutObject_Handler,
		},
		{
			MethodName: "RemoveObject",
			Handler:    _Minio_RemoveObject_Handler,
		},
		{
			MethodName: "ListBucket",
			Handler:    _Minio_ListBucket_Handler,
		},
		{
			MethodName: "ListObject",
			Handler:    _Minio_ListObject_Handler,
		},
		{
			MethodName: "GetObject",
			Handler:    _Minio_GetObject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "minio.proto",
}