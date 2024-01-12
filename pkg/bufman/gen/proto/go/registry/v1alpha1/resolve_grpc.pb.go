// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: registry/v1alpha1/resolve.proto

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
	ResolveService_GetModulePins_FullMethodName   = "/bufman.dubbo.apache.org.registry.v1alpha1.ResolveService/GetModulePins"
	ResolveService_GetGoVersion_FullMethodName    = "/bufman.dubbo.apache.org.registry.v1alpha1.ResolveService/GetGoVersion"
	ResolveService_GetSwiftVersion_FullMethodName = "/bufman.dubbo.apache.org.registry.v1alpha1.ResolveService/GetSwiftVersion"
	ResolveService_GetMavenVersion_FullMethodName = "/bufman.dubbo.apache.org.registry.v1alpha1.ResolveService/GetMavenVersion"
	ResolveService_GetNPMVersion_FullMethodName   = "/bufman.dubbo.apache.org.registry.v1alpha1.ResolveService/GetNPMVersion"
)

// ResolveServiceClient is the client API for ResolveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResolveServiceClient interface {
	// GetModulePins finds all the latest digests and respective dependencies of
	// the provided module references and picks a set of distinct modules pins.
	//
	// Note that module references with commits should still be passed to this function
	// to make sure this function can do dependency resolution.
	//
	// This function also deals with tiebreaking what ModulePin wins for the same repository.
	GetModulePins(ctx context.Context, in *GetModulePinsRequest, opts ...grpc.CallOption) (*GetModulePinsResponse, error)
	// GetGoVersion resolves the given plugin and module references to a version.
	GetGoVersion(ctx context.Context, in *GetGoVersionRequest, opts ...grpc.CallOption) (*GetGoVersionResponse, error)
	// GetSwiftVersion resolves the given plugin and module references to a version.
	GetSwiftVersion(ctx context.Context, in *GetSwiftVersionRequest, opts ...grpc.CallOption) (*GetSwiftVersionResponse, error)
	// GetMavenVersion resolves the given plugin and module references to a version.
	GetMavenVersion(ctx context.Context, in *GetMavenVersionRequest, opts ...grpc.CallOption) (*GetMavenVersionResponse, error)
	// GetNPMVersion resolves the given plugin and module references to a version.
	GetNPMVersion(ctx context.Context, in *GetNPMVersionRequest, opts ...grpc.CallOption) (*GetNPMVersionResponse, error)
}

type resolveServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResolveServiceClient(cc grpc.ClientConnInterface) ResolveServiceClient {
	return &resolveServiceClient{cc}
}

