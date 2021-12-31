package handlers

import (
	"context"

	pb "github.com/ranxx/proxy/proto/tunnels/v1"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.TunnelsServer {
	return tunnelsService{}
}

type tunnelsService struct{}

func (s tunnelsService) ListTunnel(ctx context.Context, in *pb.ListTunnelReq) (*pb.ListTunnelRsp, error) {
	var resp pb.ListTunnelRsp
	return &resp, nil
}

func (s tunnelsService) AddTunnel(ctx context.Context, in *pb.AddTunnelReq) (*pb.AddTunnelRsp, error) {
	var resp pb.AddTunnelRsp
	return &resp, nil
}
