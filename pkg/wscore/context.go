package wscore

type Context struct {
	Callback *Callback
	attr     map[string]interface{}
	MetaMsg  *Message
}

func Background(metaMsg *Message) *Context {
	return &Context{
		MetaMsg: metaMsg,
		attr:    make(map[string]interface{}),
	}
}
