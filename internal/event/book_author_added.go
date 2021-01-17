package event

// BookAuthorAdded an aggregate.Author was added into the aggregate.Book authors list
type BookAuthorAdded struct {
	BookID   string `json:"book_id"`
	AuthorID string `json:"author_id"`
	AddTime  string `json:"add_time"`
}

// ID returns the topic/event unique identifier
func (a BookAuthorAdded) ID() string {
	return "BOOK_AUTHOR_ADDED"
}
