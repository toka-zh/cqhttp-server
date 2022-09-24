package wscore

import (
	"cqhttp-server/config"
	"encoding/json"
)

// HandlerFunc 控制器

type HandlerFunc func(ctx *Context) error

type RouterGroup struct {
	router map[ReceiveFromType]map[string]HandlerFunc
}

type HandlerStruct struct {
	HandlerFunc
	level ReceiveFromType
}

func NewRouterGroup() *RouterGroup {
	return &RouterGroup{router: make(map[ReceiveFromType]map[string]HandlerFunc)}
}

func (r *RouterGroup) All(matchWords string, handlerFunc HandlerFunc) {
	r.Handle(FromAll, matchWords, handlerFunc)
}

func (r *RouterGroup) Friend(matchWords string, handlerFunc HandlerFunc) {
	r.Handle(FromFriend, matchWords, handlerFunc)
}

func (r *RouterGroup) GroupTmp(matchWords string, handlerFunc HandlerFunc) {
	r.Handle(FromGroup, matchWords, handlerFunc)
}

func (r *RouterGroup) Handle(receiveFromType ReceiveFromType, matchWords string, handlerFunc HandlerFunc) {
	if _, ok := r.router[receiveFromType]; !ok {
		r.router[receiveFromType] = make(map[string]HandlerFunc)
	}
	r.router[receiveFromType][matchWords] = handlerFunc
}

func (r *RouterGroup) UseHandle(ctx *Context) {
	handle := r.GetHandle(ctx.EventMsg.SubType, ctx.EventMsg.Message)
	if handle == nil {
		return
	}
	err := handle(ctx)
	if err != nil || ctx.Callback == nil {
		return
	}

	r.SetReceiver(ctx)
}

func (r *RouterGroup) SetReceiver(ctx *Context) {
	switch ctx.EventMsg.SubType {
	case SubFriend:
		if _, ok := config.Config.WhiteListPrivateMap[ctx.EventMsg.Sender.UserId]; !ok &&
			config.Config.WhitelistPrivateFlg {
			return
		}
		ctx.Callback.Action = "send_private_msg"
		ctx.Callback.Params.(*CallbackSender).UserId = &ctx.EventMsg.Sender.UserId
	case SubGroup: //临时回话
		if _, ok := config.Config.WhiteListPrivateMap[ctx.EventMsg.Sender.UserId]; !ok &&
			config.Config.WhitelistPrivateFlg {
			return
		}
		ctx.Callback.Action = "send_private_msg"
		ctx.Callback.Params.(*CallbackSender).UserId = &ctx.EventMsg.Sender.UserId
	case SubNormal:

		//群组
		var groupMsg *GroupMsg
		err := json.Unmarshal(ctx.MetaMsg, &groupMsg)
		if err != nil {
			return
		}

		if _, ok := config.Config.WhiteListGroupMap[groupMsg.GroupId]; !ok &&
			config.Config.WhitelistGroupFlg {
			return
		}

		ctx.Callback.Action = "send_group_msg"
		ctx.Callback.Params.(*CallbackSender).GroupId = &groupMsg.GroupId
	}
	return
}

func (r *RouterGroup) GetHandle(subType SubType, key string) HandlerFunc {
	var Index ReceiveFromType
	switch subType {
	case SubGroup:
		Index = FromGroup
	case SubFriend:
		Index = FromFriend
	case SubNormal:
		Index = FromGroup
	default: //
	}
	if _, ok := r.router[Index]; ok {
		if _, ok := r.router[Index][key]; ok {
			return r.router[Index][key]
		}
	}

	route, ok := r.router[FromAll][key]
	if !ok {
		return nil
	}

	return route
}
