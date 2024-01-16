package templates

import (
	_ "embed"
)

//go:embed files/internal/databases/mysql.tmpl
var mysqlDatabase []byte

//go:embed files/env.tmpl
var mysqlEnv []byte

//go:embed files/docker-compose.tmpl
var mysqlDockerCompose []byte

//go:embed files/cmd/main.tmpl
var mysqlMain []byte

func MysqlMainTemplate() ([]byte, string) {
	return mysqlMain, "cmd/main.go"
}

func MysqlDockerComposeTemplate() ([]byte, string) {
	return mysqlDockerCompose, "docker-compose.yml"
}

func MysqlDatabaseTemplate() ([]byte, string) {
	return mysqlDatabase, "internal/databases/mysql.go"
}

func MysqlEnvTemplate() ([]byte, string) {
	return mysqlEnv, ".env"
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
