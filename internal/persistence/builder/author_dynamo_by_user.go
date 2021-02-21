package builder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/maestre3d/newton/internal/aggregate"
	"github.com/maestre3d/newton/internal/repository"
)

type authorDynamoGetByUser struct {
	DB       *dynamodb.Client
	Criteria repository.Criteria
}

func (u *authorDynamoGetByUser) Exec(ctx context.Context) ([]*aggregate.Author, string, error) {
	return nil, "", nil
}
