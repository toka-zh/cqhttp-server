package api

import (
	"cqhttp-server/internal/handler/pixiv"
	"cqhttp-server/internal/handler/wallhaven"
	"cqhttp-server/internal/pkg/wscore"
	"github.com/gin-gonic/gin"
)

func SocketHandler(ctx *gin.Context) {
	// 注册
	router := wscore.NewRouterGroup()
	// p站插画日榜随机
	router.All("pixiv", pixiv.Rank)
	// wall_haven Sketchy等级随机
	router.All("wallhaven", wallhaven.Sketchy)
	// 色图
	router.GroupTmp("我要色色", wallhaven.X18)

	// 协议升级,接入API
	wscore.UpdateWebSocket(ctx, router)

}
