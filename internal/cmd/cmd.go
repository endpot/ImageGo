package cmd

import (
	"ImageGo/internal/controller/core"
	"context"
	"github.com/gogf/gf/v2/os/gfile"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.SetServerRoot(gfile.Join(gfile.SelfDir(), "storage", "static"))
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					core.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
