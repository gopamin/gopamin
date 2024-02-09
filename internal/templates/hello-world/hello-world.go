package templates

import (
	_ "embed"
)

//go:embed files/env.tmpl
var helloWorldEnv []byte

func HelloWorldEnvTemplate() ([]byte, string) {
	return helloWorldEnv, ".env"
}

//go:embed files/makefile.tmpl
var helloWorldMakefile []byte

func HelloWorldMakefileTemplate() ([]byte, string) {
	return helloWorldMakefile, "Makefile"
}

//go:embed files/readme.tmpl
var helloWorldReadme []byte

func HelloWorldReadmeTemplate() ([]byte, string) {
	return helloWorldReadme, "README.md"
}

//go:embed files/cmd/main.tmpl
var helloWorldMain []byte

func HelloWorldMainTemplate() ([]byte, string) {
	return helloWorldMain, "cmd/main.go"
}
