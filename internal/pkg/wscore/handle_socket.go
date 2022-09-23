package wscore

import (
	"cqhttp-server/config"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

// UpdateWebSocket WS处理器,转入任务池处理

func UpdateWebSocket(c *gin.Context, group *Group) {

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

	WsConn = &WebSocket{Conn: wsConn}

	log.Println("ws connect success")

	for {
		_, msg, err := WsConn.Conn.ReadMessage()
		if err != nil {
			panic(err)
		}

		//获取到消息byte

		MyWorker.Run(group.NewReceiver(msg, 8*time.Second))
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
		log.Println(event)
	case PostNotice:
		log.Println(event)
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
		if _, ok := config.Config.WhiteListPrivateMap[ctx.MetaMsg.Sender.UserId]; !ok &&
			config.Config.WhitelistPrivateFlg {
			return
		}
		ctx.Callback.Action = "send_private_msg"
		ctx.Callback.Params.(*CallbackSender).UserId = &ctx.MetaMsg.Sender.UserId
	case "group": //临时回话
		if _, ok := config.Config.WhiteListPrivateMap[ctx.MetaMsg.Sender.UserId]; !ok &&
			config.Config.WhitelistPrivateFlg {
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

		if _, ok := config.Config.WhiteListGroupMap[groupMsg.GroupId]; !ok &&
			config.Config.WhitelistGroupFlg {
			return
		}

		ctx.Callback.Action = "send_group_msg"
		ctx.Callback.Params.(*CallbackSender).GroupId = &groupMsg.GroupId
	}

	_ = WsConn.Conn.WriteJSON(ctx.Callback)

}
