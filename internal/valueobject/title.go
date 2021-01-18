package valueobject

import (
	"github.com/maestre3d/newton/internal/domain"
)

// Title Book display name
type Title string

const (
	titleMinLength = 1
	titleMaxLength = 256
)

// ErrTitleOutOfRange the given title character length is out of range
var ErrTitleOutOfRange = domain.NewOutOfRange("title", titleMinLength, titleMaxLength)

// NewTitle creates and validates a Title
func NewTitle(v string) (Title, error) {
	if err := ensureTitleLength(v); err != nil {
		return "", err
	}
	return Title(v), nil
}

func ensureTitleLength(v string) error {
	if len(v) < titleMinLength || len(v) > titleMaxLength {
		return ErrTitleOutOfRange
	}
	return nil
}

// Value get the current value
func (t Title) Value() string {
	return string(t)
}
