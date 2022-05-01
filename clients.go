package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	log "github.com/sirupsen/logrus"
)

func (service *ASAGIService) getECSClient() {
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Error("configuration error: %v", err.Error())
	}

	service.ecsClient = ecs.NewFromConfig(config)
}

func (service *ASAGIService) getCloudWatchClient() {
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Error("configuration error: %v", err.Error())
	}

	service.cloudwatchClient = cloudwatch.NewFromConfig(config)
}

func (service *ASAGIService) getS3Client() {
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Error("configuration error: %v", err.Error())
	}

	service.s3Client = s3.NewFromConfig(config)
}
