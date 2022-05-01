package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	logger := initializeLogger()

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "ecs",
				Usage: "provides ECS functionality",
				Subcommands: []*cli.Command{
					{
						Name:  "tasks",
						Usage: "returns all tasks and containers",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "cluster",
								Aliases:  []string{"c"},
								Usage:    "the ECS cluster",
								Required: true,
							},
						},
						Action: func(c *cli.Context) error {
							clusterName := c.String("cluster")

							service := ASAGIService{}

							service.logger = logger
							service.getECSClient()
							taskIds := service.getECSTaskIds(clusterName)
							service.getTaskContainers(clusterName, taskIds)

							return nil
						},
					},
					{
						Name:  "exec",
						Usage: "open a shell in a Fargate container",
						Action: func(c *cli.Context) error {
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

							return nil
						},
					},
				},
			},
			{
				Name:  "cloudwatch",
				Usage: "provides CloudWatch functionality",
				Subcommands: []*cli.Command{
					{
						Name:  "export",
						Usage: "exports metric widget images to a markdown file",
						Action: func(c *cli.Context) error {

							return nil
						},
					},
				},
			},
			{
				Name:  "s3",
				Usage: "provides S3 functionality",
				Subcommands: []*cli.Command{
					{
						Name:  "upload",
						Usage: "uploads files to an S3 bucket",
						Action: func(c *cli.Context) error {

							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal(err)
	}
}
