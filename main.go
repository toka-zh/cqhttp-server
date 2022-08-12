package main

import (
	"cqhttp-server/internal/core"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router.Group("")
	router.GET("/", core.SocketHandler)

	_ = router.Run(":9999")
}
