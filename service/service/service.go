package service

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/ranxx/proxy/config"
	"github.com/ranxx/proxy/constant"
	"github.com/ranxx/proxy/internal/model"
	proto "github.com/ranxx/proxy/proto/msg/v1"
	"github.com/ranxx/ztcp/conn"
	"github.com/ranxx/ztcp/conner"
	"github.com/ranxx/ztcp/dispatch"
	"github.com/ranxx/ztcp/handle"
	"github.com/ranxx/ztcp/request"
	"github.com/ranxx/ztcp/router"
	"github.com/ranxx/ztcp/server"
)

// var
var (
	Svc *Service
)

// Extra ...
type Extra struct {
	HeartBeatChan chan *proto.HeartBeat
	Account       string
	Bind          bool
}

// Service svc
type Service struct {
	*server.Server
	ip     string
	port   int
	rwlock *sync.RWMutex
	// unbind map[int64]conner.Conner // 保存没有绑定的数据
}

func serviceRouters(s *Service) *router.Root {
	rs := make([]*router.Router, 0, 10)

	rs = append(rs, router.NewRouter(constant.HeartBeat, s.HeartBeatHandle()))

	rs = append(rs, router.NewRouter(constant.ClientBind, s.ClientbindHandle()))

	root := router.NewRoot().NotFound(handle.WrapHandler(func(c context.Context, r *request.Request) {
		fmt.Println(r.M.GetMsgID(), "没注册消息", string(r.M.GetData()))
	}))

	return root.AddRouter(rs...)
}

// NewService new
func NewService(ip string, port int, opts ...server.Option) *Service {
	svc := &Service{
		ip:     ip,
		port:   port,
		rwlock: new(sync.RWMutex),
		// unbind: map[int64]conner.Conner{},
	}

	opts = append(opts, server.WithGenConner(svc.genConner))

	// 注册 observer 消息
	config.Obs.SubscribeByTopicFunc("service_write_tcp_event", svc.serviceWriteTCPEvent)
	config.Obs.SubscribeByTopicFunc("new_tunnel_tcp_event", svc.newTunnelTCPEvent)

	svc.Server = server.NewServer("tcp", fmt.Sprintf("%s:%d", ip, port), opts...)

	Svc = svc

	return svc
}

// Start ...
func (s *Service) Start() {
	s.Server.Start(func(l net.Listener) error {
		log.Println("service start listening", fmt.Sprintf("%s:%d", s.ip, s.port))
		return nil
	})
}

// genConner
func (s *Service) genConner(i int64, c net.Conn) (conner.Conner, error) {
	extra := conn.WithExtra(&Extra{HeartBeatChan: make(chan *proto.HeartBeat)})

	closeHandle := conn.WithCloseHandle(func(c conner.Conner) {
		// 发送 del event
		config.Obs.Publish("del_client_event", &model.Client{
			ClientID: c.ID(),
			Address:  c.RemoteAddr().String(),
			Acccount: c.Extra().(*Extra).Account,
		})
	})

	conner := conn.NewConn(i, c, extra, closeHandle, conn.WithDispatcher(dispatch.DefaultDispatcher(serviceRouters(s))))

	// 开启 心跳检测
	go s.HeartBeatCheck(conner)

	s.rwlock.Lock()
	defer s.rwlock.Unlock()
	// s.unbind[i] = conner

	return conner, nil
}

// HeartBeatHandle 心跳事件
func (s *Service) HeartBeatHandle() handle.Handler {
	return handle.WrapHandler(func(c context.Context, r *request.Request) {
		s.rwlock.RLock()
		defer s.rwlock.RUnlock()

		// 未绑定的client 不能进行心跳消息
		if !r.C.Extra().(*Extra).Bind {
			log.Println("service", "心跳事件", r.C.RemoteAddr(), "请绑定")
			return
		}

		req := &proto.HeartBeat{}
		if err := req.Unmarshal(r.M.GetData()); err != nil {
			log.Println("service", "心跳事件", "解析 req 失败", fmt.Sprintf("err:%s", err))
			return
		}

		r.C.Extra().(*Extra).HeartBeatChan <- req
	})
}

// HeartBeatCheck 心跳检查，如果超时或者未响应，断开这个client，并移除
func (s *Service) HeartBeatCheck(c conner.Conner) {
	for {
		hbchannel := c.Extra().(*Extra).HeartBeatChan

		// 主动发送心跳
		hb := &proto.HeartBeat{Now: time.Now().Unix()}
		c.Writer().WriteValueWithID(constant.HeartBeat, hb)

		// 接收心跳
		rhb := &proto.HeartBeat{}
		select {
		case <-time.After(time.Second * 3):
		case rhb = <-hbchannel:
		}

		// 检测时间 是否在5秒钟内，这里不保证两端时间相同，所以取±3
		if rhb.Now-hb.Now >= -3 && rhb.Now-hb.Now <= 3 {
			time.Sleep(time.Second * 3)
			continue
		}

		// 断开
		log.Println("service", c.RemoteAddr(), "心跳检测失败")
		// 删除 client
		s.Server.Del(c.ID())
		// 关闭
		c.Close()
		return
	}
}

// ClientbindHandle client bind
func (s *Service) ClientbindHandle() handle.Handler {
	return handle.WrapHandler(func(c context.Context, r *request.Request) {
		bind := &proto.ClientBind{}
		if err := bind.Unmarshal(r.M.GetData()); err != nil {
			log.Println("service", "客户端绑定事件", "解析 req 失败", fmt.Sprintf("err:%s", err))
			return
		}

		// 绑定，如果客户端不绑定，则不会处理其他信息
		extra := r.C.Extra().(*Extra)
		extra.Account = bind.Account
		extra.Bind = true

		log.Println("service", "客户端绑定事件", r.C.RemoteAddr().String(), bind.Account)

		// 发送 event
		config.Obs.Publish("new_client_event", &model.Client{
			Acccount: bind.Account,
			ClientID: r.C.ID(),
			Address:  r.C.RemoteAddr().String(),
		})
	})
}

func (s *Service) serviceWriteTCPEvent(clientID int64, body *proto.TCPBody) {
	client := s.Server.GetManager().Get(clientID)
	if client == nil {
		// 没有该
		log.Println("service", "client 不存在", clientID)
		return
	}

	// 如果没有绑定，则不关心
	if !client.Extra().(*Extra).Bind {
		log.Println("service", "client 未绑定 暂不可用", clientID)
		return
	}

	client.Writer().WriteValueWithID(constant.TCP, body)
}

func (s *Service) newTunnelTCPEvent(match *model.TransferMatch) {
	s.NoticeClientsForNewTunnel(match.Acccount)
}

// NoticeClientsForNewTunnel 新隧道 client 通知
func (s *Service) NoticeClientsForNewTunnel(accounts ...string) {
	accountMap := map[string]bool{}
	for _, account := range accounts {
		accountMap[account] = true
	}
	items := []*model.Client{}
	s.GetManager().Range(func(c conner.Conner) error {
		extra := c.Extra().(*Extra)
		if len(extra.Account) <= 0 {
			return nil
		}
		if len(accounts) > 0 && !accountMap[extra.Account] {
			return nil
		}
		items = append(items, &model.Client{
			Acccount: extra.Account,
			ClientID: c.ID(),
			Address:  c.RemoteAddr().String(),
		})
		return nil
	})
	if len(items) <= 0 {
		return
	}
	config.Obs.Publish("new_client_events", items)
}
