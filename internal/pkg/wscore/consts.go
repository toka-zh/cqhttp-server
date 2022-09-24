package wscore

const (
	PostMessage   = "message"    //消息, 例如, 群聊消息
	PostRequest   = "request"    //请求, 例如, 好友申请
	PostNotice    = "notice"     //通知, 例如, 群成员增加
	PostMetaEvent = "meta_event" //元事件, 例如, go-cqhttp 心跳包
)

type SubType string

const (
	SubFriend SubType = "friend"
	SubGroup  SubType = "group"
	SubNormal SubType = "normal"
)

type ReceiveFromType int

const (
	FromAll ReceiveFromType = iota
	FromGroup
	FromFriend
)
