# Introduction

Gopamin CLI creates new projects based on ideas promoted by [Standard Go Project Layout](https://github.com/golang-standards/project-layout) which is a **well-accepted** architecture by the Go community (It's not an official standard defined by the core Go dev team; however, it is a set of common historical and emerging project layout patterns in the Go ecosystem).  
To create this tool, I got inspirations from [Golang Blueprint](https://github.com/Melkeydev/go-blueprint) and some of its functionalities are used inside this tool.

## Clean Architecture

All boilerplates created by Gopamin CLI are based on clean architecture. Simply put, it is a software architectural pattern introduced by Robert C. Martin in order to create maintainable, scalable, and testable software systems by decoupling the application logic from the infrastructure details (This architectural pattern is referred to by other names such as Onion Architecture, Hexagonal Architecture, and Onion Architecture).

## Prerequisites

The minimum required tools for using the Gopamin CLI is Golang which can be download from [Go All Releases](https://go.dev/dl). To have full development setup though, other tools are also recommended to be installed on your local machine:

-   **[Git](https://git-scm.com/)**: By default each new project created by this tool initializes a Git repo; that's why you need to make sure Git is installed on your machine.
-   **[Docker](https://www.docker.com)**: A `Dockerfile` is included in each new project for containerization purposes.
-   **[GNU Make](https://www.gnu.org/software/make)**: This a tool which controls the generation of executables and other types of files. By default, each new project includes a `Makefile` for running simple commands (This tool is installed by default on Mac and some distributions of Linux). To check whether this tool is installed on your machine, open terminal and run `make --version`.

## Installation

To install Gopamin CLI, run the following command inside terminal:

```text
$ go install github.com/gopamin/gopamin@v1.0.5
```

To make sure it's installed correctly in your `$GOPATH`, run the following command:

```text
$ gopamin version
```

## Usage

If you got the current version correctly in the previous step, now you can create a new project by running the following command:

```text
$ gopamin new -h
```

The above command show the help on how to use different flags to scaffold different types of projects.

## Author

This project is maintained by [Bez Moradi](https://github.com/bezmoradi)

## License

Gopamin CLI is licensed under [MIT](https://github.com/gopamin/gopamin/blob/master/LICENSE)
