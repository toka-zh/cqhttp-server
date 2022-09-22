package api

import (
	"cqhttp-server/internal/handler"
	"cqhttp-server/pkg"
	"cqhttp-server/pkg/wscore"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

// SocketHandler WS处理器,转入任务池处理
func SocketHandler(c *gin.Context) {
	group := wscore.NewGroup()
	handler.Register(group)

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

	wscore.WsConn = &wscore.WebSocket{Conn: wsConn}

	log.Println("ws connect success")

	for {
		_, msg, err := wscore.WsConn.Conn.ReadMessage()
		if err != nil {
			panic(err)
		}

		//获取到消息byte

		pkg.MyWorker.Run(wscore.NewReceiver(group, msg, 8*time.Second))
	}

}
