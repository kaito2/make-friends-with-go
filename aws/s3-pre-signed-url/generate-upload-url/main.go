package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	bucketName = "sample-bucket"
	uploadKey  = "sample.png"
)

func main() {
	// Initialize a session
	sess, _ := session.NewSession(&aws.Config{
		Region:           aws.String("us-east-1"),
		Credentials:      credentials.NewStaticCredentials("test", "test", ""),
		S3ForcePathStyle: aws.Bool(true),
		Endpoint:         aws.String("http://localhost:4566"),
	})

	client := s3.New(sess)

	req, _ := client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(uploadKey),
		// NOTE: Emptly
		// Body:   strings.NewReader("EXPECTED CONTENTS"),
	})

	str, err := req.Presign(15 * time.Minute)
	if err != nil {
		log.Fatal(err)
	}

	decodedStr, err := url.QueryUnescape(str)
	if err != nil {
		fmt.Printf("Error decoding the string %v", err)
	}

	log.Println("The URL is:", decodedStr, " err:", err)
}
