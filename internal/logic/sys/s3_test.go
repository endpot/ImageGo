package sys

import (
	"context"
	"github.com/stretchr/testify/assert"
	"server/internal/service"
	"testing"
)

func TestSSysS3_UploadFile(t *testing.T) {
	ctx := context.Background()
	service.SysConfig().InitConfig(ctx)

	s := service.SysS3()
	err := s.UploadFile(ctx, "test", "test")
	assert.Nil(t, err)
}

func TestSSysS3_DownloadFile(t *testing.T) {
	ctx := context.Background()
	service.SysConfig().InitConfig(ctx)

	s := service.SysS3()
	err := s.DownloadFile(ctx, "test", "test")
	assert.Nil(t, err)
}
