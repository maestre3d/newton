package repository

import (
	"context"

	"github.com/maestre3d/newton/internal/aggregate"
	"github.com/maestre3d/newton/internal/valueobject"
)

// Book handles all persistence interactions
type Book interface {
	// Save stores, update or deletes the given record
	Save(context.Context, aggregate.Book) error
	// Get returns the current aggregate if available, returns nil if not found
	Get(context.Context, valueobject.BookID) (*aggregate.Book, error)
	// Search returns a list of the current aggregate filtering and ordering by the given criteria, returns the
	// next page token as second argument and returns nil if not found
	Search(context.Context, Criteria) ([]*aggregate.Book, string, error)
}
