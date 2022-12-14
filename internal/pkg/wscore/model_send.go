package wscore

// Request 请求上报(PostType 为 request)数据

type CallbackSender struct {
	MessageType string      `json:"message_type"`
	UserId      *int64      `json:"user_id,omitempty"`
	GroupId     *int64      `json:"group_id,omitempty"`
	Message     interface{} `json:"message"`
}

type PrivateSender struct {
	MessageType string      `json:"message_type"`
	UserId      int64       `json:"user_id"`
	Message     interface{} `json:"message"`
}

type GroupSender struct {
	GroupId int64       `json:"group_id"`
	Message interface{} `json:"message"`
}

type Callback struct {
	Action string      `json:"action"`
	Params interface{} `json:"params"`
	Echo   string      `json:"echo"`
}
