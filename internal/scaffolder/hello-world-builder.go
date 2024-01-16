package scaffolder

import "fmt"

type helloWorldProjectBuilder struct {
	project *Project
}

func (h *helloWorldProjectBuilder) build() {
	readme := []string{"hello-world-readme"}
	makefile := []string{"hello-world-makefile"}
	main := []string{"hello-world-main"}

	if h.project.Database != "" {
		readme = append(readme, "mysql-readme")
		makefile = append(makefile, "mysql-makefile")
		main = []string{"mysql-main"}
		dbSelector(h.project)
	}

	fileGenerator(main, h.project)
	fileGenerator([]string{"hello-world-load-env"}, h.project)
	fileGenerator(readme, h.project)
	fileGenerator(makefile, h.project)
	fileGenerator([]string{"hello-world-env", "mysql-env"}, h.project)

	fmt.Printf("%v project created successfully", h.project.Name)
}

type helloWorldBuilderFactory func(p *Project) boilerplateBuilder

var helloWorldBuilderMap = map[string]helloWorldBuilderFactory{
	"hello-world": func(p *Project) boilerplateBuilder {
		return &helloWorldProjectBuilder{p}
	},
}
