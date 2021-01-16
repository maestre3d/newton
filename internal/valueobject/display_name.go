package valueobject

import "errors"

// DisplayName title or simplified name that will mostly appear in every UI interaction
type DisplayName struct {
	value string
}

const (
	displayNameMinLength = 2
	displayNameMaxLength = 256
)

// ErrDisplayNameOutOfRange the given DisplayName char length was out of the specified range
var ErrDisplayNameOutOfRange = errors.New("display name is out of range [2, 256)")

// NewDisplayName creates and validates a DisplayName
func NewDisplayName(v string) (*DisplayName, error) {
	if err := ensureDisplayNameLength(v); err != nil {
		return nil, err
	}
	return &DisplayName{value: v}, nil
}

func ensureDisplayNameLength(v string) error {
	if length := len(v); length < displayNameMinLength || length > displayNameMaxLength {
		return ErrDisplayNameOutOfRange
	}
	return nil
}

// Value get the current value
func (n DisplayName) Value() string {
	return n.value
}