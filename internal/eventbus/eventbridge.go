package eventbus

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/hashicorp/go-multierror"
	"github.com/maestre3d/newton/internal/event"
	"github.com/maestre3d/newton/internal/infrastructure"
)

// EventBridge Amazon Web Services EventBridge implementation
type EventBridge struct {
	cfg infrastructure.Configuration
	c   *eventbridge.Client
	mu  sync.Mutex
}

// NewEventBridge allocates a new AWS EventBridge implementation
func NewEventBridge(cfg infrastructure.Configuration, c *eventbridge.Client) *EventBridge {
	return &EventBridge{
		cfg: cfg,
		c:   c,
		mu:  sync.Mutex{},
	}
}

func (b *EventBridge) Publish(ctx context.Context, events ...event.DomainEvent) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	entries, err := b.marshalAWSEventBulk(events)
	if err != nil {
		return err
	}
	o, err := b.c.PutEvents(ctx, &eventbridge.PutEventsInput{
		Entries: entries,
	})
	if err != nil {
		return err
	}
	return b.processError(o)
}

func (b *EventBridge) marshalAWSEventBulk(events []event.DomainEvent) ([]types.PutEventsRequestEntry, error) {
	entries := make([]types.PutEventsRequestEntry, 0)
	for _, ev := range events {
		eJSON, err := json.Marshal(ev)
		if err != nil {
			return nil, err
		}
		entries = append(entries, types.PutEventsRequestEntry{
			Detail:       aws.String(string(eJSON)),
			DetailType:   aws.String(ev.ID()),
			EventBusName: aws.String(b.cfg.AWSEventBus),
			Source:       aws.String(b.cfg.Application),
			Time:         aws.Time(time.Now().UTC()),
		},
		)
	}
	return entries, nil
}

func (b *EventBridge) processError(o *eventbridge.PutEventsOutput) error {
	errs := new(multierror.Error)
	for _, i := range o.Entries {
		if i.ErrorMessage != nil && *i.ErrorMessage != "" {
			errs = multierror.Append(errs, errors.New(*i.ErrorMessage))
		}
	}
	return errs.ErrorOrNil()
}
