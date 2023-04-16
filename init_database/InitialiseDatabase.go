package init_database

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
)

type App struct {
}

func (app *App) StartAWSSession() *session.Session {
	log.Println("Initiating aws session to create tables")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return sess
}

func (app *App) GetDynamoDatabaseClient(session *session.Session) *dynamodb.DynamoDB {
	svc := dynamodb.New(session)
	log.Println("Dynamodb client connector obtained")
	return svc
}

func (app *App) InitDatabase(svc *dynamodb.DynamoDB) {
	fmt.Println("Creating database tables")
	createUserTable(svc)
	CreateAuthenticationTable(svc)
	fmt.Println("Tables created")
}

func (app *App) GetS3Connector(session *session.Session) *s3.S3 {
	svc := s3.New(session)
	log.Println("S3 connector obtained")
	return svc
}

func (app *App) GetS3Uploader(session *session.Session) *s3manager.Uploader {
	uploader := s3manager.NewUploader(session)
	log.Println("S3 uploader obtained")
	return uploader
}
