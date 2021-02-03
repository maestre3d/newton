package persistence

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/maestre3d/newton/internal/aggregate"
	"github.com/maestre3d/newton/internal/infrastructure"
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

var (
	// authorAdjacencyPattern string pattern for author schemas using Amazon Web Services DynamoDB's
	// Adjacency List pattern.
	//
	// A more detailed information can be found here:
	// https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/bp-adjacency-graphs.html
	authorAdjacencyPattern = "author#"
)

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
	authorDb, err := attributevalue.MarshalMap(marshalAuthorDynamo(author))
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
		Key: marshalDynamoKeyWithSort(dynamoDefaultPartitionKey, dynamoDefaultSortKey,
			authorAdjacencyPattern+author.ID.Value()),
		TableName: aws.String(d.cfg.DynamoTable),
	})
	return err
}

// Get returns the current aggregate if available, returns nil if not found
func (d *AuthorDynamo) Get(ctx context.Context, id valueobject.AuthorID) (*aggregate.Author, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	exp, _ := expression.NewBuilder().WithProjection(d.newProjectionExpression()).Build()
	o, err := d.db.GetItem(ctx, &dynamodb.GetItemInput{
		Key: marshalDynamoKeyWithSort(dynamoDefaultPartitionKey, dynamoDefaultSortKey,
			authorAdjacencyPattern+id.Value()),
		TableName:                aws.String(d.cfg.DynamoTable),
		ExpressionAttributeNames: exp.Names(),
		ProjectionExpression:     exp.Projection(),
		ConsistentRead:           aws.Bool(false),
	})
	if err != nil {
		return nil, err
	} else if o.Item == nil {
		return nil, nil
	}
	author := authorDynamoSchema{}
	if err = attributevalue.UnmarshalMap(o.Item, &author); err != nil {
		return nil, err
	}
	return unmarshalAuthorDynamo(author)
}

func (d *AuthorDynamo) newProjectionExpression() expression.ProjectionBuilder {
	return expression.NamesList(expression.Name("PK"), expression.Name("SK"), expression.Name("DisplayName"), expression.Name("CreateTime"),
		expression.Name("UpdateTime"), expression.Name("Active"), expression.Name("CreateBy"), expression.Name("Image"))
}

// Search returns a list of the current aggregate filtering and ordering by the given criteria, returns the
// next page token as second argument and returns nil if not found
func (d *AuthorDynamo) Search(ctx context.Context, criteria repository.Criteria) ([]*aggregate.Author, string,
	error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	exp, _ := expression.NewBuilder().WithFilter(d.newSearchExpression()).
		WithProjection(d.newProjectionExpression()).Build()
	o, err := d.db.Scan(ctx, &dynamodb.ScanInput{
		TableName:                 aws.String(d.cfg.DynamoTable),
		Limit:                     aws.Int32(int32(criteria.Limit)),
		ExpressionAttributeNames:  exp.Names(),
		ExpressionAttributeValues: exp.Values(),
		FilterExpression:          exp.Filter(),
		ProjectionExpression:      exp.Projection(),
		ExclusiveStartKey:         marshalDynamoKey(dynamoDefaultPartitionKey, criteria.NextPage),
	})
	if err != nil {
		return nil, "", err
	} else if len(o.Items) == 0 {
		return nil, "", nil
	}
	authorsPrimitive := make([]authorDynamoSchema, 0)
	if err = attributevalue.UnmarshalListOfMaps(o.Items, &authorsPrimitive); err != nil {
		return nil, "", err
	}
	authors, err := unmarshalAuthorDynamoBulk(authorsPrimitive)
	if err != nil {
		return nil, "", err
	}
	return authors, unmarshalDynamoKey(dynamoDefaultPartitionKey, o.LastEvaluatedKey), nil
}

func (d *AuthorDynamo) newSearchExpression() expression.ConditionBuilder {
	partitionExp := expression.Name(dynamoDefaultPartitionKey).BeginsWith(authorAdjacencyPattern)
	sortExp := expression.Name(dynamoDefaultSortKey).BeginsWith(authorAdjacencyPattern)
	return expression.And(partitionExp, sortExp)
}
