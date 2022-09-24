package main

import (
	"cqhttp-server/config"
	"cqhttp-server/internal/api"
	"cqhttp-server/internal/pkg/wscore"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func WSWorker() *wscore.Pool {
	pool := wscore.New(10)
	return pool
}

func cmd() bool {
	args := os.Args
	if args[0] == "-version" || args[0] == "-v" {
		fmt.Println(config.Config.Version)
		return true
	}
	return false
}

func main() {
	// cmd模式
	if cmd() {
		return
	}

	// 注册全局变量
	wscore.MyWorker = WSWorker()

	// 异步保存图片
	//go pkg.PixivCraw(config.Static.PixivUrl)
	//go pkg.WallHavenCraw(config.Static.WHUrl)

	// 注册路由器,并升级http为ws
	router := gin.Default()
	router.GET("/ws", api.SocketHandler)

	_ = router.Run(config.Config.Port)
}
