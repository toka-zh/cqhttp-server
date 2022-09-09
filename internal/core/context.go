package core

import (
	"cqhttp-server/internal/model"
	"cqhttp-server/internal/model/post"
)

type Context struct {
	callback *model.Callback
	attr     map[string]interface{}
	metaMsg  *post.Message
}

func Background(metaMsg *post.Message) *Context {
	return &Context{
		metaMsg: metaMsg,
		attr:    make(map[string]interface{}),
	}
}
