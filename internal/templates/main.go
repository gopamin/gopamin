package templates

import (
	api "github.com/gopamin/gopamin/internal/templates/api"
	common "github.com/gopamin/gopamin/internal/templates/common"
	database "github.com/gopamin/gopamin/internal/templates/database"
	helloWorld "github.com/gopamin/gopamin/internal/templates/hello-world"
	microservice "github.com/gopamin/gopamin/internal/templates/microservice"
	webApp "github.com/gopamin/gopamin/internal/templates/web-app"
)

func Mapper() map[string]func() ([]byte, string) {
	return map[string]func() ([]byte, string){
		"env":                       common.EnvTemplate,
		"readme":                    common.ReadmeTemplate,
		"log-readme":                common.ReadmeLogTemplate,
		"slog-readme":               common.ReadmeSlogTemplate,
		"logrus-readme":             common.ReadmeLogrusTemplate,
		"zap-readme":                common.ReadmeZapTemplate,
		"makefile":                  common.MakefileTemplate,
		"gitignore":                 common.GitIgnoreTemplate,
		"license":                   common.LicenseTemplate,
		"mock-repository":           common.MockRepositoryTemplate,
		"user-test":                 common.UserTestTemplate,
		"user":                      common.UserTemplate,
		"user-repository-interface": common.UserRepositoryInterfaceTemplate,
		"user-service-interface":    common.UserServiceInterfaceTemplate,
		"user-service":              common.UserServiceTemplate,
		"user-service-test":         common.UserServiceTestTemplate,
		"router-inferface":          common.RouterInterfaceTemplate,
		"logger-interface":          common.LoggerInterfaceTemplate,
		"log-logger":                common.LogLoggerTemplate,
		"slog-logger":               common.SlogLoggerTemplate,
		"logrus-logger":             common.LogrusLoggerTemplate,
		"zap-logger":                common.ZapLoggerTemplate,
		"api-errors":                common.ApiErrorsTemplate,
		"api-response":              common.ApiResponseTemplate,
		"configs":                   common.ConfigsTemplate,
		"configs-test":              common.ConfigsTestTemplate,
		"message":                   common.MessageTemplate,
		"broker-service-interface":  common.BrokerServiceInterfaceTemplate,
		"broker-service":            common.BrokerServiceTemplate,
		"message-broker-interface":  common.MessageBrokerInterfaceTemplate,
		"tools":                     common.ToolsTemplate,
		"tools-test":                common.ToolsTestTemplate,

		"hello-world-main":           helloWorld.HelloWorldMainTemplate,
		"hello-world-readme-with-db": helloWorld.HelloWorldReadmeWithDbTemplate,
		"hello-world-main-test":      helloWorld.HelloWorldMainTestTemplate,
		"hello-world-main-with-db":   helloWorld.HelloWorldMainWithDbTemplate,

		"api-server":             api.ApiServerTemplate,
		"graphql-server":         api.GraphqlServerTemplate,
		"api-http-routes":        api.ApiHttpRoutesTemplate,
		"api-http-users":         api.ApiHttpUsersTemplate,
		"api-readme":             api.ApiReadmeTemplate,
		"api-readme-with-db":     api.ApiReadmeWithDbTemplate,
		"graphql-readme":         api.GraphqlReadmeTemplate,
		"graphql-readme-with-db": api.GraphqlReadmeWithDbTemplate,
		"api-env":                api.ApiEnvTemplate,
		"api-chi-routes":         api.ApiChiRoutesTemplate,
		"api-chi-users":          api.ApiChiUsersTemplate,
		"api-echo-routes":        api.ApiEchoRoutesTemplate,
		"api-echo-users":         api.ApiEchoUsersTemplate,
		"api-gin-routes":         api.ApiGinRoutesTemplate,
		"api-gin-users":          api.ApiGinUsersTemplate,
		"api-gorilla-routes":     api.ApiGorillaRoutesTemplate,
		"api-gorilla-users":      api.ApiGorillaUsersTemplate,
		"api-httprouter-routes":  api.ApiHttprouterRoutesTemplate,
		"api-httprouter-users":   api.ApiHttprouterUsersTemplate,
		"api-main":               api.ApiMainTemplate,
		"api-main-with-db":       api.ApiMainWithDbTemplate,
		"graphql-main":           api.GraphqlMainTemplate,
		"graphql-main-with-db":   api.GraphqlMainWithDbTemplate,
		"graphql-schema":         api.GraphqlSchemaTemplate,

		"dynamodb-env":            database.DynamodbEnvTemplate,
		"dynamodb-repository":     database.DynamodbRepositoryTemplate,
		"dynamodb-readme":         database.DynamodbReadmeTemplate,
		"dynamodb-makefile":       database.DynamodbMakefileTemplate,
		"dynamodb-docker-compose": database.DynamodbDockerComposeTemplate,
		"mongodb-env":             database.MongodbEnvTemplate,
		"mongodb-repository":      database.MongodbRepositoryTemplate,
		"mongodb-readme":          database.MongodbReadmeTemplate,
		"mongodb-makefile":        database.MongodbMakefileTemplate,
		"mongodb-docker-compose":  database.MongodbDockerComposeTemplate,
		"mysql-env":               database.MysqlEnvTemplate,
		"mysql-repository":        database.MysqlRepositoryTemplate,
		"mysql-readme":            database.MysqlReadmeTemplate,
		"mysql-makefile":          database.MysqlMakefileTemplate,
		"mysql-docker-compose":    database.MysqlDockerComposeTemplate,
		"postgres-env":            database.PostgresEnvTemplate,
		"postgres-repository":     database.PostgresRepositoryTemplate,
		"postgres-readme":         database.PostgresReadmeTemplate,
		"postgres-makefile":       database.PostgresMakefileTemplate,
		"postgres-docker-compose": database.PostgresDockerComposeTemplate,
		"redis-env":               database.RedisEnvTemplate,
		"redis-repository":        database.RedisRepositoryTemplate,
		"redis-readme":            database.RedisReadmeTemplate,
		"redis-makefile":          database.RedisMakefileTemplate,
		"redis-docker-compose":    database.RedisDockerComposeTemplate,
		"sqlite-env":              database.SqliteEnvTemplate,
		"sqlite-repository":       database.SqliteRepositoryTemplate,
		"sqlite-makefile":         database.SqliteMakefileTemplate,
		"sqlite-readme":           database.SqliteReadmeTemplate,
		"badgerdb-env":            database.BadgerdbEnvTemplate,
		"badgerdb-repository":     database.BadgerdbRepositoryTemplate,
		"badgerdb-makefile":       database.BadgerdbMakefileTemplate,
		"badgerdb-readme":         database.BadgerdbReadmeTemplate,

		"web-app-server":              webApp.WebAppServerTemplate,
		"web-app-readme":              webApp.WebAppReadmeTemplate,
		"web-app-readme-with-db":      webApp.WebAppReadmeWithDbTemplate,
		"web-app-env":                 webApp.WebAppEnvTemplate,
		"web-app-http-routes":         webApp.WebAppHttpRoutesTemplate,
		"web-app-http-users":          webApp.WebAppHttpUsersTemplate,
		"web-app-chi-routes":          webApp.WebAppChiRoutesTemplate,
		"web-app-chi-users":           webApp.WebAppChiUsersTemplate,
		"web-app-echo-routes":         webApp.WebAppEchoRoutesTemplate,
		"web-app-echo-users":          webApp.WebAppEchoUsersTemplate,
		"web-app-gin-routes":          webApp.WebAppGinRoutesTemplate,
		"web-app-gin-users":           webApp.WebAppGinUsersTemplate,
		"web-app-gorilla-routes":      webApp.WebAppGorillaRoutesTemplate,
		"web-app-gorilla-users":       webApp.WebAppGorillaUsersTemplate,
		"web-app-httprouter-routes":   webApp.WebAppHttprouterRoutesTemplate,
		"web-app-httprouter-users":    webApp.WebAppHttprouterUsersTemplate,
		"web-app-main":                webApp.WebAppMainTemplate,
		"web-app-main-with-db":        webApp.WebAppMainWithDbTemplate,
		"web-app-users-html-template": webApp.WebAppUsersHtmlTemplateTemplate,
		"web-app-user-html-template":  webApp.WebAppUserHtmlTemplateTemplate,
		"web-app-styles":              webApp.WebAppStylesTemplate,

		"microservice-main":                    microservice.MicroserviceMainTemplate,
		"redis-microservice-readme":            microservice.RedisMicroserviceReadmeTemplate,
		"redis-microservice-env":               microservice.RedisMicroserviceEnvTemplate,
		"redis-microservice-makefile":          microservice.RedisMicroserviceMakefileTemplate,
		"redis-microservice-docker-compose":    microservice.RedisMicroserviceDockerComposeTemplate,
		"redis-microservice-broker":            microservice.RedisMicroserviceBrokerTemplate,
		"kafka-microservice-readme":            microservice.KafkaMicroserviceReadmeTemplate,
		"kafka-microservice-env":               microservice.KafkaMicroserviceEnvTemplate,
		"kafka-microservice-makefile":          microservice.KafkaMicroserviceMakefileTemplate,
		"kafka-microservice-docker-compose":    microservice.KafkaMicroserviceDockerComposeTemplate,
		"kafka-microservice-broker":            microservice.KafkaMicroserviceBrokerTemplate,
		"rabbitmq-microservice-readme":         microservice.RabbitmqMicroserviceReadmeTemplate,
		"rabbitmq-microservice-env":            microservice.RabbitmqMicroserviceEnvTemplate,
		"rabbitmq-microservice-makefile":       microservice.RabbitmqMicroserviceMakefileTemplate,
		"rabbitmq-microservice-docker-compose": microservice.RabbitmqMicroserviceDockerComposeTemplate,
		"rabbitmq-microservice-broker":         microservice.RabbitmqMicroserviceBrokerTemplate,
	}
}
