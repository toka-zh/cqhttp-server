package wscore

import "time"

type WSReceiver struct {
	MetaMessage []byte
	Router      *RouterGroup
	timeout     <-chan time.Time
}

// NewReceiver 新建接受器
func (r *RouterGroup) NewReceiver(msg []byte, duration time.Duration) *WSReceiver {
	return &WSReceiver{Router: r, MetaMessage: msg, timeout: time.After(duration)}
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
