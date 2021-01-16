package event

// AuthorRemoved an aggregate.Author was permanently removed
type AuthorRemoved struct {
	AuthorID   string `json:"author_id"`
	DeleteTime string `json:"delete_time"`
}

// ID returns the event unique identifier
func (c AuthorRemoved) ID() string {
	return "AUTHOR_REMOVED"
}
