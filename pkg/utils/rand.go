package utils

import (
	"math/rand"
	"time"
)

// RandInt 获取随机数
func RandInt(n int) int {
	now := time.Now()
	rand.Seed(now.Unix())
	return rand.Intn(n)
}
