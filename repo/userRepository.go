package repo

import (
	"SwipeMeter/models"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
)

const (
	authTableName = "Authentication"
	tablename     = "Users"
)

type UserInterface interface {
	GetAll() ([]models.User, error)
	GetOne() (models.User, error)
	Destroy(id int) (bool, error)
	Update([]string) (bool, error)
	Create(user models.User) (bool, error)
	Authenticate(user models.User) bool
}

type UserRepository struct{}

//func (u *UserRepository) GetAll() ([]models.User, error) {
//}

// Get One
//func (u *UserRepository) GetOne(vars map[string]string) (models.User, error) {
//}

// Destroy
//func (u *UserRepository) Destroy(vars map[string]string) (bool, error) {
//}

// Update
//func (u *UserRepository) Update(vars map[string]string) (bool, error) {
//}

func (u *UserRepository) Create(user models.User, password string, dynamoDBSvc *dynamodb.DynamoDB) (bool, error) {

	authUser := models.CreateAuthUserObject(user.Username, password)

	authUserItem, err := dynamodbattribute.MarshalMap(authUser)
	if err != nil {
		log.Fatalf("Error Marshalling auth user object into map: %s", err)
		return false, nil
	}

	authInput := &dynamodb.PutItemInput{
		Item:      authUserItem,
		TableName: aws.String(authTableName),
	}

	_, err = dynamoDBSvc.PutItem(authInput)
	if err != nil {
		log.Fatalf("Error inserting new auth user in to the table: %s", err)
		return false, nil
	}
	log.Println("New Auth User inserted for user:", user.Username)

	av, err := dynamodbattribute.MarshalMap(user)
	fmt.Println(av)
	fmt.Println(user)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
		return false, nil
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tablename),
	}

	_, err = dynamoDBSvc.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
		return false, nil
	}
	log.Println("New User inserted for user:", user.Username)

	return true, nil
}

func (u *UserRepository) Authenticate(authUser models.AuthenticationUser, dynamoDBSvc *dynamodb.DynamoDB) (bool, error) {

	result, err := dynamoDBSvc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(authTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Username": {
				S: aws.String(authUser.Username),
			},
		},
	})
	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
		return false, nil
	}

	if result.Item == nil {
		msg := "Could not find '" + authUser.Username + "'"
		return false, errors.New(msg)
	}

	authenticatedUser := models.AuthenticationUser{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &authenticatedUser)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Failed to unmarshal Record, %v", err))
		return false, nil
	}

	if authenticatedUser.Password == authUser.Password {
		return true, nil
	}

	return false, errors.New("Authentication Failed")
}
