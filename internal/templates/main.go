package templates

import (
	apiHttp "github.com/gopamin/cli/internal/templates/api/http"
	common "github.com/gopamin/cli/internal/templates/common"
	dynamodb "github.com/gopamin/cli/internal/templates/database/dynamodb"
	mongodb "github.com/gopamin/cli/internal/templates/database/mongodb"
	mysql "github.com/gopamin/cli/internal/templates/database/mysql"
	postgres "github.com/gopamin/cli/internal/templates/database/postgres"
	redis "github.com/gopamin/cli/internal/templates/database/redis"
	sqlite "github.com/gopamin/cli/internal/templates/database/sqlite"
	helloWorld "github.com/gopamin/cli/internal/templates/hello-world"
)

func Mapper() map[string]func() ([]byte, string) {
	return map[string]func() ([]byte, string){
		"dockerfile":                common.DockerFileTemplate,
		"gitignore":                 common.GitIgnoreTemplate,
		"license":                   common.LicenseTemplate,
		"dockerignore":              common.DockerIgnoreTemplate,
		"mock-repository":           common.MockRepositoryTemplate,
		"user-test":                 common.UserTestTemplate,
		"user":                      common.UserTemplate,
		"user-repository-interface": common.UserRepositoryInterfaceTemplate,
		"user-service-interface":    common.UserServiceInterfaceTemplate,
		"user-service":              common.UserServiceTemplate,
		"user-service-test":         common.UserServiceTestTemplate,

		"hello-world-main":     helloWorld.HelloWorldMainTemplate,
		"hello-world-load-env": helloWorld.HelloWoldLoadEnvTemplate,
		"hello-world-readme":   helloWorld.HelloWorldReadmeTemplate,
		"hello-world-makefile": helloWorld.HelloWorldMakefileTemplate,
		"hello-world-env":      helloWorld.HelloWorldEnvTemplate,

		"api-http-main":     apiHttp.ApiHttpMainTemplate,
		"api-http-load-env": apiHttp.ApiHttpLoadEnvTemplate,
		"api-http-readme":   apiHttp.ApiHttpReadmeTemplate,
		"api-http-makefile": apiHttp.ApiHttpMakefileTemplate,
		"api-http-env":      apiHttp.ApiHttpEnvTemplate,
		"api-http-handlers": apiHttp.ApiHttpUsersHandlerTemplate,
		"api-http-routes":   apiHttp.ApiHttpRoutesTemplate,
		"api-http-server":   apiHttp.ApiHttpServerTemplate,
		"api-http-services": apiHttp.ApiHttpUsersServiceTemplate,

		"mysql-env":            mysql.MysqlEnvTemplate,
		"mysql-main":           mysql.MysqlMainTemplate,
		"mysql-repository":     mysql.MysqlRepositoryTemplate,
		"mysql-readme":         mysql.MysqlReadmeTemplate,
		"mysql-makefile":       mysql.MysqlMakefileTemplate,
		"mysql-docker-compose": mysql.MysqlDockerComposeTemplate,

		"postgres-env":            postgres.PostgresEnvTemplate,
		"postgres-main":           postgres.PostgresMainTemplate,
		"postgres-repository":     postgres.PostgresRepositoryTemplate,
		"postgres-readme":         postgres.PostgresReadmeTemplate,
		"postgres-makefile":       postgres.PostgresMakefileTemplate,
		"postgres-docker-compose": postgres.PostgresDockerComposeTemplate,

		"mongodb-env":            mongodb.MongodbEnvTemplate,
		"mongodb-main":           mongodb.MongodbMainTemplate,
		"mongodb-repository":     mongodb.MongodbRepositoryTemplate,
		"mongodb-readme":         mongodb.MongodbReadmeTemplate,
		"mongodb-makefile":       mongodb.MongodbMakefileTemplate,
		"mongodb-docker-compose": mongodb.MongodbDockerComposeTemplate,

		"redis-env":            redis.RedisEnvTemplate,
		"redis-main":           redis.RedisMainTemplate,
		"redis-repository":     redis.RedisRepositoryTemplate,
		"redis-readme":         redis.RedisReadmeTemplate,
		"redis-makefile":       redis.RedisMakefileTemplate,
		"redis-docker-compose": redis.RedisDockerComposeTemplate,

		"sqlite-env":        sqlite.SqliteEnvTemplate,
		"sqlite-main":       sqlite.SqliteMainTemplate,
		"sqlite-repository": sqlite.SqliteRepositoryTemplate,
		"sqlite-makefile":   sqlite.SqliteMakefileTemplate,
		"sqlite-readme":     sqlite.SqliteReadmeTemplate,

		"dynamodb-env":            dynamodb.DynamodbEnvTemplate,
		"dynamodb-main":           dynamodb.DynamodbMainTemplate,
		"dynamodb-repository":     dynamodb.DynamodbRepositoryTemplate,
		"dynamodb-readme":         dynamodb.DynamodbReadmeTemplate,
		"dynamodb-makefile":       dynamodb.DynamodbMakefileTemplate,
		"dynamodb-docker-compose": dynamodb.DynamodbDockerComposeTemplate,
	}
}
