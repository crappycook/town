package rpc

import (
	"bys/bootstrap"
	"bys/rpc/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var (
	hostStatusCLi proto.HostStatusServiceClient
)

func InitClient(cfg *bootstrap.Config) {
	MustHostStatusService(cfg.HostStatusRpc.Addr)
}

func MustHostStatusService(serviceUrl string, dialOpts ...grpc.DialOption) {
	tlsOpt := grpc.WithInsecure()
	keepAliveOpt := grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                2 * time.Minute,
		PermitWithoutStream: true,
	})

	dialOpts = append(dialOpts, tlsOpt, keepAliveOpt)

	conn, err := grpc.Dial(serviceUrl, dialOpts...)
	if err != nil {
		panic(err)
	}

	hostStatusCLi = proto.NewHostStatusServiceClient(conn)
}
