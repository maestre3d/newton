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
	if err := ensureDisplayNameLength(v); err != nil {
		return "", err
	}
	return DisplayName(v), nil
}

func ensureDisplayNameLength(v string) error {
	if length := len(v); length < displayNameMinLength || length > displayNameMaxLength {
		return ErrDisplayNameOutOfRange
	}
	return nil
}

// Value get the current value
func (n DisplayName) Value() string {
	return string(n)
}
