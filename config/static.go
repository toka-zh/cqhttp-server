package config

var SavePath = "./.download/"
var WHPath = "./.wallhaven/"

var Static = struct {
	PixivUrl string
}{
	PixivUrl: "https://www.pixiv.net/ranking.php?mode=daily&content=illust",
}
