package tunnels

import (
	"context"
	"fmt"

	"github.com/ranxx/proxy/config"
	"github.com/ranxx/proxy/constant"
	"github.com/ranxx/proxy/errors"
	"github.com/ranxx/proxy/internal/event"
	"github.com/ranxx/proxy/internal/model"
	"github.com/ranxx/proxy/pkg/jwt"
	"github.com/ranxx/proxy/pkg/token"
	msg "github.com/ranxx/proxy/proto/msg/v1"
	v1 "github.com/ranxx/proxy/proto/tunnels/v1"
	"github.com/ranxx/proxy/tunnels/store"
	"google.golang.org/grpc"
)

func init() {
	event.SubscribeNewTCPTunnelEvent(func(t ...*model.Tunnel) {})
}

// Tunnels 用户
type Tunnels struct {
	local store.Tunnels
}

// NewTunnels ...
func NewTunnels(local store.Tunnels) *Tunnels {
	return &Tunnels{local: local}
}

// ProvideGRPC grpc
func (t *Tunnels) ProvideGRPC(server *grpc.Server) {
	v1.RegisterTunnelsServer(server, t)
}

// ListTunnel 代理端口列表
func (t *Tunnels) ListTunnel(ctx context.Context, req *v1.ListTunnelReq) (*v1.ListTunnelRsp, error) {
	claim := token.GetToken(ctx)

	items, total, err := t.local.List(ctx, claim.UserID, nil, 0, 10)
	if err != nil {
		return nil, err
	}

	resp := &v1.ListTunnelRsp{Total: total, Items: make([]*v1.Tunneler, 0, len(items))}
	for _, v := range items {
		resp.Items = append(resp.Items, &v1.Tunneler{
			Id:      v.ID,
			Network: v.Network,
			Laddr:   (*msg.Addr)(v.Laddr),
			Raddr:   (*msg.Addr)(v.Raddr),
			Match:   (*v1.Match)(v.Match),
			Status:  v1.Status(v.Status),
		})
	}
	return resp, nil
}

// AddTunnel 新增代理端口
func (t *Tunnels) AddTunnel(ctx context.Context, req *v1.AddTunnelReq) (*v1.AddTunnelRsp, error) {
	claim := token.GetToken(ctx)

	// 重复的
	// 是否有重复的, 外网端口重复
	items, _, err := t.local.List(ctx, claim.UserID, nil, 0, -1)
	if err != nil {
		return nil, err
	}

	itemsMap := map[string]bool{}
	for _, v := range items {
		itemsMap[fmt.Sprintf("%d", v.Laddr.Port)] = true
	}

	tunnels := make([]*model.Tunnel, 0, len(req.Transfers))
	for _, v := range req.Transfers {
		if itemsMap[fmt.Sprintf("%d", v.Laddr.Port)] {
			continue
		}
		tunnels = append(tunnels, &model.Tunnel{
			Base:    model.Base{},
			UserID:  claim.UserID,
			Laddr:   (*model.Addr)(v.Laddr),
			Raddr:   (*model.Addr)(v.Raddr),
			Network: v.Network,
			Match:   (*model.TunnelMatch)(v.Match),
			Status:  model.TunnelStatus(v1.Status_Stop),
		})
	}

	// 创建
	if err := t.local.Create(ctx, tunnels...); err != nil {
		return nil, err
	}

	return &v1.AddTunnelRsp{}, nil
}

func (t *Tunnels) get(ctx context.Context, claim *jwt.Claim, id int64) (*model.Tunnel, error) {
	tunnel, ok, err := t.local.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.NewNotFound("隧道%d不存在", id)
	}

	if tunnel.UserID != claim.UserID {
		return nil, errors.NewErrCode(errors.ErrPermissionDenied, "没有权限")
	}

	return tunnel, nil
}

// Start 开启
func (t *Tunnels) Start(ctx context.Context, req *v1.StartTunnelReq) (*v1.StartTunnelRsp, error) {
	claim := token.GetToken(ctx)

	tunnel, err := t.get(ctx, claim, req.Id)
	if err != nil {
		return nil, err
	}

	if tunnel.Status == model.TunnelStatus(v1.Status_Run) {
		return &v1.StartTunnelRsp{}, nil
	}

	config.Obs.SyncPublishWithRet(constant.CreateTCPTunnelEvent, func(e error) {
		err = e
	}, tunnel)

	// 更新状态
	if err == nil {
		t.local.Status(ctx, req.Id, model.TunnelStatus(v1.Status_Run))
	}

	return &v1.StartTunnelRsp{}, err
}

// Stop 停止
func (t *Tunnels) Stop(ctx context.Context, req *v1.StopTunnelReq) (*v1.StopTunnelRsp, error) {
	claim := token.GetToken(ctx)

	tunnel, err := t.get(ctx, claim, req.Id)
	if err != nil {
		return nil, err
	}

	if tunnel.Status != model.TunnelStatus(v1.Status_Run) {
		return &v1.StopTunnelRsp{}, nil
	}

	config.Obs.SyncPublishWithRet(constant.StopTCPTunnelEvent, func() {}, tunnel.ID)

	t.local.Status(ctx, req.Id, model.TunnelStatus(v1.Status_Stop))

	return &v1.StopTunnelRsp{}, nil
}

// Delete 删除
func (t *Tunnels) Delete(ctx context.Context, req *v1.DeleteTunnelReq) (*v1.DeleteTunnelRsp, error) {
	claim := token.GetToken(ctx)

	tunnel, err := t.get(ctx, claim, req.Id)
	if err != nil {
		return nil, err
	}

	// 如果没有停止，需要先停止
	if tunnel.Status == model.TunnelStatus(v1.Status_Run) {
		return nil, fmt.Errorf("先暂停然后删除")
	}

	if err := t.local.Delete(ctx, tunnel.ID); err != nil {
		return nil, err
	}

	// 发送删除消息

	return &v1.DeleteTunnelRsp{}, nil
}
