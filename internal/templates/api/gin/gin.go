package templates

import (
	_ "embed"
)

//go:embed files/env.tmpl
var apiGinEnv []byte

func ApiGinEnvTemplate() ([]byte, string) {
	return apiGinEnv, ".env"
}

//go:embed files/makefile.tmpl
var apiGinMakefile []byte

func ApiGinMakefileTemplate() ([]byte, string) {
	return apiGinMakefile, "Makefile"
}

//go:embed files/readme.tmpl
var apiGinReadme []byte

func ApiGinReadmeTemplate() ([]byte, string) {
	return apiGinReadme, "README.md"
}

//go:embed files/internal/adapters/handlers/gin/routes.tmpl
var apiGinRoutes []byte

func ApiGinRoutesTemplate() ([]byte, string) {
	return apiGinRoutes, "internal/adapters/handlers/gin/routes.go"
}

//go:embed files/internal/adapters/handlers/gin/users.tmpl
var apiGinUsers []byte

func ApiGinUsersTemplate() ([]byte, string) {
	return apiGinUsers, "internal/adapters/handlers/gin/users.go"
}
