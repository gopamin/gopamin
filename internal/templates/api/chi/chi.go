package templates

import (
	_ "embed"
)

//go:embed files/env.tmpl
var apiChiEnv []byte

func ApiChiEnvTemplate() ([]byte, string) {
	return apiChiEnv, ".env"
}

//go:embed files/makefile.tmpl
var apiChiMakefile []byte

func ApiChiMakefileTemplate() ([]byte, string) {
	return apiChiMakefile, "Makefile"
}

//go:embed files/readme.tmpl
var apiChiReadme []byte

func ApiChiReadmeTemplate() ([]byte, string) {
	return apiChiReadme, "README.md"
}

//go:embed files/internal/adapters/handlers/chi/routes.tmpl
var apiChiRoutes []byte

func ApiChiRoutesTemplate() ([]byte, string) {
	return apiChiRoutes, "internal/adapters/handlers/chi/routes.go"
}

//go:embed files/internal/adapters/handlers/chi/users.tmpl
var apiChiUsers []byte

func ApiChiUsersTemplate() ([]byte, string) {
	return apiChiUsers, "internal/adapters/handlers/chi/users.go"
}
