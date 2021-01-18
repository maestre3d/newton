package event

// AuthorRemoved an aggregate.Author was permanently removed
type AuthorRemoved struct {
	AuthorID    string `json:"author_id"`
	DisplayName string `json:"display_name"`
	DeleteTime  string `json:"delete_time"`
}

// ID returns the event unique identifier
func (c AuthorRemoved) ID() string {
	return "AUTHOR_REMOVED"
}
