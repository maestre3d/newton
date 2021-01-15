package event

// DomainEvent represents something has happened within our ecosystem
type DomainEvent interface {
	// ID Event unique identifier, might be a topic
	ID() string
}
