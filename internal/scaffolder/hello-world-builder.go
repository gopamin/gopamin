package scaffolder

import "fmt"

type helloWorldBuilder interface {
	build()
}

type helloWorldProjectBuilder struct {
	project Project
}

func (h *helloWorldProjectBuilder) build() {
	fmt.Printf("%v\n", h.project.ProjectType)
}

type helloWorldDirector struct {
	builder apiBuilder
}

func (a *helloWorldDirector) construct() {
	a.builder.build()
}

type helloWorldBuilderFactory func(p Project) helloWorldBuilder

var helloWorldBuilderMap = map[string]helloWorldBuilderFactory{
	"hello-world": func(p Project) helloWorldBuilder {
		return &helloWorldProjectBuilder{p}
	},
}
