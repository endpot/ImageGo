package sys

import (
	"ImageGo/internal/service"
	"context"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gogf/gf/v2/os/gfile"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type sSysS3 struct {
	//
}

func NewSysS3() *sSysS3 {
	return &sSysS3{}
}

func init() {
	service.RegisterSysS3(NewSysS3())
}

func (s *sSysS3) UploadFileFromReader(ctx context.Context, key string, r io.ReadSeeker) error {
	client, err := s.newS3Client(ctx)
	if err != nil {
		return err
	}

	_, err = client.PutObject(&s3.PutObjectInput{
		Body:   r,
		Bucket: s.getS3Bucket(ctx),
		Key:    aws.String(key),
	})

	return err
}

func (s *sSysS3) UploadFile(ctx context.Context, key string, path string) error {
	client, err := s.newS3Client(ctx)
	if err != nil {
		return err
	}

	f, err := gfile.Open(path)
	if err != nil {
		return err
	}

	_, err = client.PutObject(&s3.PutObjectInput{
		Body:   f,
		Bucket: s.getS3Bucket(ctx),
		Key:    aws.String(key),
	})

	return err
}

func (s *sSysS3) DownloadFile(ctx context.Context, key string, path string) error {
	output, err := s.getObject(ctx, key)
	if err != nil {
		return err
	}

	f, err := gfile.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, output.Body)

	return err
}

func (s *sSysS3) DeleteFile(ctx context.Context, key string) error {
	client, err := s.newS3Client(ctx)
	if err != nil {
		return err
	}

	_, err = client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: s.getS3Bucket(ctx),
		Key:    aws.String(key),
	})

	return err
}

func (s *sSysS3) getObject(ctx context.Context, key string) (*s3.GetObjectOutput, error) {
	client, err := s.newS3Client(ctx)
	if err != nil {
		return nil, err
	}

	output, err := client.GetObject(&s3.GetObjectInput{
		Bucket: s.getS3Bucket(ctx),
		Key:    aws.String(key),
	})

	return output, err
}

func (s *sSysS3) newS3Client(_ context.Context) (*s3.S3, error) {
	s3Config := service.SysConfig().GetS3Config()

	cfg := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(s3Config.Id, s3Config.Key, ""),
		Endpoint:         aws.String(s3Config.Endpoint),
		Region:           aws.String(s3Config.Region),
		S3ForcePathStyle: aws.Bool(true),
	}

	sess, err := session.NewSession(cfg)
	if err != nil {
		return nil, err
	}

	return s3.New(sess, cfg), nil
}

func (s *sSysS3) getS3Bucket(_ context.Context) *string {
	s3Config := service.SysConfig().GetS3Config()

	return aws.String(s3Config.Bucket)
}
