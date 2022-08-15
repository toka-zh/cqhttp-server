package core

import (
	"cqhttp-server/internal/model"
	"cqhttp-server/internal/model/post"
	"cqhttp-server/internal/pkg"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type WSReceiver struct {
	MetaMessage []byte
	timeout     <-chan time.Time
}

//NewReceiver 新建接受器
func NewReceiver(msg []byte, duration time.Duration) *WSReceiver {
	return &WSReceiver{MetaMessage: msg, timeout: time.After(duration)}
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

		if eventMsg.Sender.UserId != 978766951 {
			return
		}

		if strings.Contains(eventMsg.Message, "图片") {
			callback := model.QQCallback{
				Action: "send_private_msg",
				Params: model.PrivateSender{
					//MessageType: eventMsg.SubType,
					UserId:  eventMsg.Sender.UserId,
					Message: fmt.Sprintf("[CQ:image,file=%s]", GetRandFileAbsPath("./download")),
				},
			}
			_ = WsConn.conn.WriteJSON(callback)
		}
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
			callback := model.QQCallback{
				Action: "send_group_msg",
				Params: model.GroupSender{
					GroupId: groupMsg.GroupId,
					Message: fmt.Sprintf("[CQ:image,file=%s]", GetRandFileAbsPath("./download")),
				},
			}
			_ = WsConn.conn.WriteJSON(callback)
		}

		//default:
		//	return
	}

}

func GetRandFileAbsPath(path string) string {
	return "file:///" + GetRandFile(path)
}

// GetRandFile 随机获取目录下的文件绝对路径
// todo 做一个资源池
//  	获取目录下的文件(目录可配置),然后随机获取一个文件的绝对路径
func GetRandFile(path string) string {
	dir, err := os.ReadDir(path)
	if err != nil {
		return ""
	}

	randInt := pkg.RandInt(len(dir))
	abs, _ := filepath.Abs(path + "/" + dir[randInt].Name())
	return abs
}
