package main

import (
	"github.com/PraveenPin/SwipeMeter/init_database"
	"log"
)

func main() {
	log.Println("Initiating AWS Session")

	app := &init_database.App{}

	session := app.StartAWSSession()
	dynamoDBSvc := app.GetDynamoDatabaseClient(session)
	s3Connector := app.GetS3Connector(session)

	authMgmtClient := app.GetAuth0Client()

	dispatcher := Dispatcher{}
	dispatcher.Init(dynamoDBSvc, s3Connector, authMgmtClient)

	//app.InitDatabase(dynamoDBSvc)
	//controllers.CreateUserController(dynamoDBSvc)
	//controllers.AuthenticateUser(dynamoDBSvc)
	//utils.GetAllS3Objects(s3Connector)
}
