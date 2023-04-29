package services

import (
	"context"
	"github.com/PraveenPin/SwipeMeter/repo"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

type userService struct {
	db *dynamodb.DynamoDB
}

func NewUserService(db *dynamodb.DynamoDB) *userService {
	return &userService{db: db}
}

func (u *userService) AddGroupToUser(ctx context.Context, userReq *AddGroupToUserRequest) (*AddGroupToUserResponse, error) {
	log.Println("Adding group to user value: ", userReq)

	userRepository := repo.UserRepository{}

	resp := &AddGroupToUserResponse{}
	resp.Success = true

	_, err := userRepository.AddGroupToUserMap(u.db, userReq.GetUsername(), userReq.GetGroupId())
	if err != nil {
		log.Fatalf("Error adding groups to user in grpc ", err)
		resp.Success = false
		return resp, err
	}
	return resp, nil
}
func (u *userService) RemoveGroupFromUser(ctx context.Context, userReq *RemoveGroupFromUserRequest) (*RemoveGroupFromUserResponse, error) {
	log.Println("Removing group from user value ", userReq)
	userRepository := repo.UserRepository{}

	resp := &RemoveGroupFromUserResponse{}
	resp.Success = true

	_, err := userRepository.RemoveGroupFromUserMap(u.db, userReq.GetUsername(), userReq.GetGroupId())
	if err != nil {
		log.Fatalf("Error removing groups to user in grpc ", err)
		resp.Success = false
		return resp, err
	}
	return resp, nil
}

func (u *userService) GetAllUserGroupsAndUpdateTotalScore(ctx context.Context, userReq *UserNameRequest) (*UserNameResponse, error) {
	log.Println("Get all groups for the user: ", userReq.GetUsername())
	userRepository := repo.UserRepository{}
	resp := &UserNameResponse{}
	resp.Username = userReq.GetUsername()
	resp.Groups = []string{}

	groups, err := userRepository.GetAllUserGroupsAndUpdateTotalTime(u.db, userReq.GetUsername(), userReq.GetScore())
	if err != nil {
		log.Fatalf("Error fetching all groups to user in grpc ", err)
		return resp, err
	}
	resp.Groups = groups
	return resp, nil

}

func (u *userService) mustEmbedUnimplementedUserServiceServer() {}
