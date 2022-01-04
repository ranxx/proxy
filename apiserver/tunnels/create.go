package tunnels

import (
	"github.com/gin-gonic/gin"
	"github.com/ranxx/proxy/errors"
	pb "github.com/ranxx/proxy/proto/tunnels/v1"
	"github.com/ranxx/proxy/tunnels"
	"github.com/ranxx/proxy/tunnels/store"
)

// Create 创建
func Create(ctx *gin.Context) (interface{}, error) {
	req := new(pb.ListTunnelReq)

	if err := ctx.BindJSON(req); err != nil {
		return nil, errors.NewHTTPBadRequest(err.Error())
	}

	return tunnels.NewTunnels(store.NewTunnels()).ListTunnel(ctx, req)
}
