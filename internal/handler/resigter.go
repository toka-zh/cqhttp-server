package handler

import (
	"cqhttp-server/internal/handler/pixiv"
	"cqhttp-server/internal/handler/wallhaven"
	"cqhttp-server/pkg/wscore"
)

func Register(api *wscore.Group) {
	// p站插画日榜随机
	api.Register("pixiv", pixiv.Rank)

	// wall_haven Sketchy等级随机
	api.Register("wh111", wallhaven.Sketchy)

	// todo 手动更新图,返回下载图的数量
}
