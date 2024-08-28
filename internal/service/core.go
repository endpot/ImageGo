// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"ImageGo/internal/model/entity"
	"ImageGo/internal/model/in/corein"
	"context"
)

type (
	ICoreImage interface {
		Download(ctx context.Context, in *corein.ImageDownloadInp) string
		Upload(ctx context.Context, in *corein.ImageUploadInp) (*entity.Image, error)
		Delete(ctx context.Context, in *corein.ImageDeleteInp) error
		CleanColdImages(ctx context.Context)
	}
	ICoreWeb interface {
		GetImageCount(ctx context.Context) int
		GetImageView(ctx context.Context) int
		GetNewImageCount(ctx context.Context) int
		GetUploaderCount(ctx context.Context) int
		GetRecentImageList(ctx context.Context, in *corein.WebHomeInp) ([]*entity.Image, error)
	}
)

var (
	localCoreImage ICoreImage
	localCoreWeb   ICoreWeb
)

func CoreImage() ICoreImage {
	if localCoreImage == nil {
		panic("implement not found for interface ICoreImage, forgot register?")
	}
	return localCoreImage
}

func RegisterCoreImage(i ICoreImage) {
	localCoreImage = i
}

func CoreWeb() ICoreWeb {
	if localCoreWeb == nil {
		panic("implement not found for interface ICoreWeb, forgot register?")
	}
	return localCoreWeb
}

func RegisterCoreWeb(i ICoreWeb) {
	localCoreWeb = i
}
