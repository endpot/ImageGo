package sys

import (
	"ImageGo/internal/model"
	"ImageGo/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type sSysConfig struct {
	s3Config *model.S3Config
}

func NewSysConfig() *sSysConfig {
	return &sSysConfig{}
}

func init() {
	service.RegisterSysConfig(NewSysConfig())
}

func (s *sSysConfig) InitConfig(ctx context.Context) {
	if err := s.LoadConfig(ctx); err != nil {
		g.Log().Fatalf(ctx, "InitConfig fail: %+v", err)
	}
}

func (s *sSysConfig) LoadConfig(ctx context.Context) error {
	if err := s.LoadS3(ctx); err != nil {
		return err
	}

	return nil
}

func (s *sSysConfig) LoadS3(ctx context.Context) error {
	var conf *model.S3Config
	err := g.Cfg().MustGet(ctx, "s3").Scan(&conf)
	if err != nil {
		return err
	}

	s.s3Config = conf
	return nil
}

func (s *sSysConfig) GetS3Config() *model.S3Config {
	return s.s3Config
}
