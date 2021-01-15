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
	Cover valueobject.Cover
	// Potential data replication/projection
	Authors []valueobject.AuthorID
	// Potential data replication/projection
	Categories []valueobject.CategoryID

	CreateTime time.Time
	UpdateTime time.Time
	State      bool

	MarkedAsRemoval bool
	events          []event.DomainEvent
}

var (
	// ErrBookNotFound the specified book was not found
	ErrBookNotFound = errors.New("book not found")
	// ErrBookAlreadyExists the given book already exists
	ErrBookAlreadyExists = errors.New("book already exists")
)

// NewBook creates a Book and pushes event.BookCreated event
func NewBook(id valueobject.BookID, title valueobject.Title, uploader valueobject.Username,
	year valueobject.PublishYear, cover valueobject.Cover, authors []valueobject.AuthorID,
	categories []valueobject.CategoryID) *Book {
	b := &Book{
		ID:              id,
		Title:           title,
		Uploader:        uploader,
		PublishYear:     year,
		Resource:        valueobject.BookResource{},
		Cover:           cover,
		Authors:         authors,
		Categories:      categories,
		CreateTime:      time.Now().UTC(),
		UpdateTime:      time.Now().UTC(),
		State:           true,
		MarkedAsRemoval: false,
		events:          nil,
	}
	b.events = []event.DomainEvent{event.BookCreated{
		BookID:      b.ID.Value(),
		Title:       b.Title.Value(),
		Uploader:    b.Uploader.Value(),
		PublishYear: b.PublishYear.Value(),
		Authors:     b.AuthorsPrimitive(),
		Categories:  b.CategoriesPrimitive(),
		Cover:       b.Cover.Value(),
		CreateTime:  b.CreateTime.String(),
	}}
	return b
}

// UploadFile update the book resource url
func (b *Book) UploadFile(r valueobject.BookResource) {
	b.Resource = r
	b.events = append(b.events, event.BookUploaded{
		BookID:   b.ID.Value(),
		Resource: b.Resource.Value(),
	})
}

// UploadFile mark this book for permanent removal
func (b *Book) Remove() {
	b.MarkedAsRemoval = true
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

// AuthorsPrimitive constructs an slice of authors primitives
func (b Book) AuthorsPrimitive() []string {
	if len(b.Authors) == 0 {
		return nil
	}
	a := make([]string, 0)
	for _, author := range b.Authors {
		a = append(a, author.Value())
	}
	return a
}

// CategoriesPrimitive constructs an slice of categories primitives
func (b Book) CategoriesPrimitive() []string {
	if len(b.Categories) == 0 {
		return nil
	}
	c := make([]string, 0)
	for _, category := range b.Categories {
		c = append(c, category.Value())
	}
	return c
}
