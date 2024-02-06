package templates

import (
	_ "embed"
)

//go:embed files/env.tmpl
var mongodbEnv []byte

func MongodbEnvTemplate() ([]byte, string) {
	return mongodbEnv, ".env"
}

//go:embed files/makefile.tmpl
var mongodbMakefile []byte

func MongodbMakefileTemplate() ([]byte, string) {
	return mongodbMakefile, "Makefile"
}

//go:embed files/readme.tmpl
var mongodbReadme []byte

func MongodbReadmeTemplate() ([]byte, string) {
	return mongodbReadme, "README.md"
}

//go:embed files/docker-compose.tmpl
var mongodbDockerCompose []byte

func MongodbDockerComposeTemplate() ([]byte, string) {
	return mongodbDockerCompose, "docker-compose.yml"
}

//go:embed files/cmd/main.tmpl
var mongodbMain []byte

func MongodbMainTemplate() ([]byte, string) {
	return mongodbMain, "cmd/main.go"
}

//go:embed files/internal/adapters/repositories/mongodb/repository.tmpl
var mongodbRepository []byte

func MongodbRepositoryTemplate() ([]byte, string) {
	return mongodbRepository, "internal/adapters/mongodb/repository.go"
}
