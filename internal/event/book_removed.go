package event

// BookRemoved a new book was permanently removed from our catalog
type BookRemoved struct {
	BookID     string `json:"book_id"`
	RemoveTime string `json:"remove_time"`
}

// ID Event unique identifier
//	Might be the topic name
func (r BookRemoved) ID() string {
	return "BOOK_REMOVED"
}
