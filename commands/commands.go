package commands

import (
	"github.com/spf13/cobra"
)

func createAsagiCommand() *cobra.Command {
	cmdAsagi := &cobra.Command{
		Use:   "asagi [command]",
		Short: "Execute AWS commands",
	}

	return cmdAsagi
}

func createECSCommand() *cobra.Command {
	var clusterName string
	var taskId string
	var containerName string

	cmdECS := &cobra.Command{
		Use:   "ecs [command] [options]",
		Short: "Execute an ECS command",
	}

	cmdECSTasks := &cobra.Command{
		Use:   "tasks [options]",
		Short: "Retrieve all ECS tasks and containers",
		Run: func(cmd *cobra.Command, args []string) {
			handleECSTasks(clusterName)
		},
	}

	cmdECSExec := &cobra.Command{
		Use:   "exec [options]",
		Short: "Open an interactive shell in an ECS container",
		Run: func(cmd *cobra.Command, args []string) {
			handleECSExec(clusterName, taskId, containerName)
		},
	}

	cmdECSTasks.Flags().StringVarP(&clusterName, "cluster", "c", "", "the name of the ECS cluster")
	cmdECSTasks.MarkFlagRequired("cluster")

	cmdECSExec.Flags().StringVarP(&clusterName, "cluster", "c", "", "the name of the ECS cluster")
	cmdECSExec.Flags().StringVarP(&taskId, "task", "t", "", "the Id of the ECS task")
	cmdECSExec.Flags().StringVarP(&containerName, "container", "C", "", "the name of the ECS container")
	cmdECSExec.MarkFlagRequired("cluster")
	cmdECSExec.MarkFlagRequired("task")
	cmdECSExec.MarkFlagRequired("container")

	cmdECS.AddCommand(cmdECSTasks)
	cmdECS.AddCommand(cmdECSExec)

	return cmdECS
}

func createS3Command() *cobra.Command {
	var filename string
	var bucketName string

	cmdS3 := &cobra.Command{
		Use:   "s3 [command] [options]",
		Short: "Execute an S3 command",
	}

	cmdS3Upload := &cobra.Command{
		Use:   "upload [options]",
		Short: "Upload a file to an S3 bucket",
		Run: func(cmd *cobra.Command, args []string) {
			handleS3Upload(filename, bucketName)
		},
	}

	cmdS3Upload.Flags().StringVarP(&filename, "filename", "f", "", "the name of the file to upload")
	cmdS3Upload.Flags().StringVarP(&bucketName, "bucket", "b", "", "the name of the S3 bucket")
	cmdS3Upload.MarkFlagRequired("filename")
	cmdS3Upload.MarkFlagRequired("bucket")

	cmdS3.AddCommand(cmdS3Upload)

	return cmdS3
}

func Execute() {
	cmdAsagi := createAsagiCommand()

	cmdECS := createECSCommand()
	cmdS3 := createS3Command()

	cmdAsagi.AddCommand(cmdECS)
	cmdAsagi.AddCommand(cmdS3)
	cmdAsagi.Execute()
}
