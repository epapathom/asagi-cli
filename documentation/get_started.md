# Get Started

## Building and Installation

To utilize the Asagi CLI, we can either build the CLI in the root directory of the project or install the CLI globally.

### Build the CLI

To build the CLI and run it from the output executable binary file in the root directory of the project, run the following command.

```
go build
```

### Install the CLI

To install the CLI globally, add *GOPATH* to your *PATH* by adding the following lines to your **.bashrc**.

```
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

After adding the lines, source your **.bashrc** or open a new terminal and in the root directory of the project run the following command.

```
go install
```

## Commands

The available commands of the CLI are listed in the table below.

| Command | Subcommand | Flags                             | Description                                                 | Example                                         |
| ------- | ---------- | --------------------------------- | ----------------------------------------------------------- | ----------------------------------------------- |
| -       | help       | -                                 | Lists all the available commands or subcommands of the CLI. | asagi ecs help                                  |
| s3      | upload     | -f, --file,<br />-b, --bucket | Uploads a file to an S3 bucket.                             | asagi s3 upload -f <FILE_NAME> -b <BUCKET_NAME> |
| ecs     | tasks      | -c, --cluster                     | Returns all the ECS tasks and containers of an ECS cluster. | asagi ecs tasks -c <CLUSTER_NAME>               |
| ecs     | exec       | -                                 | Executes an interactive shell in an ECS container.          | asagi ecs exec                                  |
