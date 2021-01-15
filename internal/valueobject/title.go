package valueobject

import "errors"

// Title Book display name
type Title struct {
	value string
}

const (
	titleMinLength = 1
	titleMaxLength = 256
)

// ErrTitleOutOfRange the given title character length is out of range
var ErrTitleOutOfRange = errors.New("title is out of range [1, 256)")

// NewTitle creates and validates a Title
func NewTitle(v string) (*Title, error) {
	t := new(Title)
	if err := t.ensureLength(v); err != nil {
		return nil, err
	}
	t.value = v
	return t, nil
}

func (t Title) ensureLength(v string) error {
	if len(v) < titleMinLength || len(v) > titleMaxLength {
		return ErrTitleOutOfRange
	}
	return nil
}
