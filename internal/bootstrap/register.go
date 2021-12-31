package bootstrap

import (
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

// ProvideHTTP ...
type ProvideHTTP interface {
	ProvideHTTP(router *mux.Router)
}

// ProvideGRPC ...
type ProvideGRPC interface {
	ProvideGRPC(server *grpc.Server)
}

// R ...
type R struct {
	https []ProvideHTTP
	grpcs []ProvideGRPC
}

// RegisterService ...
func RegisterService(items ...interface{}) {
	r := R{}
	for _, v := range items {
		if i, ok := v.(ProvideHTTP); ok {
			r.https = append(r.https, i)
		}
		if i, ok := v.(ProvideGRPC); ok {
			r.grpcs = append(r.grpcs, i)
		}
	}
}
