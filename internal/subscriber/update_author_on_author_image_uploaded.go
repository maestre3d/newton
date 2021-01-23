package subscriber

import (
	"context"

	"github.com/maestre3d/newton/internal/application"
	"github.com/maestre3d/newton/internal/event"
	"github.com/maestre3d/newton/internal/valueobject"
)

// UpdateAuthorOnImageUploaded projects the new Author's image url into main persistence store
func UpdateAuthorOnImageUploaded(app *application.Author, ctx context.Context, ev event.AuthorImageUploaded) error {
	id, err := valueobject.NewAuthorID(ev.AuthorID)
	if err != nil {
		return err
	}
	image, err := valueobject.NewImage(ev.Image)
	if err != nil {
		return err
	}
	return app.Modify(ctx, id, "", "", image)
}
