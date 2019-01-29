mkdir dynamodb
mockgen -package dynamodb -destination dynamodb/DynamoDBAPI_mock.go github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface DynamoDBAPI