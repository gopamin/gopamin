package templates

import (
	apiChi "github.com/gopamin/gopamin/internal/templates/api/chi"
	apiEcho "github.com/gopamin/gopamin/internal/templates/api/echo"
	apiGin "github.com/gopamin/gopamin/internal/templates/api/gin"
	apiGorilla "github.com/gopamin/gopamin/internal/templates/api/gorilla"
	apiHttp "github.com/gopamin/gopamin/internal/templates/api/http"
	apiHttprouter "github.com/gopamin/gopamin/internal/templates/api/httprouter"
	common "github.com/gopamin/gopamin/internal/templates/common"
	dynamodb "github.com/gopamin/gopamin/internal/templates/database/dynamodb"
	mongodb "github.com/gopamin/gopamin/internal/templates/database/mongodb"
	mysql "github.com/gopamin/gopamin/internal/templates/database/mysql"
	postgres "github.com/gopamin/gopamin/internal/templates/database/postgres"
	redis "github.com/gopamin/gopamin/internal/templates/database/redis"
	sqlite "github.com/gopamin/gopamin/internal/templates/database/sqlite"
	helloWorld "github.com/gopamin/gopamin/internal/templates/hello-world"
	kafkaMicroservice "github.com/gopamin/gopamin/internal/templates/microservice/kafka"
	rabbitmqMicroservice "github.com/gopamin/gopamin/internal/templates/microservice/rabbitmq"
	redisMicroservice "github.com/gopamin/gopamin/internal/templates/microservice/redis"
	webApp "github.com/gopamin/gopamin/internal/templates/web-app"
)

