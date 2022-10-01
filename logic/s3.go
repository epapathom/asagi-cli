package logic

import (
	"asagi/utils"
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3 struct {
	Client *s3.Client
}

var S3Singleton *S3

func (s *S3) Init() {
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		utils.LoggerSingleton.Logger.Error("configuration error: %v", err.Error())
	}

	s.Client = s3.NewFromConfig(config)
}

func (s S3) UploadToBucket(filename string, bucket string) {
	file, err := os.Open(filename)
	if err != nil {
		utils.LoggerSingleton.Logger.Fatalf("Failed to open file %q, %v", filename, err)
	}

	uploader := manager.NewUploader(s.Client)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &filename,
		Body:   file,
	})
	if err != nil {
		utils.LoggerSingleton.Logger.Fatalf("Failed to upload file %q, %v", filename, err)
	}
	utils.LoggerSingleton.Logger.Infof("File uploaded to: %s\n", result.Location)
}
