package wscore

import (
	"github.com/gorilla/websocket"
)

const (
	PostMessage   = "message"    //消息, 例如, 群聊消息
	PostRequest   = "request"    //请求, 例如, 好友申请
	PostNotice    = "notice"     //通知, 例如, 群成员增加
	PostMetaEvent = "meta_event" //元事件, 例如, go-cqhttp 心跳包
)

type WebSocket struct {
	Conn *websocket.Conn
}
