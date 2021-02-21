package builder

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/maestre3d/newton/internal/repository"
)

// BuildDynamoQuery builds an Amazon Web Services (AWS) DynamoDB condition builder from a Criteria struct
// to be used along with a expression.Builder from the AWS SDK v2
func BuildDynamoQuery(criteria repository.Criteria) expression.ConditionBuilder {
	conditions := make([]expression.ConditionBuilder, 0)
	for key, filter := range criteria.Query.Filters {
		conditions = append(conditions, newDynamoCondition(key, filter))
	}

	if len(conditions) > 1 && criteria.Query.IsOr() {
		return expression.Or(conditions[0], conditions[1], conditions[1:]...)
	} else if len(conditions) > 1 {
		return expression.And(conditions[0], conditions[1], conditions[1:]...)
	} else if len(conditions) == 1 {
		return conditions[0] // avoid nil references, thus, we avoid panics at runtime
	}
	return expression.ConditionBuilder{}
}

// newCondition generates an AWS DynamoDB condition builder from the given filter
func newDynamoCondition(key string, filter repository.Filter) expression.ConditionBuilder {
	switch filter.Condition {
	case repository.BeginsWithCondition:
		b := expression.Name(key).BeginsWith(filter.Value.(string))
		return addDynamoNegateCondition(b, filter)
	case repository.BetweenCondition:
		b := expression.Name(key).Between(expression.Value(filter.Value),
			expression.Value(filter.AltValue))
		return addDynamoNegateCondition(b, filter)
	case repository.EqualsCondition:
		b := expression.Name(key).Equal(expression.Value(filter.Value))
		return addDynamoNegateCondition(b, filter)
	case repository.GreaterCondition:
		b := expression.Name(key).GreaterThan(expression.Value(filter.Value))
		return addDynamoNegateCondition(b, filter)
	case repository.LessCondition:
		b := expression.Name(key).LessThan(expression.Value(filter.Value))
		return addDynamoNegateCondition(b, filter)
	case repository.EqualOrGreaterCondition:
		b := expression.Name(key).GreaterThanEqual(expression.Value(filter.Value))
		return addDynamoNegateCondition(b, filter)
	case repository.EqualOrLessCondition:
		b := expression.Name(key).LessThanEqual(expression.Value(filter.Value))
		return addDynamoNegateCondition(b, filter)
	case repository.ContainsCondition:
		b := expression.Name(key).Contains(filter.Value.(string))
		return addDynamoNegateCondition(b, filter)
	case repository.InCondition:
		b := expression.Name(key).In(expression.Value(filter.Value))
		return addDynamoNegateCondition(b, filter)
	default:
		return expression.ConditionBuilder{}
	}
}

// addDynamoNegateCondition adds a NEGATE statement to the given condition builder if filter specifies it
func addDynamoNegateCondition(b expression.ConditionBuilder, filter repository.Filter) expression.ConditionBuilder {
	if filter.Negate {
		return b.Not()
	}
	return b
}
