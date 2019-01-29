# gomockAWS

go mocks for AWS golang classes

## Project Code Example

```go
package example

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"strconv"
)

var d dynamodbiface.DynamoDBAPI

func init() {
	sess := session.Must(session.NewSession())
	d = dynamodb.New(sess)
}


func doDb(table, value1, value2 string) (value int64, err error){
	f := &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]*dynamodb.AttributeValue{
			"field1": {
				S: aws.String(value1),
			},

			"field2":
			{
				S: aws.String(value2),
			},
		},
	}
	o, err := d.GetItem(f)
	n := o.Item["value"].N
	return strconv.ParseInt(*n,10,64)
}

```

## Test Code Example

```go
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

```

