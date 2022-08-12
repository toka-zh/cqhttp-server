package core

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

func SocketHandler(c *gin.Context) {
	upGrader := websocket.Upgrader{
		CheckOrigin:     func(r *http.Request) bool { return true },
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		closeErr := ws.Close()
		if closeErr != nil {
			panic(err)
		}
	}()

	for {
		// 解析消息
		receiver := NewReceiver(ws)
		receiver.MsgHandler()
	}
}
