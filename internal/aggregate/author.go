package aggregate

import (
	"errors"
	"time"

	"github.com/maestre3d/newton/internal/domain"
	"github.com/maestre3d/newton/internal/event"
	"github.com/maestre3d/newton/internal/valueobject"
)

// Author creator or collaborator of an specific media/lecture (e.g. book, video, document, ...)
type Author struct {
	ID          valueobject.AuthorID
	DisplayName valueobject.DisplayName
	CreateBy    valueobject.Username
	Image       valueobject.Image

	Metadata valueobject.Metadata
	Events   []event.DomainEvent
}

var (
	// ErrAuthorNotFound the given Author was not found
	ErrAuthorNotFound = domain.NewNotFound("author")
	// ErrAuthorAlreadyExists the given Author already exists
	ErrAuthorAlreadyExists = domain.NewAlreadyExists("author")
	// ErrAuthorCannotParse the current Author could not be parsed successfully
	ErrAuthorCannotParse = errors.New("author can not be parsed")
)

// NewAuthor creates and pushes Events into an aggregate.Author
func NewAuthor(id valueobject.AuthorID, name valueobject.DisplayName, createBy valueobject.Username,
	image valueobject.Image) *Author {
	currentTime := time.Now().UTC()
	return &Author{
		ID:          id,
		DisplayName: name,
		CreateBy:    createBy,
		Image:       image,
		Metadata: valueobject.Metadata{
			CreateTime:     currentTime,
			UpdateTime:     currentTime,
			State:          true,
			MarkAsMutation: false,
			MarkAsRemoval:  false,
		},
		Events: []event.DomainEvent{
			event.AuthorCreated{
				AuthorID:    id.Value(),
				DisplayName: name.Value(),
				CreateBy:    createBy.Value(),
				Image:       image.Value(),
				CreateTime:  currentTime.String(),
			},
		},
	}
}

// Update perform a bulk modification
func (a *Author) Update(name valueobject.DisplayName, createBy valueobject.Username, image valueobject.Image) {
	currentTime := time.Now().UTC()
	a.DisplayName = name
	a.CreateBy = createBy
	a.Image = image
	a.Metadata.UpdateTime = currentTime
	a.Metadata.MarkAsMutation = true
	a.Events = append(a.Events, event.AuthorUpdated{
		AuthorID:    a.ID.Value(),
		DisplayName: a.DisplayName.Value(),
		CreateBy:    a.CreateBy.Value(),
		Image:       a.Image.Value(),
		UpdateTime:  currentTime.String(),
	})
}

// ChangeState overrides the current state to the given value
func (a *Author) ChangeState(s bool) {
	currentTime := time.Now().UTC()
	a.Metadata.State = s
	a.Metadata.UpdateTime = currentTime
	if s {
		a.Events = append(a.Events, event.AuthorRestored{
			AuthorID:    a.ID.Value(),
			RestoreTime: currentTime.String(),
		})
		return
	}

	a.Events = append(a.Events, event.AuthorDeactivated{
		AuthorID:   a.ID.Value(),
		DeleteTime: currentTime.String(),
	})
}

// Remove mark aggregate as removed
func (a *Author) Remove() {
	a.Metadata.MarkAsRemoval = true
	a.Events = append(a.Events, event.AuthorRemoved{
		AuthorID:   a.ID.Value(),
		DeleteTime: time.Now().UTC().String(),
	})
}

// PullEvents retrieves and flushes all occurred Events
func (a *Author) PullEvents() []event.DomainEvent {
	memo := a.Events
	a.Events = []event.DomainEvent{}
	return memo
}
