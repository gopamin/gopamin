package scaffolder

import "fmt"

type apiBuilder interface {
	build()
}

type graphqlBuilder struct {
	project Project
}

func (g *graphqlBuilder) build() {
	fmt.Printf("%v -> %v \n", g.project.ProjectType, g.project.Platform)
}

type echoBuilder struct {
	project Project
}

func (e *echoBuilder) build() {
	fmt.Printf("%v -> %v \n", e.project.ProjectType, e.project.Platform)
}

type chiBuilder struct {
	project Project
}

func (c *chiBuilder) build() {
	fmt.Printf("%v -> %v \n", c.project.ProjectType, c.project.Platform)
}

type ginBuilder struct {
	project Project
}

func (g *ginBuilder) build() {
	fmt.Printf("%v -> %v \n", g.project.ProjectType, g.project.Platform)
}

type fiberBuilder struct {
	project Project
}

func (f *fiberBuilder) build() {
	fmt.Printf("%v -> %v \n", f.project.ProjectType, f.project.Platform)
}

type httprouterBuilder struct {
	project Project
}

func (h *httprouterBuilder) build() {
	fmt.Printf("%v -> %v \n", h.project.ProjectType, h.project.Platform)
}

type gorillaBuilder struct {
	project Project
}

func (g *gorillaBuilder) build() {
	fmt.Printf("%v -> %v \n", g.project.ProjectType, g.project.Platform)
}

type builtInBuilder struct {
	project Project
}

func (b *builtInBuilder) build() {
	fmt.Printf("%v -> %v \n", b.project.ProjectType, b.project.Platform)
}

type apiDirector struct {
	builder apiBuilder
}

func (a *apiDirector) construct() {
	a.builder.build()
}

type apiBuilderFactory func(p Project) apiBuilder

var apiBuilderMap = map[string]apiBuilderFactory{
	"graphql": func(p Project) apiBuilder {
		return &graphqlBuilder{p}
	},
	"echo": func(p Project) apiBuilder {
		return &echoBuilder{p}
	},
	"chi": func(p Project) apiBuilder {
		return &chiBuilder{p}
	},
	"gin": func(p Project) apiBuilder {
		return &ginBuilder{p}
	},
	"fiber": func(p Project) apiBuilder {
		return &fiberBuilder{p}
	},
	"httprouter": func(p Project) apiBuilder {
		return &httprouterBuilder{p}
	},
	"gorilla": func(p Project) apiBuilder {
		return &gorillaBuilder{p}
	},
	"built-in": func(p Project) apiBuilder {
		return &builtInBuilder{p}
	},
}
