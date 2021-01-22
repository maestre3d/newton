package command

import (
	"context"
	"io"

	"github.com/maestre3d/newton/internal/application"
	"github.com/maestre3d/newton/internal/valueobject"
)

// CreateAuthor requests an aggregate.Author insertion
type CreateAuthor struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	CreateBy    string `json:"create_by"`
	Image       string `json:"image"`
}

// CreateAuthorHandle executes CreateAuthor command
func CreateAuthorHandle(app *application.Author, ctx context.Context, cmd CreateAuthor) error {
	id, err := valueobject.NewAuthorID(cmd.ID)
	if err != nil {
		return err
	}
	name, err := valueobject.NewDisplayName(cmd.DisplayName)
	if err != nil {
		return err
	}
	createBy, err := valueobject.NewUsername(cmd.CreateBy)
	if err != nil {
		return err
	}
	var image valueobject.Image
	if cmd.Image != "" {
		image, err = valueobject.NewImage(cmd.Image)
		if err != nil {
			return err
		}
	}
	return app.Create(ctx, id, name, createBy, image)
}

// UpdateAuthor requests an aggregate.Author mutation
type UpdateAuthor struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	CreateBy    string `json:"create_by"`
	Image       string `json:"image"`
}

// UpdateAuthorHandle executes UpdateAuthor command
func UpdateAuthorHandle(app *application.Author, ctx context.Context, cmd UpdateAuthor) error {
	id, err := valueobject.NewAuthorID(cmd.ID)
	if err != nil {
		return err
	}
	name, err := valueobject.NewDisplayName(cmd.DisplayName)
	if err != nil {
		return err
	}
	createBy, err := valueobject.NewUsername(cmd.CreateBy)
	if err != nil {
		return err
	}
	var image valueobject.Image
	if cmd.Image != "" {
		image, err = valueobject.NewImage(cmd.Image)
		if err != nil {
			return err
		}
	}
	return app.Modify(ctx, id, name, createBy, image)
}

// ChangeAuthorState requests an aggregate.Author removal or restoring
type ChangeAuthorState struct {
	ID    string `json:"id"`
	State bool   `json:"state"`
}

// ChangeAuthorStateHandle executes ChangeAuthorState command
func ChangeAuthorStateHandle(app *application.Author, ctx context.Context, cmd ChangeAuthorState) error {
	id, err := valueobject.NewAuthorID(cmd.ID)
	if err != nil {
		return err
	}
	return app.ChangeState(ctx, id, cmd.State)
}

// DeleteAuthor requests an aggregate.Author permanent removal
type DeleteAuthor struct {
	ID    string `json:"id"`
	State bool   `json:"state"`
}

// DeleteAuthorHandle executes DeleteAuthor command
func DeleteAuthorHandle(app *application.Author, ctx context.Context, cmd DeleteAuthor) error {
	id, err := valueobject.NewAuthorID(cmd.ID)
	if err != nil {
		return err
	}
	return app.Remove(ctx, id)
}

// UploadAuthorPicture requests an aggregate.Author image upload
type UploadAuthorPicture struct {
	ID       string
	Filename string
	Size     int64
	Image    io.Reader
}

// UploadAuthorPictureHandle executes UploadAuthorPicture command
func UploadAuthorPictureHandle(app *application.Author, ctx context.Context, cmd UploadAuthorPicture) error {
	id, err := valueobject.NewAuthorID(cmd.ID)
	if err != nil {
		return err
	}
	return app.UploadPicture(ctx, id, valueobject.NewFile(cmd.Filename, cmd.Size, cmd.Image))
}
