# Introduction

CLI tool for the Gopamin framework. Gopamin CLI create new project based on ideas promoted by [Standard Go Project Layout](https://github.com/golang-standards/project-layout) which is a well-accepted architecture by the Go community (It's not an official standard defined by the core Go dev team; however, it is a set of common historical and emerging project layout patterns in the Go ecosystem).  
To create this tool, I got inspirations from [Golang Blueprint](https://github.com/Melkeydev/go-blueprint) and some functionalities are used inside the Gopamin CLI.  
**THIS PROJECT IS UNDER ACTIVE DEVELOPMENT. SO YOU MIGHT SEE SOME BREAKING CHANGES!**

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

## Author
This project is initiated by [Bez Moradi](https://github.com/bezmoradi)

## License
Gopamin CLI is licensed under [MIT](https://github.com/gopamin/cli/blob/master/LICENSE)
