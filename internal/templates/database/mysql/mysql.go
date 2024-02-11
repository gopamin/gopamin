package templates

import (
	_ "embed"
)

//go:embed files/env.tmpl
var mysqlEnv []byte

func MysqlEnvTemplate() ([]byte, string) {
	return mysqlEnv, ".env"
}

//go:embed files/docker-compose.tmpl
var mysqlDockerCompose []byte

func MysqlDockerComposeTemplate() ([]byte, string) {
	return mysqlDockerCompose, "docker-compose.yml"
}

//go:embed files/makefile.tmpl
var mysqlMakefile []byte

func MysqlMakefileTemplate() ([]byte, string) {
	return mysqlMakefile, "Makefile"
}

//go:embed files/readme.tmpl
var mysqlReadme []byte

func MysqlReadmeTemplate() ([]byte, string) {
	return mysqlReadme, "README.md"
}

//go:embed files/internal/adapters/repositories/mysql/repository.tmpl
var mysqlRepository []byte

func MysqlRepositoryTemplate() ([]byte, string) {
	return mysqlRepository, "internal/adapters/repositories/mysql/repository.go"
}
