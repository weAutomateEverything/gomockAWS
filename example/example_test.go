package example

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/weAutomateEverything/gomockAWS/dynamodbMock"
	"testing"
)

func TestBusiness(t *testing.T) {
	mock := gomock.NewController(t)

	dynamoMock := dynamodbMock.NewMockDynamoDBAPI(mock)

	//overwrite the dynamo db client in example.go with the mock
	d = dynamoMock

	//What to expect and return
	dynamoMock.EXPECT().GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("testTable"),
		Key: map[string]*dynamodb.AttributeValue{
			"field1": {
				S: aws.String("value1"),
			},

			"field2":
			{
				S: aws.String("value2"),
			},
		},
	}).Return(&dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			"value": {
				N: aws.String("100"),
			},
		},
	}, nil)

	out, err := doDb("testTable", "value1", "value2")
	assert.Nil(t,err)
	assert.Equal(t,int64(100),out)
}
