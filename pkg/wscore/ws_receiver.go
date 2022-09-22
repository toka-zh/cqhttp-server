package wscore

import (
	"cqhttp-server/config"
	"encoding/json"
	"time"
)

var WsConn *WebSocket

type WSReceiver struct {
	MetaMessage []byte
	Group       *Group
	timeout     <-chan time.Time
}

// NewReceiver 新建接受器
func NewReceiver(group *Group, msg []byte, duration time.Duration) *WSReceiver {
	return &WSReceiver{Group: group, MetaMessage: msg, timeout: time.After(duration)}
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

func (w WSReceiver) eventHandler() {
	var event *Event
	err := json.Unmarshal(w.MetaMessage, &event)
	if err != nil {
		return
	}

	switch event.PostType {
	case PostMessage:
		w.postHandler(err)
	case PostRequest:
	case PostNotice:

	case PostMetaEvent:

	default:
		return
	}

}

func (w WSReceiver) postHandler(err error) {
	var eventMsg *Message
	err = json.Unmarshal(w.MetaMessage, &eventMsg)
	if err != nil {
		return
	}
	// 私聊
	//todo 注册
	f := w.Group.GetHandler(eventMsg.Message)
	if f == nil {
		return
	}

	//todo 传入eventMsg 返回callback
	ctx := Background(eventMsg)
	err = f(ctx)
	if err != nil {
		return
	}
	if ctx.Callback == nil {
		return
	}

	switch eventMsg.SubType {
	case "friend":
		if _, ok := config.Config.WhiteListPrivateMap[ctx.MetaMsg.Sender.UserId]; !ok {
			return
		}
		ctx.Callback.Action = "send_private_msg"
		ctx.Callback.Params.(*CallbackSender).UserId = &ctx.MetaMsg.Sender.UserId

	case "normal":
		//群组
		var groupMsg *GroupMsg
		err = json.Unmarshal(w.MetaMessage, &groupMsg)
		if err != nil {
			return
		}

		if _, ok := config.Config.WhiteListGroupMap[groupMsg.GroupId]; !ok {
			return
		}

		ctx.Callback.Action = "send_group_msg"
		ctx.Callback.Params.(*CallbackSender).GroupId = &groupMsg.GroupId
	}

	_ = WsConn.Conn.WriteJSON(ctx.Callback)

}

type Group struct {
	router map[string]HandlerFunc
}

func NewGroup() *Group {
	return &Group{router: make(map[string]HandlerFunc)}
}

func (g *Group) Register(key string, handlerFunc HandlerFunc) {
	g.router[key] = handlerFunc
}

func (g *Group) GetHandler(key string) HandlerFunc {
	return g.router[key]
}

// HandlerFunc 控制器
type HandlerFunc func(ctx *Context) error
