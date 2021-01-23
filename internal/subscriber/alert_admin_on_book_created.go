package subscriber

import (
	"context"

	"github.com/maestre3d/newton/internal/application"
	"github.com/maestre3d/newton/internal/event"
)

// AlertAdminOnBookCreated notifies admin through either email or push notification when an
// aggregate.Book was created
func AlertAdminOnBookCreated(app *application.Author, ctx context.Context, ev event.BookCreated) error {
	return nil
}
