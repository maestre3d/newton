package persistence

import "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

func marshalDynamoNextPage(pk, nextPage string) map[string]types.AttributeValue {
	if nextPage == "" {
		return nil
	}
	return map[string]types.AttributeValue{
		pk: &types.AttributeValueMemberS{Value: nextPage},
	}
}

func unmarshalDynamoNextPage(pk string, lastKey map[string]types.AttributeValue) string {
	k, ok := lastKey[pk].(*types.AttributeValueMemberS)
	if !ok {
		return ""
	}
	return k.Value
}
