package valueobject

import (
	"github.com/maestre3d/newton/internal/domain"
)

// AuthorID aggregate.Author unique identifier
//	Prefer nano ids over UUIDs for performance purposes
type AuthorID string

const (
	authorIDMinLength = 16
	authorIDMaxLength = 128
)

// ErrAuthorIDOutOfRange the given author id character length is out of range, use gonanoid.New(16)
var ErrAuthorIDOutOfRange = domain.NewOutOfRange("author_id", authorIDMinLength, authorIDMaxLength)

// NewAuthorID creates and validates an AuthorID
func NewAuthorID(v string) (AuthorID, error) {
	if err := ensureAuthorIDLength(v); err != nil {
		return "", err
	}
	return AuthorID(v), nil
}

func ensureAuthorIDLength(v string) error {
	if len(v) < authorIDMinLength || len(v) > authorIDMaxLength {
		return ErrAuthorIDOutOfRange
	}

	return nil
}

// Value get the current value
func (i AuthorID) Value() string {
	return string(i)
}
