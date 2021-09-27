package client

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	DynamoDBConnection *dynamodb.DynamoDB
)

func init() {
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           "devuser",
	}))
	DynamoDBConnection = dynamodb.New(session)
}
