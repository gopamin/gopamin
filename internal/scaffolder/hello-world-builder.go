package scaffolder

import "fmt"

type helloWorldProjectBuilder struct {
	project *Project
}

func (h *helloWorldProjectBuilder) build() {
	readme := []string{"hello-world-readme"}
	makefile := []string{"hello-world-makefile"}
	env := []string{"hello-world-env"}

	if h.project.Database == "" {
		fileGenerator([]string{"hello-world-main"}, h.project)
	} else {
		fileGenerator([]string{"hello-world-main-with-db"}, h.project)
		readme = append(readme, h.project.Database+"-readme")
		makefile = append(makefile, h.project.Database+"-makefile")
		env = append(env, h.project.Database+"-env")
		dbSelector(h.project)
	}

	fileGenerator(env, h.project)
	fileGenerator(readme, h.project)
	fileGenerator(makefile, h.project)
	fileGenerator([]string{"load-env"}, h.project)

	fmt.Printf("%v project created successfully", h.project.Name)
}

type helloWorldBuilderFactory func(p *Project) boilerplateBuilder

var helloWorldBuilderMap = map[string]helloWorldBuilderFactory{
	"hello-world": func(p *Project) boilerplateBuilder {
		return &helloWorldProjectBuilder{p}
	},
}
