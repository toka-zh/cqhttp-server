package wscore

// MetaEvent 元事件上报(PostType 为 meta_event)数据
type MetaEvent struct {
	MetaEventType string `json:"meta_event_type"`
}
