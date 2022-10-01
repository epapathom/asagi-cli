package logic

import (
	"asagi/utils"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

type ECS struct {
	Client *ecs.Client
}

var ECSSingleton *ECS

func (e *ECS) Init() {
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		utils.LoggerSingleton.Logger.Error("configuration error: %v", err.Error())
	}

	e.Client = ecs.NewFromConfig(config)
}

func (e *ECS) GetECSTaskIds(cluster string) (taskIds []string) {
	input := &ecs.ListTasksInput{
		Cluster: aws.String(cluster),
	}

	response, err := e.Client.ListTasks(context.TODO(), input)
	if err != nil {
		utils.LoggerSingleton.Logger.Fatal(err)
	}

	for _, taskArn := range response.TaskArns {
		taskArnSplit := strings.Split(taskArn, "/")
		taskId := taskArnSplit[len(taskArnSplit)-1]
		taskIds = append(taskIds, taskId)
	}

	return
}

func (e *ECS) GetTaskContainers(clusterName string, taskIds []string) {
	input := &ecs.DescribeTasksInput{
		Tasks:   taskIds,
		Cluster: aws.String(clusterName),
	}

	response, err := e.Client.DescribeTasks(context.TODO(), input)
	if err != nil {
		utils.LoggerSingleton.Logger.Fatal(err)
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

func (e *ECS) RunECSExec(clusterName string, taskId string, containerName string) {
	commandString := fmt.Sprintf("aws ecs execute-command --cluster %s --task %s --container %s --interactive --command '/bin/sh'", clusterName, taskId, containerName)

	command := exec.Command("bash", "-c", commandString)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		utils.LoggerSingleton.Logger.Fatal(err)
	}
}
