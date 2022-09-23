package wscore

// Message 消息上报(PostType 为 message)数据
type Message struct {
	SubType    string `json:"sub_type"`    //表示消息的子类型
	MessageId  int    `json:"message_id"`  //消息 ID
	UserId     int64  `json:"user_id"`     //发送者 QQ 号
	Message    string `json:"message"`     //消息链
	RawMessage string `json:"raw_message"` //CQ 码格式的消息
	Font       int    `json:"font"`        //字体
	Sender     Sender `json:"sender"`      //发送者信息

	//TargetId    int64  `json:"target_id"`
	//GroupId     string `json:"group_id"`
	//MessageType string `json:"message_type"`
}

type Sender struct {
	Age      int    `json:"age"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	UserId   int64  `json:"user_id"`
}

type GroupMsg struct {
	GroupId int64 `json:"group_id"`
}
