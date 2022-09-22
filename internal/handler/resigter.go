package handler

import (
	"cqhttp-server/internal/handler/pixiv"
	"cqhttp-server/pkg/wscore"
)

func Register(api *wscore.Group) {
	api.Register("pixiv图片", pixiv.Rank)
}
