package builder

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/maestre3d/newton/internal/aggregate"
	"github.com/maestre3d/newton/internal/infrastructure"
	"github.com/maestre3d/newton/internal/persistence/model"
	"github.com/maestre3d/newton/internal/repository"
)

type authorDynamoList struct {
	DB       *dynamodb.Client
	Config   infrastructure.Configuration
	Criteria repository.Criteria

	b   expression.Builder
	exp expression.Expression
}

func (l *authorDynamoList) Exec(ctx context.Context) ([]*aggregate.Author, string, error) {
	var err error
	l.exp, err = expression.NewBuilder().WithFilter(l.generateSearch()).
		WithProjection(model.NewAuthorDynamoProjection()).Build()
	if err != nil {
		return nil, "", err
	}
	return l.exec(ctx)
}

func (l *authorDynamoList) generateSearch() expression.ConditionBuilder {
	partitionExp := expression.Name(model.DynamoPartitionKey).BeginsWith(model.AuthorAdjacencyPattern)
	sortExp := expression.Name(model.DynamoSortKey).BeginsWith(model.AuthorAdjacencyPattern)

	other := make([]expression.ConditionBuilder, 0)
	if l.Criteria.ActiveOnly {
		other = append(other, expression.Name(model.ActiveFlagKey).Equal(expression.Value(l.Criteria.ActiveOnly)))
	}
	if len(l.Criteria.Query.Filters) >= 1 {
		other = append(other, BuildDynamoQuery(l.Criteria))
	}

	return expression.And(partitionExp, sortExp, other...)
}

func (l *authorDynamoList) exec(ctx context.Context) ([]*aggregate.Author, string, error) {
	o, err := l.DB.Scan(ctx, l.newScanInput())
	if err != nil {
		return nil, "", err
	} else if len(o.Items) == 0 {
		return nil, "", nil
	}
	return l.marshalAggregate(o)
}

func (l *authorDynamoList) newScanInput() *dynamodb.ScanInput {
	return &dynamodb.ScanInput{
		TableName:                 aws.String(l.Config.DynamoTable),
		ExpressionAttributeNames:  l.exp.Names(),
		ExpressionAttributeValues: l.exp.Values(),
		FilterExpression:          l.exp.Filter(),
		ProjectionExpression:      l.exp.Projection(),
		ExclusiveStartKey: model.MarshalDynamoKeyWithSort(model.DynamoPartitionKey, model.DynamoSortKey,
			model.AuthorAdjacencyPattern+l.Criteria.NextPage),
	}
}

func (l *authorDynamoList) marshalAggregate(o *dynamodb.ScanOutput) ([]*aggregate.Author, string, error) {
	authorsDynamo := make([]model.AuthorDynamo, 0)
	if err := attributevalue.UnmarshalListOfMaps(o.Items, &authorsDynamo); err != nil {
		return nil, "", err
	}
	authors, err := model.UnmarshalAuthorDynamoBulk(authorsDynamo)
	if err != nil {
		return nil, "", err
	}
	if len(authors) > l.Criteria.Limit {
		return authors[:l.Criteria.Limit], string(authors[l.Criteria.Limit-1].ID), nil // enables local pagination
	}
	return authors, l.marshalPaginationKey(o), nil
}

func (l *authorDynamoList) marshalPaginationKey(o *dynamodb.ScanOutput) string {
	key := model.UnmarshalDynamoKey(model.DynamoPartitionKey,
		o.LastEvaluatedKey)
	cleanedKey := strings.Split(key, model.AuthorAdjacencyPattern)
	return strings.Join(cleanedKey, "")
}
