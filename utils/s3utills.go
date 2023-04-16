package utils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

const (
	S3Bucket = "swipemeter"
)

func GetAllS3Objects(s3connector *s3.S3) {
	resp, err := s3connector.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(S3Bucket),
		Prefix: aws.String("dps/"),
	})
	if err != nil {
		log.Fatalf("Unable to list items in bucket %q, %v", S3Bucket, err)
	}

	for _, item := range resp.Contents {
		fmt.Println("Name:         ", *item.Key)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size:         ", *item.Size)
		fmt.Println("Storage class:", *item.StorageClass)
		fmt.Println("")
	}
}
