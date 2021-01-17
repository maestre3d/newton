package aggregate

import (
	"time"

	"github.com/maestre3d/newton/internal/event"
	"github.com/maestre3d/newton/internal/valueobject"
)

// BookAuthors aggregate.Author list item of an aggregate.Book
type BookAuthors struct {
	BookID   valueobject.BookID
	AuthorID valueobject.AuthorID
	AddTime  time.Time

	events []event.DomainEvent
}

// NewBookAuthors creates and pushes domain events into an aggregate.BookAuthors
func NewBookAuthors(id valueobject.BookID, authorID valueobject.AuthorID) *BookAuthors {
	addTime := time.Now().UTC()
	return &BookAuthors{
		BookID:   id,
		AuthorID: authorID,
		AddTime:  addTime,
		events: []event.DomainEvent{
			event.BookAuthorAdded{
				BookID:   id.Value(),
				AuthorID: authorID.Value(),
				AddTime:  addTime.String(),
			},
		},
	}
}
