package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

func (service ASAGIService) getECSTaskIds(cluster string) (taskIds []string) {
	input := &ecs.ListTasksInput{
		Cluster: aws.String(cluster),
	}

	response, err := service.ecsClient.ListTasks(context.TODO(), input)
	if err != nil {
		service.logger.Fatal(err)
	}

	for _, taskArn := range response.TaskArns {
		taskArnSplit := strings.Split(taskArn, "/")
		taskId := taskArnSplit[len(taskArnSplit)-1]
		taskIds = append(taskIds, taskId)
	}

	return
}

func (service ASAGIService) getTaskContainers(clusterName string, taskIds []string) {
	input := &ecs.DescribeTasksInput{
		Tasks:   taskIds,
		Cluster: aws.String(clusterName),
	}

	response, err := service.ecsClient.DescribeTasks(context.TODO(), input)
	if err != nil {
		service.logger.Fatal(err)
	}

	fmt.Printf("ECS cluster: %s\n", clusterName)

	for _, task := range response.Tasks {
		taskArn := task.TaskArn
		taskArnSplit := strings.Split(*taskArn, "/")
		taskId := taskArnSplit[len(taskArnSplit)-1]

		fmt.Printf("\nTask ID: %s\n\n", taskId)
		fmt.Println("Containers:")

		for index, container := range task.Containers {
			containerName := container.Name
			fmt.Printf("%d. %s\n", index+1, *containerName)
		}
	}
}

func (service ASAGIService) runECSExec(clusterName string, taskId string, containerName string) {
	commandString := fmt.Sprintf("aws ecs execute-command --cluster %s --task %s --container %s --interactive --command '/bin/sh'", clusterName, taskId, containerName)

	command := exec.Command("bash", "-c", commandString)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		service.logger.Fatal(err)
	}
}
