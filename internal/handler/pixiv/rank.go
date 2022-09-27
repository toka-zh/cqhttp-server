package pixiv

import (
	"cqhttp-server/config"
	"cqhttp-server/internal/pkg/wscore"
	"cqhttp-server/pkg/craw"
	"cqhttp-server/pkg/utils"
	"fmt"
)

func Rank(ctx *wscore.Context) error {
	path := utils.GetRandFileAbsPath(config.SavePath)
	if path == "" {
		_ = craw.PixivCraw(config.Static.PixivUrl)
		path = utils.GetRandFileAbsPath(config.SavePath)
	}
	callback := &wscore.Callback{
		Params: &wscore.CallbackSender{
			Message: fmt.Sprintf("[CQ:image,file=%s]", path),
		},
	}
	ctx.Callback = callback
	return nil
}
