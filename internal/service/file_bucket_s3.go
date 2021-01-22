package service

import (
	"bytes"
	"context"
	"net/http"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
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

	buffer := make([]byte, file.Size)
	if _, err := file.File.Read(buffer); err != nil {
		return "", err
	}
	_, err := s.c.PutObject(ctx, &s3.PutObjectInput{
		Bucket:        aws.String(s.cfg.BucketName),
		ACL:           types.ObjectCannedACLPrivate,
		Key:           aws.String(key),
		Body:          bytes.NewReader(buffer),
		StorageClass:  types.StorageClassIntelligentTiering,
		ContentType:   aws.String(http.DetectContentType(buffer)),
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

func (s *FileBucketS3) getMIMEType(key string) string {
	switch {
	case strings.HasSuffix(key, ".jpg") || strings.HasSuffix(key, ".jpeg"):
		return "Content-Type: image/jpeg"
	case strings.HasSuffix(key, ".png"):
		return "image/png"
	case strings.HasSuffix(key, ".svg"):
		return "image/svg+xml"
	case strings.HasSuffix(key, ".webp"):
		return "image/webp"
	case strings.HasSuffix(key, ".pdf"):
		return "application/pdf"
	default:
		return ""
	}
}
