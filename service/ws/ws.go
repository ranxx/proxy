package ws

import (
	"context"
	"net/http"
	"sync"

	"nhooyr.io/websocket"
)

// Ws websocket
type Ws struct {
	*websocket.Conn
}

// NewWebSocket ws
func NewWebSocket(w http.ResponseWriter, r *http.Request) (*Ws, error) {
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return nil, err
	}
	return &Ws{conn}, nil
}

// Manage ...
type Manage struct {
	conns  map[int64]*Ws
	rwlock *sync.RWMutex
}

// Global ...
var Global *Manage

func init() {
	Global = &Manage{
		conns:  map[int64]*Ws{},
		rwlock: new(sync.RWMutex),
	}
}

// Append ...
func (m *Manage) Append(ctx context.Context, id int64, ws *Ws) *Manage {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	m.conns[id] = ws
	return m
}

// Get ...
func (m *Manage) Get(ctx context.Context, id int64) *Ws {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	return m.conns[id]
}
