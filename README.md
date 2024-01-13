# Introduction

CLI tool for the Gopamin framework. Gopamin CLI create new project based on ideas promoted by [Standard Go Project Layout](https://github.com/golang-standards/project-layout) which is a well-accepted architecture by the Go community (It's not an official standard defined by the core Go dev team; however, it is a set of common historical and emerging project layout patterns in the Go ecosystem).  
To create this tool, I got inspirations from [Golang Blueprint](https://github.com/Melkeydev/go-blueprint) and some functionalities are used inside the Gopamin CLI.  
**THIS PROJECT IS UNDER ACTIVE DEVELOPMENT. SO YOU MIGHT SEE SOME BREAKING CHANGES!**

## Prerequisites

The minimum required tool for using the Gopamin CLI is Golang which can be download from [Go All Releases](https://go.dev/dl)  
To have full development setup though, other tools are also recommended to be installed on your local machine:  
**[Git](https://git-scm.com/)**: By default each new project created by this CLI initialize a Git repo; that's why you need to make sure Git is installed on your machine.  
**[Docker](https://www.docker.com)**: A `Dockerfile` is included to each new project to Dockerize it.  
**[GNU Make](https://www.gnu.org/software/make)**: This a tool which controls the generation of executables and other types of files. By default, each new project includes a `Makefile` for running simple commands (This tool is installed by default on Mac and some distributions of Linux). To check whether this tool is installed on your machine, open terminal and run `make --version`.

## Installation

To install Gopamin CLI, run the following command inside terminal:

```text
$ go install -v github.com/gopamin/cli/cmd/gopamin@latest
```

## Usage

In order to create a new Go boilerplate project, first make sure the Gopamin CLI is installed correctly in your `$GOPATH` by running the following command:

```text
$ gopamin version
```

If you got the current version correctly, now you can create a new project by running the following command:

```text
$ gopamin create
```

## How to Get Boilerplates without Installing This Tool

-   If you have difficulty installing the Gopamin CLI tool or for any reason not interested in installing it, you can refer to the [examples](https://github.com/gopamin/cli/blob/master/examples) folder and download whatever boilerplate you're looking for.

## Possible Project Types

### Hello World

After being asked about "Project type?", you can select "Hello World" which will create a simple Go application with the following structure:

```text
.
├── .dockerignore
├── .env
├── .git
├── .gitignore
├── Dockerfile
├── LICENSE
├── Makefile
├── README.md
├── cmd
│   └── main.go
├── configs
│   └── load-env.go
├── go.mod
└── go.sum
```

-   `.dockerignore`: This file is used to ignore those files that shouldn't be added to the Docker image. By default, `Dockerfile`, `.git`, and `.env` file are added
-   `.env`: This file includes a sample key/value
-   `.git`: After creating a project with Gopamin CLI, by default a Git repository will be initialized in it
-   `.gitignore`: By default, `.env` file is added to `.gitignore`
-   `Dockerfile`: Basic Docker commands are added to this file to make the app Dockerized
-   `LICENSE`: By default, "MIT License" is added to newly-created projects using Gopamin CLI. Feel free to remove/update it
-   `Makefile`: Basic commands are added to the `Makefile` for running the application
-   `README.md`: Some default text is added to this file. Feel free to update it
-   `cmd/main.go`: This file is the starting point of the Go application
-   `configs/load-env.go`: This file reads environment variables from `.env` file in the root of the project
-   `go.mod`: This is the default `go.mod` file for Go programs
-   `go.sum`: This is the default `go.sum` file for Go programs

## Author

This project is initiated by [Bez Moradi](https://github.com/bezmoradi)

## License

Gopamin CLI is licensed under [MIT](https://github.com/gopamin/cli/blob/master/LICENSE)
