package wscore

import (
	"github.com/gorilla/websocket"
	"time"
)

var WsConn *WebSocket

type WebSocket struct {
	Conn *websocket.Conn
}

type WSReceiver struct {
	MetaMessage []byte
	Group       *Group
	timeout     <-chan time.Time
}

// NewReceiver 新建接受器
func (g *Group) NewReceiver(msg []byte, duration time.Duration) *WSReceiver {
	return &WSReceiver{Group: g, MetaMessage: msg, timeout: time.After(duration)}
}

func (w WSReceiver) Task() {
	go func() {
		w.eventHandler()
	}()

	select {
	case <-w.timeout:
		return
	}

}

// HandlerFunc 控制器
type HandlerFunc func(ctx *Context) error

type Group struct {
	router map[string]HandlerFunc
}

func NewGroup() *Group {
	return &Group{router: make(map[string]HandlerFunc)}
}

func (g *Group) Handle(key string, handlerFunc HandlerFunc) {
	g.router[key] = handlerFunc
}

func (g *Group) GetHandler(key string) HandlerFunc {
	return g.router[key]
}
