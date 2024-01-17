package templates

import (
	_ "embed"
)

//go:embed files/internal/databases/mongodb.tmpl
var mongodbDatabase []byte

//go:embed files/env.tmpl
var mongodbEnv []byte

//go:embed files/makefile.tmpl
var mongodbMakefile []byte

//go:embed files/readme.tmpl
var mongodbReadme []byte

//go:embed files/docker-compose.tmpl
var mongodbDockerCompose []byte

//go:embed files/cmd/main.tmpl
var mongodbMain []byte

func MongodbMainTemplate() ([]byte, string) {
	return mongodbMain, "cmd/main.go"
}

func MongodbDockerComposeTemplate() ([]byte, string) {
	return mongodbDockerCompose, "docker-compose.yml"
}

func MongodbDatabaseTemplate() ([]byte, string) {
	return mongodbDatabase, "internal/databases/mongodb.go"
}

func MongodbEnvTemplate() ([]byte, string) {
	return mongodbEnv, ".env"
}

func MongodbMakefileTemplate() ([]byte, string) {
	return mongodbMakefile, "Makefile"
}

func MongodbReadmeTemplate() ([]byte, string) {
	return mongodbReadme, "README.md"
}
