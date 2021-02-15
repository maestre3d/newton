package persistence

import "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

var (
	dynamoDefaultPartitionKey = "PK"
	dynamoDefaultSortKey      = "SK"
)

// marshalDynamoKey parses the given field into an AWS DynamoDB Partition Key
func marshalDynamoKey(pk, key string) map[string]types.AttributeValue {
	if key == "" {
		return nil
	}
	return map[string]types.AttributeValue{
		pk: &types.AttributeValueMemberS{Value: key},
	}
}

// marshalDynamoKeyWithSort parses the given fields into AWS DynamoDB Partition and Sort Keys
func marshalDynamoKeyWithSort(pk, sk, key string) map[string]types.AttributeValue {
	if key == "" {
		return nil
	}
	return map[string]types.AttributeValue{
		pk: &types.AttributeValueMemberS{Value: key},
		sk: &types.AttributeValueMemberS{Value: key},
	}
}

// unmarshalDynamoKey parses the given AWS DynamoDB Partition Key into a primtive string
func unmarshalDynamoKey(pk string, keyMap map[string]types.AttributeValue) string {
	k, ok := keyMap[pk].(*types.AttributeValueMemberS)
	if !ok {
		return ""
	}
	return k.Value
}

// unmarshalDynamoKeyWithSort parses the given AWS DynamoDB Partition and Sort Key into primtive strings
func unmarshalDynamoKeyWithSort(pk, sk string, keyMap map[string]types.AttributeValue) (string, string) {
	k, ok := keyMap[pk].(*types.AttributeValueMemberS)
	if !ok {
		return "", ""
	}
	s, ok := keyMap[sk].(*types.AttributeValueMemberS)
	if !ok {
		return k.Value, ""
	}
	return k.Value, s.Value
}
