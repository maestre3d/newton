package eventbus

import (
	"context"

	"github.com/maestre3d/newton/internal/event"
	"go.uber.org/zap"
)

// Bus Newton Event bus, receives and delivers side-effects from one to many components in the entire ecosystem
type Bus interface {
	// Publish propagates side-effects
	Publish(context.Context, ...event.DomainEvent) error
}

// NewBus wraps the given Bus (referenced as `root`) with observability tooling such as
// logging, tracing and monitoring
func NewBus(root Bus, l *zap.Logger) Bus {
	return logger{
		log:  l,
		next: root,
	}
}
