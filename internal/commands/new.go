package commands

import (
	"github.com/gopamin/gopamin/internal/scaffolder"
	"github.com/spf13/cobra"
)

var (
	database    string
	platform    string
	name        string
	projectType string
)

var newCmd = &cobra.Command{
	Example: `To create a simple hello world app, run the following command:
gopamin new -n HelloWorld -t hello-world
To create a hello world app with MySQL integration, run the following command:
 gopamin new -n HelloWorld -t hello-world -d mysql
To create a simple web application, run the following command:
 gopamin new -n HelloWorld -t web-app
To create a RESTful API using the echo framework, run the following command:
 gopamin new -n HelloWorld -t api -p echo
To create a RESTful API using the built-in http package and MongoDB integration, run the following command:
 gopamin new -n HelloWorld -t api -p http -d mongodb
To create a Microservice using Kafka message broker, run the following command:
 gopamin new -n HelloWorld -t microservice -p kafka`,
	Use:   "new",
	Short: "Create a new project",
	Run: func(cmd *cobra.Command, args []string) {
		if argsValidator() {
			scaffolder.New(projectType, platform, name, database)
		}
	},
}

func init() {
	newCmd.Flags().StringVarP(&name, "name", "n", "", `Name of project (Keep in mind the all white spaces will be replaced by dash). 
If you want to use space-separated words, place them inside double quotes like "my demo app" then it will be converted to "my-demo-app".`)

	newCmd.Flags().StringVarP(&projectType, "type", "t", "", `Type of the project. Available types are:
 - hello-world (A simple "Hello World" app without any extra functionalities)
 - web-app (A minital web applicaion which serves an HTML and CSS files)
 - microservice (If chosen, you must also add the "-p" flag to specify the type of the third-party platform you want to use)
 - api (If chosen, you must also add the "-p" flag to specify the type of the third-party platform you want to use)`)

	newCmd.Flags().StringVarP(&platform, "platform", "p", "", `If the chosen "-t" is either "api" or "microservice", you must also add the "-p" flag to specify the underlying third-party platform.
Available values for the "api" type are:
 - echo
 - chi
 - gin
 - httprouter
 - gorilla
 - http (The build-in HTTP client will be used)
Available values for the "microservice" type are:
- redis
- kafka
- rabbitmq`)

	newCmd.Flags().StringVarP(&database, "database", "d", "", `Type of the database. Available values are:
 - mysql
 - postgres
 - mongodb
 - sqlite
 - dynamodb
 - redis`)
	rootCmd.AddCommand(newCmd)
}
