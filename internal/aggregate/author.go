package aggregate

import (
	"time"

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
	events   []event.DomainEvent
}

// NewAuthor creates and pushes events into an aggregate.Author
func NewAuthor(id valueobject.AuthorID, name valueobject.DisplayName, createBy valueobject.Username,
	image valueobject.Image) *Author {
	currentTime := time.Now().UTC()
	return &Author{
		ID:          id,
		DisplayName: name,
		CreateBy:    createBy,
		Image:       image,
		Metadata: valueobject.Metadata{
			CreateTime:    currentTime,
			UpdateTime:    currentTime,
			State:         true,
			MarkAsRemoval: false,
		},
		events: []event.DomainEvent{
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
func (a *Author) Update(id valueobject.AuthorID, name valueobject.DisplayName, createBy valueobject.Username,
	image valueobject.Image) {
	currentTime := time.Now().UTC()
	a.ID = id
	a.DisplayName = name
	a.CreateBy = createBy
	a.Image = image
	a.Metadata.UpdateTime = currentTime
	a.events = append(a.events, event.AuthorUpdated{
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
		a.events = append(a.events, event.AuthorRestored{
			AuthorID:    a.ID.Value(),
			RestoreTime: currentTime.String(),
		})
		return
	}

	a.events = append(a.events, event.AuthorDeactivated{
		AuthorID:   a.ID.Value(),
		DeleteTime: currentTime.String(),
	})
}

// Remove mark aggregate as removed
func (a *Author) Remove() {
	a.Metadata.MarkAsRemoval = true
	a.events = append(a.events, event.AuthorRemoved{
		AuthorID:   a.ID.Value(),
		DeleteTime: time.Now().UTC().String(),
	})
}

// PullEvents retrieves and flushes all occurred events
func (a *Author) PullEvents() []event.DomainEvent {
	memo := a.events
	a.events = []event.DomainEvent{}
	return memo
}
