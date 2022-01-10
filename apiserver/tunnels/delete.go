package tunnels

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ranxx/proxy/errors"
	pb "github.com/ranxx/proxy/proto/tunnels/v1"
	"github.com/ranxx/proxy/tunnels"
	"github.com/ranxx/proxy/tunnels/store"
)

// Delete 删除
func Delete(ctx *gin.Context) (interface{}, error) {
	req := new(pb.DeleteTunnelReq)
	err := (error)(nil)

	id, ok := ctx.Params.Get("id")
	if !ok {
		return nil, fmt.Errorf("缺少id参数")
	}

	req.Id, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errors.NewHTTPBadRequest(err.Error())
	}

	return tunnels.NewTunnels(store.NewTunnels()).Delete(ctx, req)
}
