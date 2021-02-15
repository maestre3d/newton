package valueobject

import (
	"strings"

	"github.com/maestre3d/newton/internal/domain"
)

// BookResource an aggregate.Book external resource URL, references a file in format pdf
type BookResource string

const (
	bookResourceValidExtension          = "pdf"
	bookResourceValidExtensionUppercase = "PDF"
)

var (
	// ErrBookResourceInvalidURL the given book resource url is not compliant
	ErrBookResourceInvalidURL = domain.NewInvalidFormat("book_resource", "url")
	// ErrBookResourceInvalidExtension the given book resource url has a forbidden file format
	ErrBookResourceInvalidExtension = domain.NewInvalidFormat("book_resource", "pdf")
	// ErrBookResourceOutOfRange the given book resource url char length is out of range
	ErrBookResourceOutOfRange = domain.NewOutOfRange("book_resource", urlMinLength, urlMaxLength)
)

// NewBookResource creates and validates a BookResource
func NewBookResource(v string) (BookResource, error) {
	r := BookResource(v)
	if err := ensureValidURL(v, ErrBookResourceInvalidURL); err != nil {
		return "", err
	} else if err := ensureURLLength(v, ErrBookResourceOutOfRange); err != nil {
		return "", err
	} else if err := r.ensureValidExtension(); err != nil {
		return "", err
	}
	return r, nil
}

func (r BookResource) ensureValidExtension() error {
	if !strings.HasSuffix(string(r), "."+bookResourceValidExtension) && !strings.HasSuffix(string(r),
		"."+bookResourceValidExtensionUppercase) {
		return ErrBookResourceInvalidExtension
	}
	return nil
}

// Value get the current value
func (r BookResource) Value() string {
	return string(r)
}
