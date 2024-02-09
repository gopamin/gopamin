package templates

import (
	_ "embed"
)

//go:embed files/env.tmpl
var apiHttpEnv []byte

func ApiHttpEnvTemplate() ([]byte, string) {
	return apiHttpEnv, ".env"
}

//go:embed files/makefile.tmpl
var apiHttpMakefile []byte

func ApiHttpMakefileTemplate() ([]byte, string) {
	return apiHttpMakefile, "Makefile"
}

//go:embed files/readme.tmpl
var apiHttpReadme []byte

func ApiHttpReadmeTemplate() ([]byte, string) {
	return apiHttpReadme, "README.md"
}

//go:embed files/cmd/main.tmpl
var apiHttpMain []byte

func ApiHttpMainTemplate() ([]byte, string) {
	return apiHttpMain, "cmd/main.go"
}

//go:embed files/internal/adapters/handlers/built-in-http/routes.tmpl
var apiHttpRoutes []byte

func ApiHttpRoutesTemplate() ([]byte, string) {
	return apiHttpRoutes, "internal/adapters/handlers/built-in-http/routes.go"
}

//go:embed files/internal/adapters/handlers/built-in-http/users.tmpl
var apiHttpUsers []byte

func ApiHttpUsersTemplate() ([]byte, string) {
	return apiHttpUsers, "internal/adapters/handlers/built-in-http/users.go"
}
