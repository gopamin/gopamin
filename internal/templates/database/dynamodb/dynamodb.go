package templates

import (
	_ "embed"
)

//go:embed files/internal/databases/dynamodb.tmpl
var dynamodbDatabase []byte

//go:embed files/env.tmpl
var dynamodbEnv []byte

//go:embed files/makefile.tmpl
var dynamodbMakefile []byte

//go:embed files/readme.tmpl
var dynamodbReadme []byte

//go:embed files/docker-compose.tmpl
var dynamodbDockerCompose []byte

//go:embed files/cmd/main.tmpl
var dynamodbMain []byte

func DynamodbMainTemplate() ([]byte, string) {
	return dynamodbMain, "cmd/main.go"
}

func DynamodbDockerComposeTemplate() ([]byte, string) {
	return dynamodbDockerCompose, "docker-compose.yml"
}

func DynamodbDatabaseTemplate() ([]byte, string) {
	return dynamodbDatabase, "internal/databases/dynamodb.go"
}

func DynamodbEnvTemplate() ([]byte, string) {
	return dynamodbEnv, ".env"
}

func DynamodbMakefileTemplate() ([]byte, string) {
	return dynamodbMakefile, "Makefile"
}

func DynamodbReadmeTemplate() ([]byte, string) {
	return dynamodbReadme, "README.md"
}
