package templates

import (
	_ "embed"
)

//go:embed files/configs/load-env.template
var helloWorldLoadEnv []byte

func HelloWoldLoadEnvTemplate() ([]byte, string) {
	return helloWorldLoadEnv, "configs/load-env.go"
}

//go:embed files/env.template
var helloWorldEnv []byte

func HelloWorldEnvTemplate() ([]byte, string) {
	return helloWorldEnv, ".env"
}

//go:embed files/makefile.template
var helloWorldMakefile []byte

func HelloWorldMakefileTemplate() ([]byte, string) {
	return helloWorldMakefile, "Makefile"
}

//go:embed files/readme.template
var helloWorldReadme []byte

func HelloWorldReadmeTemplate() ([]byte, string) {
	return helloWorldReadme, "README.md"
}

//go:embed files/cmd/main.go.template
var helloWorldMain []byte

func HelloWorldMainTemplate() ([]byte, string) {
	return helloWorldMain, "cmd/main.go"
}
