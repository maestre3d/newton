package application

import (
	"context"
	"errors"
	"io"

	"github.com/maestre3d/newton/internal/aggregate"
	"github.com/maestre3d/newton/internal/event"
	"github.com/maestre3d/newton/internal/repository"
	"github.com/maestre3d/newton/internal/valueobject"
)

// Author performs all the aggregate.Author use cases atomically
type Author struct {
	repo repository.Author
	bus  event.Bus
}

// NewAuthor allocates a new Author use case performer
func NewAuthor(r repository.Author, b event.Bus) *Author {
	return &Author{
		repo: r,
		bus:  b,
	}
}

// GetByID retrieves an aggregate.Author by its unique identifier
func (a *Author) GetByID(ctx context.Context, id valueobject.AuthorID) (*aggregate.Author, error) {
	if id.Value() == "" {
		return nil, aggregate.ErrAuthorNotFound
	}
	author, err := a.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if author == nil {
		return nil, aggregate.ErrAuthorNotFound
	}

	return author, nil
}

// SearchAll retrieves an aggregate.Author lists by the given criteria along with a pagination token
func (a *Author) SearchAll(ctx context.Context, criteria repository.Criteria) ([]*aggregate.Author, string, error) {
	authors, nextPage, err := a.repo.Search(ctx, criteria)
	if err != nil {
		return nil, "", err
	} else if len(authors) == 0 {
		return nil, "", aggregate.ErrAuthorNotFound
	}
	return authors, nextPage, nil
}

// Create creates and persists an aggregate.Author
func (a *Author) Create(ctx context.Context, id valueobject.AuthorID, name valueobject.DisplayName,
	createBy valueobject.Username, image valueobject.Image) error {
	if author, _ := a.GetByID(ctx, id); author != nil {
		return aggregate.ErrAuthorAlreadyExists
	}

	author := aggregate.NewAuthor(id, name, createBy, image)
	if err := a.repo.Save(ctx, *author); err != nil {
		return err
	} else if err := a.bus.Publish(ctx, author.PullEvents()...); a.bus != nil && err != nil {
		go func() { // rollback
			author.Metadata.MarkAsRemoval = true
			_ = a.repo.Save(ctx, *author)
		}()
		return err
	}
	return nil
}

// Modify mutates the given aggregate.Author state
func (a *Author) Modify(ctx context.Context, id valueobject.AuthorID, name valueobject.DisplayName,
	createBy valueobject.Username, image valueobject.Image) error {
	author, err := a.GetByID(ctx, id)
	if err != nil {
		return err
	} else if name.Value() == "" && createBy.Value() == "" && image.Value() == "" {
		return nil
	}

	memo := author
	if name.Value() != "" {
		author.DisplayName = name
	}
	if createBy.Value() != "" {
		author.CreateBy = createBy
	}
	if image.Value() != "" {
		author.Image = image
	}
	author.Update(author.DisplayName, author.CreateBy, author.Image)

	if err = a.repo.Save(ctx, *author); err != nil {
		return err
	} else if err := a.bus.Publish(ctx, author.PullEvents()...); a.bus != nil && err != nil {
		go func() { // rollback
			_ = a.repo.Save(ctx, *memo)
		}()
		return err
	}
	return nil
}

// ChangeState restores or deactivates the given aggregate.Author
func (a *Author) ChangeState(ctx context.Context, id valueobject.AuthorID, s bool) error {
	author, err := a.GetByID(ctx, id)
	if err != nil {
		return err
	} else if author.Metadata.State == s {
		return nil
	}

	memo := author.Metadata.State
	author.ChangeState(s)
	if err = a.repo.Save(ctx, *author); err != nil {
		return err
	} else if err := a.bus.Publish(ctx, author.PullEvents()...); a.bus != nil && err != nil {
		go func() { // rollback
			author.Metadata.State = memo
			_ = a.repo.Save(ctx, *author)
		}()
		return err
	}
	return nil
}

// Remove permanently deletes the given aggregate.Author
func (a *Author) Remove(ctx context.Context, id valueobject.AuthorID) error {
	author, err := a.GetByID(ctx, id)
	if err != nil && errors.Is(err, aggregate.ErrAuthorNotFound) {
		return nil
	} else if err != nil {
		return err
	}

	memo := author
	author.Remove()
	if err = a.repo.Save(ctx, *author); err != nil {
		return err
	} else if err := a.bus.Publish(ctx, author.PullEvents()...); a.bus != nil && err != nil {
		go func() { // rollback
			_ = a.repo.Save(ctx, *memo)
		}()
		return err
	}
	return nil
}

// UploadPicture stores the given image into specific static buckets and CDNs
func (a *Author) UploadPicture(ctx context.Context, file io.ReadCloser, size int64, name string) error {
	return nil
}
