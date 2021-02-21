package model

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// MarshalDynamoKey parses the given field into an AWS DynamoDB Partition Key
func MarshalDynamoKey(pk, key string) map[string]types.AttributeValue {
	if key == "" {
		return nil
	}
	return map[string]types.AttributeValue{
		pk: &types.AttributeValueMemberS{Value: key},
	}
}

// MarshalDynamoKeyWithSort parses the given fields into AWS DynamoDB Partition and Sort Keys
func MarshalDynamoKeyWithSort(pk, sk, key string) map[string]types.AttributeValue {
	if key == "" {
		return nil
	}
	return map[string]types.AttributeValue{
		pk: &types.AttributeValueMemberS{Value: key},
		sk: &types.AttributeValueMemberS{Value: key},
	}
}

// UnmarshalDynamoKey parses the given AWS DynamoDB Partition Key into a primtive string
func UnmarshalDynamoKey(pk string, keyMap map[string]types.AttributeValue) string {
	k, ok := keyMap[pk].(*types.AttributeValueMemberS)
	if !ok {
		return ""
	}
	return k.Value
}

// UnmarshalDynamoKeyWithSort parses the given AWS DynamoDB Partition and Sort Key into primtive strings
func UnmarshalDynamoKeyWithSort(pk, sk string, keyMap map[string]types.AttributeValue) (string, string) {
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
