package scaffolder

import "fmt"

type webAppProjectBuilder struct {
	project *Project
}

func (b *webAppProjectBuilder) build() {
	fileGenerator([]string{"load-env"}, b.project)
	fileGenerator([]string{"web-app-index-page"}, b.project)
	fileGenerator([]string{"web-app-styles"}, b.project)
	fileGenerator([]string{"web-app-main"}, b.project)
	fileGenerator([]string{"web-app-readme"}, b.project)
	fileGenerator([]string{"web-app-env"}, b.project)
	fileGenerator([]string{"web-app-makefile"}, b.project)

	fmt.Printf("%v project created successfully", b.project.Name)
}

type webAppBuilderFactory func(p *Project) boilerplateBuilder

var webAppBuilderMap = map[string]webAppBuilderFactory{
	"web-app": func(p *Project) boilerplateBuilder {
		return &webAppProjectBuilder{p}
	},
}
