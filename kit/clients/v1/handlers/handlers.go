package handlers

import (
	"context"

	pb "github.com/ranxx/proxy/proto/clients/v1"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.ClientsServer {
	return clientsService{}
}

type clientsService struct{}

func (s clientsService) ListClient(ctx context.Context, in *pb.ListClientReq) (*pb.ListClientRsp, error) {
	var resp pb.ListClientRsp
	return &resp, nil
}
