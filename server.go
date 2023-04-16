package main

import (
	"SwipeMeter/init_database"
	"SwipeMeter/routes"
	"log"
)

func main() {
	log.Println("Initiating AWS Session")

	app := &init_database.App{}

	session := app.StartAWSSession()
	dynamoDBSvc := app.GetDynamoDatabaseClient(session)
	s3Connector := app.GetS3Connector(session)

	dispatcher := routes.Dispatcher{}
	dispatcher.Init(dynamoDBSvc, s3Connector)

	//app.InitDatabase(dynamoDBSvc)
	//controllers.CreateUser(dynamoDBSvc)
	//controllers.AuthenticateUser(dynamoDBSvc)
	//utils.GetAllS3Objects(s3Connector)
}
