package users

import (
	"github.com/gin-gonic/gin"
	"github.com/ranxx/proxy/errors"
	pb "github.com/ranxx/proxy/proto/users/v1"
	"github.com/ranxx/proxy/users"
	"github.com/ranxx/proxy/users/store"
)

// Login 登录
func Login(ctx *gin.Context) (interface{}, error) {
	req := new(pb.LoginReq)

	if err := ctx.BindJSON(req); err != nil {
		return nil, errors.NewHTTPBadRequest(err.Error())
	}

	// 调用 svc
	return users.NewUsers(store.NewUsers()).Login(ctx, req)
}
