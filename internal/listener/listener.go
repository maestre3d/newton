package listener

import (
	"context"

	"github.com/maestre3d/newton/internal/event"
)

// Listener subscribes to a domain event in order to execute independent tasks asynchronously
type Listener interface {
	// SubscribedTo returns the event.DomainEvent this specific listener is subscribed to
	SubscribedTo() event.DomainEvent
	// On executes the required tasks while receiving the subscribed event.DomainEvent as argument
	On(context.Context, event.DomainEvent) error
}
