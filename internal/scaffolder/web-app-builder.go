package scaffolder

import "fmt"

type webAppProjectBuilder struct {
	project *Project
}

func (w *webAppProjectBuilder) build() {
	fmt.Printf("%v\n", w.project.ProjectType)
}

type webAppBuilderFactory func(p *Project) boilerplateBuilder

var webAppBuilderMap = map[string]webAppBuilderFactory{
	"web-app": func(p *Project) boilerplateBuilder {
		return &webAppProjectBuilder{p}
	},
}
