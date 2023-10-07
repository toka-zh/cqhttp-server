package main

import (
	"cqhttp-server/config"
	"cqhttp-server/internal/api"
	_ "cqhttp-server/internal/pkg/cron"
	"cqhttp-server/internal/pkg/wscore"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	run()
}

func WSWorker() *wscore.Pool {
	pool := wscore.New(10)
	return pool
}

func run() {
	fmt.Println(config.Config.Version, "is running...")
	// 注册全局变量
	wscore.MyWorker = WSWorker()

	// 注册路由器,并升级http为ws
	router := gin.Default()
	router.GET("/ws", api.SocketHandler)

	_ = router.Run(config.Config.Port)
}
