// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package claim

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

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	ClaimEth(ctx context.Context, in *MsgClaimEth, opts ...grpc.CallOption) (*MsgClaimEthResponse, error)
	ClaimArkeo(ctx context.Context, in *MsgClaimArkeo, opts ...grpc.CallOption) (*MsgClaimArkeoResponse, error)
	TransferClaim(ctx context.Context, in *MsgTransferClaim, opts ...grpc.CallOption) (*MsgTransferClaimResponse, error)
	AddClaim(ctx context.Context, in *MsgAddClaim, opts ...grpc.CallOption) (*MsgAddClaimResponse, error)
	ClaimThorchain(ctx context.Context, in *MsgClaimThorchain, opts ...grpc.CallOption) (*MsgClaimThorchainResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ClaimEth(ctx context.Context, in *MsgClaimEth, opts ...grpc.CallOption) (*MsgClaimEthResponse, error) {
	out := new(MsgClaimEthResponse)
	err := c.cc.Invoke(ctx, "/arkeo.claim.Msg/ClaimEth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ClaimArkeo(ctx context.Context, in *MsgClaimArkeo, opts ...grpc.CallOption) (*MsgClaimArkeoResponse, error) {
	out := new(MsgClaimArkeoResponse)
	err := c.cc.Invoke(ctx, "/arkeo.claim.Msg/ClaimArkeo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) TransferClaim(ctx context.Context, in *MsgTransferClaim, opts ...grpc.CallOption) (*MsgTransferClaimResponse, error) {
	out := new(MsgTransferClaimResponse)
	err := c.cc.Invoke(ctx, "/arkeo.claim.Msg/TransferClaim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddClaim(ctx context.Context, in *MsgAddClaim, opts ...grpc.CallOption) (*MsgAddClaimResponse, error) {
	out := new(MsgAddClaimResponse)
	err := c.cc.Invoke(ctx, "/arkeo.claim.Msg/AddClaim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ClaimThorchain(ctx context.Context, in *MsgClaimThorchain, opts ...grpc.CallOption) (*MsgClaimThorchainResponse, error) {
	out := new(MsgClaimThorchainResponse)
	err := c.cc.Invoke(ctx, "/arkeo.claim.Msg/ClaimThorchain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	ClaimEth(context.Context, *MsgClaimEth) (*MsgClaimEthResponse, error)
	ClaimArkeo(context.Context, *MsgClaimArkeo) (*MsgClaimArkeoResponse, error)
	TransferClaim(context.Context, *MsgTransferClaim) (*MsgTransferClaimResponse, error)
	AddClaim(context.Context, *MsgAddClaim) (*MsgAddClaimResponse, error)
	ClaimThorchain(context.Context, *MsgClaimThorchain) (*MsgClaimThorchainResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) ClaimEth(context.Context, *MsgClaimEth) (*MsgClaimEthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimEth not implemented")
}
func (UnimplementedMsgServer) ClaimArkeo(context.Context, *MsgClaimArkeo) (*MsgClaimArkeoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimArkeo not implemented")
}
func (UnimplementedMsgServer) TransferClaim(context.Context, *MsgTransferClaim) (*MsgTransferClaimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferClaim not implemented")
}
func (UnimplementedMsgServer) AddClaim(context.Context, *MsgAddClaim) (*MsgAddClaimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddClaim not implemented")
}
func (UnimplementedMsgServer) ClaimThorchain(context.Context, *MsgClaimThorchain) (*MsgClaimThorchainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimThorchain not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_ClaimEth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClaimEth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClaimEth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arkeo.claim.Msg/ClaimEth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClaimEth(ctx, req.(*MsgClaimEth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ClaimArkeo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClaimArkeo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClaimArkeo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arkeo.claim.Msg/ClaimArkeo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClaimArkeo(ctx, req.(*MsgClaimArkeo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_TransferClaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransferClaim)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TransferClaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arkeo.claim.Msg/TransferClaim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TransferClaim(ctx, req.(*MsgTransferClaim))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddClaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddClaim)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddClaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arkeo.claim.Msg/AddClaim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddClaim(ctx, req.(*MsgAddClaim))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ClaimThorchain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClaimThorchain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClaimThorchain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arkeo.claim.Msg/ClaimThorchain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClaimThorchain(ctx, req.(*MsgClaimThorchain))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "arkeo.claim.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClaimEth",
			Handler:    _Msg_ClaimEth_Handler,
		},
		{
			MethodName: "ClaimArkeo",
			Handler:    _Msg_ClaimArkeo_Handler,
		},
		{
			MethodName: "TransferClaim",
			Handler:    _Msg_TransferClaim_Handler,
		},
		{
			MethodName: "AddClaim",
			Handler:    _Msg_AddClaim_Handler,
		},
		{
			MethodName: "ClaimThorchain",
			Handler:    _Msg_ClaimThorchain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "arkeo/claim/tx.proto",
}