package cron

import (
	"cqhttp-server/config"
	craw2 "cqhttp-server/internal/craw"
)

func SyncPicture() {
	syncDailyPixivDaily()
	syncWallHavenDaily()
	sync18X()
}

func syncDailyPixivDaily() {
	target := "https://www.pixiv.net/ranking.php?mode=daily&content=illust"
	_ = craw2.PixivCraw(target)
}

func syncWallHavenDaily() {
	target := "https://wallhaven.cc/search?categories=010&purity=110&topRange=1d&sorting=toplist&order=desc&page="
	_ = craw2.WallHavenCraw(target+"1", config.WHPath)
	_ = craw2.WallHavenCraw(target+"2", config.WHPath)
	_ = craw2.WallHavenCraw(target+"3", config.WHPath)
}

func sync18X() {
	target := "https://wallhaven.cc/search?categories=111&purity=001&topRange=1d&sorting=toplist&order=desc&page="
	//var err error
	_ = craw2.WallHavenCraw(target+"1", config.X18Path)
	_ = craw2.WallHavenCraw(target+"2", config.X18Path)
	_ = craw2.WallHavenCraw(target+"3", config.X18Path)
	_ = craw2.WallHavenCraw(target+"4", config.X18Path)
	_ = craw2.WallHavenCraw(target+"5", config.X18Path)
}
