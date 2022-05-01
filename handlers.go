package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func handleECSTasks(context *cli.Context, logger logrus.Logger) {
	clusterName := context.String("cluster")

	service := ASAGIService{}

	service.logger = logger
	service.getECSClient()
	taskIds := service.getECSTaskIds(clusterName)
	service.getTaskContainers(clusterName, taskIds)
}

func handleECSExec(context *cli.Context, logger logrus.Logger) {
	var clusterName string
	var taskId string
	var containerName string

	fmt.Print("ECS cluster name: ")
	fmt.Scanln(&clusterName)

	fmt.Print("ECS task ID: ")
	fmt.Scanln(&taskId)

	fmt.Print("ECS container name: ")
	fmt.Scanln(&containerName)

	service := ASAGIService{}

	service.logger = logger
	service.runECSExec(clusterName, taskId, containerName)
}

func handleS3Upload(context *cli.Context, logger logrus.Logger) {
	filename := context.String("filename")
	bucket := context.String("bucket")

	service := ASAGIService{}

	service.logger = logger
	service.getS3Client()
	service.uploadToBucket(filename, bucket)
}
