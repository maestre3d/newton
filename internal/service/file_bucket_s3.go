package service

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/maestre3d/newton/internal/domain"
	"github.com/maestre3d/newton/internal/infrastructure"
	"github.com/maestre3d/newton/internal/valueobject"
)

// FileBucketS3 FileBucket Amazon Web Services S3 (Simple Storage Service) concrete service implementation
type FileBucketS3 struct {
	cfg infrastructure.Configuration
	c   *s3.Client
	mu  sync.Mutex
}

// NewFileBucketS3 allocates a new FileBucketS3 implementation
func NewFileBucketS3(cfg infrastructure.Configuration, client *s3.Client) *FileBucketS3 {
	return &FileBucketS3{
		cfg: cfg,
		c:   client,
		mu:  sync.Mutex{},
	}
}

// Upload stores the given file into the preferred static bucket, returns file url
func (s *FileBucketS3) Upload(ctx context.Context, key string, file *valueobject.File) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	uploader := manager.NewUploader(s.c, func(u *manager.Uploader) {
		u.PartSize = 10 * domain.MebiByte
	})
	_, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:        aws.String(s.cfg.BucketName),
		ACL:           types.ObjectCannedACLPrivate,
		Key:           aws.String(key),
		Body:          file.File,
		StorageClass:  types.StorageClassIntelligentTiering,
		ContentType:   aws.String(s.getMIMEType(file.Extension)),
		ContentLength: file.Size,
	})
	if err != nil {
		return "", err
	}
	return s.cfg.StaticCDN + "/" + key, nil
}

// Delete permanently removes the specified file by its key
func (s *FileBucketS3) Delete(ctx context.Context, key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, err := s.c.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.cfg.BucketName),
		Key:    aws.String(key),
	})
	return err
}

// getMIMEType returns a valid MIME Type from the given extension (must be lowercase)
//	For further reference, take a read of the article https://www.iana.org/assignments/media-types/media-types.xhtml#image
func (s *FileBucketS3) getMIMEType(ext string) string {
	switch ext {
	case "jpg":
		return "image/jpeg"
	case "jpeg":
		return "image/jpeg"
	case "png":
		return "image/png"
	case "svg":
		return "image/svg+xml"
	case "webp":
		return "image/webp"
	case "pdf":
		return "application/pdf"
	default:
		return "application/octet-stream"
	}
}
