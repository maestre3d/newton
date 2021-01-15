package event

// BookUploaded an aggregate.Book file was uploaded
type BookUploaded struct {
	BookID   string `json:"book_id"`
	Resource string `json:"resource"`
}

func (b BookUploaded) ID() string {
	return "BOOK_UPLOADED"
}
