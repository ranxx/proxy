package handlers

import (
	"context"

	pb "github.com/ranxx/proxy/proto/transfers/v1"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.TransfersServer {
	return transfersService{}
}

type transfersService struct{}

func (s transfersService) ListTransfer(ctx context.Context, in *pb.ListTransferReq) (*pb.ListTransferRsp, error) {
	var resp pb.ListTransferRsp
	return &resp, nil
}

func (s transfersService) AddTransfer(ctx context.Context, in *pb.AddTransferReq) (*pb.AddTransferRsp, error) {
	var resp pb.AddTransferRsp
	return &resp, nil
}
