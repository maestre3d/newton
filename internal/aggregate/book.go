package aggregate

import (
	"time"

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
	Cover string
	// Potential data replication/projection
	Authors []string
	// Potential data replication/projection
	Categories []string
	CreateTime time.Time
	UpdateTime time.Time
}
