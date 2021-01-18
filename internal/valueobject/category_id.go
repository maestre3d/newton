package valueobject

import (
	"github.com/maestre3d/newton/internal/domain"
)

// CategoryID aggregate.Book unique identifier
//	Prefer nano ids over UUIDs for performance purposes
type CategoryID string

const (
	categoryIDMinLength = 16
	categoryIDMaxLength = 128
)

// ErrCategoryIDOutOfRange the given category id character length is out of range, use gonanoid.New(16)
var ErrCategoryIDOutOfRange = domain.NewOutOfRange("category_id", categoryIDMinLength, categoryIDMaxLength)

// NewCategoryID creates and validates a CategoryID
func NewCategoryID(v string) (CategoryID, error) {
	if err := ensureCategoryIDLength(v); err != nil {
		return "", err
	}
	return CategoryID(v), nil
}

func ensureCategoryIDLength(v string) error {
	if len(v) < categoryIDMinLength || len(v) > categoryIDMaxLength {
		return ErrCategoryIDOutOfRange
	}

	return nil
}

// Value get the current value
func (i CategoryID) Value() string {
	return string(i)
}
