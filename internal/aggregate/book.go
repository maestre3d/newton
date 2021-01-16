package aggregate

import (
	"errors"
	"time"

	"github.com/maestre3d/newton/internal/event"
	"github.com/maestre3d/newton/internal/valueobject"
)

// Book eBook or document uploaded by users
type Book struct {
	ID valueobject.BookID
	// Book's display name
	Title valueobject.Title
	// User's username who uploaded this book, since usernames are unique and immutable, high-cardinality is ensured
	Uploader    valueobject.Username
	PublishYear valueobject.PublishYear
	// S3 url media reference
	Resource valueobject.BookResource
	// S3 url Book's cover image reference
	Image valueobject.Image

	Metadata valueobject.Metadata
	events   []event.DomainEvent
}

var (
	// ErrBookNotFound the specified book was not found
	ErrBookNotFound = errors.New("book not found")
	// ErrBookAlreadyExists the given book already exists
	ErrBookAlreadyExists = errors.New("book already exists")
)

// NewBook creates a Book and pushes event.BookCreated event
func NewBook(id valueobject.BookID, title valueobject.Title, uploader valueobject.Username,
	year valueobject.PublishYear, image valueobject.Image) *Book {
	currentTime := time.Now().UTC()
	return &Book{
		ID:          id,
		Title:       title,
		Uploader:    uploader,
		PublishYear: year,
		Resource:    valueobject.BookResource{},
		Image:       image,
		Metadata: valueobject.Metadata{
			CreateTime:    currentTime,
			UpdateTime:    currentTime,
			State:         true,
			MarkAsRemoval: false,
		},
		events: []event.DomainEvent{event.BookCreated{
			BookID:      id.Value(),
			Title:       title.Value(),
			Uploader:    uploader.Value(),
			PublishYear: year.Value(),
			Image:       image.Value(),
			CreateTime:  currentTime.String(),
		}},
	}
}

// UploadResourceFile update the book resource url
func (b *Book) UploadResourceFile(r valueobject.BookResource) {
	b.Resource = r
	b.events = append(b.events, event.BookUploaded{
		BookID:   b.ID.Value(),
		Resource: b.Resource.Value(),
	})
}

// Remove mark this book for permanent removal
func (b *Book) Remove() {
	b.Metadata.MarkAsRemoval = true
	b.events = append(b.events, event.BookRemoved{
		BookID:     b.ID.Value(),
		RemoveTime: time.Now().UTC().String(),
	})
}

// PullEvents retrieves all current events that happened inside aggregate and flush them
func (b *Book) PullEvents() []event.DomainEvent {
	memo := b.events
	b.events = []event.DomainEvent{}
	return memo
}
