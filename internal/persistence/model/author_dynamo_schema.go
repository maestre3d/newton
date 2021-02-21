package model

import (
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/maestre3d/newton/internal/aggregate"
	"github.com/maestre3d/newton/internal/event"
	"github.com/maestre3d/newton/internal/valueobject"
)

const (
	// AuthorAdjacencyPattern string pattern for author schemas using Amazon Web Services DynamoDB's
	// Adjacency List pattern.
	//
	// A more detailed information can be found here:
	// https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/bp-adjacency-graphs.html
	AuthorAdjacencyPattern = "author#"
	// AuthorIdKey persistence author primary key name
	AuthorIdKey   = "author_id"
	totalBooksKey = "TotalBooks"
)

// AuthorDynamo Amazon Web Services DynamoDB aggregate.Author specific schema
type AuthorDynamo struct {
	PK          string
	SK          string
	DisplayName string
	CreateBy    string
	TotalBooks  int64
	Image       string `dynamodbav:"Image,omitempty"`
	CreateTime  string
	UpdateTime  string
	Active      bool
}

// MarshalAuthorDynamo parses the given aggregate.Author into an Amazon Web Services DynamoDB schema
func MarshalAuthorDynamo(a aggregate.Author) *AuthorDynamo {
	return &AuthorDynamo{
		PK:          AuthorAdjacencyPattern + a.ID.Value(),
		SK:          AuthorAdjacencyPattern + a.ID.Value(),
		DisplayName: a.DisplayName.Value(),
		CreateBy:    a.CreateBy.Value(),
		TotalBooks:  int64(a.TotalBooks),
		Image:       a.Image.Value(),
		CreateTime:  a.Metadata.CreateTime.Format(time.RFC3339),
		UpdateTime:  a.Metadata.UpdateTime.Format(time.RFC3339),
		Active:      a.Metadata.State,
	}
}

// UnmarshalAuthorDynamo parses the given authorDynamo into an aggregate.Author
func UnmarshalAuthorDynamo(a AuthorDynamo) (*aggregate.Author, error) {
	createTime, _ := time.Parse(time.RFC3339, a.CreateTime)
	updateTime, _ := time.Parse(time.RFC3339, a.UpdateTime)
	total, err := valueobject.NewTotalBooks(a.TotalBooks) // avoid negative uint overflow at runtime
	if err != nil {
		return nil, err
	}
	cleanedID := strings.Split(a.PK, AuthorAdjacencyPattern) // avpid use of strings.Trim() since it removes
	// each character given
	return &aggregate.Author{
		ID:          valueobject.AuthorID(strings.Join(cleanedID, "")),
		DisplayName: valueobject.DisplayName(a.DisplayName),
		CreateBy:    valueobject.Username(a.CreateBy),
		TotalBooks:  total,
		Image:       valueobject.Image(a.Image),
		Metadata: valueobject.Metadata{
			CreateTime: createTime,
			UpdateTime: updateTime,
			State:      a.Active,
		},
		Events: make([]event.DomainEvent, 0),
	}, nil
}

// UnmarshalAuthorDynamoBulk parses the given authorDynamo list into a list of aggregate.Author(s)
func UnmarshalAuthorDynamoBulk(as []AuthorDynamo) ([]*aggregate.Author, error) {
	authors := make([]*aggregate.Author, 0)
	for _, aDb := range as {
		a, err := UnmarshalAuthorDynamo(aDb)
		if err != nil {
			return nil, err
		}
		authors = append(authors, a)
	}
	return authors, nil
}

// NewAuthorDynamoProjection creates a new AWS DynamoDB projection builder for an Author
func NewAuthorDynamoProjection() expression.ProjectionBuilder {
	return expression.NamesList(expression.Name(DynamoPartitionKey), expression.Name(DynamoSortKey), expression.Name(DisplayNameKey),
		expression.Name(CreateTimeKey), expression.Name(UpdateTimeKey), expression.Name(ActiveFlagKey), expression.Name(CreateByKey),
		expression.Name(ImageKey), expression.Name(totalBooksKey))
}
