package wallhaven

import (
	"cqhttp-server/config"
	wscore2 "cqhttp-server/internal/pkg/wscore"
	"cqhttp-server/pkg/craw"
	"cqhttp-server/pkg/utils"
	"fmt"
)

func Sketchy(ctx *wscore2.Context) error {
	path := utils.GetRandFileAbsPath(config.WHPath)
	if path == "" {
		_ = craw.WallHavenCraw(config.Static.WHUrl, config.WHPath)
		path = utils.GetRandFileAbsPath(config.WHPath)
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
	path := utils.GetRandFileAbsPath(config.X18Path)
	if path == "" {
		_ = craw.WallHavenCraw(config.Static.WHUrl, config.X18Path)
		path = utils.GetRandFileAbsPath(config.X18Path)
	}
	callback := &wscore2.Callback{
		Params: &wscore2.CallbackSender{
			Message: fmt.Sprintf("[CQ:image,file=%s]", path),
		},
	}
	ctx.Callback = callback
	return nil
}
