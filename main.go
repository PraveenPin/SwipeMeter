package main

import (
	"SwipeMeter/controllers"
	"SwipeMeter/init_database"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello World")
	log.Println("Initiating AWS Session")
	dynamoDBSvc := init_database.GetDynamoDatabaseClient()
	//init_database.InitDatabase(dynamoDBSvc)
	//controllers.CreateUser(dynamoDBSvc)
	controllers.AuthenticateUser(dynamoDBSvc)
}
