package clients

import (
	"github.com/gin-gonic/gin"
	"github.com/ranxx/proxy/clients"
	"github.com/ranxx/proxy/clients/store"
	"github.com/ranxx/proxy/errors"
	pb "github.com/ranxx/proxy/proto/clients/v1"
)

// List 列表
func List(ctx *gin.Context) (interface{}, error) {
	req := new(pb.ListClientReq)

	if err := ctx.BindQuery(req); err != nil {
		return nil, errors.NewHTTPBadRequest(err.Error())
	}

	return clients.NewClients(store.NewClients()).ListClient(ctx, req)
}
