package global

import (
	"ImageGo/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func Init(ctx context.Context) {
	if err := gtime.SetTimeZone("Asia/Shanghai"); err != nil {
		g.Log().Fatalf(ctx, "SetTimeZone fail: %+v", err)
	}

	service.SysConfig().InitConfig(ctx)

	service.SysCron().InitCron(ctx)
}
