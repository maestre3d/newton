package valueobject

import (
	"errors"
)

// BookID aggregate.Book unique identifier
//	Prefer nano ids over UUIDs for performance purposes
type BookID struct {
	value string
}

const (
	bookIDMinLength = 16
	bookIDMaxLength = 128
)

// ErrBookIDOutOfRange the given book id character length is out of range, use gonanoid.New(16)
var ErrBookIDOutOfRange = errors.New("book id is out of range [16, 128)")

// NewBookID creates and validates a BookID
func NewBookID(v string) (*BookID, error) {
	id := new(BookID)
	if err := id.ensureLength(v); err != nil {
		return nil, err
	}
	id.value = v
	return id, nil
}

func (i BookID) ensureLength(v string) error {
	if len(v) < bookIDMinLength || len(v) > bookIDMaxLength {
		return ErrBookIDOutOfRange
	}

	return nil
}
