package templates

import (
	_ "embed"
)

//go:embed files/env.tmpl
var apiEchoEnv []byte

func ApiEchoEnvTemplate() ([]byte, string) {
	return apiEchoEnv, ".env"
}

//go:embed files/makefile.tmpl
var apiEchoMakefile []byte

func ApiEchoMakefileTemplate() ([]byte, string) {
	return apiEchoMakefile, "Makefile"
}

//go:embed files/readme.tmpl
var apiEchoReadme []byte

func ApiEchoReadmeTemplate() ([]byte, string) {
	return apiEchoReadme, "README.md"
}

//go:embed files/internal/adapters/handlers/echo/routes.tmpl
var apiEchoRoutes []byte

func ApiEchoRoutesTemplate() ([]byte, string) {
	return apiEchoRoutes, "internal/adapters/handlers/echo/routes.go"
}

//go:embed files/internal/adapters/handlers/echo/users.tmpl
var apiEchoUsers []byte

func ApiEchoUsersTemplate() ([]byte, string) {
	return apiEchoUsers, "internal/adapters/handlers/echo/users.go"
}
