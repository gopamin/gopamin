package templates

import (
	_ "embed"
)

//go:embed files/env.tmpl
var apiGorillaEnv []byte

func ApiGorillaEnvTemplate() ([]byte, string) {
	return apiGorillaEnv, ".env"
}

//go:embed files/makefile.tmpl
var apiGorillaMakefile []byte

func ApiGorillaMakefileTemplate() ([]byte, string) {
	return apiGorillaMakefile, "Makefile"
}

//go:embed files/readme.tmpl
var apiGorillaReadme []byte

func ApiGorillaReadmeTemplate() ([]byte, string) {
	return apiGorillaReadme, "README.md"
}

//go:embed files/internal/adapters/handlers/gorilla/routes.tmpl
var apiGorillaRoutes []byte

func ApiGorillaRoutesTemplate() ([]byte, string) {
	return apiGorillaRoutes, "internal/adapters/handlers/gorilla/routes.go"
}

//go:embed files/internal/adapters/handlers/gorilla/users.tmpl
var apiGorillaUsers []byte

func ApiGorillaUsersTemplate() ([]byte, string) {
	return apiGorillaUsers, "internal/adapters/handlers/gorilla/users.go"
}
