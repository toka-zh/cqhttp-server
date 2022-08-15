package main

import (
	"cqhttp-server/pkg"
	"testing"
)

func TestCraw(t *testing.T) {
	//target := "https://www.pixiv.net/ranking.php?mode=weekly&content=illust"
	target := "https://www.pixiv.net/ranking.php?mode=monthly&content=illust"
	pkg.Craw(target)
}