func (c *resolveServiceClient) GetModulePins(ctx context.Context, in *GetModulePinsRequest, opts ...grpc.CallOption) (*GetModulePinsResponse, error) {
	out := new(GetModulePinsResponse)
	err := c.cc.Invoke(ctx, ResolveService_GetModulePins_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resolveServiceClient) GetGoVersion(ctx context.Context, in *GetGoVersionRequest, opts ...grpc.CallOption) (*GetGoVersionResponse, error) {
	out := new(GetGoVersionResponse)
	err := c.cc.Invoke(ctx, ResolveService_GetGoVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resolveServiceClient) GetSwiftVersion(ctx context.Context, in *GetSwiftVersionRequest, opts ...grpc.CallOption) (*GetSwiftVersionResponse, error) {
	out := new(GetSwiftVersionResponse)
	err := c.cc.Invoke(ctx, ResolveService_GetSwiftVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resolveServiceClient) GetMavenVersion(ctx context.Context, in *GetMavenVersionRequest, opts ...grpc.CallOption) (*GetMavenVersionResponse, error) {
	out := new(GetMavenVersionResponse)
	err := c.cc.Invoke(ctx, ResolveService_GetMavenVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resolveServiceClient) GetNPMVersion(ctx context.Context, in *GetNPMVersionRequest, opts ...grpc.CallOption) (*GetNPMVersionResponse, error) {
	out := new(GetNPMVersionResponse)
	err := c.cc.Invoke(ctx, ResolveService_GetNPMVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResolveServiceServer is the server API for ResolveService service.
// All implementations must embed UnimplementedResolveServiceServer
// for forward compatibility
type ResolveServiceServer interface {
	// GetModulePins finds all the latest digests and respective dependencies of
	// the provided module references and picks a set of distinct modules pins.
	//
	// Note that module references with commits should still be passed to this function
	// to make sure this function can do dependency resolution.
	//
	// This function also deals with tiebreaking what ModulePin wins for the same repository.
	GetModulePins(context.Context, *GetModulePinsRequest) (*GetModulePinsResponse, error)
	// GetGoVersion resolves the given plugin and module references to a version.
	GetGoVersion(context.Context, *GetGoVersionRequest) (*GetGoVersionResponse, error)
	// GetSwiftVersion resolves the given plugin and module references to a version.
	GetSwiftVersion(context.Context, *GetSwiftVersionRequest) (*GetSwiftVersionResponse, error)
	// GetMavenVersion resolves the given plugin and module references to a version.
	GetMavenVersion(context.Context, *GetMavenVersionRequest) (*GetMavenVersionResponse, error)
	// GetNPMVersion resolves the given plugin and module references to a version.
	GetNPMVersion(context.Context, *GetNPMVersionRequest) (*GetNPMVersionResponse, error)
	mustEmbedUnimplementedResolveServiceServer()
}

// UnimplementedResolveServiceServer must be embedded to have forward compatible implementations.
type UnimplementedResolveServiceServer struct {
}

func (UnimplementedResolveServiceServer) GetModulePins(context.Context, *GetModulePinsRequest) (*GetModulePinsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModulePins not implemented")
}
func (UnimplementedResolveServiceServer) GetGoVersion(context.Context, *GetGoVersionRequest) (*GetGoVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoVersion not implemented")
}
func (UnimplementedResolveServiceServer) GetSwiftVersion(context.Context, *GetSwiftVersionRequest) (*GetSwiftVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSwiftVersion not implemented")
}
func (UnimplementedResolveServiceServer) GetMavenVersion(context.Context, *GetMavenVersionRequest) (*GetMavenVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMavenVersion not implemented")
}
func (UnimplementedResolveServiceServer) GetNPMVersion(context.Context, *GetNPMVersionRequest) (*GetNPMVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNPMVersion not implemented")
}
func (UnimplementedResolveServiceServer) mustEmbedUnimplementedResolveServiceServer() {}

// UnsafeResolveServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResolveServiceServer will
// result in compilation errors.
type UnsafeResolveServiceServer interface {
	mustEmbedUnimplementedResolveServiceServer()
}

func RegisterResolveServiceServer(s grpc.ServiceRegistrar, srv ResolveServiceServer) {
	s.RegisterService(&ResolveService_ServiceDesc, srv)
}

func _ResolveService_GetModulePins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModulePinsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResolveServiceServer).GetModulePins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResolveService_GetModulePins_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResolveServiceServer).GetModulePins(ctx, req.(*GetModulePinsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResolveService_GetGoVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResolveServiceServer).GetGoVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResolveService_GetGoVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResolveServiceServer).GetGoVersion(ctx, req.(*GetGoVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResolveService_GetSwiftVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSwiftVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResolveServiceServer).GetSwiftVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResolveService_GetSwiftVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResolveServiceServer).GetSwiftVersion(ctx, req.(*GetSwiftVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResolveService_GetMavenVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMavenVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResolveServiceServer).GetMavenVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResolveService_GetMavenVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResolveServiceServer).GetMavenVersion(ctx, req.(*GetMavenVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResolveService_GetNPMVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNPMVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResolveServiceServer).GetNPMVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResolveService_GetNPMVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResolveServiceServer).GetNPMVersion(ctx, req.(*GetNPMVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ResolveService_ServiceDesc is the grpc.ServiceDesc for ResolveService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResolveService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bufman.dubbo.apache.org.registry.v1alpha1.ResolveService",
	HandlerType: (*ResolveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetModulePins",
			Handler:    _ResolveService_GetModulePins_Handler,
		},
		{
			MethodName: "GetGoVersion",
			Handler:    _ResolveService_GetGoVersion_Handler,
		},
		{
			MethodName: "GetSwiftVersion",
			Handler:    _ResolveService_GetSwiftVersion_Handler,
		},
		{
			MethodName: "GetMavenVersion",
			Handler:    _ResolveService_GetMavenVersion_Handler,
		},
		{
			MethodName: "GetNPMVersion",
			Handler:    _ResolveService_GetNPMVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry/v1alpha1/resolve.proto",
}

const (
	LocalResolveService_GetLocalModulePins_FullMethodName = "/bufman.dubbo.apache.org.registry.v1alpha1.LocalResolveService/GetLocalModulePins"
)

// LocalResolveServiceClient is the client API for LocalResolveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocalResolveServiceClient interface {
	// GetLocalModulePins gets the latest pins for the specified local module references.
	// It also includes all of the modules transitive dependencies for the specified references.
	//
	// We want this for two reasons:
	//
	//  1. It makes it easy to say "we know we're looking for owner/repo on this specific remote".
	//     While we could just do this in GetModulePins by being aware of what our remote is
	//     (something we probably still need to know, DNS problems aside, which are more
	//     theoretical), this helps.
	//  2. Having a separate method makes us able to say "do not make decisions about what
	//     wins between competing pins for the same repo". This should only be done in
	//     GetModulePins, not in this function, i.e. only done at the top level.
	GetLocalModulePins(ctx context.Context, in *GetLocalModulePinsRequest, opts ...grpc.CallOption) (*GetLocalModulePinsResponse, error)
}

type localResolveServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocalResolveServiceClient(cc grpc.ClientConnInterface) LocalResolveServiceClient {
	return &localResolveServiceClient{cc}
}

func (c *localResolveServiceClient) GetLocalModulePins(ctx context.Context, in *GetLocalModulePinsRequest, opts ...grpc.CallOption) (*GetLocalModulePinsResponse, error) {
	out := new(GetLocalModulePinsResponse)
	err := c.cc.Invoke(ctx, LocalResolveService_GetLocalModulePins_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocalResolveServiceServer is the server API for LocalResolveService service.
// All implementations must embed UnimplementedLocalResolveServiceServer
// for forward compatibility
type LocalResolveServiceServer interface {
	// GetLocalModulePins gets the latest pins for the specified local module references.
	// It also includes all of the modules transitive dependencies for the specified references.
	//
	// We want this for two reasons:
	//
	//  1. It makes it easy to say "we know we're looking for owner/repo on this specific remote".
	//     While we could just do this in GetModulePins by being aware of what our remote is
	//     (something we probably still need to know, DNS problems aside, which are more
	//     theoretical), this helps.
	//  2. Having a separate method makes us able to say "do not make decisions about what
	//     wins between competing pins for the same repo". This should only be done in
	//     GetModulePins, not in this function, i.e. only done at the top level.
	GetLocalModulePins(context.Context, *GetLocalModulePinsRequest) (*GetLocalModulePinsResponse, error)
	mustEmbedUnimplementedLocalResolveServiceServer()
}

// UnimplementedLocalResolveServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLocalResolveServiceServer struct {
}

func (UnimplementedLocalResolveServiceServer) GetLocalModulePins(context.Context, *GetLocalModulePinsRequest) (*GetLocalModulePinsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocalModulePins not implemented")
}
func (UnimplementedLocalResolveServiceServer) mustEmbedUnimplementedLocalResolveServiceServer() {}

// UnsafeLocalResolveServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocalResolveServiceServer will
// result in compilation errors.
type UnsafeLocalResolveServiceServer interface {
	mustEmbedUnimplementedLocalResolveServiceServer()
}

func RegisterLocalResolveServiceServer(s grpc.ServiceRegistrar, srv LocalResolveServiceServer) {
	s.RegisterService(&LocalResolveService_ServiceDesc, srv)
}

func _LocalResolveService_GetLocalModulePins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocalModulePinsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalResolveServiceServer).GetLocalModulePins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocalResolveService_GetLocalModulePins_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalResolveServiceServer).GetLocalModulePins(ctx, req.(*GetLocalModulePinsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LocalResolveService_ServiceDesc is the grpc.ServiceDesc for LocalResolveService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocalResolveService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bufman.dubbo.apache.org.registry.v1alpha1.LocalResolveService",
	HandlerType: (*LocalResolveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLocalModulePins",
			Handler:    _LocalResolveService_GetLocalModulePins_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry/v1alpha1/resolve.proto",
}
