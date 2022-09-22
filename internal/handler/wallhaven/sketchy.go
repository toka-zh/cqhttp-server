package wallhaven

import (
	"cqhttp-server/config"
	"cqhttp-server/pkg"
	"cqhttp-server/pkg/wscore"
	"fmt"
)

func Sketchy(ctx *wscore.Context) error {
	path := pkg.GetRandFileAbsPath(config.WHPath)
	if path == "" {
		pkg.WallHavenCraw(config.Static.PixivUrl)
		path = pkg.GetRandFileAbsPath(config.WHPath)
	}
	callback := &wscore.Callback{
		Params: &wscore.CallbackSender{
			Message: fmt.Sprintf("[CQ:image,file=%s]", path),
		},
	}
	ctx.Callback = callback
	return nil
}
