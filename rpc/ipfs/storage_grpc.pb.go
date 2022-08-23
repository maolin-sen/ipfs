// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: storage.proto

package ipfs

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

// IPFSClient is the client API for IPFS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IPFSClient interface {
	Stage(ctx context.Context, in *APIID, opts ...grpc.CallOption) (*Cid, error)
	StageCid(ctx context.Context, in *StageCidRequest, opts ...grpc.CallOption) (*Empty, error)
	Unpin(ctx context.Context, in *UnpinRequest, opts ...grpc.CallOption) (*Empty, error)
	Pin(ctx context.Context, in *PinRequest, opts ...grpc.CallOption) (*PinResponse, error)
	Get(ctx context.Context, in *Cid, opts ...grpc.CallOption) (*Empty, error)
	Replace(ctx context.Context, in *ReplaceRequest, opts ...grpc.CallOption) (*ReplaceResponse, error)
	IsPinned(ctx context.Context, in *IsPinnedRequest, opts ...grpc.CallOption) (*IsPinnedResponse, error)
	GCStaged(ctx context.Context, in *GCStagedRequest, opts ...grpc.CallOption) (*GCStagedResponse, error)
	PinnedCids(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type iPFSClient struct {
	cc grpc.ClientConnInterface
}

func NewIPFSClient(cc grpc.ClientConnInterface) IPFSClient {
	return &iPFSClient{cc}
}

func (c *iPFSClient) Stage(ctx context.Context, in *APIID, opts ...grpc.CallOption) (*Cid, error) {
	out := new(Cid)
	err := c.cc.Invoke(ctx, "/ipfs.IPFS/Stage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPFSClient) StageCid(ctx context.Context, in *StageCidRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/ipfs.IPFS/StageCid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPFSClient) Unpin(ctx context.Context, in *UnpinRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/ipfs.IPFS/Unpin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPFSClient) Pin(ctx context.Context, in *PinRequest, opts ...grpc.CallOption) (*PinResponse, error) {
	out := new(PinResponse)
	err := c.cc.Invoke(ctx, "/ipfs.IPFS/Pin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPFSClient) Get(ctx context.Context, in *Cid, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/ipfs.IPFS/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPFSClient) Replace(ctx context.Context, in *ReplaceRequest, opts ...grpc.CallOption) (*ReplaceResponse, error) {
	out := new(ReplaceResponse)
	err := c.cc.Invoke(ctx, "/ipfs.IPFS/Replace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPFSClient) IsPinned(ctx context.Context, in *IsPinnedRequest, opts ...grpc.CallOption) (*IsPinnedResponse, error) {
	out := new(IsPinnedResponse)
	err := c.cc.Invoke(ctx, "/ipfs.IPFS/IsPinned", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPFSClient) GCStaged(ctx context.Context, in *GCStagedRequest, opts ...grpc.CallOption) (*GCStagedResponse, error) {
	out := new(GCStagedResponse)
	err := c.cc.Invoke(ctx, "/ipfs.IPFS/GCStaged", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPFSClient) PinnedCids(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/ipfs.IPFS/PinnedCids", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IPFSServer is the server API for IPFS service.
// All implementations must embed UnimplementedIPFSServer
// for forward compatibility
type IPFSServer interface {
	Stage(context.Context, *APIID) (*Cid, error)
	StageCid(context.Context, *StageCidRequest) (*Empty, error)
	Unpin(context.Context, *UnpinRequest) (*Empty, error)
	Pin(context.Context, *PinRequest) (*PinResponse, error)
	Get(context.Context, *Cid) (*Empty, error)
	Replace(context.Context, *ReplaceRequest) (*ReplaceResponse, error)
	IsPinned(context.Context, *IsPinnedRequest) (*IsPinnedResponse, error)
	GCStaged(context.Context, *GCStagedRequest) (*GCStagedResponse, error)
	PinnedCids(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedIPFSServer()
}

// UnimplementedIPFSServer must be embedded to have forward compatible implementations.
type UnimplementedIPFSServer struct {
}

func (UnimplementedIPFSServer) Stage(context.Context, *APIID) (*Cid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stage not implemented")
}
func (UnimplementedIPFSServer) StageCid(context.Context, *StageCidRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StageCid not implemented")
}
func (UnimplementedIPFSServer) Unpin(context.Context, *UnpinRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unpin not implemented")
}
func (UnimplementedIPFSServer) Pin(context.Context, *PinRequest) (*PinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pin not implemented")
}
func (UnimplementedIPFSServer) Get(context.Context, *Cid) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedIPFSServer) Replace(context.Context, *ReplaceRequest) (*ReplaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Replace not implemented")
}
func (UnimplementedIPFSServer) IsPinned(context.Context, *IsPinnedRequest) (*IsPinnedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsPinned not implemented")
}
func (UnimplementedIPFSServer) GCStaged(context.Context, *GCStagedRequest) (*GCStagedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GCStaged not implemented")
}
func (UnimplementedIPFSServer) PinnedCids(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PinnedCids not implemented")
}
func (UnimplementedIPFSServer) mustEmbedUnimplementedIPFSServer() {}

// UnsafeIPFSServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IPFSServer will
// result in compilation errors.
type UnsafeIPFSServer interface {
	mustEmbedUnimplementedIPFSServer()
}

func RegisterIPFSServer(s grpc.ServiceRegistrar, srv IPFSServer) {
	s.RegisterService(&IPFS_ServiceDesc, srv)
}

func _IPFS_Stage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(APIID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPFSServer).Stage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ipfs.IPFS/Stage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPFSServer).Stage(ctx, req.(*APIID))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPFS_StageCid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StageCidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPFSServer).StageCid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ipfs.IPFS/StageCid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPFSServer).StageCid(ctx, req.(*StageCidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPFS_Unpin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnpinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPFSServer).Unpin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ipfs.IPFS/Unpin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPFSServer).Unpin(ctx, req.(*UnpinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPFS_Pin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPFSServer).Pin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ipfs.IPFS/Pin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPFSServer).Pin(ctx, req.(*PinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPFS_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPFSServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ipfs.IPFS/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPFSServer).Get(ctx, req.(*Cid))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPFS_Replace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPFSServer).Replace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ipfs.IPFS/Replace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPFSServer).Replace(ctx, req.(*ReplaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPFS_IsPinned_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsPinnedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPFSServer).IsPinned(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ipfs.IPFS/IsPinned",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPFSServer).IsPinned(ctx, req.(*IsPinnedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPFS_GCStaged_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GCStagedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPFSServer).GCStaged(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ipfs.IPFS/GCStaged",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPFSServer).GCStaged(ctx, req.(*GCStagedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPFS_PinnedCids_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPFSServer).PinnedCids(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ipfs.IPFS/PinnedCids",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPFSServer).PinnedCids(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// IPFS_ServiceDesc is the grpc.ServiceDesc for IPFS service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IPFS_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ipfs.IPFS",
	HandlerType: (*IPFSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Stage",
			Handler:    _IPFS_Stage_Handler,
		},
		{
			MethodName: "StageCid",
			Handler:    _IPFS_StageCid_Handler,
		},
		{
			MethodName: "Unpin",
			Handler:    _IPFS_Unpin_Handler,
		},
		{
			MethodName: "Pin",
			Handler:    _IPFS_Pin_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _IPFS_Get_Handler,
		},
		{
			MethodName: "Replace",
			Handler:    _IPFS_Replace_Handler,
		},
		{
			MethodName: "IsPinned",
			Handler:    _IPFS_IsPinned_Handler,
		},
		{
			MethodName: "GCStaged",
			Handler:    _IPFS_GCStaged_Handler,
		},
		{
			MethodName: "PinnedCids",
			Handler:    _IPFS_PinnedCids_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storage.proto",
}
