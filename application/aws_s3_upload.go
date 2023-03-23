package application

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"github.com/yuorei/anime-ranking/graph/model"
)

func AWSS3Upload(input model.NewAnimeRankingInput) (*s3manager.UploadOutput, error) {
	// The session the S3 Uploader will use
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String(os.Getenv("S3_REGION"))},
		Profile: "default",
	})
	if err != nil {
		return nil, err
	}
	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	bucketName := os.Getenv("S3_BUCKET_NAME")
	arr1 := strings.Split(input.AnimeImage.Filename, ".")
	uu, _ := uuid.NewRandom()
	objectKey := uu.String() + "." + arr1[1]

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   input.AnimeImage.File,
	})
	if err != nil {
		return nil, err
	}

	fmt.Printf("file uploaded to, %s\n", result.Location)
	return result, nil
}
