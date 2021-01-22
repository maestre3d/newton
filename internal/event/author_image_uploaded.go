package event

// AuthorImageUploaded an aggregate.Author picture was uploaded to the respective bucket
type AuthorImageUploaded struct {
	AuthorID string `json:"author_id"`
	Image    string `json:"image"`
}

// ID returns the event id
func (a AuthorImageUploaded) ID() string {
	return "AUTHOR_IMAGE_UPLOADED"
}
