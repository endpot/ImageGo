package core

import (
	"ImageGo/internal/dao"
	"ImageGo/internal/model/entity"
	"ImageGo/internal/model/in/corein"
	"ImageGo/internal/service"
	"context"
	"time"
)

type sCoreWeb struct{}

func NewCoreWeb() *sCoreWeb {
	return &sCoreWeb{}
}

func init() {
	service.RegisterCoreWeb(NewCoreWeb())
}

func (s *sCoreWeb) GetImageCount(ctx context.Context) int {
	count, _ := dao.Image.Ctx(ctx).Count()

	return count
}

func (s *sCoreWeb) GetImageView(ctx context.Context) int {
	count, _ := dao.Image.Ctx(ctx).Sum(dao.Image.Columns().Views)

	return int(count)
}

func (s *sCoreWeb) GetNewImageCount(ctx context.Context) int {
	count, _ := dao.Image.Ctx(ctx).WhereGTE(dao.Image.Columns().CreatedAt, time.Now().Truncate(24*time.Hour)).Count()

	return count
}

func (s *sCoreWeb) GetUploaderCount(ctx context.Context) int {
	count, _ := dao.Image.Ctx(ctx).Distinct().CountColumn(dao.Image.Columns().UploaderIp)

	return count
}

func (s *sCoreWeb) GetRecentImageList(ctx context.Context, in *corein.WebHomeInp) ([]*entity.Image, error) {
	columns := dao.Image.Columns()

	var imageList []*entity.Image
	err := dao.Image.Ctx(ctx).Where(columns.Nsfw, in.Nsfw).OrderDesc(columns.Id).Limit(randomPictureCount).Scan(&imageList)
	if err != nil {
		return nil, err
	}

	return imageList, nil
}
