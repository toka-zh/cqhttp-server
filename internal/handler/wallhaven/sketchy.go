package wallhaven

import (
	"cqhttp-server/config"
	wscore2 "cqhttp-server/internal/pkg/wscore"
	"cqhttp-server/pkg"
	"cqhttp-server/pkg/craw"
	"fmt"
)

func Sketchy(ctx *wscore2.Context) error {
	path := pkg.GetRandFileAbsPath(config.WHPath)
	if path == "" {
		craw.WallHavenCraw(config.Static.WHUrl, config.WHPath)
		path = pkg.GetRandFileAbsPath(config.WHPath)
	}
	callback := &wscore2.Callback{
		Params: &wscore2.CallbackSender{
			Message: fmt.Sprintf("[CQ:image,file=%s]", path),
		},
	}
	ctx.Callback = callback
	return nil
}

func X18(ctx *wscore2.Context) error {
	path := pkg.GetRandFileAbsPath(config.X18Path)
	if path == "" {
		craw.WallHavenCraw(config.Static.WHUrl, config.X18Path)
		path = pkg.GetRandFileAbsPath(config.X18Path)
	}
	callback := &wscore2.Callback{
		Params: &wscore2.CallbackSender{
			Message: fmt.Sprintf("[CQ:image,file=%s]", path),
		},
	}
	ctx.Callback = callback
	return nil
}
