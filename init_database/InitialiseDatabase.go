package init_database

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

func GetDynamoDatabaseClient() *dynamodb.DynamoDB {
	log.Println("Initiating aws session to create tables")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	return svc
}

func InitDatabase(svc *dynamodb.DynamoDB) {
	fmt.Println("Creating database tables")
	createUserTable(svc)
	CreateAuthenticationTable(svc)
	fmt.Println("Tables created")
}
