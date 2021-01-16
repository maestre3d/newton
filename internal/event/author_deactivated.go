package event

// AuthorDeactivated an aggregate.Author was deactivated
type AuthorDeactivated struct {
	AuthorID   string `json:"author_id"`
	DeleteTime string `json:"delete_time"`
}

// ID returns the event unique identifier
func (c AuthorDeactivated) ID() string {
	return "AUTHOR_DEACTIVATED"
}