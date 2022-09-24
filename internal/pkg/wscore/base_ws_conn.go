package wscore

import "github.com/gorilla/websocket"

var WsConn *WebSocket

type WebSocket struct {
	Conn *websocket.Conn
}
