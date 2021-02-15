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

// NewUsername creates and validates a Username
func NewUsername(v string) (Username, error) {
	u := Username(v)
	if err := u.ensureLength(); err != nil {
		return "", err
	}
	return u, nil
}

func (u Username) ensureLength() error {
	if len(u) < usernameMinLength || len(u) > usernameMaxLength {
		return ErrUsernameOutOfRange
	}
	return nil
}

// Value get the current value
func (u Username) Value() string {
	return string(u)
}
