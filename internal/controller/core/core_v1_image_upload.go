package core

import (
	"ImageGo/internal/service"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"

	"ImageGo/api/core/v1"
)

func (c *ControllerV1) ImageUpload(ctx context.Context, req *v1.ImageUploadReq) (res *v1.ImageUploadRes, err error) {
	inp := req.ImageUploadInp
	inp.Ip = g.RequestFromCtx(ctx).GetClientIp()

	uploadRes, err := service.CoreImage().Upload(ctx, &inp)
	if err != nil {
		return nil, err
	}

	hostWithSchema := fmt.Sprintf("%s://%s", g.RequestFromCtx(ctx).GetSchema(), g.RequestFromCtx(ctx).Host)
	res = &v1.ImageUploadRes{
		Name:       uploadRes.Name,
		Link:       fmt.Sprintf("%s/image/%s/%s", hostWithSchema, uploadRes.Code, uploadRes.Name),
		DeleteLink: fmt.Sprintf("%s/delete/%s", hostWithSchema, uploadRes.DeleteCode),
	}

	return res, nil
}
