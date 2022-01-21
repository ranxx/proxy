package tcp

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/ranxx/proxy/config"
	"github.com/ranxx/proxy/constant"
	"github.com/ranxx/proxy/internal/event"
	"github.com/ranxx/proxy/internal/model"
	"github.com/ranxx/proxy/pkg/utils"
	msg "github.com/ranxx/proxy/proto/msg/v1"
	"github.com/ranxx/ztcp/conn"
	"github.com/ranxx/ztcp/conner"
	"github.com/ranxx/ztcp/server"
)

// Server ...
type Server struct {
	logPrefix string
	config    model.Tunnel
	once      *sync.Once
	srv       *server.Server
	rwlock    *sync.RWMutex
	clientID  int64
	clientIds []int64
	cancel    []func()
}

// NewServer ...
func NewServer(c model.Tunnel, logPrefix string) *Server {
	s := &Server{
		config:    c,
		logPrefix: fmt.Sprintf("%s %s", logPrefix, utils.TunnelAddrInfo(c.Laddr, c.Raddr)),
		once:      new(sync.Once),
		rwlock:    new(sync.RWMutex),
		clientID:  -1,
		clientIds: make([]int64, 0, 3),
		cancel:    make([]func(), 0, 3),
	}

	genConner := server.WithGenConner(func(i int64, c net.Conn) (conner.Conner, error) {
		return s.genConner(i, c)
	})

	// 注册 observer 消息
	s.cancel = append(s.cancel, event.SubscribeNewClientEvent(s.NewClientEvent))
	s.cancel = append(s.cancel, event.SubscribeBatchNewClientEvent(s.BatchNewClientEvent))
	s.cancel = append(s.cancel, event.SubscribeDelClientEvent(s.DelClientEvent))

	s.srv = server.NewServer("tcp", fmt.Sprintf("%s:%d", c.Laddr.Ip, c.Laddr.Port), genConner)
	return s
}

func (s *Server) genConner(id int64, c net.Conn) (conner.Conner, error) {
	// 如果没有client id，不允许连接
	if s.clientID < 0 {
		log.Println(s.logPrefix, "conn", "未绑定 client", "连接失败")
		c.Close()
		return nil, fmt.Errorf("不允许连接")
	}

	conner := conn.NewConn(id, c, conn.WithStop(true))
	// 首选设置超时时间，因为这个时候并不知道conn是client请求还是用户请求，客户端请求必定有数据
	conner.SetReadDeadline(time.Now().Add(time.Second * 2))
	message, headData, err := conner.Reader().ReadHeader()
	conner.SetDeadline(time.Time{})

	if err != nil || int64(len(headData)) != conner.Reader().Packer().GetHeadLength() {
		conner.Reader().Write(headData)
		s.noticeClientBind(conner)
		return conner, nil
	}

	// 判断是否为 transfer
	if message.GetMsgID() != constant.Transfer {
		// 通知 client bind
		conner.Reader().Write(headData)
		s.noticeClientBind(conner)
		return conner, nil
	}

	// 读取 body
	if _, err := conner.Reader().ReadBody(message); err != nil || message.GetDataLength() != uint32(len(message.GetData())) {
		// 通知 client bind
		conner.Reader().Write(headData)
		conner.Reader().Write(message.GetData())
		s.noticeClientBind(conner)
		return conner, nil
	}

	log.Println(s.logPrefix, "client", id)

	// bind
	bind := new(msg.TransferBind)
	bind.Unmarshal(message.GetData())

	lconner := s.srv.GetManager().Get(bind.Id)
	if lconner == nil {
		panic(fmt.Sprintf("不存在的conner %d", bind.Id))
	}

	cancel := s.clientClose(lconner, conner)

	go func() {
		defer cancel()
		if rn, err := io.Copy(conner.NetConn(), lconner.Reader()); err != nil {
			log.Println(s.logPrefix, "ioCopy11", rn, conner.ID(), lconner.ID(), string(lconner.Reader().Bytes()), err)
		}
	}()
	go func() {
		defer cancel()
		if rn, err := io.Copy(lconner.NetConn(), conner.NetConn()); err != nil {
			log.Println(s.logPrefix, "ioCopy22", rn, conner.ID(), lconner.ID(), string(lconner.Reader().Bytes()), err)
		}
	}()
	return conner, nil
}

