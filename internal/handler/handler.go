package handler

import (
	"cqhttp-server/internal/craw"
	"cqhttp-server/internal/pkg/wscore"
	"cqhttp-server/pkg/utils"
	"fmt"
)

func SetPicture(ctx *wscore.Context, pictureId string, drawUrl string) {

	// 共通 发送图片消息
	path := utils.GetRandFileAbsPath(pictureId)
	if path == "" {
		_ = craw.WallHavenCraw(drawUrl, pictureId)
		path = utils.GetRandFileAbsPath(pictureId)
		// todo 抽出错误处理
	}
	callback := &wscore.Callback{
		Params: &wscore.CallbackSender{
			Message: fmt.Sprintf("[CQ:image,file=%s]", path),
		},
	}

	// 将图片路径保存到上下文中
	ctx.Callback = callback
}
