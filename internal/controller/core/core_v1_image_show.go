package core

import (
	"ImageGo/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"ImageGo/api/core/v1"
)

func (c *ControllerV1) ImageShow(ctx context.Context, req *v1.ImageShowReq) (res *v1.ImageShowRes, err error) {
	filePath := service.CoreImage().Download(ctx, &req.ImageDownloadInp)
	g.RequestFromCtx(ctx).Response.ServeFile(filePath)
	return nil, nil
}
