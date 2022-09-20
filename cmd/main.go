package main

import (
	"cqhttp-server/config"
	"cqhttp-server/internal/core"
	pkg2 "cqhttp-server/pkg"
	"github.com/gin-gonic/gin"
	"os"
)

func WSWorker() *core.Pool {
	pool := core.New(10)
	return pool
}

func cmd() bool {
	args := os.Args
	if args[0] == "-version" || args[0] == "-v" {
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
	core.MyWorker = WSWorker()
	go pkg2.PixivCraw(config.PixivUrl)

	// 注册路由器,并升级http为ws
	router := gin.Default()
	router.GET("/ws", core.SocketHandler)

	_ = router.Run(":9999")
}
