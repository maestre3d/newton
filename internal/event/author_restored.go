package event

// AuthorRestored an aggregate.Author was restored from soft-removal
type AuthorRestored struct {
	AuthorID    string `json:"author_id"`
	DisplayName string `json:"display_name"`
	RestoreTime string `json:"restore_time"`
}

// ID returns the event unique identifier
func (c AuthorRestored) ID() string {
	return "AUTHOR_RESTORED"
}
