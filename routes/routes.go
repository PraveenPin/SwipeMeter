package routes

import (
	"fmt"
	"github.com/PraveenPin/SwipeMeter/controllers"
	"github.com/PraveenPin/SwipeMeter/services"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

const PORT = ":8080"
const GRPC_PORT = ":9000"

type Dispatcher struct{}

func HomeEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world :)")
}

func (r *Dispatcher) StartGRPCServer(db *dynamodb.DynamoDB) {
	//start a grpc server
	log.Println("Starting GRPC server on port", GRPC_PORT)
	lis, err := net.Listen("tcp", GRPC_PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s", GRPC_PORT)

	grpcServer := grpc.NewServer()
	userServiceServer := services.NewUserService(db)
	services.RegisterUserServiceServer(grpcServer, userServiceServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("Started GRPC server on port", GRPC_PORT)
}

func (r *Dispatcher) Init(db *dynamodb.DynamoDB, s3 *s3.S3) {
	//start grpc server
	go r.StartGRPCServer(db)

	log.Println("Initialize the router")
	router := mux.NewRouter()
	userController := controllers.NewUserController(db, s3, nil)

	router.StrictSlash(true)
	router.HandleFunc("/", HomeEndpoint).Methods("GET")
	// User Resource
	//userRoutes := router.PathPrefix("/users").Subrouter()
	router.HandleFunc("/login", userController.AuthenticateUser).Methods("POST")
	router.HandleFunc("/signup", userController.CreateUser).Methods("POST")

	//Authenticate
	//router.HandleFunc("/authenticateToken", userController.AuthenticateToken).Methods("POST")

	// bind the routes
	http.Handle("/", router)

	log.Println("Add the listener to port ", PORT)

	//serve
	http.ListenAndServe(PORT, nil)
}

func profile(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("test"))
}
