package persistence

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/maestre3d/newton/internal/aggregate"
	"github.com/maestre3d/newton/internal/infrastructure"
	"github.com/maestre3d/newton/internal/persistence/builder"
	"github.com/maestre3d/newton/internal/persistence/model"
	"github.com/maestre3d/newton/internal/repository"
	"github.com/maestre3d/newton/internal/valueobject"
)

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
	if author.Metadata.MarkAsRemoval {
		return d.deleteItem(ctx, author)
	}
	return d.putItem(ctx, author)
}

func (d *AuthorDynamo) putItem(ctx context.Context, author aggregate.Author) error {
	authorDb, err := attributevalue.MarshalMap(model.MarshalAuthorDynamo(author))
	if err != nil {
		return err
	}
	_, err = d.db.PutItem(ctx, &dynamodb.PutItemInput{
		Item:      authorDb,
		TableName: aws.String(d.cfg.DynamoTable),
	})
	return err
}

func (d *AuthorDynamo) deleteItem(ctx context.Context, author aggregate.Author) error {
	_, err := d.db.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		Key: model.MarshalDynamoKeyWithSort(model.DynamoPartitionKey, model.DynamoSortKey,
			model.AuthorAdjacencyPattern+author.ID.Value()),
		TableName: aws.String(d.cfg.DynamoTable),
	})
	return err
}

// Get returns the current aggregate if available, returns nil if not found
func (d *AuthorDynamo) Get(ctx context.Context, id valueobject.AuthorID) (*aggregate.Author, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	build := builder.BuildAuthorDynamo(d.db, d.cfg, d.constructCriteriaGetUser(id))
	authors, _, err := build.Exec(ctx)
	if err != nil {
		return nil, err
	} else if len(authors) < 1 {
		return nil, nil
	}
	return authors[0], nil
}

func (d *AuthorDynamo) constructCriteriaGetUser(id valueobject.AuthorID) repository.Criteria {
	return repository.Criteria{
		Query: repository.Query{
			Filters: map[string]repository.Filter{
				model.AuthorIdKey: {Condition: repository.EqualsCondition, Value: id.Value()},
			},
		},
	}
}

// Search returns a list of the current aggregate filtering and ordering by the given criteria, returns the
// next page token as second argument and returns nil if not found
func (d *AuthorDynamo) Search(ctx context.Context, criteria repository.Criteria) ([]*aggregate.Author, string,
	error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	build := builder.BuildAuthorDynamo(d.db, d.cfg, criteria)
	return build.Exec(ctx)
}
