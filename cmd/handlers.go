package cmd

import (
	"asagi/logic"
	"fmt"

	"github.com/urfave/cli/v2"
)

func handleECSTasks(context *cli.Context) {
	clusterName := context.String("cluster")

	taskIds := logic.ECSSingleton.GetECSTaskIds(clusterName)
	logic.ECSSingleton.GetTaskContainers(clusterName, taskIds)
}

func handleECSExec(context *cli.Context) {
	var clusterName string
	var taskId string
	var containerName string

	fmt.Print("ECS cluster name: ")
	fmt.Scanln(&clusterName)

	fmt.Print("ECS task ID: ")
	fmt.Scanln(&taskId)

	fmt.Print("ECS container name: ")
	fmt.Scanln(&containerName)

	logic.ECSSingleton.RunECSExec(clusterName, taskId, containerName)
}

func handleS3Upload(context *cli.Context) {
	filename := context.String("filename")
	bucket := context.String("bucket")

	logic.S3Singleton.UploadToBucket(filename, bucket)
}
