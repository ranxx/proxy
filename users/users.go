package users

import (
	"context"
	"time"

	gjwt "github.com/golang-jwt/jwt"
	"github.com/ranxx/proxy/config"
	"github.com/ranxx/proxy/errors"
	"github.com/ranxx/proxy/internal/model"
	"github.com/ranxx/proxy/pkg/jwt"
	v1 "github.com/ranxx/proxy/proto/users/v1"
	"github.com/ranxx/proxy/users/store"
	"google.golang.org/grpc"
)

// Users 用户
type Users struct {
	local store.Users
}

// NewUsers ...
func NewUsers(local store.Users) *Users {
	return &Users{local: local}
}

// ProvideGRPC grpc
func (u *Users) ProvideGRPC(server *grpc.Server) {
	v1.RegisterUsersServer(server, u)
}

// Register 注册
func (u *Users) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterRsp, error) {
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
func (u *Users) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginRsp, error) {
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

	key := []byte(config.GetConfig().Auth.Key)

	token, err := jwt.Generate(key, &jwt.Claim{
		StandardClaims: gjwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(config.Cfg.Auth.ExpireInterval) * time.Second).Unix(),
			Issuer:    "proxy",
		},
		UserID: user.ID,
	}, nil)

	return &v1.LoginRsp{Token: token}, nil
}
