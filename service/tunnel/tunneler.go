package tunnel

import (
	"fmt"
	"sync"

	"github.com/ranxx/proxy/config"
	"github.com/ranxx/proxy/internal/model"
	proto "github.com/ranxx/proxy/proto/msg/v1"
	"github.com/ranxx/proxy/service/tunnel/tcp"
	"github.com/ranxx/ztcp/pkg/index"
)

// Tunneler ...
type Tunneler interface {
	Start() error

	Close()
}

var (
	// Mgr ...
	Mgr *Manager
)

// Manager ...
type Manager struct {
	index  *index.Index64
	data   map[int64]Tunneler
	rwlock *sync.RWMutex
}

func init() {
	m := &Manager{
		index:  index.NewIndexI64(),
		data:   map[int64]Tunneler{},
		rwlock: &sync.RWMutex{},
	}
	config.Obs.SubscribeByTopicFunc("create_tcp_tunnel_client_event", m.createTCPTunnelClientEvent)
	Mgr = m
}

// NewManager ...
func NewManager() *Manager {
	if Mgr == nil {
		Mgr = &Manager{
			index:  index.NewIndexI64(),
			data:   map[int64]Tunneler{},
			rwlock: &sync.RWMutex{},
		}
	}
	config.Obs.SubscribeByTopicFunc("create_tcp_tunnel_client_event", Mgr.createTCPTunnelClientEvent)
	for _, tunnel := range config.Cfg.Tunnels {
		ter := NewTunnelerServer(model.TransferConfig{
			Laddr:   tunnel.Laddr,
			Raddr:   tunnel.Raddr,
			Match:   tunnel.Match,
			Network: tunnel.Network,
		})
		Mgr.Add(ter)
		config.Obs.Publish("new_tunnel_tcp_event", &tunnel.Match)
		go ter.Start()
	}
	return Mgr
}

// Add ...
func (m *Manager) Add(t Tunneler) {
	i := m.index.NewIndex()
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	m.data[i] = t
}

// Range ...
func (m *Manager) Range(fc func(t Tunneler) error) {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	for _, v := range m.data {
		if err := fc(v); err != nil {
			return
		}
	}
}

// createTCPTunnelClientEvent 等待创建 tunnel tcp client
func (m *Manager) createTCPTunnelClientEvent(body *proto.TCPBody) {
	tcp.NewClient("tunnel-tcp", body).Start()
}

// NewTunnelerServer ...
func NewTunnelerServer(cfg model.TransferConfig) Tunneler {
	var ter Tunneler
	switch cfg.Network {
	case proto.NetworkType_TCP:
		ter = tcp.NewServer(cfg, "tunnel-tcp")
	default:
		fmt.Println("不存在", cfg)
	}
	return ter
}
