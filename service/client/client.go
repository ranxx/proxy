package client

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/ranxx/proxy/config"
	"github.com/ranxx/proxy/constant"
	"github.com/ranxx/proxy/internal/event"

	msg "github.com/ranxx/proxy/proto/msg/v1"
	"github.com/ranxx/ztcp/conn"
	"github.com/ranxx/ztcp/conner"
	"github.com/ranxx/ztcp/dispatch"
	"github.com/ranxx/ztcp/handle"
	"github.com/ranxx/ztcp/request"
	"github.com/ranxx/ztcp/router"
)

// Extra ...
type Extra struct {
}

func clientRouters(c *Client) *router.Root {
	rs := make([]*router.Router, 0, 10)

	rs = append(rs, router.NewRouter(constant.HeartBeat, c.HeartBeatHandle()))

	rs = append(rs, router.NewRouter(constant.TCP, c.TCPHandle()))

	root := router.NewRoot().NotFound(handle.WrapHandler(func(c context.Context, r *request.Request) {
		fmt.Println(r.M.GetMsgID(), "没注册消息", string(r.M.GetData()))
	}))

	return root.AddRouter(rs...)
}

// Client 客户端
// 具有重连机制，重连之后所有连接会全部断开
type Client struct {
	logPrefix string
	ip        string
	port      int
	conn      conner.Conner
	once      *sync.Once
	opts      []conn.Option
}

// NewClient ...
func NewClient(ip string, port int, opts ...conn.Option) *Client {
	return &Client{
		logPrefix: "client",
		ip:        ip,
		port:      port,
		once:      new(sync.Once),
		opts:      opts,
	}
}

// Start 开启
func (c *Client) Start() error {
	// 开启连接
	cc, err := c.dail()
	if err != nil {
		log.Println("client", "dail", fmt.Sprintf("%s:%d", c.ip, c.port), fmt.Sprintf("err:%s", err))
		return err
	}

	extra := conn.WithExtra(&Extra{})

	c.conn = conn.NewConn(0, cc, extra, conn.WithDispatcher(dispatch.DefaultDispatcher(clientRouters(c))))

	// 发送客户端绑定事件
	c.conn.Writer().WriteValueWithID(constant.ClientBind, &msg.ClientBind{
		Id: config.Cfg.Client.UserID,
	})

	c.conn.Start()

	return nil
}

func (c *Client) dail() (net.Conn, error) {
	dial := net.Dialer{Timeout: time.Second * 10}
	conn, err := dial.Dial("tcp", fmt.Sprintf("%s:%d", c.ip, c.port))
	return conn, err
}

// TCPHandle tcp 消息
func (c *Client) TCPHandle() handle.Handler {
	return handle.WrapHandler(func(c context.Context, r *request.Request) {
		req := &msg.TCPBody{}
		if err := req.Unmarshal(r.M.GetData()); err != nil {
			log.Println("client", "tcp事件", "解析 req 失败", fmt.Sprintf("err:%s", err))
			return
		}
		event.PublishCreateTCPTunnelClientEvent(req)
	})
}

// HeartBeatHandle 心跳事件
func (c *Client) HeartBeatHandle() handle.Handler {
	return handle.WrapHandler(func(c context.Context, r *request.Request) {
		req := &msg.HeartBeat{}
		if err := req.Unmarshal(r.M.GetData()); err != nil {
			log.Println("client", "心跳事件", "解析 req 失败", fmt.Sprintf("err:%s", err))
			return
		}
		log.Println("client", "收到心跳", req.Now)
		req.Now = time.Now().Unix()
		r.C.Writer().WriteValueWithID(constant.HeartBeat, req)
	})
}
