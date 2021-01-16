package valueobject

import (
	"errors"
	"strings"
)

// Image an aggregate Image image, referenced as external resource with its url
type Image struct {
	value string
}

const (
	ImageValidExtension0          = ".jpg"
	ImageValidExtension0Uppercase = ".JPG"
	ImageValidExtension1          = ".jpeg"
	ImageValidExtension1Uppercase = ".JPEG"
	ImageValidExtension2          = ".png"
	ImageValidExtension2Uppercase = ".PNG"
	ImageValidExtension3          = ".webp"
	ImageValidExtension3Uppercase = ".WEBP"
)

var (
	// ErrImageInvalidURL the given Image url is not compliant
	ErrImageInvalidURL = errors.New("image is an invalid url")
	// ErrImageInvalidExtension the given Image url has a forbidden file format
	ErrImageInvalidExtension = errors.New("image contains an invalid extension, expected [jpg, jpeg, png, webp]")
	// ErrImageOutOfRange the given Image url char length is out of range
	ErrImageOutOfRange = errors.New("image is out of range [5, 2000)")
)

// NewImage creates and validates an Image
func NewImage(v string) (*Image, error) {
	if err := ensureValidURL(v, ErrImageInvalidURL); err != nil {
		return nil, err
	} else if err := ensureURLLength(v, ErrImageOutOfRange); err != nil {
		return nil, err
	} else if err := ensureImageValidExtension(v); err != nil {
		return nil, err
	}
	return &Image{value: v}, nil
}

func ensureImageValidExtension(v string) error {
	isInvalidExtension := !strings.HasSuffix(v, ImageValidExtension0) &&
		!strings.HasSuffix(v, ImageValidExtension0Uppercase) && !strings.HasSuffix(v, ImageValidExtension1) &&
		!strings.HasSuffix(v, ImageValidExtension1Uppercase) && !strings.HasSuffix(v, ImageValidExtension2) &&
		!strings.HasSuffix(v, ImageValidExtension2Uppercase) && !strings.HasSuffix(v, ImageValidExtension3) &&
		!strings.HasSuffix(v, ImageValidExtension3Uppercase)
	if isInvalidExtension {
		return ErrImageInvalidExtension
	}
	return nil
}

// Value get the current value
func (c Image) Value() string {
	return c.value
}
