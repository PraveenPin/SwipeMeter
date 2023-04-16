package main

import (
	"SwipeMeter/init_database"
	"SwipeMeter/utils"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello World")
	log.Println("Initiating AWS Session")
	session := init_database.StartAWSSession()
	//dynamoDBSvc := init_database.GetDynamoDatabaseClient(session)
	s3Connector := init_database.GetS3Connector(session)
	//init_database.InitDatabase(dynamoDBSvc)
	//controllers.CreateUser(dynamoDBSvc)
	//controllers.AuthenticateUser(dynamoDBSvc)
	utils.GetAllS3Objects(s3Connector)
}
