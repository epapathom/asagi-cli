package cmd

import (
	"asagi/utils"
	"os"

	"github.com/urfave/cli/v2"
)

func Execute() {
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
								Name:     "file",
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
							handleS3Upload(context)

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
							handleECSTasks(context)

							return nil
						},
					},
					{
						Name:  "exec",
						Usage: "open a shell in a Fargate container",
						Action: func(context *cli.Context) error {
							handleECSExec(context)

							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		utils.LoggerSingleton.Logger.Fatal(err)
	}
}
