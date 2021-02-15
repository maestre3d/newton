package repository

// Criteria Newton DSL query lang
type Criteria struct {
	Limit      int    `json:"limit"`
	NextPage   string `json:"next_page"`
	ActiveOnly bool   `json:"active_only"`
}

// NewCriteria creates and ensures criteria
func NewCriteria(l int, nextPage string) *Criteria {
	if l <= 0 {
		l = 100
	}
	return &Criteria{
		Limit:      l,
		NextPage:   nextPage,
		ActiveOnly: true,
	}
}
