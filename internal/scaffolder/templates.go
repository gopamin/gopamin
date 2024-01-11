package scaffolder

import (
	_ "embed"
)

//go:embed templates/cmd/main.go.template
var main []byte

//go:embed templates/configs/load-env.template
var loadEnv []byte

//go:embed templates/env.template
var env []byte

//go:embed templates/dockerfile.template
var dockerFile []byte

//go:embed templates/makefile.template
var makeFile []byte

//go:embed templates/readme.template
var readme []byte

//go:embed templates/license.template
var license []byte

//go:embed templates/gitignore.template
var gitIgnore []byte

func mainTemplate() ([]byte, string) {
	return main, "cmd/main.go"
}

func readmeTemplate() ([]byte, string) {
	return readme, "README.md"
}

func dockerFileTemplate() ([]byte, string) {
	return dockerFile, "Dockerfile"
}

func licenseTemplate() ([]byte, string) {
	return license, "LICENSE"
}

func gitIgnoreTemplate() ([]byte, string) {
	return gitIgnore, ".gitignore"
}

func makefileTemplate() ([]byte, string) {
	return makeFile, "Makefile"
}

func envTemplate() ([]byte, string) {
	return env, ".env"
}

func loadEnvTemplate() ([]byte, string) {
	return loadEnv, "configs/load-env.go"
}

func templateMapper() map[string]func() ([]byte, string) {
	return map[string]func() ([]byte, string){
		"main":       mainTemplate,
		"readme":     readmeTemplate,
		"gitignore":  gitIgnoreTemplate,
		"makefile":   makefileTemplate,
		"env":        envTemplate,
		"load-env":   loadEnvTemplate,
		"license":    licenseTemplate,
		"dockerfile": dockerFileTemplate,
	}
}
