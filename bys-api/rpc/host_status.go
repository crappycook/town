package rpc

import (
	"bys/rpc/proto"
	"context"
)

// GetHostStatus sample
func GetHostStatus(ctx context.Context) (*proto.GetStatusResponse, error) {
	res, err := hostStatusCLi.GetStatus(ctx, &proto.GetStatusRequest{})
	if err != nil {
		return nil, err
	}
	return res, err
}
