package tunnel

import (
	"context"
	"fmt"
	"sync"

	"github.com/ranxx/proxy/internal/event"
	"github.com/ranxx/proxy/internal/model"
	proto "github.com/ranxx/proxy/proto/msg/v1"
	tunnelsV1 "github.com/ranxx/proxy/proto/tunnels/v1"
	"github.com/ranxx/proxy/service/tunnel/tcp"
	"github.com/ranxx/proxy/tunnels/store"
	"github.com/ranxx/ztcp/pkg/index"
)

// Tunneler ...
type Tunneler interface {
	Start() error

	Close()

	Info() model.Tunnel
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
	event.SubscribeCreateTCPTunnelClientEvent(m.createTCPTunnelClientEvent)
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
	event.SubscribeCreateTCPTunnelClientEvent(Mgr.createTCPTunnelClientEvent)
	event.SubscribeCreateTCPTunnelEvent(Mgr.createTCPTunnelEvent)
	event.SubscribeStopTCPTunnelEvent(Mgr.stopTCPTunnelEvent)
	Mgr.initTunnels()
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

// initTunnels 初始化
func (m *Manager) initTunnels() {
	list, _, err := store.NewTunnels().List(context.Background(), -1, []model.TunnelStatus{model.TunnelStatus(tunnelsV1.Status_Run)}, 0, -1)
	if err != nil {
		fmt.Println("启动失败")
	}
	m.createTCPTunnelEvent(list...)
}

// createTCPTunnelClientEvent 等待创建 tunnel tcp client
func (m *Manager) createTCPTunnelClientEvent(body *proto.TCPBody) {
	tcp.NewClient("tunnel-tcp", body).Start()
}

func (m *Manager) createTCPTunnelEvent(tc ...*model.Tunnel) error {
	for _, t := range tc {
		ter := tcp.NewServer(*t, "tunnel-tcp")
		Mgr.Add(ter)
		if err := ter.Start(); err != nil {
			return err
		}
	}
	// 发送新tunnel
	event.PublishNewTCPTunnelEvent(tc...)
	return nil
}

func (m *Manager) stopTCPTunnelEvent(id int64) {
	m.Range(func(t Tunneler) error {
		if t.Info().ID != id {
			return nil
		}
		t.Close()
		return fmt.Errorf("nil")
	})
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	tunnler, ok := m.data[id]
	if !ok {
		return
	}
	tunnler.Close()
	delete(m.data, id)
}

// NewTunnelerServer ...
func NewTunnelerServer(cfg model.Tunnel) Tunneler {
	var ter Tunneler
	switch cfg.Network {
	case proto.NetworkType_TCP:
		ter = tcp.NewServer(cfg, "tunnel-tcp")
	default:
		fmt.Println("不存在", cfg)
	}
	return ter
}
