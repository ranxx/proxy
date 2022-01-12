package clients

import (
	"context"
	"fmt"

	"github.com/ranxx/proxy/clients/store"
	"github.com/ranxx/proxy/config"
	"github.com/ranxx/proxy/constant"
	"github.com/ranxx/proxy/errors"
	"github.com/ranxx/proxy/internal/event"
	"github.com/ranxx/proxy/internal/model"
	"github.com/ranxx/proxy/pkg/jwt"
	"github.com/ranxx/proxy/pkg/token"
	v1 "github.com/ranxx/proxy/proto/clients/v1"
	"google.golang.org/grpc"
)

func init() {
	event.SubscribeNewClientEvent(func(c *model.Client) {
		c.Status = model.ClientStatus(v1.ClientStatus_Run)
		store.NewClients().Update(context.Background(), c)
	})
	event.SubscribeStopClientEvent(func(c *model.Client) {
		c.Status = model.ClientStatus(v1.ClientStatus_Stop)
		store.NewClients().Update(context.Background(), c)
	})
	event.SubscribeExitEvent(func() {
		store.NewClients().StopAll(context.TODO())
	})
}

// Clients 用户
type Clients struct {
	local store.Clients
}

// NewClients ...
func NewClients(local store.Clients) *Clients {
	return &Clients{local: local}
}

// ProvideGRPC grpc
func (c *Clients) ProvideGRPC(server *grpc.Server) {
	v1.RegisterClientsServer(server, c)
}

// ListClient 客户端列表
func (c *Clients) ListClient(ctx context.Context, req *v1.ListClientReq) (*v1.ListClientRsp, error) {
	claim := token.GetToken(ctx)

	items, total, err := c.local.List(ctx, claim.UserID, nil, req.Offset, req.Limit)
	if err != nil {
		return nil, err
	}

	resp := &v1.ListClientRsp{
		Total: total,
		Items: make([]*v1.Client, 0, len(items)),
	}

	for _, v := range items {
		resp.Items = append(resp.Items, &v1.Client{
			Id:      v.ID,
			Machine: v.Machine,
			Status:  v1.ClientStatus(v.Status),
		})
	}

	return resp, nil
}

func (c *Clients) get(ctx context.Context, claim *jwt.Claim, id int64) (*model.Client, error) {
	item, exist, err := c.local.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("不存在的client")
	}

	if item.UserID != claim.UserID {
		return nil, errors.NewErrCode(errors.ErrPermissionDenied, "没有权限")
	}

	return item, nil
}

// StopClient 停止
func (c *Clients) StopClient(ctx context.Context, req *v1.StopClientReq) (*v1.StopClientRsp, error) {
	claim := token.GetToken(ctx)

	item, err := c.get(ctx, claim, req.Id)
	if err != nil {
		return nil, err
	}

	config.GetObs().SyncPublish(constant.StopClientEvent, item)

	item.Status = model.ClientStatus(v1.ClientStatus_Stop)
	// 修改状态
	return &v1.StopClientRsp{}, c.local.Update(ctx, item)
}

// DeleteClient 删除
func (c *Clients) DeleteClient(ctx context.Context, req *v1.DeleteClientReq) (*v1.DeleteClientRsp, error) {
	claim := token.GetToken(ctx)

	item, err := c.get(ctx, claim, req.Id)
	if err != nil {
		return nil, err
	}

	if item.Status == model.ClientStatus(v1.ClientStatus_Run) {
		return nil, fmt.Errorf("不能删除正在运行的client")
	}

	return &v1.DeleteClientRsp{}, c.local.Delete(ctx, req.Id)
}
