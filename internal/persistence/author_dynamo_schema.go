package persistence

import (
	"strings"
	"time"

	"github.com/maestre3d/newton/internal/aggregate"
	"github.com/maestre3d/newton/internal/event"
	"github.com/maestre3d/newton/internal/valueobject"
)

// authorDynamo Amazon Web Services DynamoDB aggregate.Author specific schema
type authorDynamo struct {
	PK          string
	SK          string
	DisplayName string
	CreateBy    string
	Image       string `dynamodbav:"Image,omitempty"`
	CreateTime  string
	UpdateTime  string
	Active      bool
}

// marshalAuthorDynamo parses the given aggregate.Author into an Amazon Web Services DynamoDB schema
func marshalAuthorDynamo(a aggregate.Author) *authorDynamo {
	return &authorDynamo{
		PK:          authorAdjacencyPattern + a.ID.Value(),
		SK:          authorAdjacencyPattern + a.ID.Value(),
		DisplayName: a.DisplayName.Value(),
		CreateBy:    a.CreateBy.Value(),
		Image:       a.Image.Value(),
		CreateTime:  a.Metadata.CreateTime.Format(time.RFC3339),
		UpdateTime:  a.Metadata.UpdateTime.Format(time.RFC3339),
		Active:      a.Metadata.State,
	}
}

// unmarshalAuthorDynamo parses the given authorDynamo into an aggregate.Author
func unmarshalAuthorDynamo(a authorDynamo) (*aggregate.Author, error) {
	createTime, _ := time.Parse(time.RFC3339, a.CreateTime)
	updateTime, _ := time.Parse(time.RFC3339, a.UpdateTime)
	return &aggregate.Author{
		ID:          valueobject.AuthorID(strings.Trim(a.PK, authorAdjacencyPattern)),
		DisplayName: valueobject.DisplayName(a.DisplayName),
		CreateBy:    valueobject.Username(a.CreateBy),
		Image:       valueobject.Image(a.Image),
		Metadata: valueobject.Metadata{
			CreateTime: createTime,
			UpdateTime: updateTime,
			State:      a.Active,
		},
		Events: make([]event.DomainEvent, 0),
	}, nil
}

// unmarshalAuthorDynamoBulk parses the given authorDynamo list into a list of aggregate.Author(s)
func unmarshalAuthorDynamoBulk(as []authorDynamo) ([]*aggregate.Author, error) {
	authors := make([]*aggregate.Author, 0)
	for _, aDb := range as {
		a, err := unmarshalAuthorDynamo(aDb)
		if err != nil {
			return nil, err
		}
		authors = append(authors, a)
	}
	return authors, nil
}
