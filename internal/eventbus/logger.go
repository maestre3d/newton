package eventbus

import (
	"context"

	"github.com/maestre3d/newton/internal/event"
	"go.uber.org/zap"
)

type logger struct {
	log  *zap.Logger
	next Bus
}

func (l logger) Publish(ctx context.Context, evs ...event.DomainEvent) (err error) {
	defer func() {
		fields := []zap.Field{
			zap.Any("events", evs),
			zap.Int("total_events", len(evs)),
		}
		if err != nil {
			fields = append(fields, zap.Error(err))
			l.log.Error("failed to publish events", fields...)
			return
		}
		l.log.Info("published events", fields...)
	}()

	err = l.next.Publish(ctx, evs...)
	return
}
