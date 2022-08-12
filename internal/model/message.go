package model

type QQMessage struct {
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	Time        int    `json:"time"`
	SelfId      int64  `json:"self_id"`
	SubType     string `json:"sub_type"`
	Sender      Sender `json:"sender"`
	MessageId   int    `json:"message_id"`
	UserId      int    `json:"user_id"`
	TargetId    int64  `json:"target_id"`
	Message     string `json:"message"`
	RawMessage  string `json:"raw_message"`
	Font        int    `json:"font"`
}
type Sender struct {
	Age      int    `json:"age"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	UserId   int    `json:"user_id"`
}
