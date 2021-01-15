package valueobject

import "errors"

// Username aggregate.User unique name identifier, it is immutable if used as single username
//	Could be also used as preferred username.
//	Hence, immutability wont be available when used as is
type Username struct {
	value string
}

const (
	usernameMinLength = 1
	usernameMaxLength = 128
)

// ErrUsernameOutOfRange the given username char length is out of range
var ErrUsernameOutOfRange = errors.New("username is out of range [1, 128)")

func NewUsername(v string) (*Username, error) {
	if err := ensureUsernameLength(v); err != nil {
		return nil, err
	}
	return &Username{value: v}, nil
}

func ensureUsernameLength(v string) error {
	if len(v) < usernameMinLength || len(v) > usernameMaxLength {
		return ErrUsernameOutOfRange
	}
	return nil
}
