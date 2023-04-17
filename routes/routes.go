package routes

import (
	"fmt"
	"github.com/PraveenPin/SwipeMeter/controllers"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Dispatcher struct{}

func HomeEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world :)")
}

func (r *Dispatcher) Init(db *dynamodb.DynamoDB, s3 *s3.S3) {
	log.Println("Initialize the router")
	router := mux.NewRouter()
	userController := &controllers.UserController{}
	userController.SetDynamoDbClient(db)
	userController.SetS3ConnectorClient(s3)

	router.StrictSlash(true)
	router.HandleFunc("/", HomeEndpoint).Methods("GET")
	// User Resource
	//userRoutes := router.PathPrefix("/users").Subrouter()
	router.HandleFunc("/login", userController.AuthenticateUser).Methods("POST")
	router.HandleFunc("/signup", userController.CreateUser).Methods("POST")

	//Authenticate
	//userRoutes.HandleFunc("/authenticate", UserController.Authenticate).Methods("POST")

	// bind the routes
	http.Handle("/", router)

	log.Println("Add the listener to port 8080")

	//serve
	http.ListenAndServe(":8080", nil)
}

func profile(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("test"))
}
