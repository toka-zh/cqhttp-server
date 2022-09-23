package pixiv

import (
	"cqhttp-server/config"
	wscore2 "cqhttp-server/internal/pkg/wscore"
	"cqhttp-server/pkg"
	"cqhttp-server/pkg/craw"
	"fmt"
)

func Rank(ctx *wscore2.Context) error {
	path := pkg.GetRandFileAbsPath(config.SavePath)
	if path == "" {
		craw.PixivCraw(config.Static.PixivUrl)
		path = pkg.GetRandFileAbsPath(config.SavePath)
	}
	callback := &wscore2.Callback{
		Params: &wscore2.CallbackSender{
			Message: fmt.Sprintf("[CQ:image,file=%s]", path),
		},
	}
	ctx.Callback = callback
	return nil
}
