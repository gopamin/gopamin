package templates

import (
	_ "embed"
)

//go:embed files/env.tmpl
var webAppEnv []byte

func WebAppEnvTemplate() ([]byte, string) {
	return webAppEnv, ".env"
}

//go:embed files/makefile.tmpl
var webAppMakefile []byte

func WebAppMakefileTemplate() ([]byte, string) {
	return webAppMakefile, "Makefile"
}

//go:embed files/readme.tmpl
var webAppReadme []byte

func WebAppReadmeTemplate() ([]byte, string) {
	return webAppReadme, "README.md"
}
