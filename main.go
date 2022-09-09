package main

import (
	"cqhttp-server/internal/core"
	"github.com/gin-gonic/gin"
)

func WSWorker() *core.Pool {
	pool := core.New(10)
	//.SetTimeout(10 * time.Second)
	return pool
}

func main() {
	core.MyWorker = WSWorker()
	// 注册路由器

	router := gin.Default()
	router.GET("/ws", core.SocketHandler)

	_ = router.Run(":9999")
}
