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
