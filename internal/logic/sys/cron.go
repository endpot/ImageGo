package sys

import (
	"ImageGo/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
)

type sSysCron struct{}

func NewSysCron() *sSysCron {
	return &sSysCron{}
}

func init() {
	service.RegisterSysCron(NewSysCron())
}

func (s *sSysCron) InitCron(ctx context.Context) {
	_, err := gcron.AddSingleton(ctx, "0 0 0 * * *", service.CoreImage().CleanColdImages, "CleanColdImages")
	if err != nil {
		g.Log().Fatal(ctx, "InitCron fail: %+v", err)
	}
}
