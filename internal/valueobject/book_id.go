package valueobject

import (
	"github.com/maestre3d/newton/internal/domain"
)

// BookID aggregate.Book unique identifier
//	Prefer nano ids over UUIDs for performance purposes
type BookID string

const (
	bookIDMinLength = 16
	bookIDMaxLength = 128
)

// ErrBookIDOutOfRange the given book id character length is out of range, use gonanoid.New(16)
var ErrBookIDOutOfRange = domain.NewOutOfRange("book_id", bookIDMinLength, bookIDMaxLength)

// NewBookID creates and validates a BookID
func NewBookID(v string) (BookID, error) {
	id := BookID(v)
	if err := id.ensureLength(); err != nil {
		return "", err
	}
	return id, nil
}

func (i BookID) ensureLength() error {
	if length := len(i); length < bookIDMinLength || length > bookIDMaxLength {
		return ErrBookIDOutOfRange
	}

	return nil
}

// Value get the current value
func (i BookID) Value() string {
	return string(i)
}
