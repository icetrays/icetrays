// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package consensus

import (
	context "context"
	"github.com/icetrays/icetrays/consensus/pb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RemoteExecuteClient is the client API for RemoteExecute service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RemoteExecuteClient interface {
	Execute(ctx context.Context, in *pb.Instruction, opts ...grpc.CallOption) (*pb.Empty, error)
}

type remoteExecuteClient struct {
	cc grpc.ClientConnInterface
}

func NewRemoteExecuteClient(cc grpc.ClientConnInterface) RemoteExecuteClient {
	return &remoteExecuteClient{cc}
}

func (c *remoteExecuteClient) Execute(ctx context.Context, in *pb.Instruction, opts ...grpc.CallOption) (*pb.Empty, error) {
	out := new(pb.Empty)
	err := c.cc.Invoke(ctx, "/pb.RemoteExecute/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteExecuteServer is the server API for RemoteExecute service.
// All implementations must embed UnimplementedRemoteExecuteServer
// for forward compatibility
type RemoteExecuteServer interface {
	Execute(context.Context, *pb.Instruction) (*pb.Empty, error)
	mustEmbedUnimplementedRemoteExecuteServer()
}

// UnimplementedRemoteExecuteServer must be embedded to have forward compatible implementations.
type UnimplementedRemoteExecuteServer struct {
}

func (UnimplementedRemoteExecuteServer) Execute(context.Context, *pb.Instruction) (*pb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedRemoteExecuteServer) mustEmbedUnimplementedRemoteExecuteServer() {}

// UnsafeRemoteExecuteServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemoteExecuteServer will
// result in compilation errors.
type UnsafeRemoteExecuteServer interface {
	mustEmbedUnimplementedRemoteExecuteServer()
}

func RegisterRemoteExecuteServer(s grpc.ServiceRegistrar, srv RemoteExecuteServer) {
	s.RegisterService(&RemoteExecute_ServiceDesc, srv)
}

func _RemoteExecute_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Instruction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteExecuteServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RemoteExecute/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteExecuteServer).Execute(ctx, req.(*pb.Instruction))
	}
	return interceptor(ctx, in, info, handler)
}

// RemoteExecute_ServiceDesc is the grpc.ServiceDesc for RemoteExecute service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RemoteExecute_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RemoteExecute",
	HandlerType: (*RemoteExecuteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _RemoteExecute_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "consensus/pb/fs.proto",
}