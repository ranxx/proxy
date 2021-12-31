// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: aa4445e0ef
// Version Date: 2021-11-05T10:39:36Z

package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/ranxx/proxy/proto/users/v1"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC UsersServer.
func MakeGRPCServer(endpoints Endpoints, options ...grpctransport.ServerOption) pb.UsersServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
	}
	serverOptions = append(serverOptions, options...)
	return &grpcServer{
		// users

		register: grpctransport.NewServer(
			endpoints.RegisterEndpoint,
			DecodeGRPCRegisterRequest,
			EncodeGRPCRegisterResponse,
			serverOptions...,
		),
		login: grpctransport.NewServer(
			endpoints.LoginEndpoint,
			DecodeGRPCLoginRequest,
			EncodeGRPCLoginResponse,
			serverOptions...,
		),
	}
}

// grpcServer implements the UsersServer interface
type grpcServer struct {
	register grpctransport.Handler
	login    grpctransport.Handler
}

// Methods for grpcServer to implement UsersServer interface

func (s *grpcServer) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterRsp, error) {
	_, rep, err := s.register.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.RegisterRsp), nil
}

func (s *grpcServer) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRsp, error) {
	_, rep, err := s.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.LoginRsp), nil
}

// Server Decode

// DecodeGRPCRegisterRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC register request to a user-domain register request. Primarily useful in a server.
func DecodeGRPCRegisterRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.RegisterReq)
	return req, nil
}

// DecodeGRPCLoginRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC login request to a user-domain login request. Primarily useful in a server.
func DecodeGRPCLoginRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.LoginReq)
	return req, nil
}

// Server Encode

// EncodeGRPCRegisterResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain register response to a gRPC register reply. Primarily useful in a server.
func EncodeGRPCRegisterResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.RegisterRsp)
	return resp, nil
}

// EncodeGRPCLoginResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain login response to a gRPC login reply. Primarily useful in a server.
func EncodeGRPCLoginResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.LoginRsp)
	return resp, nil
}

// Helpers

func metadataToContext(ctx context.Context, md metadata.MD) context.Context {
	for k, v := range md {
		if v != nil {
			// The key is added both in metadata format (k) which is all lower
			// and the http.CanonicalHeaderKey of the key so that it can be
			// accessed in either format
			ctx = context.WithValue(ctx, k, v[0])
			ctx = context.WithValue(ctx, http.CanonicalHeaderKey(k), v[0])
		}
	}

	return ctx
}
