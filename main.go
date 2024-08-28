package main

import (
	"ImageGo/internal/global"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "ImageGo/internal/packed"

	_ "ImageGo/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"ImageGo/internal/cmd"
)

func main() {
	ctx := gctx.GetInitCtx()
	global.Init(ctx)
	cmd.Main.Run(ctx)
}
