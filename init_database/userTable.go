package init_database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"log"
)

func createUserTable(dynamoDBSvc *dynamodb.DynamoDB) {

	// Create table Movies
	tableName := "Users"
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Username"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("Creationdate"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Username"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("Creationdate"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}

	_, err := dynamoDBSvc.CreateTable(input)
	if err != nil {
		log.Fatalf("Got error calling CreateTable: %s", err)
	}

	log.Println("Created the table", tableName)
}
