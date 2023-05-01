package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/PraveenPin/SwipeMeter/models"
	"github.com/PraveenPin/SwipeMeter/repo"
	"github.com/auth0/go-auth0/management"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"gopkg.in/auth0.v5"
	"log"
)

const FILE_NAME = "UserService:"

type UserService struct {
	db          *dynamodb.DynamoDB
	s3Connector *s3.S3
	s3Uploader  *s3manager.Uploader
	authClient  *management.Management
	userRepo    *repo.UserRepository
}

func NewUserService(db *dynamodb.DynamoDB, s3Connector *s3.S3, s3Uploader *s3manager.Uploader, authClient *management.Management, userRepo *repo.UserRepository) *UserService {
	return &UserService{db: db, s3Connector: s3Connector, s3Uploader: s3Uploader, authClient: authClient, userRepo: userRepo}
}

func (u *UserService) AddGroupToUser(ctx context.Context, userReq *AddGroupToUserRequest) (*AddGroupToUserResponse, error) {
	log.Println(FILE_NAME, "Adding group to user value: ", userReq)

	resp := &AddGroupToUserResponse{}
	resp.Success = true

	_, err := u.userRepo.AddGroupToUserMap(u.db, userReq.GetUsername(), userReq.GetGroupId())
	if err != nil {
		log.Fatalf(FILE_NAME, "Error adding groups to user in grpc ", err)
		resp.Success = false
		return resp, err
	}
	return resp, nil
}
func (u *UserService) RemoveGroupFromUser(ctx context.Context, userReq *RemoveGroupFromUserRequest) (*RemoveGroupFromUserResponse, error) {
	log.Println(FILE_NAME, "Removing group from user value ", userReq)

	resp := &RemoveGroupFromUserResponse{}
	resp.Success = true

	_, err := u.userRepo.RemoveGroupFromUserMap(u.db, userReq.GetUsername(), userReq.GetGroupId())
	if err != nil {
		log.Fatalf(FILE_NAME, "Error removing groups to user in grpc ", err)
		resp.Success = false
		return resp, err
	}
	return resp, nil
}

func (u *UserService) GetAllUserGroupsAndUpdateTotalScore(ctx context.Context, userReq *UserNameRequest) (*UserNameResponse, error) {
	log.Println("Get all groups for the user: ", userReq.GetUsername())

	resp := &UserNameResponse{}
	resp.Username = userReq.GetUsername()
	resp.Groups = []string{}

	groups, err := u.userRepo.GetAllUserGroupsAndUpdateTotalTime(u.db, userReq.GetUsername(), userReq.GetScore())
	if err != nil {
		log.Fatalf(FILE_NAME, "Error fetching all groups to user in grpc ", err)
		return resp, err
	}
	resp.Groups = groups
	return resp, nil

}

func (u *UserService) CreateUserService(newUser models.User) (bool, error) {
	_, create_err := u.userRepo.Create(newUser, u.db)

	if create_err != nil {
		log.Fatal(FILE_NAME, "Error %v creating user with", create_err, newUser)
		return false, create_err
	}

	return true, nil
}

func (u *UserService) CreateAuthUserService(newSignUpUser models.SignUpUser) (bool, error) {

	log.Println(FILE_NAME, "New request to add user to auth0 account", newSignUpUser)

	//Create a new user object
	var user = &management.User{
		Email:       auth0.String(newSignUpUser.Email),
		Username:    auth0.String(newSignUpUser.Username),
		Password:    auth0.String(newSignUpUser.Password),
		Connection:  auth0.String("Username-Password-Authentication"),
		VerifyEmail: auth0.Bool(false),
	}

	err := u.authClient.User.Create(user)
	if err != nil {
		msg := fmt.Sprintf(FILE_NAME, "Error adding user to auth0 database", err)
		return false, errors.New(msg)
	}

	log.Println(FILE_NAME, "Fetching users", err)
	// Query the list of users
	ul, err := u.authClient.User.List()
	if err != nil {
		fmt.Println("Failed to list users:", err)
	}

	// Print out the list of users
	for _, user := range ul.Users {
		fmt.Println(user)
	}

	return true, nil
}

func (u *UserService) mustEmbedUnimplementedUserServiceServer() {}
