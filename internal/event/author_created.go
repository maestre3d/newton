package event

// AuthorCreated an aggregate.Author was aggregated
type AuthorCreated struct {
	AuthorID    string `json:"author_id"`
	DisplayName string `json:"display_name"`
	CreateBy    string `json:"create_by"`
	Image       string `json:"image"`
	CreateTime  string `json:"create_time"`
}

// ID returns the event unique identifier
func (c AuthorCreated) ID() string {
	return "AUTHOR_CREATED"
}
