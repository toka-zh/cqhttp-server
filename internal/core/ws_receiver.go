package core

import (
	"cqhttp-server/internal/model"
	"cqhttp-server/internal/model/post"
	"cqhttp-server/internal/pkg"
	pkg2 "cqhttp-server/pkg"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

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
	var event *post.Event
	err := json.Unmarshal(w.MetaMessage, &event)
	if err != nil {
		return
	}

	switch event.PostType {
	case Message:
		w.postHandler(err)
	case Request:
	case Notice:

	case MetaEvent:

	default:
		return
	}

}

func (w WSReceiver) postHandler(err error) {
	var eventMsg *post.Message
	err = json.Unmarshal(w.MetaMessage, &eventMsg)
	if err != nil {
		return
	}

	switch eventMsg.SubType {
	case "friend":
		//todo 注册
		f := w.Group.GetHandler(eventMsg.Message)

		//todo 传入eventMsg 返回callback
		ctx := Background(eventMsg)
		err := f(ctx)
		if err != nil {
			return
		}

		if ctx.callback != nil {
			_ = WsConn.conn.WriteJSON(ctx.callback)
		}

		//if eventMsg.Sender.UserId != 978766951 {
		//	return
		//}
		//
		//if strings.Contains(eventMsg.Message, "图片") {
		//	callback := model.Callback{
		//		Action: "send_private_msg",
		//		Params: model.PrivateSender{
		//			//MessageType: eventMsg.SubType,
		//			UserId:  eventMsg.Sender.UserId,
		//			Message: fmt.Sprintf("[CQ:image,file=%s]", pkg.GetRandFileAbsPath("./download")),
		//		},
		//	}
		//	_ = WsConn.conn.WriteJSON(callback)
		//}
	//case "group":
	default:
		var groupMsg *post.GroupMsg
		err = json.Unmarshal(w.MetaMessage, &groupMsg)
		if err != nil {
			return
		}
		if eventMsg.Sender.UserId != 978766951 {
			return
		}

		if strings.Contains(eventMsg.Message, "图片") {
			path := pkg.GetRandFileAbsPath("./download")
			if path == "" {
				pkg2.Craw("https://www.pixiv.net/ranking.php?mode=monthly&content=illust")
				path = pkg.GetRandFileAbsPath("./download")
			}
			callback := model.Callback{
				Action: "send_group_msg",
				Params: model.GroupSender{
					GroupId: groupMsg.GroupId,
					Message: fmt.Sprintf("[CQ:image,file=%s]", path),
				},
			}
			_ = WsConn.conn.WriteJSON(callback)
		}

		//default:
		//	return
	}

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
