// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: registry/v1alpha1/jsonschema.proto

package registryv1alpha1

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
	JSONSchemaService_GetJSONSchema_FullMethodName = "/bufman.dubbo.apache.org.registry.v1alpha1.JSONSchemaService/GetJSONSchema"
)

// JSONSchemaServiceClient is the client API for JSONSchemaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JSONSchemaServiceClient interface {
	// GetJSONSchema allows users to get an (approximate) json schema for a
	// protobuf type.
	GetJSONSchema(ctx context.Context, in *GetJSONSchemaRequest, opts ...grpc.CallOption) (*GetJSONSchemaResponse, error)
}

type jSONSchemaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJSONSchemaServiceClient(cc grpc.ClientConnInterface) JSONSchemaServiceClient {
	return &jSONSchemaServiceClient{cc}
}

func (c *jSONSchemaServiceClient) GetJSONSchema(ctx context.Context, in *GetJSONSchemaRequest, opts ...grpc.CallOption) (*GetJSONSchemaResponse, error) {
	out := new(GetJSONSchemaResponse)
	err := c.cc.Invoke(ctx, JSONSchemaService_GetJSONSchema_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JSONSchemaServiceServer is the server API for JSONSchemaService service.
// All implementations must embed UnimplementedJSONSchemaServiceServer
// for forward compatibility
type JSONSchemaServiceServer interface {
	// GetJSONSchema allows users to get an (approximate) json schema for a
	// protobuf type.
	GetJSONSchema(context.Context, *GetJSONSchemaRequest) (*GetJSONSchemaResponse, error)
	mustEmbedUnimplementedJSONSchemaServiceServer()
}

// UnimplementedJSONSchemaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJSONSchemaServiceServer struct {
}

func (UnimplementedJSONSchemaServiceServer) GetJSONSchema(context.Context, *GetJSONSchemaRequest) (*GetJSONSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJSONSchema not implemented")
}
func (UnimplementedJSONSchemaServiceServer) mustEmbedUnimplementedJSONSchemaServiceServer() {}

// UnsafeJSONSchemaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JSONSchemaServiceServer will
// result in compilation errors.
type UnsafeJSONSchemaServiceServer interface {
	mustEmbedUnimplementedJSONSchemaServiceServer()
}

func RegisterJSONSchemaServiceServer(s grpc.ServiceRegistrar, srv JSONSchemaServiceServer) {
	s.RegisterService(&JSONSchemaService_ServiceDesc, srv)
}

func _JSONSchemaService_GetJSONSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJSONSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JSONSchemaServiceServer).GetJSONSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JSONSchemaService_GetJSONSchema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JSONSchemaServiceServer).GetJSONSchema(ctx, req.(*GetJSONSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JSONSchemaService_ServiceDesc is the grpc.ServiceDesc for JSONSchemaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JSONSchemaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bufman.dubbo.apache.org.registry.v1alpha1.JSONSchemaService",
	HandlerType: (*JSONSchemaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetJSONSchema",
			Handler:    _JSONSchemaService_GetJSONSchema_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry/v1alpha1/jsonschema.proto",
}
