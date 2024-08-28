// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"ImageGo/internal/model"
	"context"
	"io"
)

type (
	ISysConfig interface {
		InitConfig(ctx context.Context)
		LoadConfig(ctx context.Context) error
		LoadS3(ctx context.Context) error
		GetS3Config() *model.S3Config
	}
	ISysCron interface {
		InitCron(ctx context.Context)
	}
	ISysS3 interface {
		UploadFileFromReader(ctx context.Context, key string, r io.ReadSeeker) error
		UploadFile(ctx context.Context, key string, path string) error
		DownloadFile(ctx context.Context, key string, path string) error
		DeleteFile(ctx context.Context, key string) error
	}
)

var (
	localSysConfig ISysConfig
	localSysCron   ISysCron
	localSysS3     ISysS3
)

func SysConfig() ISysConfig {
	if localSysConfig == nil {
		panic("implement not found for interface ISysConfig, forgot register?")
	}
	return localSysConfig
}

func RegisterSysConfig(i ISysConfig) {
	localSysConfig = i
}

func SysCron() ISysCron {
	if localSysCron == nil {
		panic("implement not found for interface ISysCron, forgot register?")
	}
	return localSysCron
}

func RegisterSysCron(i ISysCron) {
	localSysCron = i
}

func SysS3() ISysS3 {
	if localSysS3 == nil {
		panic("implement not found for interface ISysS3, forgot register?")
	}
	return localSysS3
}

func RegisterSysS3(i ISysS3) {
	localSysS3 = i
}
