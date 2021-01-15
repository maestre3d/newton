package service

import (
	"context"

	"github.com/maestre3d/newton/internal/aggregate"
)

// BookNotifier publishes or alerts somebody and/or something (e.g. admin, external system)
type BookNotifier interface {
	Notify(ctx context.Context, b aggregate.Book) error
}
