package query

import (
	"context"

	"github.com/maestre3d/newton/internal/aggregate"
	"github.com/maestre3d/newton/internal/application"
	"github.com/maestre3d/newton/internal/repository"
	"github.com/maestre3d/newton/internal/valueobject"
)

// GetAuthorByID retrieves an aggregate.Author by its unique identifier
type GetAuthorByID struct {
	ID string `json:"id"`
}

// GetAuthorByIDHandle executes GetAuthorByID query
func GetAuthorByIDHandle(app *application.Author, ctx context.Context, q GetAuthorByID) (*authorResponse, error) {
	if q.ID == "" {
		return nil, aggregate.ErrAuthorNotFound
	}
	id, err := valueobject.NewAuthorID(q.ID)
	if err != nil {
		return nil, err
	}
	author, err := app.GetByID(ctx, id)
	if err != nil {
		return nil, err
	} else if author.Metadata.State == false {
		return nil, aggregate.ErrAuthorNotFound
	}
	return marshalAuthorResponse(*author), nil
}

// ListAuthors retrieves a list of aggregate.Author(s)
type ListAuthors struct {
	Criteria repository.Criteria
}

// GetAuthorByIDHandle executes ListAuthors query
func ListAuthorsHandle(app *application.Author, ctx context.Context, q ListAuthors) (*authorsResponse, error) {
	authors, nextPage, err := app.SearchAll(ctx, q.Criteria)
	if err != nil {
		return nil, err
	}
	return marshalAuthorsResponse(authors, nextPage), nil
}
