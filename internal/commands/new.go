package commands

import (
	"github.com/gopamin/cli/internal/scaffolder"
	"github.com/spf13/cobra"
)

var (
	database    string
	platform    string
	name        string
	projectType string
)

var newCmd = &cobra.Command{
	Example: "gopamin new -n HelloWorld",
	Use:     "new",
	Short:   "Create a new project",
	Run: func(cmd *cobra.Command, args []string) {
		if argsValidator() {
			scaffolder.New(projectType, platform, name, database)
		}
	},
}

func init() {
	newCmd.Flags().StringVarP(&name, "name", "n", "", `Name of project (Keep in mind the all white spaces will be replaced by dash). 
If you want to space-separated words, place them inside double quotes like "my demo app" then it will be converted to "my-demo-app".`)

	newCmd.Flags().StringVarP(&projectType, "type", "t", "", `Type of the project. Available types are as follows:
 - hello-world (A simple "Hello World" app without any extra functionalities)
 - web-app (A minital web applicaion which serves an HTML file)
 - microservice NOT_IMPLEMENTED_YET!!! (If chosen, you also MUST add the "-p" flag to specify the type of the microservice you want to create)
 - api (If chosen, you also MUST add the "-p" flag to specify the type of the third-party platform you want to use)`)

	newCmd.Flags().StringVarP(&platform, "platform", "p", "", `If the chosen "--type" is either "api" or "microservice", you also Must add the "-p" flag to specify the underlying third-party platform.
Available values for the "api" type are as follows:
 - echo
 - chi
 - gin
 - httprouter
 - gorilla
 - http (The build-in HTTP client will be used)
 - graphql NOT_IMPLEMENTED_YET!!!
Available values for the "microservice" type are as follows:
- grpc NOT_IMPLEMENTED_YET!!!
- kafka NOT_IMPLEMENTED_YET!!!
- rabbitmq NOT_IMPLEMENTED_YET!!!`)

	newCmd.Flags().StringVarP(&database, "database", "d", "", `Type of the database. Available values are as follows:
 - mysql
 - postgres
 - mongodb
 - sqlite
 - dynamodb
 - redis`)
	rootCmd.AddCommand(newCmd)
}
