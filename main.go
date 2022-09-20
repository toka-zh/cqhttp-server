package main

import (
	"cqhttp-server/internal/core"
	"github.com/gin-gonic/gin"
	"os"
)

func WSWorker() *core.Pool {
	pool := core.New(10)
	//.SetTimeout(10 * time.Second)
	return pool
}

func main() {
	args := os.Args
	if args[0] == "-version" || args[0] == "-v" {
		return
	}

	core.MyWorker = WSWorker()
	// 注册路由器

	router := gin.Default()
	router.GET("/ws", core.SocketHandler)

	_ = router.Run(":9999")
}
