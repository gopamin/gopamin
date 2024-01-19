package commands

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func argsValidator() bool {
	switch {
	case name == "":
		fmt.Println(`The -n flag is required. For more help, type "gopamin new -h"`)
		return false
	case projectType == "":
		fmt.Println(`The -t flag is required. For more help, type "gopamin new -h"`)
		return false
	case noPlatfromForApiOrMicroservice():
		fmt.Println(`The -p flag is required for projects of type "api" or "microservice". For more help, type "gopamin new -h"`)
		return false
	}

	if isProjectNameTaken() {
		fmt.Println("This project name is already taken in the current directory")
		return false
	}

	if !validateType() {
		fmt.Println(`The specified value for the -t flag is wrong. For more help, type "gopamin new -h"`)
		return false
	}

	if projectType == "api" && !restFrameworkValidator() {
		fmt.Println(`The specified value for the -p flag for "api" type of apps is wrong. For more help, type "gopamin new -h"`)
		return false
	}

	if projectType == "microservice" && !microserviceFrameworkValidator() {
		fmt.Println(`The specified value for the -p flag for "microservice" type of apps is wrong. For more help, type "gopamin new -h"`)
		return false
	}

	if database != "" && !databaseValidator() {
		fmt.Println(`The specified value for the -d flag is wrong. For more help, type "gopamin new -h"`)
		return false
	}

	return true
}

func noPlatfromForApiOrMicroservice() bool {
	if (projectType == "microservice" || projectType == "api") && platform == "" {
		return true
	}
	return false
}

func validateType() bool {
	if projectType == "hello-world" ||
		projectType == "web-app" ||
		projectType == "microservice" ||
		projectType == "api" {
		return true
	}

	return false
}

func restFrameworkValidator() bool {
	if platform == "echo" ||
		platform == "chi" ||
		platform == "gin" ||
		platform == "fiber" ||
		platform == "httprouter" ||
		platform == "http" ||
		platform == "graphql" ||
		platform == "gorilla" {
		return true
	}

	return false
}

func microserviceFrameworkValidator() bool {
	if platform == "kafka" ||
		platform == "rabbitmq" ||
		platform == "grpc" {
		return true
	}

	return false
}

func databaseValidator() bool {
	if database == "mysql" ||
		database == "postgres" ||
		database == "mongodb" ||
		database == "sqlite" ||
		database == "dynamodb" ||
		database == "redis" {
		return true
	}

	return false
}

func isProjectNameTaken() bool {
	name = strings.Replace(strings.TrimSpace(name), " ", "-", -1)

	if _, err := os.Stat(name); err == nil {
		dirEntries, err := os.ReadDir(name)
		if err != nil {
			log.Println("Could not read the directory")
		}
		if len(dirEntries) > 0 {
			return true
		}
	}

	return false
}
