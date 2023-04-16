package init_database

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
)

func StartAWSSession() *session.Session {
	log.Println("Initiating aws session to create tables")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return sess
}

func GetDynamoDatabaseClient(session *session.Session) *dynamodb.DynamoDB {

	log.Println("Obtaining dynamo db client connector to create tables")
	svc := dynamodb.New(session)

	return svc
}

func InitDatabase(svc *dynamodb.DynamoDB) {
	fmt.Println("Creating database tables")
	createUserTable(svc)
	CreateAuthenticationTable(svc)
	fmt.Println("Tables created")
}

func GetS3Connector(session *session.Session) *s3.S3 {
	log.Println("Obtaining s3 connector ")
	svc := s3.New(session)
	return svc
}

func GetS3Uploader(session *session.Session) *s3manager.Uploader {
	log.Println("Obtaining s3 uploader")
	uploader := s3manager.NewUploader(session)
	return uploader
}
