package handler

import (
	"cqhttp-server/internal/handler/pixiv"
	"cqhttp-server/internal/handler/wallhaven"
	"cqhttp-server/internal/pkg/wscore"
	"github.com/gin-gonic/gin"
)

func SocketHandler(ctx *gin.Context) {
	// 注册
	group := wscore.NewGroup()
	// p站插画日榜随机
	group.Handle("pixiv", pixiv.Rank)

	// wall_haven Sketchy等级随机
	group.Handle("wh111", wallhaven.Sketchy)

	group.Handle("我要色色", wallhaven.X18)

	wscore.UpdateWebSocket(ctx, group)

}
