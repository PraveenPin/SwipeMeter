package controllers

import (
	"encoding/json"
	"github.com/PraveenPin/SwipeMeter/models"
	"github.com/PraveenPin/SwipeMeter/repo"
	"github.com/PraveenPin/SwipeMeter/services"
	"github.com/PraveenPin/SwipeMeter/utils"
	"github.com/auth0/go-auth0/management"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"net/http"
	"time"
)

const FILE_NAME = "UserController:"

var response *utils.Response

type UserController struct {
	userService *services.UserService
}

func NewUserController(dynamodbSVC *dynamodb.DynamoDB, s3Connector *s3.S3, s3Uploader *s3manager.Uploader, authClient *management.Management) *UserController {
	userRepository := &repo.UserRepository{}
	userService := services.NewUserService(dynamodbSVC, s3Connector, s3Uploader, authClient, userRepository)
	return &UserController{userService}
}

func (u *UserController) CreateUserController(w http.ResponseWriter, r *http.Request) {
	log.Println(FILE_NAME, "Create User Request: ", r)
	decoder := json.NewDecoder(r.Body)

	newSignUpUser := models.SignUpUser{}
	err := decoder.Decode(&newSignUpUser)
	if err != nil {
		response.Format(w, r, true, 417, err)
		return
	}

	log.Println(FILE_NAME, "User object is :", newSignUpUser)

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

	_, create_auth_err := u.userService.CreateAuthUserService(newSignUpUser)
	if create_auth_err != nil {
		log.Fatalf(FILE_NAME, create_auth_err)
		response.Format(w, r, true, 418, create_auth_err)
		return
	}

	log.Println(FILE_NAME, "Auth User Successfully created")

	newUser := models.User{
		Username:     newSignUpUser.Username,
		Email:        newSignUpUser.Email,
		Creationdate: time.Now().Format(time.RFC850),
		Totaltime:    0.0,
	}

	_, create_err := u.userService.CreateUserService(newUser)

	if create_err != nil {
		log.Fatalf(FILE_NAME, "Error %v creating user with", create_err, newUser)
		response.Format(w, r, true, 418, create_err)
		return
	}

	log.Println(FILE_NAME, "New User created:", newUser)
	response.Format(w, r, false, 201, newUser)

}
