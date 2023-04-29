package repo

import (
	"errors"
	"fmt"
	"github.com/PraveenPin/SwipeMeter/models"
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
	logPrefix := "UserRepository:Create:"
	authUser := models.CreateAuthUserObject(user.Username, password)

	authUserItem, err := dynamodbattribute.MarshalMap(authUser)
	if err != nil {
		log.Fatalf(logPrefix, "Error Marshalling auth user object into map: %s", err)
		return false, nil
	}

	authInput := &dynamodb.PutItemInput{
		Item:      authUserItem,
		TableName: aws.String(authTableName),
	}

	_, err = dynamoDBSvc.PutItem(authInput)
	if err != nil {
		log.Fatalf(logPrefix, "Error inserting new auth user in to the table: %s", err)
		return false, nil
	}
	log.Println(logPrefix, "New Auth User inserted for user:", user.Username)

	av, err := dynamodbattribute.MarshalMap(user)
	fmt.Println(av)
	fmt.Println(user)
	if err != nil {
		log.Fatalf(logPrefix, "Got error marshalling new movie item: %s", err)
		return false, nil
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tablename),
	}

	_, err = dynamoDBSvc.PutItem(input)
	if err != nil {
		log.Fatalf(logPrefix, "Got error calling PutItem: %s", err)
		return false, nil
	}
	log.Println(logPrefix, "New User inserted for user:", user.Username)

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

func (u *UserRepository) AddGroupToUserMap(db *dynamodb.DynamoDB, userName string, groupId string) (bool, error) {

	key := map[string]*dynamodb.AttributeValue{
		"Username": {
			S: aws.String(userName),
		},
	}
	log.Println("Find the user with username :", userName)
	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tablename),
		Key:       key,
	},
	)

	if err != nil {
		msg := fmt.Sprintf("Got error calling GetItem: %s", err)
		return false, errors.New(msg)
	}

	if result.Item == nil {
		msg := "Could not find '" + userName + "'"
		return false, errors.New(msg)
	}

	user := models.User{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		msg := fmt.Sprintf("Failed to unmarshal Record, %v", err)
		return false, errors.New(msg)
	}

	log.Println("Previous list of groups in user", user)
	user.Groups = append(user.Groups, groupId)
	log.Println("Using groupsList.L", user)

	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		msg := fmt.Sprintf("Got error marshalling new user item: %s", err)
		return false, errors.New(msg)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tablename),
	}

	_, err = db.PutItem(input)
	if err != nil {
		msg := fmt.Sprintf("Got error calling PutItem: %s", err)
		return false, errors.New(msg)
	}
	log.Println("New Group:", groupId, "added to the user ", user.Username)
	return true, nil
}

func (u *UserRepository) RemoveGroupFromUserMap(db *dynamodb.DynamoDB, userName string, groupId string) (bool, error) {
	log.Println("Find the user with username :", userName)
	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tablename),
		Key: map[string]*dynamodb.AttributeValue{
			"Username": {
				S: aws.String(userName),
			},
		},
	})
	if err != nil {
		msg := fmt.Sprintf("Got error calling GetItem: %s", err)
		return false, errors.New(msg)
	}

	if result.Item == nil {
		msg := "Could not find '" + userName + "'"
		return false, errors.New(msg)
	}

	user := models.User{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		msg := fmt.Sprintf("Failed to unmarshal Record, %v", err)
		return false, errors.New(msg)
	}

	//remove group from groups list for user
	log.Println("Previous list of groups in user", user)
	//search for same group in the groups list
	found := -1
	for index, group := range user.Groups {
		if group == groupId {
			found = index
			break
		}
	}

	if found == -1 {
		errMsg := fmt.Sprintf("Group does not exist in the user list")
		return false, errors.New(errMsg)
	} else {
		user.Groups[found] = user.Groups[len(user.Groups)-1]
		user.Groups = user.Groups[:len(user.Groups)-1]
	}
	log.Println("New list of groups in user", user)

	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		msg := fmt.Sprintf("Got error marshalling new user item: %s", err)
		return false, errors.New(msg)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tablename),
	}

	_, err = db.PutItem(input)
	if err != nil {
		msg := fmt.Sprintf("Got error calling PutItem: %s", err)
		return false, errors.New(msg)
	}
	log.Println("New Group:", groupId, "added to the user ", user.Username)
	return true, nil
}

func (u *UserRepository) GetAllUserGroupsAndUpdateTotalTime(db *dynamodb.DynamoDB, userName string, score float32) ([]string, error) {
	log.Println("Find the user with username :", userName)
	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tablename),
		Key: map[string]*dynamodb.AttributeValue{
			"Username": {
				S: aws.String(userName),
			},
		},
	})
	if err != nil {
		msg := fmt.Sprintf("Got error calling GetItem: %s", err)
		return []string{}, errors.New(msg)
	}

	if result.Item == nil {
		msg := "Could not find '" + userName + "'"
		return []string{}, errors.New(msg)
	}

	user := models.User{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		msg := fmt.Sprintf("Failed to unmarshal Record, %v", err)
		return []string{}, errors.New(msg)
	}

	log.Println("Old Total Time:", user.Totaltime)
	user.Totaltime += score
	log.Println("New Total Time:", user.Totaltime)

	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		msg := fmt.Sprintf("Got error marshalling user item: %s", err)
		return []string{}, errors.New(msg)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tablename),
	}

	_, err = db.PutItem(input)
	if err != nil {
		msg := fmt.Sprintf("Got error calling PutItem: %s", err)
		return []string{}, errors.New(msg)
	}
	log.Println("Score:", score, "added to the user ", user.Username, user.Totaltime)

	return user.Groups, nil
}
