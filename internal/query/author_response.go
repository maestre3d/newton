package query

import (
	"time"

	"github.com/maestre3d/newton/internal/aggregate"
)

// AuthorResponse aggregate.Author query response
type AuthorResponse struct {
	AuthorID    string `json:"author_id"`
	DisplayName string `json:"display_name"`
	CreateBy    string `json:"create_by"`
	Image       string `json:"image"`
	LastUpdate  string `json:"last_update"`
}

// MarshalAuthorResponse parses the given aggregate.Author into an AuthorResponse
func MarshalAuthorResponse(a aggregate.Author) *AuthorResponse {
	return &AuthorResponse{
		AuthorID:    a.ID.Value(),
		DisplayName: a.DisplayName.Value(),
		CreateBy:    a.CreateBy.Value(),
		Image:       a.Image.Value(),
		LastUpdate:  a.Metadata.UpdateTime.Format(time.RFC3339),
	}
}
