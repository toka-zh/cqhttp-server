package wscore

// Event 上报事件通用数据
type Event struct {
	PostType string `json:"post_type"` //表示该上报的类型: 消息, 请求, 通知, 或元事件 message, request, notice, meta_event
	Time     int    `json:"time"`      //事件发生的时间戳
	SelfId   int64  `json:"self_id"`   //收到事件的机器人的 QQ 号
}
