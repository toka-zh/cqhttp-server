package wscore

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

// UpdateWebSocket WS处理器,转入任务池处理

func UpdateWebSocket(c *gin.Context, group *RouterGroup) {

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
	ctx := Background(w.MetaMessage)
	// 私聊
	//todo 注册

	w.Router.UseHandle(ctx)

	//todo 传入eventMsg 返回callback

	_ = WsConn.Conn.WriteJSON(ctx.Callback)

}
