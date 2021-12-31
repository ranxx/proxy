package handlers

import (
	"context"

	pb "github.com/ranxx/proxy/proto/users/v1"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.UsersServer {
	return usersService{}
}

type usersService struct{}

func (s usersService) Register(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterRsp, error) {
	var resp pb.RegisterRsp
	return &resp, nil
}

func (s usersService) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginRsp, error) {
	var resp pb.LoginRsp
	return &resp, nil
}
