// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package core

import (
	"context"

	"ImageGo/api/core/v1"
)

type ICoreV1 interface {
	ImageUpload(ctx context.Context, req *v1.ImageUploadReq) (res *v1.ImageUploadRes, err error)
	ImageShow(ctx context.Context, req *v1.ImageShowReq) (res *v1.ImageShowRes, err error)
	ImageDelete(ctx context.Context, req *v1.ImageDeleteReq) (res *v1.ImageDeleteRes, err error)
	WebHome(ctx context.Context, req *v1.WebHomeReq) (res *v1.WebHomeRes, err error)
}
