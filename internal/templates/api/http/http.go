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

//go:embed files/configs/load-env.tmpl
var apiHttpLoadEnv []byte

func ApiHttpLoadEnvTemplate() ([]byte, string) {
	return apiHttpLoadEnv, "configs/load-env.go"
}

//go:embed files/cmd/main.tmpl
var apiHttpMain []byte

func ApiHttpMainTemplate() ([]byte, string) {
	return apiHttpMain, "cmd/main.go"
}
