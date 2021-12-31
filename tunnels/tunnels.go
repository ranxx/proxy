package tunnels

import (
	"context"
	"fmt"

	"github.com/ranxx/proxy/internal/model"
	msg "github.com/ranxx/proxy/proto/msg/v1"
	v1 "github.com/ranxx/proxy/proto/tunnels/v1"
	"github.com/ranxx/proxy/tunnels/store"
	"google.golang.org/grpc"
)

// Tunnels 用户
type Tunnels struct {
	local store.Tunnels
}

// NewTunnels ...
func NewTunnels(local store.Tunnels) *Tunnels {
	return &Tunnels{local: local}
}

// // ProvideHTTP http
// func (t *Tunnels) ProvideHTTP(router *mux.Router) {
// 	// 提供 router
// 	h := svc.MakeHTTPHandler(svc.Endpoints{
// 		ListTunnelEndpoint: svc.MakeListTunnelEndpoint(t),
// 		AddTunnelEndpoint:  svc.MakeAddTunnelEndpoint(t),
// 	},
// 		components.EncodeHTTPGenericResponse,
// 		httptransport.ServerErrorEncoder(components.ServerErrorEncoder),
// 	)
// 	router.PathPrefix("/tunnels/v1").Handler(http.StripPrefix(constant.APIPrefix+"/tunnels/v1", h))
// }

// ProvideGRPC grpc
func (t *Tunnels) ProvideGRPC(server *grpc.Server) {
	v1.RegisterTunnelsServer(server, t)
}

// ListTunnel 代理端口列表
func (t *Tunnels) ListTunnel(ctx context.Context, req *v1.ListTunnelReq) (*v1.ListTunnelRsp, error) {
	items, total, err := t.local.List(ctx, "ran.star@foxmail.com", 0, 10)
	if err != nil {
		return nil, err
	}

	resp := &v1.ListTunnelRsp{Total: total, Items: make([]*v1.Tunneler, 0, len(items))}
	for _, v := range items {
		resp.Items = append(resp.Items, &v1.Tunneler{
			Network: v.Network,
			Laddr:   (*msg.Addr)(v.Laddr),
			Raddr:   (*msg.Addr)(v.Raddr),
			Match:   (*v1.Match)(v.Match),
		})
	}
	return resp, nil
}

// AddTunnel 新增代理端口
func (t *Tunnels) AddTunnel(ctx context.Context, req *v1.AddTunnelReq) (*v1.AddTunnelRsp, error) {
	// // 本地
	// user, exists, err := u.local.Get(ctx, req.Email)
	// if err != nil {
	// 	return nil, err
	// }

	// if !exists {
	// 	return nil, errors.NewErrCode(errors.ErrNotFound, "用户不存在")
	// }

	// // 比对 passwd
	// if user.Passwd != req.Passwd {
	// 	return nil, errors.NewErrCode(errors.ErrPasswd, "密码错误")
	// }

	// return &v1.LoginRsp{Token: "xxx"}, nil

	account := "ran.star@foxmail.com"

	// 重复的
	// 是否有重复的, 外网端口重复
	items, _, err := t.local.List(ctx, account, 0, -1)
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
			Account: account,
			Laddr:   (*model.Addr)(v.Laddr),
			Raddr:   (*model.Addr)(v.Raddr),
			Network: v.Network,
			Match:   (*model.TunnelMatch)(v.Match),
			Status:  model.TunnelStatusStop,
		})
	}

	// 创建
	if err := t.local.Create(ctx, tunnels...); err != nil {
		return nil, err
	}

	return &v1.AddTunnelRsp{}, nil
}
