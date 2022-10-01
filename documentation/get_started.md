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
