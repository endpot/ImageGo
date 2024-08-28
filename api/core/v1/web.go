package v1

import (
	"ImageGo/internal/model/in/corein"
	"github.com/gogf/gf/v2/frame/g"
)

type WebHomeReq struct {
	g.Meta `path:"/" tags:"Web" method:"get" summary:"Homepage"`
	corein.WebHomeInp
}

type WebHomeRes struct {
	g.Meta `mime:"text/html"`
}
