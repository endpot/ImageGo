package core

import (
	"ImageGo/internal/service"
	"context"

	"ImageGo/api/core/v1"
)

func (c *ControllerV1) ImageDelete(ctx context.Context, req *v1.ImageDeleteReq) (res *v1.ImageDeleteRes, err error) {
	return nil, service.CoreImage().Delete(ctx, &req.ImageDeleteInp)
}
