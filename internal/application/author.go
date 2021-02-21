package application

import (
	"context"
	"errors"

	"github.com/maestre3d/newton/internal/aggregate"
	"github.com/maestre3d/newton/internal/domain"
	"github.com/maestre3d/newton/internal/eventbus"
	"github.com/maestre3d/newton/internal/repository"
	"github.com/maestre3d/newton/internal/service"
	"github.com/maestre3d/newton/internal/valueobject"
)

// Author performs all the aggregate.Author use cases atomically
type Author struct {
	repo   repository.Author
	bucket service.FileBucket
	bus    eventbus.Bus
}

// NewAuthor allocates a new Author use case performer
func NewAuthor(r repository.Author, b eventbus.Bus, bucket service.FileBucket) *Author {
	return &Author{
		repo:   r,
		bus:    b,
		bucket: bucket,
	}
}

var (
	maxPictureSize = 2 * domain.MebiByte
)

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
	} else if err = a.bus.Publish(ctx, author.PullEvents()...); a.bus != nil && err != nil {
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
	}
	wasNotUpdated := (name.Value() == "" || name.Value() == author.DisplayName.Value()) &&
		(createBy.Value() == "" || createBy.Value() == author.CreateBy.Value()) &&
		(image.Value() == "" || image.Value() == author.Image.Value())
	if wasNotUpdated {
		return nil
	}

	memo := author
	if name.Value() != "" && name != author.DisplayName {
		author.DisplayName = name
	}
	if createBy.Value() != "" && createBy != author.CreateBy {
		author.CreateBy = createBy
	}
	if image.Value() != "" && image != author.Image {
		author.Image = image
	}

	author.Update(author.DisplayName, author.CreateBy, author.TotalBooks, author.Image)
	if err = a.repo.Save(ctx, *author); err != nil {
		return err
	} else if err = a.bus.Publish(ctx, author.PullEvents()...); a.bus != nil && err != nil {
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
	} else if err = a.bus.Publish(ctx, author.PullEvents()...); a.bus != nil && err != nil {
		go func() { // rollback
			_ = a.repo.Save(ctx, *memo)
		}()
		return err
	}
	return nil
}

// UploadPicture stores the given image into specific static buckets and CDNs
func (a *Author) UploadPicture(ctx context.Context, id valueobject.AuthorID, file *valueobject.File) error {
	if file.Size > maxPictureSize {
		return aggregate.ErrImageSizeOutOfRange // stop here to avoid useless data fetching
	} else if file.Extension != "png" && file.Extension != "jpeg" && file.Extension != "jpg" && file.Extension != "webp" {
		return valueobject.ErrImageInvalidExtension
	}

	author, err := a.GetByID(ctx, id)
	if err != nil {
		return err
	}

	key := "newton/authors/" + id.Value() + "." + file.Extension
	fileUrl, err := a.bucket.Upload(ctx, key, file)
	if err != nil {
		return err
	}

	author.UploadPicture(fileUrl)
	if err = a.bus.Publish(ctx, author.PullEvents()...); err != nil {
		go func() {
			_ = a.bucket.Delete(ctx, key)
		}()
		return err
	}

	return nil
}
