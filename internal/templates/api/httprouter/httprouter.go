package templates

import (
	_ "embed"
)

//go:embed files/env.tmpl
var apiHttprouterEnv []byte

func ApiHttprouterEnvTemplate() ([]byte, string) {
	return apiHttprouterEnv, ".env"
}

//go:embed files/makefile.tmpl
var apiHttprouterMakefile []byte

func ApiHttprouterMakefileTemplate() ([]byte, string) {
	return apiHttprouterMakefile, "Makefile"
}

//go:embed files/readme.tmpl
var apiHttprouterReadme []byte

func ApiHttprouterReadmeTemplate() ([]byte, string) {
	return apiHttprouterReadme, "README.md"
}

//go:embed files/cmd/main.tmpl
var apiHttprouterMain []byte

func ApiHttprouterMainTemplate() ([]byte, string) {
	return apiHttprouterMain, "cmd/main.go"
}

//go:embed files/internal/adapters/handlers/httprouter/routes.tmpl
var apiHttprouterRoutes []byte

func ApiHttprouterRoutesTemplate() ([]byte, string) {
	return apiHttprouterRoutes, "internal/adapters/handlers/httprouter/routes.go"
}

//go:embed files/internal/adapters/handlers/httprouter/users.tmpl
var apiHttprouterUsers []byte

func ApiHttprouterUsersTemplate() ([]byte, string) {
	return apiHttprouterUsers, "internal/adapters/handlers/httprouter/users.go"
}
