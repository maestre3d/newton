package subscriber

import (
	"context"

	"github.com/maestre3d/newton/internal/event"
	"go.uber.org/zap"
)

// logger wrapper struct, logs Subscriber executions
type logger struct {
	logger *zap.Logger
	next   Subscriber
}

func (l logger) SubscribedTo() event.DomainEvent {
	return l.next.SubscribedTo()
}

func (l logger) Action() string {
	return l.next.Action()
}

func (l logger) On(ctx context.Context, args event.DomainEvent) (err error) {
	defer func() {
		fields := []zap.Field{
			zap.String("event_id", args.ID()),
			zap.String("event_processor", l.Action()),
		}
		if err != nil {
			fields = append(fields, zap.Error(err))
			l.logger.Error("failed to process event", fields...)
			return
		}
		l.logger.Info("processed event successfully", fields...)
	}()

	err = l.next.On(ctx, args)
	return
}
