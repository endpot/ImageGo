package core

import (
	"ImageGo/internal/dao"
	"ImageGo/internal/model/entity"
	"ImageGo/internal/model/in/corein"
	"ImageGo/internal/service"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"slices"
	"time"
)

type sCoreImage struct{}

func NewCoreImage() *sCoreImage {
	return &sCoreImage{}
}

func init() {
	service.RegisterCoreImage(NewCoreImage())
}

func (s *sCoreImage) Download(ctx context.Context, in *corein.ImageDownloadInp) string {
	fallbackImagePath := gfile.Join(gfile.SelfDir(), "resource", "public", "resource", "image", "404.jpg")

	var img *entity.Image
	columns := dao.Image.Columns()
	if err := dao.Image.Ctx(ctx).Where(columns.Code, in.Code).Scan(&img); err != nil || img == nil {
		g.Log().Errorf(ctx, "get images failed, code: %s", in.Code)
		return fallbackImagePath
	}

	_, err := dao.Image.Ctx(ctx).Where(columns.Id, img.Id).Increment(columns.Views, 1)
	if err != nil {
		g.Log().Errorf(ctx, "increment views failed, code: %s", in.Code)
	}

	filename := img.SaveName
	localPath := s.getLocalPath(ctx, filename)
	if gfile.Exists(localPath) {
		return localPath
	}

	if err := s.fetchImageFromRemote(ctx, filename); err != nil {
		g.Log().Errorf(ctx, "fetch image failed, code: %s", in.Code)
		return fallbackImagePath
	}

	return localPath
}

func (s *sCoreImage) Upload(ctx context.Context, in *corein.ImageUploadInp) (*entity.Image, error) {
	if !s.validateImageFile(ctx, in.Image) {
		return nil, gerror.New("invalid image file")
	}

	originalName := in.Image.Filename
	originalExtension := gfile.ExtName(originalName)
	imageSize := in.Image.Size
	nsfw := gconv.Int(in.Nsfw)

	// code and deleteCode
	code, err := s.generateUniqueCode(ctx)
	if err != nil {
		return nil, err
	}
	deleteCode := code + s.generateCode(ctx)

	// image dimension
	imageWidth, imageHeight, err := s.getImageDimension(ctx, in.Image)
	if err != nil {
		return nil, err
	}

	// image fingerprint
	fingerprint, err := s.getImageFingerprint(ctx, in.Image)
	if err != nil {
		return nil, err
	}

	var exitsImage *entity.Image
	saveName := fmt.Sprintf("%s.%s", code, originalExtension)
	err = dao.Image.Ctx(ctx).Where(dao.Image.Columns().Fingerprint, fingerprint).Scan(&exitsImage)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	} else if err == nil && exitsImage != nil {
		saveName = exitsImage.SaveName
	}

	// save to local
	in.Image.Filename = saveName
	_, err = in.Image.Save(gfile.Dir(s.getLocalPath(ctx, saveName)))
	if err != nil {
		return nil, err
	}

	// save to remote
	if err = s.pushImageToRemote(ctx, saveName); err != nil {
		return nil, err
	}

	imageData := &entity.Image{
		Code:        code,
		DeleteCode:  deleteCode,
		Name:        originalName,
		Ext:         originalExtension,
		Width:       imageWidth,
		Height:      imageHeight,
		Nsfw:        nsfw,
		UploaderIp:  in.Ip,
		Fingerprint: fingerprint,
		SaveName:    saveName,
		Size:        imageSize,
		Views:       0,
		CreatedAt:   nil,
		UpdatedAt:   nil,
	}

	if _, err = dao.Image.Ctx(ctx).Data(imageData).Insert(); err != nil {
		return nil, err
	}

	return imageData, nil
}

