package main

import (
	"cqhttp-server/config"
	"cqhttp-server/internal/api"
	"cqhttp-server/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func WSWorker() *pkg.Pool {
	pool := pkg.New(10)
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
	pkg.MyWorker = WSWorker()

	// 异步保存图片
	go pkg.PixivCraw(config.Static.PixivUrl)

	// 注册路由器,并升级http为ws
	router := gin.Default()
	router.GET("/ws", api.SocketHandler)

	_ = router.Run(config.Config.Port)
}
