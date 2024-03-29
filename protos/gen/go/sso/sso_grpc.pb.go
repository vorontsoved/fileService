// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: sso/sso.proto

package ssov1

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

// UploadFileClient is the client API for UploadFile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UploadFileClient interface {
	Upload(ctx context.Context, in *FileUploadRequest, opts ...grpc.CallOption) (*FileUploadResponse, error)
}

type uploadFileClient struct {
	cc grpc.ClientConnInterface
}

func NewUploadFileClient(cc grpc.ClientConnInterface) UploadFileClient {
	return &uploadFileClient{cc}
}

func (c *uploadFileClient) Upload(ctx context.Context, in *FileUploadRequest, opts ...grpc.CallOption) (*FileUploadResponse, error) {
	out := new(FileUploadResponse)
	err := c.cc.Invoke(ctx, "/uploadFile.UploadFile/Upload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UploadFileServer is the server API for UploadFile service.
// All implementations must embed UnimplementedUploadFileServer
// for forward compatibility
type UploadFileServer interface {
	Upload(context.Context, *FileUploadRequest) (*FileUploadResponse, error)
	mustEmbedUnimplementedUploadFileServer()
}

// UnimplementedUploadFileServer must be embedded to have forward compatible implementations.
type UnimplementedUploadFileServer struct {
}

func (UnimplementedUploadFileServer) Upload(context.Context, *FileUploadRequest) (*FileUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedUploadFileServer) mustEmbedUnimplementedUploadFileServer() {}

// UnsafeUploadFileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UploadFileServer will
// result in compilation errors.
type UnsafeUploadFileServer interface {
	mustEmbedUnimplementedUploadFileServer()
}

func RegisterUploadFileServer(s grpc.ServiceRegistrar, srv UploadFileServer) {
	s.RegisterService(&UploadFile_ServiceDesc, srv)
}

func _UploadFile_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadFileServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uploadFile.UploadFile/Upload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadFileServer).Upload(ctx, req.(*FileUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UploadFile_ServiceDesc is the grpc.ServiceDesc for UploadFile service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UploadFile_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "uploadFile.UploadFile",
	HandlerType: (*UploadFileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upload",
			Handler:    _UploadFile_Upload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sso/sso.proto",
}
