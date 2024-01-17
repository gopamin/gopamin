package scaffolder

import "fmt"

type helloWorldProjectBuilder struct {
	project *Project
}

func (h *helloWorldProjectBuilder) build() {
	readme := []string{"hello-world-readme"}
	makefile := []string{"hello-world-makefile"}
	main := []string{"hello-world-main"}
	env := []string{"hello-world-env"}

	if h.project.Database != "" {
		readme = append(readme, h.project.Database+"-readme")
		makefile = append(makefile, h.project.Database+"-makefile")
		env = append(env, h.project.Database+"-env")
		main = []string{h.project.Database + "-main"}
		dbSelector(h.project)
	}

	fileGenerator(env, h.project)
	fileGenerator(main, h.project)
	fileGenerator(readme, h.project)
	fileGenerator(makefile, h.project)
	fileGenerator([]string{"hello-world-load-env"}, h.project)

	fmt.Printf("%v project created successfully", h.project.Name)
}

type helloWorldBuilderFactory func(p *Project) boilerplateBuilder

var helloWorldBuilderMap = map[string]helloWorldBuilderFactory{
	"hello-world": func(p *Project) boilerplateBuilder {
		return &helloWorldProjectBuilder{p}
	},
}
