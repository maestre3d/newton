package subscriber

import (
	"context"

	"github.com/maestre3d/newton/internal/event"
	"go.uber.org/zap"
)

// Subscriber listens to a domain event in order to execute independent tasks asynchronously
type Subscriber interface {
	// SubscribedTo returns the event.DomainEvent this current Subscriber is subscribed to
	SubscribedTo() event.DomainEvent
	Action() string
	// On executes the required tasks based on the subscribed event.DomainEvent
	On(context.Context, event.DomainEvent) error
}

// NewSubscriber wraps the given Subscriber (referenced as `root`) with observability tooling such as
// logging, tracing and monitoring
func NewSubscriber(root Subscriber, l *zap.Logger) Subscriber {
	return logger{
		logger: l,
		next:   root,
	}
}
