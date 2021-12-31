// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: aa4445e0ef
// Version Date: 2021-11-05T10:39:36Z

// Package grpc provides a gRPC client for the Tunnels service.
package grpc

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	"github.com/ranxx/proxy/kit/tunnels/v1/svc"
	pb "github.com/ranxx/proxy/proto/tunnels/v1"
)

// New returns an service backed by a gRPC client connection. It is the
// responsibility of the caller to dial, and later close, the connection.
func New(conn *grpc.ClientConn, options ...ClientOption) (pb.TunnelsServer, error) {
	var cc clientConfig

	for _, f := range options {
		err := f(&cc)
		if err != nil {
			return nil, errors.Wrap(err, "cannot apply option")
		}
	}

	clientOptions := []grpctransport.ClientOption{
		grpctransport.ClientBefore(
			contextValuesToGRPCMetadata(cc.headers)),
	}
	var listtunnelEndpoint endpoint.Endpoint
	{
		listtunnelEndpoint = grpctransport.NewClient(
			conn,
			"v1.Tunnels",
			"ListTunnel",
			EncodeGRPCListTunnelRequest,
			DecodeGRPCListTunnelResponse,
			pb.ListTunnelRsp{},
			clientOptions...,
		).Endpoint()
	}

	var addtunnelEndpoint endpoint.Endpoint
	{
		addtunnelEndpoint = grpctransport.NewClient(
			conn,
			"v1.Tunnels",
			"AddTunnel",
			EncodeGRPCAddTunnelRequest,
			DecodeGRPCAddTunnelResponse,
			pb.AddTunnelRsp{},
			clientOptions...,
		).Endpoint()
	}

	return svc.Endpoints{
		ListTunnelEndpoint: listtunnelEndpoint,
		AddTunnelEndpoint:  addtunnelEndpoint,
	}, nil
}

// GRPC Client Decode

// DecodeGRPCListTunnelResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC listtunnel reply to a user-domain listtunnel response. Primarily useful in a client.
func DecodeGRPCListTunnelResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.ListTunnelRsp)
	return reply, nil
}

// DecodeGRPCAddTunnelResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC addtunnel reply to a user-domain addtunnel response. Primarily useful in a client.
func DecodeGRPCAddTunnelResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.AddTunnelRsp)
	return reply, nil
}

// GRPC Client Encode

// EncodeGRPCListTunnelRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain listtunnel request to a gRPC listtunnel request. Primarily useful in a client.
func EncodeGRPCListTunnelRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ListTunnelReq)
	return req, nil
}

// EncodeGRPCAddTunnelRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain addtunnel request to a gRPC addtunnel request. Primarily useful in a client.
func EncodeGRPCAddTunnelRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddTunnelReq)
	return req, nil
}

type clientConfig struct {
	headers []string
}

// ClientOption is a function that modifies the client config
type ClientOption func(*clientConfig) error

func CtxValuesToSend(keys ...string) ClientOption {
	return func(o *clientConfig) error {
		o.headers = keys
		return nil
	}
}

func contextValuesToGRPCMetadata(keys []string) grpctransport.ClientRequestFunc {
	return func(ctx context.Context, md *metadata.MD) context.Context {
		var pairs []string
		for _, k := range keys {
			if v, ok := ctx.Value(k).(string); ok {
				pairs = append(pairs, k, v)
			}
		}

		if pairs != nil {
			*md = metadata.Join(*md, metadata.Pairs(pairs...))
		}

		return ctx
	}
}