package scaffolder

import "fmt"

type webAppBuilder interface {
	build()
}

type webAppProjectBuilder struct {
	project Project
}

func (w *webAppProjectBuilder) build() {
	fmt.Printf("%v\n", w.project.ProjectType)
}

type webAppDirector struct {
	builder webAppBuilder
}

func (w *webAppDirector) construct() {
	w.builder.build()
}

type webAppBuilderFactory func(p Project) webAppBuilder

var webAppBuilderMap = map[string]webAppBuilderFactory{
	"web-app": func(p Project) webAppBuilder {
		return &webAppProjectBuilder{p}
	},
}
