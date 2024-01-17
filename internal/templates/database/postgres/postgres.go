package templates

import (
	_ "embed"
)

//go:embed files/internal/databases/postgres.tmpl
var postgresDatabase []byte

//go:embed files/env.tmpl
var postgresEnv []byte

//go:embed files/makefile.tmpl
var postgresMakefile []byte

//go:embed files/readme.tmpl
var postgresReadme []byte

//go:embed files/docker-compose.tmpl
var postgresDockerCompose []byte

//go:embed files/cmd/main.tmpl
var postgresMain []byte

func PostgresMainTemplate() ([]byte, string) {
	return postgresMain, "cmd/main.go"
}

func PostgresDockerComposeTemplate() ([]byte, string) {
	return postgresDockerCompose, "docker-compose.yml"
}

func PostgresDatabaseTemplate() ([]byte, string) {
	return postgresDatabase, "internal/databases/postgres.go"
}

func PostgresEnvTemplate() ([]byte, string) {
	return postgresEnv, ".env"
}

func PostgresMakefileTemplate() ([]byte, string) {
	return postgresMakefile, "Makefile"
}

func PostgresReadmeTemplate() ([]byte, string) {
	return postgresReadme, "README.md"
}
