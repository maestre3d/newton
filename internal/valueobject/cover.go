package valueobject

import (
	"errors"
	"strings"
)

// Cover an aggregate cover image, referenced as external resource with its url
type Cover struct {
	value string
}

const (
	coverValidExtension0          = ".jpg"
	coverValidExtension0Uppercase = ".JPG"
	coverValidExtension1          = ".jpeg"
	coverValidExtension1Uppercase = ".JPEG"
	coverValidExtension2          = ".png"
	coverValidExtension2Uppercase = ".PNG"
	coverValidExtension3          = ".webp"
	coverValidExtension3Uppercase = ".WEBP"
)

var (
	// ErrCoverInvalidURL the given cover url is not compliant
	ErrCoverInvalidURL = errors.New("cover is an invalid url")
	// ErrCoverInvalidExtension the given cover url has a forbidden file format
	ErrCoverInvalidExtension = errors.New("cover contains an invalid extension, expected [jpg, jpeg, png, webp]")
	// ErrCoverOutOfRange the given cover url char length is out of range
	ErrCoverOutOfRange = errors.New("cover is out of range [5, 2000)")
)

// NewCover creates and validates a Cover
func NewCover(v string) (*Cover, error) {
	if err := ensureValidURL(v, ErrCoverInvalidURL); err != nil {
		return nil, err
	} else if err := ensureURLLength(v, ErrCoverOutOfRange); err != nil {
		return nil, err
	} else if err := ensureCoverValidExtension(v); err != nil {
		return nil, err
	}
	return &Cover{value: v}, nil
}

func ensureCoverValidExtension(v string) error {
	isInvalidExtension := !strings.HasSuffix(v, coverValidExtension0) &&
		!strings.HasSuffix(v, coverValidExtension0Uppercase) && !strings.HasSuffix(v, coverValidExtension1) &&
		!strings.HasSuffix(v, coverValidExtension1Uppercase) && !strings.HasSuffix(v, coverValidExtension2) &&
		!strings.HasSuffix(v, coverValidExtension2Uppercase) && !strings.HasSuffix(v, coverValidExtension3) &&
		!strings.HasSuffix(v, coverValidExtension3Uppercase)
	if isInvalidExtension {
		return ErrCoverInvalidExtension
	}
	return nil
}

// Value get the current value
func (c Cover) Value() string {
	return c.value
}
