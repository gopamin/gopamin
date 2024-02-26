package scaffolder

import "fmt"

type helloWorldProjectBuilder struct {
	project *Project
}

func (b *helloWorldProjectBuilder) build() {
	readme := []string{"hello-world-readme"}
	makefile := []string{"hello-world-makefile"}
	env := []string{"hello-world-env"}

	if b.project.Database == "" {
		fileGenerator([]string{"hello-world-main"}, b.project)
	} else {
		fileGenerator([]string{"hello-world-main-with-db"}, b.project)
		readme = append(readme, b.project.Database+"-readme")
		makefile = append(makefile, b.project.Database+"-makefile")
		env = append(env, b.project.Database+"-env")
		dbSelector(b.project)
	}

	fileGenerator(env, b.project)
	fileGenerator(readme, b.project)
	fileGenerator(makefile, b.project)
	fileGenerator([]string{"configs"}, b.project)
	fileGenerator([]string{"tools"}, b.project)

	fmt.Printf("%v project created successfully", b.project.Name)
}

type helloWorldBuilderFactory func(p *Project) boilerplateBuilder

var helloWorldBuilderMap = map[string]helloWorldBuilderFactory{
	"hello-world": func(p *Project) boilerplateBuilder {
		return &helloWorldProjectBuilder{p}
	},
}
