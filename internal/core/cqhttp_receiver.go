package core

import (
	"cqhttp-server/internal/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

type WSReceiver struct {
	ws *websocket.Conn
}

type HandlerFunc func(string) (interface{}, error)

func NewReceiver(ws *websocket.Conn) *WSReceiver {
	return &WSReceiver{ws}
}

func (r *WSReceiver) Background(msg []byte) WSContext {
	return WSContext{msg}
}

func (r *WSReceiver) MsgHandler() {
	var qqMessage model.QQMessage

	_, msg, err := r.ws.ReadMessage()
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(msg, &qqMessage)
	if err != nil {
		return
	}
	// todo 注册解析器
	if qqMessage.Message != "" {
		fmt.Println(qqMessage.Message)
	}
}

func (r *WSReceiver) handle(keyword string, handler HandlerFunc) {
	// 解析数据,返回结果
	data, err := handler(keyword)
	if err != nil {

	}

	// 结果回写
	_ = r.ws.WriteJSON(data)
}
