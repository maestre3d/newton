package controller

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/hashicorp/go-multierror"
	"github.com/maestre3d/newton/internal/event"
	"github.com/maestre3d/newton/internal/infrastructure"
	"github.com/maestre3d/newton/internal/subscriber"
	"github.com/maestre3d/newton/pkg/middleware"
	"go.uber.org/zap"
)

// EventAWS Amazon Web Services event controller that complies with lambda.Start,
// this implementation will lookup for each events.SQSMessage and execute the given subscriber.Subscriber
type EventAWS struct {
	logger *zap.Logger
	sub    subscriber.Subscriber
}

// NewEventAWS allocates a new EventAWS controller implementation
func NewEventAWS(logger *zap.Logger, s subscriber.Subscriber) EventAWS {
	return EventAWS{logger: logger, sub: s}
}

// Handle used by lambda.Start, executes required marshaling processes and calls subscriber.Subscriber.On() method
func (a EventAWS) Handle(ctx context.Context, evs events.SQSEvent) error {
	errs := new(multierror.Error)
	for _, msg := range evs.Records {
		if err := a.logHandleSQSMessage(ctx, msg); err != nil {
			errs = multierror.Append(errs, err)
			continue
		}
	}
	return errs.ErrorOrNil()
}

func (a EventAWS) logHandleSQSMessage(ctx context.Context, msg events.SQSMessage) error {
	middleware.LogSQSMessage(a.logger, msg)
	return a.handleSQSMessage(ctx, msg)
}

func (a EventAWS) handleSQSMessage(ctx context.Context, msg events.SQSMessage) error {
	entity := events.SNSEntity{}
	if err := infrastructure.UnmarshalSQSToSNS(msg.Body, &entity); err != nil {
		return err
	}
	return a.logHandleSNSEntity(ctx, entity)
}

func (a EventAWS) logHandleSNSEntity(ctx context.Context, e events.SNSEntity) error {
	middleware.LogSNSEntity(a.logger, e)
	return a.handleSNSEntity(ctx, e)
}

func (a EventAWS) handleSNSEntity(ctx context.Context, e events.SNSEntity) error {
	ev := a.sub.SubscribedTo()
	if err := infrastructure.UnmarshalSNSToEvent(e.Message, ev); err != nil {
		return err
	}
	return a.logExecSubscriber(ctx, ev)
}

func (a EventAWS) logExecSubscriber(ctx context.Context, e event.DomainEvent) error {
	middleware.LogDomainEvent(a.logger, e)
	return a.execSubscriber(ctx, e)
}

func (a EventAWS) execSubscriber(ctx context.Context, e event.DomainEvent) error {
	return a.sub.On(ctx, e)
}
