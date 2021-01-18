package persistence

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/maestre3d/newton/internal/aggregate"
	"github.com/maestre3d/newton/internal/infrastructure"
	"github.com/maestre3d/newton/internal/repository"
	"github.com/maestre3d/newton/internal/valueobject"
)

const authorDynamoPartitionKey = "author_id"

// AuthorDynamo AWS DynamoDB repository.Author implementation/adapter
type AuthorDynamo struct {
	cfg infrastructure.Configuration
	db  *dynamodb.Client
	mu  sync.RWMutex
}

// NewAuthorDynamo allocates a new repository.Author AWS DynamoDB implementation
func NewAuthorDynamo(cfg infrastructure.Configuration, db *dynamodb.Client) *AuthorDynamo {
	return &AuthorDynamo{
		cfg: cfg,
		db:  db,
		mu:  sync.RWMutex{},
	}
}

// Save stores, update or deletes the given record
func (d *AuthorDynamo) Save(ctx context.Context, author aggregate.Author) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	var exec func(context.Context, aggregate.Author) error
	if author.Metadata.MarkAsRemoval {
		exec = d.deleteItem
	} else {
		exec = d.putItem
	}
	return exec(ctx, author)
}

func (d *AuthorDynamo) putItem(ctx context.Context, author aggregate.Author) error {
	_, err := d.db.PutItem(ctx, &dynamodb.PutItemInput{
		Item:      marshalAuthorDynamo(author),
		TableName: aws.String(d.cfg.DynamoTable),
	})
	return err
}

func (d *AuthorDynamo) deleteItem(ctx context.Context, author aggregate.Author) error {
	_, err := d.db.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			authorDynamoPartitionKey: &types.AttributeValueMemberS{Value: author.ID.Value()},
		},
		TableName: aws.String(d.cfg.DynamoTable),
	})
	return err
}

// Get returns the current aggregate if available, returns nil if not found
func (d *AuthorDynamo) Get(ctx context.Context, id valueobject.AuthorID) (*aggregate.Author, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	o, err := d.db.GetItem(ctx, &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			authorDynamoPartitionKey: &types.AttributeValueMemberS{Value: id.Value()},
		},
		TableName:      aws.String(d.cfg.DynamoTable),
		ConsistentRead: aws.Bool(false),
	})
	if err != nil {
		return nil, err
	} else if o.Item == nil {
		return nil, nil
	}
	return unmarshalAuthorDynamo(o.Item)
}

// Search returns a list of the current aggregate filtering and ordering by the given criteria, returns the
// next page token as second argument and returns nil if not found
func (d *AuthorDynamo) Search(ctx context.Context, criteria repository.Criteria) ([]*aggregate.Author, string,
	error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	o, err := d.db.Scan(ctx, &dynamodb.ScanInput{
		TableName:         aws.String(d.cfg.DynamoTable),
		Limit:             aws.Int32(int32(criteria.Limit)),
		ExclusiveStartKey: marshalDynamoNextPage(authorDynamoPartitionKey, criteria.NextPage),
	})
	if err != nil {
		return nil, "", err
	} else if len(o.Items) == 0 {
		return nil, "", nil
	}
	authors, err := unmarshalAuthorDynamoBulk(o.Items)
	if err != nil {
		return nil, "", err
	}
	return authors, unmarshalDynamoNextPage(authorDynamoPartitionKey, o.LastEvaluatedKey), nil
}
