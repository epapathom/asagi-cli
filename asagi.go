package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	logger := initializeLogger()

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "s3",
				Usage: "provides S3 functionality",
				Subcommands: []*cli.Command{
					{
						Name:  "upload",
						Usage: "uploads files to an S3 bucket",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "filename",
								Aliases:  []string{"f"},
								Usage:    "the file to upload",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "bucket",
								Aliases:  []string{"b"},
								Usage:    "the S3 bucket",
								Required: true,
							},
						},
						Action: func(context *cli.Context) error {
							handleS3Upload(context, logger)

							return nil
						},
					},
				},
			},
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
						Action: func(context *cli.Context) error {
							handleECSTasks(context, logger)

							return nil
						},
					},
					{
						Name:  "exec",
						Usage: "open a shell in a Fargate container",
						Action: func(context *cli.Context) error {
							handleECSExec(context, logger)

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
