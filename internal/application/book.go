package application

import (
	"context"

	"github.com/maestre3d/newton/internal/aggregate"
	"github.com/maestre3d/newton/internal/event"
	"github.com/maestre3d/newton/internal/repository"
	"github.com/maestre3d/newton/internal/valueobject"
)

// Book performs all the aggregate.Author use cases atomically
type Book struct {
	repo repository.Book
	bus  event.Bus
}

func (b Book) Create(ctx context.Context, id valueobject.BookID, title valueobject.Title, uploader valueobject.Username,
	year valueobject.PublishYear, image valueobject.Image) error {
	if book, _ := b.GetByID(ctx, id); book != nil {
		return aggregate.ErrBookAlreadyExists
	}

	book := aggregate.NewBook(id, title, uploader, year, image)
	if err := b.repo.Save(ctx, *book); err != nil {
		return err
	} else if err := b.bus.Publish(ctx, book.PullEvents()...); err != nil {
		// rollback
		rCtx := ctx
		go func() {
			book.Metadata.MarkAsRemoval = true // avoids memory alloc, aggregate.Book.Remove() pushes a domain event
			_ = b.repo.Save(rCtx, *book)
		}()
		return err
	}
	return nil
}

func (b Book) GetByID(ctx context.Context, id valueobject.BookID) (*aggregate.Book, error) {
	book, err := b.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if book == nil {
		return nil, aggregate.ErrBookNotFound
	}

	return book, nil
}
