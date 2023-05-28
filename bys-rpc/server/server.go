package server

import (
	"bysrpc/rpc/proto"
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type HostStatusServer struct {
	proto.UnimplementedHostStatusServiceServer
}

func NewHostStatusServer() *HostStatusServer {
	return &HostStatusServer{}
}

func (s *HostStatusServer) Run() {
	addr := "localhost:10086" // TODO
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	policy := keepalive.EnforcementPolicy{
		MinTime:             10 * time.Second,
		PermitWithoutStream: true,
	}
	rpcServer := grpc.NewServer(
		grpc.MaxConcurrentStreams(4096),
		grpc.KeepaliveEnforcementPolicy(policy),
	)

	proto.RegisterHostStatusServiceServer(rpcServer, s)

	fmt.Println("Service Registered.")
	if err = rpcServer.Serve(listener); err != nil {
		panic(err)
	}
}

func (s *HostStatusServer) GetStatus(ctx context.Context, req *proto.GetStatusRequest) (resp *proto.GetStatusResponse, err error) {
	fmt.Println("HostStatusServer.GetStatus called.")
	resp = new(proto.GetStatusResponse)
	resp.Hostname = "localhost"
	resp.TimestampMs = time.Now().UnixMilli()
	return
}
