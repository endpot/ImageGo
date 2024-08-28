package core

import (
	"ImageGo/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"ImageGo/api/core/v1"
)

func (c *ControllerV1) WebHome(ctx context.Context, req *v1.WebHomeReq) (res *v1.WebHomeRes, err error) {
	recentImageList, _ := service.CoreWeb().GetRecentImageList(ctx, &req.WebHomeInp)

	return nil, g.RequestFromCtx(ctx).Response.WriteTpl("home.tpl", g.Map{
		"title":           "Home",
		"imageCount":      service.CoreWeb().GetImageCount(ctx),
		"imageView":       service.CoreWeb().GetImageView(ctx),
		"newImageCount":   service.CoreWeb().GetNewImageCount(ctx),
		"uploaderCount":   service.CoreWeb().GetUploaderCount(ctx),
		"recentImageList": recentImageList,
	})
}
