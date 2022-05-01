package main

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (service *ASAGIService) uploadToBucket(filename string, bucket string) {
	file, err := os.Open(filename)
	if err != nil {
		service.logger.Fatalf("Failed to open file %q, %v", filename, err)
	}

	uploader := manager.NewUploader(service.s3Client)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &filename,
		Body:   file,
	})
	if err != nil {
		service.logger.Fatalf("Failed to upload file %q, %v", filename, err)
	}
	service.logger.Infof("File uploaded to: %s\n", result.Location)
}
