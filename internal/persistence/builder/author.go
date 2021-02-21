package builder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/maestre3d/newton/internal/aggregate"
	"github.com/maestre3d/newton/internal/infrastructure"
	"github.com/maestre3d/newton/internal/persistence/model"
	"github.com/maestre3d/newton/internal/repository"
)

// Author fetching data strategy for Author(s)
type Author interface {
	Exec(ctx context.Context) ([]*aggregate.Author, string, error)
}

// BuildAuthorDynamo allocates an specific fetching strategy based on criteria filter(s) at runtime
//	Strategy pattern
func BuildAuthorDynamo(db *dynamodb.Client, cfg infrastructure.Configuration, criteria repository.Criteria) Author {
	isGetFetch := criteria.Query.FilterExists(model.AuthorIdKey) && criteria.Query.Filters[model.AuthorIdKey].Condition == repository.EqualsCondition

	switch {
	case isGetFetch:
		return &authorDynamoGet{DB: db, Config: cfg, ID: criteria.Query.Filters[model.AuthorIdKey].Value.(string)}
	case criteria.Query.FilterExists(model.CreateByKey):
		return &authorDynamoGetByUser{DB: db, Criteria: criteria}
	default:
		return &authorDynamoList{DB: db, Config: cfg, Criteria: criteria}
	}
}
