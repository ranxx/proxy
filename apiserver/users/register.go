package users

import (
	"github.com/gin-gonic/gin"
	"github.com/ranxx/proxy/errors"
	pb "github.com/ranxx/proxy/proto/users/v1"
	"github.com/ranxx/proxy/users"
	"github.com/ranxx/proxy/users/store"
)

// Register 注册
func Register(ctx *gin.Context) (interface{}, error) {
	req := new(pb.RegisterReq)

	if err := ctx.BindJSON(req); err != nil {
		return nil, errors.NewHTTPBadRequest(err.Error())
	}

	// 调用 svc
	return users.NewUsers(store.NewUsers()).Register(ctx, req)
}
