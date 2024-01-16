package scaffolder

import "fmt"

type helloWorldProjectBuilder struct {
	project *Project
}

func (h *helloWorldProjectBuilder) build() {
	fileGenerator("hello-world-readme", h.project)
	fileGenerator("hello-world-makefile", h.project)
	fileGenerator("hello-world-env", h.project)

	fileGenerator("hello-world-load-env", h.project)
	fileGenerator("hello-world-main", h.project)

	fmt.Printf("%v project created successfully", h.project.Name)
}

type helloWorldBuilderFactory func(p *Project) boilerplateBuilder

var helloWorldBuilderMap = map[string]helloWorldBuilderFactory{
	"hello-world": func(p *Project) boilerplateBuilder {
		return &helloWorldProjectBuilder{p}
	},
}
