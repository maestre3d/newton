package event

// BookCreated a new book was aggregated into our catalog
type BookCreated struct {
	BookID      string `json:"book_id"`
	Title       string `json:"title"`
	Uploader    string `json:"uploader"`
	PublishYear int    `json:"publish_year"`
	Image       string `json:"image"`
	CreateTime  string `json:"create_time"`
}

// ID Event unique identifier
//	Might be the topic name
func (c BookCreated) ID() string {
	return "BOOK_CREATED"
}
