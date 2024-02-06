package templates

import (
	_ "embed"
)

//go:embed files/env.tmpl
var postgresEnv []byte

func PostgresEnvTemplate() ([]byte, string) {
	return postgresEnv, ".env"
}

//go:embed files/makefile.tmpl
var postgresMakefile []byte

func PostgresMakefileTemplate() ([]byte, string) {
	return postgresMakefile, "Makefile"
}

//go:embed files/readme.tmpl
var postgresReadme []byte

func PostgresReadmeTemplate() ([]byte, string) {
	return postgresReadme, "README.md"
}

//go:embed files/docker-compose.tmpl
var postgresDockerCompose []byte

func PostgresDockerComposeTemplate() ([]byte, string) {
	return postgresDockerCompose, "docker-compose.yml"
}

//go:embed files/cmd/main.tmpl
var postgresMain []byte

func PostgresMainTemplate() ([]byte, string) {
	return postgresMain, "cmd/main.go"
}

//go:embed files/internal/adapters/repositories/postgres/repository.tmpl
var postgresRepository []byte

func PostgresRepositoryTemplate() ([]byte, string) {
	return postgresRepository, "internal/adapters/repositories/postgres/repository.go"
}
