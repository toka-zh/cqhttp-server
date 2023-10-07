package api

import (
	"cqhttp-server/internal/handler"
	"cqhttp-server/internal/pkg/wscore"
	"github.com/gin-gonic/gin"
)

func SocketHandler(ctx *gin.Context) {
	// 注册
	router := wscore.NewRouterGroup()
	// p站插画日榜随机
	router.All("pixiv", handler.Rank)
	// wall_haven Sketchy等级随机
	router.All("picture", handler.Sketchy)
	// 色图
	router.GroupTmp("18x", handler.X18)

	// 协议升级,接入API
	wscore.UpdateWebSocket(ctx, router)

}
