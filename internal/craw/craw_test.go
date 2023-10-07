package craw

import (
	"cqhttp-server/config"
	_ "cqhttp-server/config"
	"fmt"
	"testing"
)

func TestCrawPixiv(t *testing.T) {
	target := "https://www.pixiv.net/ranking.php?mode=daily&content=illust"
	err := PixivCraw(target)
	fmt.Println(err)
}

func TestCrawWallDaily(t *testing.T) {
	target := "https://wallhaven.cc/search?categories=111&purity=010&topRange=1d&sorting=toplist&order=desc&page=1"
	err := WallHavenCraw(target, config.WHPath)
	fmt.Println(err)
}

func TestCrawX18AllDaily(t *testing.T) {
	target := "https://wallhaven.cc/search?categories=111&purity=001&topRange=1d&sorting=toplist&order=desc&page="
	var err error
	_ = WallHavenCraw(target+"1", config.X18Path)
	_ = WallHavenCraw(target+"2", config.X18Path)
	_ = WallHavenCraw(target+"3", config.X18Path)
	_ = WallHavenCraw(target+"4", config.X18Path)
	_ = WallHavenCraw(target+"5", config.X18Path)
	fmt.Println(err)
}