func (s *sCoreImage) Delete(ctx context.Context, in *corein.ImageDeleteInp) error {
	columns := dao.Image.Columns()

	var img *entity.Image
	if err := dao.Image.Ctx(ctx).Where(columns.DeleteCode, in.Code).Scan(&img); err != nil {
		g.Log().Errorf(ctx, "get images failed, delete_code: %s", in.Code)
		return nil
	}

	if img == nil {
		g.Log().Debug(ctx, "image not exist, delete_code: %s", in.Code)
		return nil
	}

	count, err := dao.Image.Ctx(ctx).Where(columns.SaveName, img.SaveName).Count()
	if err != nil {
		g.Log().Errorf(ctx, "count duplicate images failed, delete_code: %s", in.Code)
		return err
	}

	_, err = dao.Image.Ctx(ctx).Where(columns.Id, img.Id).Delete()
	if err != nil {
		g.Log().Errorf(ctx, "delete images failed, delete_code: %s", in.Code)
		return err
	}

	if count <= 1 {
		if err = s.deleteFromLocal(ctx, img.SaveName); err != nil {
			g.Log().Debug(ctx, "delete local image failed, delete_code: %s", in.Code)
		}

		if err = s.deleteFromRemote(ctx, img.SaveName); err != nil {
			g.Log().Debug(ctx, "delete remote image failed, delete_code: %s", in.Code)
		}
	}

	return nil
}

func (s *sCoreImage) CleanColdImages(ctx context.Context) {
	g.Log().Infof(ctx, "start to clean cold images")

	lastWeek := time.Now().Add(-7 * 24 * time.Hour)

	var hotImages []*entity.Image
	if err := dao.Image.Ctx(ctx).WhereGTE(dao.Image.Columns().UpdatedAt, lastWeek).Scan(&hotImages); err != nil {
		g.Log().Errorf(ctx, "get hot images failed, lastWeek: %s", lastWeek)
		return
	}

	safeImageFileMap := make(map[string]bool)
	for _, img := range hotImages {
		safeImageFileMap[img.SaveName] = true
	}

	cacheImageDir := gfile.Dir(s.getLocalPath(ctx, "empty"))
	localFiles, err := gfile.ScanDir(cacheImageDir, "*.*", false)
	if err != nil {
		g.Log().Errorf(ctx, "get local files failed")
		return
	}

	for _, file := range localFiles {
		filename := gfile.Basename(file)
		if !safeImageFileMap[filename] {
			if err = s.deleteFromLocal(ctx, file); err != nil {
				g.Log().Errorf(ctx, "delete local image failed, file: %s", file)
			}
		}
	}
}

func (s *sCoreImage) validateImageFile(_ context.Context, file *ghttp.UploadFile) bool {
	return file.Size <= maximumImageSize && slices.Contains(validImageExtensions, gfile.ExtName(file.Filename))
}

func (s *sCoreImage) generateCode(_ context.Context) string {
	return grand.Str(codeString, codeLength)
}

func (s *sCoreImage) generateUniqueCode(ctx context.Context) (string, error) {
	for {
		code := s.generateCode(ctx)

		if count, err := dao.Image.Ctx(ctx).Where(dao.Image.Columns().Code, code).Count(); err != nil {
			return "", err
		} else if count == 0 {
			return code, nil
		}
	}
}

func (s *sCoreImage) getLocalPath(_ context.Context, filename string) string {
	return gfile.Join(gfile.SelfDir(), "storage", "cache", "images", filename)
}

func (s *sCoreImage) getRemotePath(_ context.Context, filename string) string {
	return gfile.Join("upload", "images", filename)
}

func (s *sCoreImage) getImageDimension(_ context.Context, file *ghttp.UploadFile) (int, int, error) {
	f, err := file.Open()
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()

	im, _, err := image.DecodeConfig(f)
	if err != nil {
		return 0, 0, err
	}

	return im.Width, im.Height, nil
}

func (s *sCoreImage) getImageFingerprint(_ context.Context, file *ghttp.UploadFile) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	fileBytes, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	return gmd5.EncryptBytes(fileBytes)
}

func (s *sCoreImage) fetchImageFromRemote(ctx context.Context, filename string) error {
	localPath := s.getLocalPath(ctx, filename)
	remotePath := s.getRemotePath(ctx, filename)

	return service.SysS3().DownloadFile(ctx, remotePath, localPath)
}

func (s *sCoreImage) pushImageToRemote(ctx context.Context, filename string) error {
	localPath := s.getLocalPath(ctx, filename)
	remotePath := s.getRemotePath(ctx, filename)

	return service.SysS3().UploadFile(ctx, remotePath, localPath)
}

func (s *sCoreImage) deleteFromLocal(ctx context.Context, filename string) error {
	return gfile.Remove(s.getLocalPath(ctx, filename))
}

func (s *sCoreImage) deleteFromRemote(ctx context.Context, filename string) error {
	remotePath := s.getRemotePath(ctx, filename)

	return service.SysS3().DeleteFile(ctx, remotePath)
}
