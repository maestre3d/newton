package event

// AuthorUpdated an aggregate.Author was modified
type AuthorUpdated struct {
	AuthorID    string `json:"author_id"`
	DisplayName string `json:"display_name"`
	CreateBy    string `json:"create_by"`
	Image       string `json:"image"`
	UpdateTime  string `json:"update_time"`
}

// ID returns the event unique identifier
func (c AuthorUpdated) ID() string {
	return "AUTHOR_UPDATED"
}
