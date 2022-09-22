package main

import (
	"cqhttp-server/config"
	_ "cqhttp-server/config"
	"cqhttp-server/pkg"
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestCraw(t *testing.T) {
	target := "https://www.pixiv.net/ranking.php?mode=daily&content=illust"
	//target := "https://www.pixiv.net/ranking.php?mode=monthly&content=illust"
	err := pkg.PixivCraw(target)
	fmt.Println(err)
}

func TestCrawX(t *testing.T) {
	target := "https://wallhaven.cc/search?categories=111&purity=010&topRange=1M&sorting=toplist&order=desc&page=2"
	//target := "https://www.pixiv.net/ranking.php?mode=monthly&content=illust"
	err := pkg.WallHavenCraw(target)
	fmt.Println(err)
}

func TestViper(t *testing.T) {
	viper.SetConfigFile("./config.yml")
	viper.ReadInConfig()
	getString := viper.GetString("port")
	fmt.Println(getString)
}

func TestViperConf(t *testing.T) {
	//fmt.Println(viper.GetString("port"))
	//fmt.Println(config.Config.Port)
	fmt.Println(config.Static.PixivUrl)
}
