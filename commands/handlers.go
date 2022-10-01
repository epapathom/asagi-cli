package commands

import (
	"asagi/logic"
)

func handleECSTasks(clusterName string) {
	taskIds := logic.ECSSingleton.GetECSTaskIds(clusterName)
	logic.ECSSingleton.GetTaskContainers(clusterName, taskIds)
}

func handleECSExec(clusterName string, taskId string, containerName string) {
	logic.ECSSingleton.RunECSExec(clusterName, taskId, containerName)
}

func handleS3Upload(filename string, bucketName string) {
	logic.S3Singleton.UploadToBucket(filename, bucketName)
}
