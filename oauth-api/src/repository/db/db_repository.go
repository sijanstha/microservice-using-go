package db

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/sijanstha/oauth-api/src/client"
	"github.com/sijanstha/oauth-api/src/domain/access_token"
	"github.com/sijanstha/oauth-api/src/utils/errors"
	"strconv"
)

const (
	tableName = "AccessTokens"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(token *access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(token access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func (r *dbRepository) UpdateExpirationTime(accessToken access_token.AccessToken) *errors.RestErr {
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				N: aws.String(strconv.Itoa(int(accessToken.Expires))),
			},
			":access_token": {
				S: aws.String(accessToken.AccessToken),
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"access_token": {
				S: aws.String(accessToken.AccessToken),
			},
		},
		ConditionExpression: aws.String("access_token = :access_token"),
		ReturnValues:        aws.String("NONE"),
		UpdateExpression:    aws.String("set expires = :r"),
	}

	_, err := client.DynamoDBConnection.UpdateItem(input)
	if err != nil {
		return errors.NewNotFoundError(fmt.Sprintf("Access token not found: %s", accessToken.AccessToken))
	}

	return nil
}

func (r *dbRepository) Create(accessToken *access_token.AccessToken) *errors.RestErr {
	av, err := dynamodbattribute.MarshalMap(accessToken)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Got error marshalling new movie item: %s", err.Error()))
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = client.DynamoDBConnection.PutItem(input)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Got error calling PutItem: %s", err.Error()))
	}
	return nil
}

func (r *dbRepository) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {
	result, err := client.DynamoDBConnection.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"access_token": {
				S: aws.String(accessTokenId),
			},
		},
	})
	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("Got error calling GetItem: %s", err.Error()))
	}

	if result.Item == nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("Could not find access token with id: %s", accessTokenId))
	}

	var accessToken = &access_token.AccessToken{}

	err = dynamodbattribute.UnmarshalMap(result.Item, accessToken)
	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	return accessToken, nil
}
