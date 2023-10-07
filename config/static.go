package config

// static value

var SavePath = "./.download/pixiv/"
var WHPath = "./.download/wallhaven/"
var X18Path = "./.download/x18/"

var Static = struct {
	PixivUrl string
	WHUrl    string
}{
	PixivUrl: "https://www.pixiv.net/ranking.php?mode=daily&content=illust",
	WHUrl:    "https://wallhaven.cc/search?categories=010&purity=010&topRange=1M&sorting=toplist&order=desc",
}
