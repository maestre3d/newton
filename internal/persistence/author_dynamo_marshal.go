package persistence

import (
	"time"

	"github.com/maestre3d/newton/internal/event"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/maestre3d/newton/internal/aggregate"
	"github.com/maestre3d/newton/internal/valueobject"
)

type authorDynamo map[string]types.AttributeValue

func marshalAuthorDynamo(a aggregate.Author) authorDynamo {
	return authorDynamo{
		"author_id":    &types.AttributeValueMemberS{Value: a.ID.Value()},
		"display_name": &types.AttributeValueMemberS{Value: a.DisplayName.Value()},
		"create_by":    &types.AttributeValueMemberS{Value: a.CreateBy.Value()},
		"image":        &types.AttributeValueMemberS{Value: a.Image.Value()},
		"create_time":  &types.AttributeValueMemberS{Value: a.Metadata.CreateTime.Format(time.RFC3339)},
		"update_time":  &types.AttributeValueMemberS{Value: a.Metadata.UpdateTime.Format(time.RFC3339)},
		"active":       &types.AttributeValueMemberBOOL{Value: a.Metadata.State},
	}
}

func unmarshalAuthorDynamo(a authorDynamo) (*aggregate.Author, error) {
	if err := ensureAuthorTypesDynamo(a); err != nil {
		return nil, err
	}
	createTime, _ := time.Parse(time.RFC3339, a["create_time"].(*types.AttributeValueMemberS).Value)
	updateTime, _ := time.Parse(time.RFC3339, a["update_time"].(*types.AttributeValueMemberS).Value)
	return &aggregate.Author{
		ID:          valueobject.AuthorID(a["author_id"].(*types.AttributeValueMemberS).Value),
		DisplayName: valueobject.DisplayName(a["display_name"].(*types.AttributeValueMemberS).Value),
		CreateBy:    valueobject.Username(a["create_by"].(*types.AttributeValueMemberS).Value),
		Image:       valueobject.Image(a["image"].(*types.AttributeValueMemberS).Value),
		Metadata: valueobject.Metadata{
			CreateTime: createTime,
			UpdateTime: updateTime,
			State:      a["active"].(*types.AttributeValueMemberBOOL).Value,
		},
		Events: make([]event.DomainEvent, 0),
	}, nil
}

// ensureAuthorTypesDynamo protects malformed AWS DynamoDB aggregate.Author attributes
func ensureAuthorTypesDynamo(a authorDynamo) error {
	_, ok := a["author_id"].(*types.AttributeValueMemberS)
	if !ok {
		return aggregate.ErrAuthorCannotParse
	}
	_, ok = a["display_name"].(*types.AttributeValueMemberS)
	if !ok {
		return aggregate.ErrAuthorCannotParse
	}
	_, ok = a["create_by"].(*types.AttributeValueMemberS)
	if !ok {
		return aggregate.ErrAuthorCannotParse
	}
	_, ok = a["image"].(*types.AttributeValueMemberS)
	if !ok {
		return aggregate.ErrAuthorCannotParse
	}
	_, ok = a["create_time"].(*types.AttributeValueMemberS)
	if !ok {
		return aggregate.ErrAuthorCannotParse
	}
	_, ok = a["update_time"].(*types.AttributeValueMemberS)
	if !ok {
		return aggregate.ErrAuthorCannotParse
	}
	_, ok = a["active"].(*types.AttributeValueMemberBOOL)
	if !ok {
		return aggregate.ErrAuthorCannotParse
	}
	return nil
}

func unmarshalAuthorDynamoBulk(as []map[string]types.AttributeValue) ([]*aggregate.Author, error) {
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
