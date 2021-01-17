package valueobject

import (
	"errors"
	"strings"
)

// BookResource an aggregate.Book external resource URL, references a file in format pdf
type BookResource string

const (
	bookResourceValidExtension          = "pdf"
	bookResourceValidExtensionUppercase = "PDF"
)

var (
	// ErrBookResourceInvalidURL the given book resource url is not compliant
	ErrBookResourceInvalidURL = errors.New("book resource is an invalid url")
	// ErrBookResourceInvalidExtension the given book resource url has a forbidden file format
	ErrBookResourceInvalidExtension = errors.New("book resource contains an invalid extension, expected [pdf]")
	// ErrBookResourceOutOfRange the given book resource url char length is out of range
	ErrBookResourceOutOfRange = errors.New("book resource is out of range [5, 2000)")
)

// NewBookResource creates and validates a BookResource
func NewBookResource(v string) (BookResource, error) {
	if err := ensureValidURL(v, ErrBookResourceInvalidURL); err != nil {
		return "", err
	} else if err := ensureURLLength(v, ErrBookResourceOutOfRange); err != nil {
		return "", err
	} else if err := ensureBookResourceValidExtension(v); err != nil {
		return "", err
	}
	return BookResource(v), nil
}

func ensureBookResourceValidExtension(v string) error {
	if !strings.HasSuffix(v, "."+bookResourceValidExtension) && !strings.HasSuffix(v,
		"."+bookResourceValidExtensionUppercase) {
		return ErrBookResourceInvalidExtension
	}
	return nil
}

// Value get the current value
func (r BookResource) Value() string {
	return string(r)
}
