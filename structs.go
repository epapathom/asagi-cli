package main

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	logger "github.com/sirupsen/logrus"
)

type ASAGIService struct {
	logger           logger.Logger
	ecsClient        *ecs.Client
	cloudwatchClient *cloudwatch.Client
	s3Client         *s3.Client
}
