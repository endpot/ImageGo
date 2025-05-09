package core

import (
	"context"

	"ImageGo/api/core/v1"
)

func (c *ControllerV1) Health(ctx context.Context, req *v1.HealthReq) (res *v1.HealthRes, err error) {
	return nil, nil
}
