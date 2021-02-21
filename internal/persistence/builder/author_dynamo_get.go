package builder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/maestre3d/newton/internal/aggregate"
	"github.com/maestre3d/newton/internal/infrastructure"
	"github.com/maestre3d/newton/internal/persistence/model"
)

type authorDynamoGet struct {
	DB     *dynamodb.Client
	Config infrastructure.Configuration
	ID     string
}

func (g *authorDynamoGet) Exec(ctx context.Context) ([]*aggregate.Author, string, error) {
	exp, _ := expression.NewBuilder().WithProjection(model.NewAuthorDynamoProjection()).Build()
	o, err := g.DB.GetItem(ctx, &dynamodb.GetItemInput{
		Key: model.MarshalDynamoKeyWithSort(model.DynamoPartitionKey, model.DynamoSortKey,
			model.AuthorAdjacencyPattern+g.ID),
		TableName:                aws.String(g.Config.DynamoTable),
		ExpressionAttributeNames: exp.Names(),
		ProjectionExpression:     exp.Projection(),
		ConsistentRead:           aws.Bool(false),
	})
	if err != nil {
		return nil, "", err
	} else if o.Item == nil {
		return nil, "", nil
	}

	return g.marshalAggregate(o)
}

func (g *authorDynamoGet) marshalAggregate(o *dynamodb.GetItemOutput) ([]*aggregate.Author, string, error) {
	authorDynamo := model.AuthorDynamo{}
	if err := attributevalue.UnmarshalMap(o.Item, &authorDynamo); err != nil {
		return nil, "", err
	}
	author, err := model.UnmarshalAuthorDynamo(authorDynamo)
	if err != nil {
		return nil, "", err
	}
	return []*aggregate.Author{author}, "", nil
}
