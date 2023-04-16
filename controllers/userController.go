package controllers

import (
	"SwipeMeter/models"
	"SwipeMeter/repo"
	"SwipeMeter/utils"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"net/http"
)

var response *utils.Response

type UserControllerInterface interface {
	setDynamoDbClient(db *dynamodb.DynamoDB)
	setS3ConnectorClient(s3 *s3.S3)
	setS3UploaderClient(s3 *s3manager.Uploader)
	getDynamoDbClient()
	getS3ConnectorClient()
	getS3UploaderClient()
}

type UserController struct {
	dynamodbSVC *dynamodb.DynamoDB
	s3Connector *s3.S3
	s3Uploader  *s3manager.Uploader
}

func (u *UserController) SetDynamoDbClient(db *dynamodb.DynamoDB) {
	u.dynamodbSVC = db
}

func (u *UserController) SetS3ConnectorClient(s3 *s3.S3) {
	u.s3Connector = s3
}

func (u *UserController) setS3UploaderClient(s3 *s3manager.Uploader) {
	u.s3Uploader = s3
}

func (u *UserController) getDynamoDbClient() *dynamodb.DynamoDB {
	return u.dynamodbSVC
}

func (u *UserController) getS3ConnectorClient() *s3.S3 {
	return u.s3Connector
}

func (u *UserController) getS3UploaderClient() *s3manager.Uploader {
	return u.s3Uploader
}

// Get All Users
//func (u *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
//
//	var uRepo = repo.UserRepository{}
//	res,err := uRepo.GetAll()
//
//	// Error occured
//	if err != nil {
//		response.Format(w, r, true, 400, err)
//	}
//
//	response.Format(w, r, false, 200, res)
//}

//Get One By ID
//func (u *UserController) GetOne(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	var uRepo = repo.UserRepository{}
//	var user,err = uRepo.GetOne(vars)
//
//	res, err := json.Marshal(user)
//
//	if err != nil {
//		w.Write([]byte("Error"))
//	}
//
//	w.Write([]byte(res))
//
//}

// Destroy user
//func (u *UserController) Destroy(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	var uRepo = repo.UserRepository{}
//	var user,err = uRepo.Destroy(vars)
//
//	res, err := json.Marshal(user)
//
//	if err != nil {
//		w.Write([]byte("Error"))
//	}
//
//	w.Write([]byte(res))
//
//}

// Update User
//func (u *UserController) Update(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	var uRepo = repo.UserRepository{}
//	var user,err = uRepo.Update(vars)
//
//	res, err := json.Marshal(user)
//
//	if err != nil {
//		w.Write([]byte("Error"))
//	}
//
//	w.Write([]byte(res))
//
//}

// CreateUser User
func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userRepository = repo.UserRepository{}
	log.Println("Create User Request: ", r)
	decoder := json.NewDecoder(r.Body)

	newUser := models.User{}
	err := decoder.Decode(&newUser)
	if err != nil {
		response.Format(w, r, true, 417, err)
		return
	}

	log.Println("User object is :", newUser)

	//config := validator.Config{
	//	TagName:         "validate",
	//	ValidationFuncs: validator.BakedInValidators,
	//}

	//validate = validator.New(config)

	//errs := validate.Struct(user)

	//if errs != nil {
	//	response.Format(w, r, true, 417, errs)
	//	return
	//}
	//newUser := models.CreateUserObject("praveenpin-1", "2023-04-15", "praveen123pinjala@gmail.com", 11.1, "")
	created, create_err := userRepository.Create(newUser, "1234", u.getDynamoDbClient())

	if create_err != nil {
		log.Fatal("Error %v creating user with", create_err, newUser)
		response.Format(w, r, true, 418, create_err)
		return
	}

	if created {
		log.Println("New User created:", newUser)
		response.Format(w, r, false, 200, newUser)
		return
	}

	return

}

func (u *UserController) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Authenticate Request: ", r)

	var userRepository = &repo.UserRepository{}
	newAuthUser := models.AuthenticationUser{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&newAuthUser)

	isSuccessFull, err := userRepository.Authenticate(newAuthUser, u.getDynamoDbClient())

	fmt.Println("Successfully Authenticated? :", isSuccessFull)

	if err != nil {
		log.Fatalf(err.Error())
		response.Format(w, r, true, 418, err)
		return
	}

	response.Format(w, r, false, 200, newAuthUser.Username)
	return
}
