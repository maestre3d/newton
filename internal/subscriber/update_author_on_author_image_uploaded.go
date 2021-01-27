package subscriber

import (
	"context"

	"github.com/maestre3d/newton/internal/application"
	"github.com/maestre3d/newton/internal/event"
	"github.com/maestre3d/newton/internal/valueobject"
)

// UpdateAuthorOnImageUploaded projects the new Author's image url into main persistence store
type UpdateAuthorOnImageUploaded struct {
	app *application.Author
}

// NewUpdateAuthorOnImageUploaded allocates a new UpdateAuthorOnImageUploaded subscriber
func NewUpdateAuthorOnImageUploaded(app *application.Author) *UpdateAuthorOnImageUploaded {
	return &UpdateAuthorOnImageUploaded{app: app}
}

func (u UpdateAuthorOnImageUploaded) SubscribedTo() event.DomainEvent {
	return &event.AuthorImageUploaded{}
}

func (u UpdateAuthorOnImageUploaded) Action() string {
	return "UPDATE_AUTHOR"
}

func (u UpdateAuthorOnImageUploaded) On(ctx context.Context, arg event.DomainEvent) error {
	ev := arg.(*event.AuthorImageUploaded)
	id, err := valueobject.NewAuthorID(ev.AuthorID)
	if err != nil {
		return err
	}
	image, err := valueobject.NewImage(ev.Image)
	if err != nil {
		return err
	}
	return u.app.Modify(ctx, id, "", "", image)
}
