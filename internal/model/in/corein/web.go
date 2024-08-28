package corein

import "github.com/gogf/gf/v2/net/ghttp"

type ImageDownloadInp struct {
	Code string `in:"path" v:"required"`
}

type ImageUploadInp struct {
	Ip    string            `json:"-"`
	Image *ghttp.UploadFile `type:"file" v:"required"`
	Nsfw  bool
}

type ImageDeleteInp struct {
	Code string `in:"path" v:"required"`
}
