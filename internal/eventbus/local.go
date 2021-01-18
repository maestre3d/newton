package eventbus

import (
	"context"
	"sync"

	"github.com/maestre3d/newton/internal/event"
	"go.uber.org/zap"
)

// Local in memory event bus implementation
type Local struct {
	logger *zap.Logger
	mu     sync.Mutex
}

// NewLocal allocates a new Local event bus implementation
func NewLocal(logger *zap.Logger) *Local {
	return &Local{
		logger: logger,
		mu:     sync.Mutex{},
	}
}

// Publish propagates side-effects
func (l *Local) Publish(ctx context.Context, events ...event.DomainEvent) error {
	for _, ev := range events {
		_ = l.publish(ctx, ev)
	}
	return nil
}

func (l *Local) publish(_ context.Context, ev event.DomainEvent) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logger.Info("published event", zap.String("topic", ev.ID()))
	return nil
}
