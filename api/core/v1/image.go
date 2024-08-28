package v1

import (
	"ImageGo/internal/model/in/corein"
	"github.com/gogf/gf/v2/frame/g"
)

type ImageUploadReq struct {
	g.Meta `path:"/api/upload" tags:"Image" method:"post" summary:"Upload Image"`
	corein.ImageUploadInp
}

type ImageUploadRes struct {
	Name       string `json:"name"`
	Link       string `json:"link"`
	DeleteLink string `json:"delete_link"`
}

type ImageShowReq struct {
	g.Meta `path:"/image/{code}/{name}" method:"get" sm:"Show image"`
	corein.ImageDownloadInp
}

type ImageShowRes struct {
	g.Meta `mime:"image/jpeg"`
}

type ImageDeleteReq struct {
	g.Meta `path:"/delete/{code}" method:"get" sm:"Delete image"`
	corein.ImageDeleteInp
}

type ImageDeleteRes struct {
	g.Meta `mime:"text/html"`
}