func Mapper() map[string]func() ([]byte, string) {
	return map[string]func() ([]byte, string){
		"dockerfile":                  common.DockerFileTemplate,
		"gitignore":                   common.GitIgnoreTemplate,
		"license":                     common.LicenseTemplate,
		"dockerignore":                common.DockerIgnoreTemplate,
		"mock-repository":             common.MockRepositoryTemplate,
		"user-test":                   common.UserTestTemplate,
		"user":                        common.UserTemplate,
		"user-repository-interface":   common.UserRepositoryInterfaceTemplate,
		"user-service-interface":      common.UserServiceInterfaceTemplate,
		"user-service":                common.UserServiceTemplate,
		"user-service-test":           common.UserServiceTestTemplate,
		"router-inferface":            common.RouterInterfaceTemplate,
		"api-errors":                  common.ApiErrorsTemplate,
		"api-response":                common.ApiResponseTemplate,
		"configs":                     common.ConfigsTemplate,
		"configs-test":                common.ConfigsTestTemplate,
		"server":                      common.ServerTemplate,
		"api-main":                    common.ApiMainTemplate,
		"api-main-with-db":            common.ApiMainWithDbTemplate,
		"web-app-main-with-db":        common.WebAppMainWithDbTemplate,
		"hello-world-main":            common.HelloWorldMainTemplate,
		"hello-world-main-test":       common.HelloWorldMainTestTemplate,
		"hello-world-main-with-db":    common.HelloWorldMainWithDbTemplate,
		"web-app-users-html-template": common.WebAppUsersHtmlTemplateTemplate,
		"web-app-user-html-template":  common.WebAppUserHtmlTemplateTemplate,
		"web-app-styles":              common.WebAppStylesTemplate,
		"web-app-main":                common.WebAppMainTemplate,
		"message":                     common.MessageTemplate,
		"broker-service-interface":    common.BrokerServiceInterfaceTemplate,
		"broker-service":              common.BrokerServiceTemplate,
		"message-broker-interface":    common.MessageBrokerInterfaceTemplate,
		"microservice-main":           common.MicroserviceMainTemplate,
		"tools":                       common.ToolsTemplate,
		"tools-test":                  common.ToolsTestTemplate,

		"hello-world-readme":   helloWorld.HelloWorldReadmeTemplate,
		"hello-world-makefile": helloWorld.HelloWorldMakefileTemplate,
		"hello-world-env":      helloWorld.HelloWorldEnvTemplate,

		"api-http-readme":   apiHttp.ApiHttpReadmeTemplate,
		"api-http-makefile": apiHttp.ApiHttpMakefileTemplate,
		"api-http-env":      apiHttp.ApiHttpEnvTemplate,
		"api-http-routes":   apiHttp.ApiHttpRoutesTemplate,
		"api-http-users":    apiHttp.ApiHttpUsersTemplate,

		"api-chi-readme":   apiChi.ApiChiReadmeTemplate,
		"api-chi-makefile": apiChi.ApiChiMakefileTemplate,
		"api-chi-env":      apiChi.ApiChiEnvTemplate,
		"api-chi-routes":   apiChi.ApiChiRoutesTemplate,
		"api-chi-users":    apiChi.ApiChiUsersTemplate,

		"api-echo-readme":   apiEcho.ApiEchoReadmeTemplate,
		"api-echo-makefile": apiEcho.ApiEchoMakefileTemplate,
		"api-echo-env":      apiEcho.ApiEchoEnvTemplate,
		"api-echo-routes":   apiEcho.ApiEchoRoutesTemplate,
		"api-echo-users":    apiEcho.ApiEchoUsersTemplate,

		"api-gin-readme":   apiGin.ApiGinReadmeTemplate,
		"api-gin-makefile": apiGin.ApiGinMakefileTemplate,
		"api-gin-env":      apiGin.ApiGinEnvTemplate,
		"api-gin-routes":   apiGin.ApiGinRoutesTemplate,
		"api-gin-users":    apiGin.ApiGinUsersTemplate,

		"api-gorilla-readme":   apiGorilla.ApiGorillaReadmeTemplate,
		"api-gorilla-makefile": apiGorilla.ApiGorillaMakefileTemplate,
		"api-gorilla-env":      apiGorilla.ApiGorillaEnvTemplate,
		"api-gorilla-routes":   apiGorilla.ApiGorillaRoutesTemplate,
		"api-gorilla-users":    apiGorilla.ApiGorillaUsersTemplate,

		"api-httprouter-readme":   apiHttprouter.ApiHttprouterReadmeTemplate,
		"api-httprouter-makefile": apiHttprouter.ApiHttprouterMakefileTemplate,
		"api-httprouter-env":      apiHttprouter.ApiHttprouterEnvTemplate,
		"api-httprouter-routes":   apiHttprouter.ApiHttprouterRoutesTemplate,
		"api-httprouter-users":    apiHttprouter.ApiHttprouterUsersTemplate,

		"mysql-env":            mysql.MysqlEnvTemplate,
		"mysql-repository":     mysql.MysqlRepositoryTemplate,
		"mysql-readme":         mysql.MysqlReadmeTemplate,
		"mysql-makefile":       mysql.MysqlMakefileTemplate,
		"mysql-docker-compose": mysql.MysqlDockerComposeTemplate,

		"postgres-env":            postgres.PostgresEnvTemplate,
		"postgres-repository":     postgres.PostgresRepositoryTemplate,
		"postgres-readme":         postgres.PostgresReadmeTemplate,
		"postgres-makefile":       postgres.PostgresMakefileTemplate,
		"postgres-docker-compose": postgres.PostgresDockerComposeTemplate,

		"mongodb-env":            mongodb.MongodbEnvTemplate,
		"mongodb-repository":     mongodb.MongodbRepositoryTemplate,
		"mongodb-readme":         mongodb.MongodbReadmeTemplate,
		"mongodb-makefile":       mongodb.MongodbMakefileTemplate,
		"mongodb-docker-compose": mongodb.MongodbDockerComposeTemplate,

		"redis-env":            redis.RedisEnvTemplate,
		"redis-repository":     redis.RedisRepositoryTemplate,
		"redis-readme":         redis.RedisReadmeTemplate,
		"redis-makefile":       redis.RedisMakefileTemplate,
		"redis-docker-compose": redis.RedisDockerComposeTemplate,

		"sqlite-env":        sqlite.SqliteEnvTemplate,
		"sqlite-repository": sqlite.SqliteRepositoryTemplate,
		"sqlite-makefile":   sqlite.SqliteMakefileTemplate,
		"sqlite-readme":     sqlite.SqliteReadmeTemplate,

		"dynamodb-env":            dynamodb.DynamodbEnvTemplate,
		"dynamodb-repository":     dynamodb.DynamodbRepositoryTemplate,
		"dynamodb-readme":         dynamodb.DynamodbReadmeTemplate,
		"dynamodb-makefile":       dynamodb.DynamodbMakefileTemplate,
		"dynamodb-docker-compose": dynamodb.DynamodbDockerComposeTemplate,

		"web-app-readme":            webApp.WebAppReadmeTemplate,
		"web-app-env":               webApp.WebAppEnvTemplate,
		"web-app-makefile":          webApp.WebAppMakefileTemplate,
		"web-app-http-routes":       webApp.WebAppHttpRoutesTemplate,
		"web-app-http-users":        webApp.WebAppHttpUsersTemplate,
		"web-app-chi-routes":        webApp.WebAppChiRoutesTemplate,
		"web-app-chi-users":         webApp.WebAppChiUsersTemplate,
		"web-app-echo-routes":       webApp.WebAppEchoRoutesTemplate,
		"web-app-echo-users":        webApp.WebAppEchoUsersTemplate,
		"web-app-gin-routes":        webApp.WebAppGinRoutesTemplate,
		"web-app-gin-users":         webApp.WebAppGinUsersTemplate,
		"web-app-gorilla-routes":    webApp.WebAppGorillaRoutesTemplate,
		"web-app-gorilla-users":     webApp.WebAppGorillaUsersTemplate,
		"web-app-httprouter-routes": webApp.WebAppHttprouterRoutesTemplate,
		"web-app-httprouter-users":  webApp.WebAppHttprouterUsersTemplate,

		"redis-microservice-readme":         redisMicroservice.RedisMicroserviceReadmeTemplate,
		"redis-microservice-env":            redisMicroservice.RedisMicroserviceEnvTemplate,
		"redis-microservice-makefile":       redisMicroservice.RedisMicroserviceMakefileTemplate,
		"redis-microservice-docker-compose": redisMicroservice.RedisMicroserviceDockerComposeTemplate,
		"redis-microservice-broker":         redisMicroservice.RedisMicroserviceBrokerTemplate,

		"kafka-microservice-readme":         kafkaMicroservice.KafkaMicroserviceReadmeTemplate,
		"kafka-microservice-env":            kafkaMicroservice.KafkaMicroserviceEnvTemplate,
		"kafka-microservice-makefile":       kafkaMicroservice.KafkaMicroserviceMakefileTemplate,
		"kafka-microservice-docker-compose": kafkaMicroservice.KafkaMicroserviceDockerComposeTemplate,
		"kafka-microservice-broker":         kafkaMicroservice.KafkaMicroserviceBrokerTemplate,

		"rabbitmq-microservice-readme":         rabbitmqMicroservice.RabbitmqMicroserviceReadmeTemplate,
		"rabbitmq-microservice-env":            rabbitmqMicroservice.RabbitmqMicroserviceEnvTemplate,
		"rabbitmq-microservice-makefile":       rabbitmqMicroservice.RabbitmqMicroserviceMakefileTemplate,
		"rabbitmq-microservice-docker-compose": rabbitmqMicroservice.RabbitmqMicroserviceDockerComposeTemplate,
		"rabbitmq-microservice-broker":         rabbitmqMicroservice.RabbitmqMicroserviceBrokerTemplate,
	}
}
