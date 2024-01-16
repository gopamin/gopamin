package templates

import (
	apiHttp "github.com/gopamin/cli/internal/templates/api/http"
	common "github.com/gopamin/cli/internal/templates/common"
	helloWorld "github.com/gopamin/cli/internal/templates/hello-world"
)

func Mapper() map[string]func() ([]byte, string) {
	return map[string]func() ([]byte, string){
		"hello-world-main":     helloWorld.HelloWorldMainTemplate,
		"hello-world-load-env": helloWorld.HelloWoldLoadEnvTemplate,
		"hello-world-readme":   helloWorld.HelloWorldReadmeTemplate,
		"hello-world-makefile": helloWorld.HelloWorldMakefileTemplate,
		"hello-world-env":      helloWorld.HelloWorldEnvTemplate,

		"api-http-main":     apiHttp.ApiHttpMainTemplate,
		"api-http-load-env": apiHttp.ApiHttpLoadEnvTemplate,
		"api-http-readme":   apiHttp.ApiHttpReadmeTemplate,
		"api-http-makefile": apiHttp.ApiHttpMakefileTemplate,
		"api-http-env":      apiHttp.ApiHttpEnvTemplate,

		"dockerfile":   common.DockerFileTemplate,
		"gitignore":    common.GitIgnoreTemplate,
		"license":      common.LicenseTemplate,
		"dockerignore": common.DockerIgnoreTemplate,
	}
}
