package wscore

import "encoding/json"

type Context struct {
	Callback *Callback
	attr     map[string]interface{}
	MetaMsg  []byte
	EventMsg *Message
}

func Background(metaMsg []byte) *Context {

	var eventMsg *Message
	err := json.Unmarshal(metaMsg, &eventMsg)
	if err != nil {
		return nil
	}
	return &Context{
		MetaMsg:  metaMsg,
		EventMsg: eventMsg,
		attr:     make(map[string]interface{}),
	}
}
