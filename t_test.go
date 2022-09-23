package main

import (
	"cqhttp-server/config"
	_ "cqhttp-server/config"
	"cqhttp-server/pkg/craw"
	"fmt"
	"testing"
)

func TestCrawPixiv(t *testing.T) {
	target := "https://www.pixiv.net/ranking.php?mode=daily&content=illust"
	err := craw.PixivCraw(target)
	fmt.Println(err)
}

func TestCrawWallDaily(t *testing.T) {
	target := "https://wallhaven.cc/search?categories=111&purity=010&topRange=1d&sorting=toplist&order=desc&page=1"
	err := craw.WallHavenCraw(target, config.WHPath)
	fmt.Println(err)
}

func TestCrawX18AllDaily(t *testing.T) {
	target := "https://wallhaven.cc/search?categories=111&purity=001&topRange=1d&sorting=toplist&order=desc&page="
	var err error
	craw.WallHavenCraw(target+"1", config.X18Path)
	craw.WallHavenCraw(target+"2", config.X18Path)
	craw.WallHavenCraw(target+"3", config.X18Path)
	craw.WallHavenCraw(target+"4", config.X18Path)
	craw.WallHavenCraw(target+"5", config.X18Path)

	fmt.Println(err)
}
