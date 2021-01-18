package valueobject

import (
	"github.com/maestre3d/newton/internal/domain"
)

// Username aggregate.User unique name identifier, it is immutable if used as single username
//	Could be also used as preferred username.
//	Hence, immutability wont be available when used as is
type Username string

const (
	usernameMinLength = 1
	usernameMaxLength = 128
)

// ErrUsernameOutOfRange the given username char length is out of range
var ErrUsernameOutOfRange = domain.NewOutOfRange("username", usernameMinLength, usernameMaxLength)

func NewUsername(v string) (Username, error) {
	if err := ensureUsernameLength(v); err != nil {
		return "", err
	}
	return Username(v), nil
}

func ensureUsernameLength(v string) error {
	if len(v) < usernameMinLength || len(v) > usernameMaxLength {
		return ErrUsernameOutOfRange
	}
	return nil
}

// Value get the current value
func (u Username) Value() string {
	return string(u)
}
