package controllers

import (
	"SwipeMeter/models"
	"SwipeMeter/repo"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

type UserController struct {
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
// func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
func CreateUser(dynamoDbSvc *dynamodb.DynamoDB) {

	var userRepository = repo.UserRepository{}

	//decoder := json.NewDecoder(r.Body)

	//err := decoder.Decode(&user)

	//if err != nil {
	//	response.Format(w, r, true, 417, err)
	//	return
	//}

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
	newUser := models.CreateUserObject("praveenpin-1", "2023-04-15", "praveen123pinjala@gmail.com", 11.1, "")
	created, create_err := userRepository.Create(newUser, "1234", dynamoDbSvc)

	if create_err != nil {
		log.Fatal("Error %v creating user with", create_err, newUser)
		return
	}

	if created {
		log.Println("New User created:", newUser)
		return
	}

}

// func (u *UserController) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
func AuthenticateUser(dynamoDbSvc *dynamodb.DynamoDB) {

	//var authuser models.AuthenticationUser
	var userRepository = repo.UserRepository{}

	newAuthUser := models.CreateAuthUserObject("praveenpin-1", "1234")

	//decoder := json.NewDecoder(r.Body)

	//decoder.Decode(&authuser)

	isSuccessFull, err := userRepository.Authenticate(newAuthUser, dynamoDbSvc)

	fmt.Println("Successfully Authenticated? :", isSuccessFull)

	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	//
	//response.Format(w, r, false, 200, user)
	//return

}
