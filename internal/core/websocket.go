package core

import (
	"cqhttp-server/internal/model"
	"cqhttp-server/internal/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var WsConn *WebSocket

const (
	Message   = "message"    //消息, 例如, 群聊消息
	Request   = "request"    //请求, 例如, 好友申请
	Notice    = "notice"     //通知, 例如, 群成员增加
	MetaEvent = "meta_event" //元事件, 例如, go-cqhttp 心跳包
)

type WebSocket struct {
	conn *websocket.Conn
}

// SocketHandler 加入任务池
func SocketHandler(c *gin.Context) {
	group := NewGroup()
	Register(group)

	// 注册websocket
	upGrader := websocket.Upgrader{
		CheckOrigin:     func(r *http.Request) bool { return true },
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	wsConn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		closeErr := wsConn.Close()
		if closeErr != nil {
			panic(err)
		}
	}()

	WsConn = &WebSocket{conn: wsConn}

	for {
		_, msg, err := WsConn.conn.ReadMessage()
		if err != nil {
			panic(err)
		}

		//获取到消息byte

		MyWorker.Run(NewReceiver(group, msg, 8*time.Second))
	}

}

func Register(api *Group) {
	api.Register("图片", func(ctx *Context) error {
		callback := &model.Callback{
			Action: "send_private_msg",
			Params: model.PrivateSender{
				//MessageType: eventMsg.SubType,
				UserId:  ctx.metaMsg.Sender.UserId,
				Message: fmt.Sprintf("[CQ:image,file=%s]", pkg.GetRandFileAbsPath("./download")),
			},
		}
		ctx.callback = callback
		return nil
	})
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
