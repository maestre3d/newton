package query

import (
	"time"

	"github.com/maestre3d/newton/internal/aggregate"
)

// AuthorResponse aggregate.Author query response
type authorResponse struct {
	AuthorID    string `json:"author_id"`
	DisplayName string `json:"display_name"`
	CreateBy    string `json:"create_by"`
	TotalBooks  int    `json:"total_books"`
	Image       string `json:"image"`
	LastUpdate  string `json:"last_update"`
}

// marshalAuthorResponse parses the given aggregate.Author into an AuthorResponse
func marshalAuthorResponse(a aggregate.Author) *authorResponse {
	return &authorResponse{
		AuthorID:    a.ID.Value(),
		DisplayName: a.DisplayName.Value(),
		CreateBy:    a.CreateBy.Value(),
		TotalBooks:  int(a.TotalBooks.Value()),
		Image:       a.Image.Value(),
		LastUpdate:  a.Metadata.UpdateTime.Format(time.RFC3339),
	}
}

type authorsResponse struct {
	Items    []*authorResponse `json:"items"`
	Count    int               `json:"count"`
	NextPage string            `json:"next_page"`
}

// marshalAuthorsResponse parses the given aggregate.Author slice into an AuthorsResponse
func marshalAuthorsResponse(as []*aggregate.Author, nextPage string) *authorsResponse {
	res := &authorsResponse{
		Items:    make([]*authorResponse, 0),
		Count:    len(as),
		NextPage: nextPage,
	}
	for _, a := range as {
		res.Items = append(res.Items, marshalAuthorResponse(*a))
	}
	return res
}
