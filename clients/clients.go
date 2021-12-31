package clients

import (
	"context"

	"github.com/ranxx/proxy/clients/store"
	"github.com/ranxx/proxy/errors"
	"github.com/ranxx/proxy/internal/model"
	v1 "github.com/ranxx/proxy/proto/users/v1"
	"google.golang.org/grpc"
)

// Clients 用户
type Clients struct {
	local store.Clients
}

// NewUsers ...
func NewUsers(local store.Clients) *Clients {
	return &Clients{local: local}
}

// // ProvideHTTP http
// func (u *Clients) ProvideHTTP(router *mux.Router) {
// 	// 提供 router
// 	h := svc.MakeHTTPHandler(svc.Endpoints{
// 		RegisterEndpoint: svc.MakeRegisterEndpoint(u),
// 		LoginEndpoint:    svc.MakeLoginEndpoint(u),
// 	},
// 		components.EncodeHTTPGenericResponse,
// 		httptransport.ServerErrorEncoder(components.ServerErrorEncoder),
// 	)
// 	router.PathPrefix("/clients/v1").Handler(http.StripPrefix("/api/clients/v1", h))
// }

// ProvideGRPC grpc
func (u *Clients) ProvideGRPC(server *grpc.Server) {
	v1.RegisterUsersServer(server, u)
}

// Register 注册
func (u *Clients) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterRsp, error) {
	// 本地
	_, exists, err := u.local.Get(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.NewErrCode(errors.ErrExists, "用户已存在")
	}

	// 不存在，则创建
	return &v1.RegisterRsp{}, u.local.Create(ctx, &model.User{Email: req.Email, Passwd: req.Passwd})
}

// Login 登录
func (u *Clients) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginRsp, error) {
	// 本地
	user, exists, err := u.local.Get(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.NewErrCode(errors.ErrNotFound, "用户不存在")
	}

	// 比对 passwd
	if user.Passwd != req.Passwd {
		return nil, errors.NewErrCode(errors.ErrPasswd, "密码错误")
	}

	return &v1.LoginRsp{Token: "xxx"}, nil
}
