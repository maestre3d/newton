package subscriber

import (
	"context"

	"github.com/maestre3d/newton/internal/event"
)

// Subscriber listens to a domain event in order to execute independent tasks asynchronously
type Subscriber interface {
	// SubscribedTo returns the event.DomainEvent this current Subscriber is subscribed to
	SubscribedTo() event.DomainEvent
	// On executes the required tasks based on the subscribed event.DomainEvent
	On(context.Context, event.DomainEvent) error
}
