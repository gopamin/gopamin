## Table of Contents

-   [Introduction](#introduction)
-   [An Intro to The Clean Architecture](#an-intro-to-the-clean-architecture)
-   [Prerequisites](#prerequisites)
-   [Supported Application Types](#supported-application-types)
-   [Supported Databases](#supported-databases)
-   [Installation](#installation)
-   [Usage](#usage)
-   [Guides](#guides)
-   [Inspirations](#inspirations)
-   [Author](#author)
-   [License](#license)

## Introduction

Gopamin is a CLI which creates new projects based on ideas promoted by [Standard Go Project Layout](https://github.com/golang-standards/project-layout) which is a **well-accepted** architecture by the Go community (It's not an official standard defined by the core Go team; however, it is a set of common historical and emerging project layout patterns in the Go ecosystem).

## An Intro to The Clean Architecture

All boilerplates created by Gopamin are based on [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html). Simply put, it is a software architectural pattern introduced by Robert C. Martin in order to create maintainable, scalable, and testable software systems by decoupling the application logic from the infrastructure details (Although there are some minor differences, this architectural pattern is referred to by other names such as Ports & Adapters Architecture, Hexagonal Architecture, and Onion Architecture).

## Prerequisites

The minimum required tools for using the Gopamin CLI tool is Golang which can be downloaded from [Go All Releases](https://go.dev/dl). To have full development setup though, other tools are also recommended to be installed on your local machine:

-   **[Git](https://git-scm.com/)**: By default each new project created by this tool initializes a Git repo; that's why you need to make sure Git is installed on your machine.
-   **[Docker](https://www.docker.com)**: A `Dockerfile` is included in each new project for containerization purposes. If you choose to create a new project with database integration, a `docker-compose.yml` will be included in the root of the project for running the database of your choice.
-   **[GNU Make](https://www.gnu.org/software/make)**: This a tool which controls the generation of executables and other types of files. By default, each new project includes a `Makefile` for running some most-used commands like running an application (This tool is installed by default on Mac and some distributions of Linux). To check whether this tool is installed on your machine, open terminal and run `make --version` (If you do not have this tool installed on your machine though, still you can use this tools without any limitations).

## Supported Application Types

You can create a range of different Golang applications using the Gopamin tool; from simple Hello World apps to Microservices. Supported applications types are as follows:

-   Hello World
-   Web Application
-   RESTful API
    -   Echo
    -   Chi
    -   Gin
    -   Httprouter
    -   Gorilla
    -   HTTP (The build-in HTTP package will be used)
-   Microservices
    -   Redis
    -   Kafka
    -   RabbitMQ

## Supported Databases

Based on the type of the project you want to scaffold, by passing the `-d` flag, you can also add database integration. Supported databases are as follows:

-   MySQL
-   PostgreSQL
-   MongoDB
-   SQLite
-   DynamoDB
-   Redis

Based on the database type you choose, a `docker-compose.yml` file will be added as well to make running the database instance way easier for development purposes.

## Installation

To install the Gopamin tool, run the following command inside terminal:

```text
$ go install github.com/gopamin/gopamin@latest
```

To make sure it's installed correctly in your `$GOPATH`, run the following command:

```text
$ gopamin version
```

If the installation process goes well, from now on you can run the `gopamin` command from anywhere on your file system.

## Usage

If you got the current version correctly in the previous step, now you can use this tools to scaffold new Golang projects.

### The `-h` Flag

The `-h` flag which is short for `--help` can be used to show you a guide on how to create a new project:

```text
$ gopamin new -h
```

The above command shows the help on how to use different flags to scaffold different types of projects. In the following sections, these flags will be introduced.

### The `-n` Flag

The `-n` flag which is short for `--name` should be used for choosing a name for your project. For example:

```text
$ gopamin new -n my-hello-world-app
```

Of course the above command will not proceed because the flag for specifying the type is not provided.

### The `-t` Flag

The `-t` flag which is short for `--type` should be used for choosing the project type you want to scaffold. For example, to create a simple Hello World app we have:

```text
$ gopamin -n my-hello-world-app -t hello-world
```

The supported values for the `-t` flag are `hello-world`, `web-app`, `api`, and `microservice`.

### The `-p` Flag

The `-p` flag which is short for `--platform` must be used for the projects of type `web-app`, `api`, and `microservice`. Supported values for projects of either `web-app` or `api` type are `echo`, `chi`, `gin`, `httprouter`, `gorilla`, and `http`. For example, to create an API which uses the Echo framework under the hood we have:

```text
$ gopamin new -n my-rest-api -t api -p echo
```

Supported values for the `-p` flag for the projects of type `microservice` are `redis`, `kafka`, and `rabbitmq`. For example, to create a microservice with Kafka integration we have:

```text
$ gopamin new -n my-kafka-microservice -t microservice -p kafka
```

### The `-d` Flag

The `-d` flag which is short for `--database` should be used to add database integration. Available values for this flag are `mysql`, `postgres`, `mongodb`, `sqlite`, `dynamodb`, and `redis`. For example, to create a web application with MySQL integration we have:

```text
$ gopamin new -n my-web-app -t web-app -p http -d mysql
```

## Guides

Each new project includes a `README.md` file in the root path which provides you with some guides on how to use that specific project.

## Inspirations

To create this tool, I got inspirations from [Golang Blueprint](https://github.com/Melkeydev/go-blueprint) and some of its functionalities are used inside this tool.

## Author

This project is maintained by [Bez Moradi](https://github.com/bezmoradi)

## License

Gopamin is licensed under [MIT](https://github.com/gopamin/gopamin/blob/master/LICENSE)
