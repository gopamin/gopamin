package scaffolder

import "fmt"

type graphqlBuilder struct {
	project *Project
}

func (g *graphqlBuilder) build() {
	fmt.Printf("%v -> %v \n", g.project.ProjectType, g.project.Platform)
}

type echoBuilder struct {
	project *Project
}

func (e *echoBuilder) build() {
	fmt.Printf("%v -> %v \n", e.project.ProjectType, e.project.Platform)
}

type chiBuilder struct {
	project *Project
}

func (c *chiBuilder) build() {
	fmt.Printf("%v -> %v \n", c.project.ProjectType, c.project.Platform)
}

type ginBuilder struct {
	project *Project
}

func (g *ginBuilder) build() {
	fmt.Printf("%v -> %v \n", g.project.ProjectType, g.project.Platform)
}

type fiberBuilder struct {
	project *Project
}

func (f *fiberBuilder) build() {
	fmt.Printf("%v -> %v \n", f.project.ProjectType, f.project.Platform)
}

type httprouterBuilder struct {
	project *Project
}

func (h *httprouterBuilder) build() {
	fmt.Printf("%v -> %v \n", h.project.ProjectType, h.project.Platform)
}

type gorillaBuilder struct {
	project *Project
}

func (g *gorillaBuilder) build() {
	fmt.Printf("%v -> %v \n", g.project.ProjectType, g.project.Platform)
}

type httpBuilder struct {
	project *Project
}

func (h *httpBuilder) build() {
	fileGenerator("api-http-readme", h.project)
	fileGenerator("api-http-makefile", h.project)
	fileGenerator("api-http-env", h.project)

	fileGenerator("api-http-load-env", h.project)
	fileGenerator("api-http-main", h.project)
	fmt.Printf("%v project created successfully", h.project.Name)
}

type apiBuilderFactory func(p *Project) boilerplateBuilder

var apiBuilderMap = map[string]apiBuilderFactory{
	"graphql": func(p *Project) boilerplateBuilder {
		return &graphqlBuilder{p}
	},
	"echo": func(p *Project) boilerplateBuilder {
		return &echoBuilder{p}
	},
	"chi": func(p *Project) boilerplateBuilder {
		return &chiBuilder{p}
	},
	"gin": func(p *Project) boilerplateBuilder {
		return &ginBuilder{p}
	},
	"fiber": func(p *Project) boilerplateBuilder {
		return &fiberBuilder{p}
	},
	"httprouter": func(p *Project) boilerplateBuilder {
		return &httprouterBuilder{p}
	},
	"gorilla": func(p *Project) boilerplateBuilder {
		return &gorillaBuilder{p}
	},
	"http": func(p *Project) boilerplateBuilder {
		return &httpBuilder{p}
	},
}
