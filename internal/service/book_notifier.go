package service

import (
	"context"

	"github.com/maestre3d/newton/internal/aggregate"
)

// BookNotifier publishes or alerts somebody and/or something (e.g. admin, external system)
//	Infrastructure service
type BookNotifier interface {
	// Notify push notification when an aggregate.Book state was mutated
	Notify(context.Context, aggregate.Book) error
}
