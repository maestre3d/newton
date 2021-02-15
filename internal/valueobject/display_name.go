package valueobject

import (
	"github.com/maestre3d/newton/internal/domain"
)

// DisplayName title or simplified name that will mostly appear in every UI interaction
type DisplayName string

const (
	displayNameMinLength = 2
	displayNameMaxLength = 256
)

// ErrDisplayNameOutOfRange the given DisplayName char length was out of the specified range
var ErrDisplayNameOutOfRange = domain.NewOutOfRange("display_name", displayNameMinLength, displayNameMaxLength)

// NewDisplayName creates and validates a DisplayName
func NewDisplayName(v string) (DisplayName, error) {
	name := DisplayName(v)
	if err := name.ensureLength(); err != nil {
		return "", err
	}
	return name, nil
}

func (n DisplayName) ensureLength() error {
	if length := len(n); length < displayNameMinLength || length > displayNameMaxLength {
		return ErrDisplayNameOutOfRange
	}
	return nil
}

// Value get the current value
func (n DisplayName) Value() string {
	return string(n)
}
