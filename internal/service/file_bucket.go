package service

import (
	"context"

	"github.com/maestre3d/newton/internal/valueobject"
)

// FileBucket preferred static bucket interaction service
// (Amazon Web Services S3, GCP Firebase, ...)
type FileBucket interface {
	// Upload stores the given file into the preferred static bucket, returns file url
	Upload(context.Context, string, *valueobject.File) (string, error)
	// Delete permanently removes the specified file by its key
	Delete(context.Context, string) error
}
