package main

import (
	"cqhttp-server/pkg"
	"fmt"
	"testing"
)

func TestCraw(t *testing.T) {
	target := "https://www.pixiv.net/ranking.php?mode=weekly&content=illust"
	//target := "https://www.pixiv.net/ranking.php?mode=monthly&content=illust"
	err := pkg.Craw(target)
	fmt.Println(err)
}