func (s *Server) clientClose(c1, c2 conner.Conner) func() {
	once := new(sync.Once)
	return func() {
		once.Do(func() {
			c1.Close()
			c2.Close()
		})
	}
}

func (s *Server) noticeClientBind(conner conner.Conner) {
	body := &msg.TCPBody{
		Rid:   conner.ID(),
		Laddr: (*msg.Addr)(s.config.Raddr),
		Raddr: &msg.Addr{Ip: config.GetConfig().Server.IP, Port: s.config.Laddr.Port},
		Body:  nil,
		Type:  0,
	}
	log.Println(s.logPrefix, "新连接", conner.ID(), conner.RemoteAddr().String())
	// 如果没有 clientid，则不会被发送
	event.PublishNewTCPConnectionEvent(s.clientID, body)
}

// Start 开启服务
func (s *Server) Start() error {
	err := make(chan error)
	go func() {
		err <- s.srv.Start(func(l net.Listener) error {
			err <- nil
			log.Println(s.logPrefix, "running")
			return nil
		})
	}()
	return <-err
}

// Close 关闭
func (s *Server) Close() {
	s.once.Do(func() {
		s.srv.Close()
		// 删除自己
		for _, cancel := range s.cancel {
			cancel()
		}
	})
}

// Info 信息
func (s *Server) Info() model.Tunnel {
	return s.config
}

// BatchNewClientEvent 新tunnel时，通知 client 事件
func (s *Server) BatchNewClientEvent(clis []*model.Client) {
	for _, cli := range clis {
		if s.config.UserID != cli.UserID {
			continue
		}
		s.NewClientEvent(cli)
	}
}

// NewClientEvent 新client事件
func (s *Server) NewClientEvent(cli *model.Client) {
	if s.config.UserID != cli.UserID {
		return
	}
	s.newClientEvent(cli)
}

// matching 匹配 client
func (s *Server) newClientEvent(cli *model.Client) {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()
	prefix := s.config.Match.MachinePrefix
	if strings.HasPrefix(cli.Machine, prefix) {
		if s.clientID >= 0 {
			s.clientIds = append(s.clientIds, cli.ID)
			log.Println(s.logPrefix, "client", "新增绑定", cli.ID)
			return
		}
		s.clientID = cli.ID
		log.Println(s.logPrefix, "client", "绑定", s.clientID)
		return
	}

	for _, v := range s.config.Match.Machines {
		if v != cli.Machine {
			continue
		}
		if s.clientID >= 0 {
			s.clientIds = append(s.clientIds, cli.ID)
			log.Println(s.logPrefix, "client", "新增绑定", cli.ID)
			continue
		}
		s.clientID = cli.ID
		log.Println(s.logPrefix, "client", "绑定切换", s.clientID)
	}
}

// DelClientEvent 删除 client 事件
func (s *Server) DelClientEvent(cli *model.Client) {
	s.delClientEvent(cli)
}

func (s *Server) delClientEvent(cli *model.Client) {
	if s.clientID == cli.ID {
		log.Println(s.logPrefix, "client", "解除当前绑定", s.clientID)
		s.clientID = -1
	}
	s.rwlock.Lock()
	defer s.rwlock.Unlock()
	clientIds := make([]int64, 0, len(s.clientIds))
	del := make([]int64, 0, len(s.clientIds))
	for _, id := range s.clientIds {
		if id != cli.ID {
			clientIds = append(clientIds, id)
		}
		del = append(del, id)
	}
	if len(del) > 0 {
		log.Println(s.logPrefix, "client", "解除绑定", del)
	}
	s.clientIds = clientIds
	if s.clientID == -1 && len(s.clientIds) > 0 {
		s.clientID = s.clientIds[0]
		s.clientIds = s.clientIds[1:]
		log.Println(s.logPrefix, "client", "切换绑定", s.clientID)
	}
}
