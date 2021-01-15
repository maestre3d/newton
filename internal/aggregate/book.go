package aggregate

import (
	"time"

	"github.com/maestre3d/newton/internal/valueobject"
)

// Book eBook or document uploaded by users
type Book struct {
	ID valueobject.BookID
	// Book's display name
	Title string
	// User's username who uploaded this book, since usernames are unique and immutable, high-cardinality is ensured
	Uploader string
	// Potential data replication/projection. Use IN statement query to lookup inside this field
	Authors     []string
	PublishYear int
	// Potential data replication/projection
	Categories []string
	// S3 url Book's cover image reference
	Cover string
	// S3 url media reference
	ExternalLink string
	CreateTime   time.Time
	UpdateTime   time.Time
}
