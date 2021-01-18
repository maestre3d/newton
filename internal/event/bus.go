package event

import "context"

// Bus Newton Event bus, receives and delivers side-effects from one to many components in the entire ecosystem
type Bus interface {
	// Publish propagates side-effects
	Publish(context.Context, ...DomainEvent) error
}
