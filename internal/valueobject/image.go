package valueobject

import (
	"strings"

	"github.com/maestre3d/newton/internal/domain"
)

// Image an aggregate Image image, referenced as external resource with its url
type Image string

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
	ErrImageInvalidURL = domain.NewInvalidFormat("image", "url")
	// ErrImageInvalidExtension the given Image url has a forbidden file format
	ErrImageInvalidExtension = domain.NewInvalidFormat("image", "jpg", "jpeg", "png", "webp")
	// ErrImageOutOfRange the given Image url char length is out of range
	ErrImageOutOfRange = domain.NewOutOfRange("image", urlMinLength, urlMaxLength)
)

// NewImage creates and validates an Image
func NewImage(v string) (Image, error) {
	if err := ensureValidURL(v, ErrImageInvalidURL); err != nil {
		return "", err
	} else if err := ensureURLLength(v, ErrImageOutOfRange); err != nil {
		return "", err
	} else if err := ensureImageValidExtension(v); err != nil {
		return "", err
	}
	return Image(v), nil
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
func (i Image) Value() string {
	return string(i)
}
