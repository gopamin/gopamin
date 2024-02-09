package scaffolder

import "fmt"

// type graphqlBuilder struct {
// 	project *Project
// }

// func (g *graphqlBuilder) build() {
// 	fmt.Printf("%v -> %v \n", g.project.ProjectType, g.project.Platform)
// }

type echoBuilder struct {
	project *Project
}

func (e *echoBuilder) build() {
	fileGenerator([]string{"api-echo-main"}, e.project)
	fileGenerator([]string{"api-echo-env"}, e.project)
	fileGenerator([]string{"api-echo-readme"}, e.project)
	fileGenerator([]string{"api-echo-makefile"}, e.project)
	fileGenerator([]string{"load-env"}, e.project)
	fileGenerator([]string{"api-server"}, e.project)
	fileGenerator([]string{"api-echo-routes"}, e.project)
	fileGenerator([]string{"api-echo-users"}, e.project)
	fileGenerator([]string{"api-errors"}, e.project)
	fileGenerator([]string{"api-response"}, e.project)

	fileGenerator([]string{"mock-repository"}, e.project)
	fileGenerator([]string{"user-test"}, e.project)
	fileGenerator([]string{"user"}, e.project)
	fileGenerator([]string{"user-repository-interface"}, e.project)
	fileGenerator([]string{"router-inferface"}, e.project)
	fileGenerator([]string{"user-service-interface"}, e.project)
	fileGenerator([]string{"user-service"}, e.project)
	fileGenerator([]string{"user-service-test"}, e.project)
	goGetPackages(e.project.Path, []string{"github.com/labstack/echo/v4"})

	fmt.Printf("%v project created successfully", e.project.Name)
}

type chiBuilder struct {
	project *Project
}

func (c *chiBuilder) build() {
	fileGenerator([]string{"api-chi-main"}, c.project)
	fileGenerator([]string{"api-chi-env"}, c.project)
	fileGenerator([]string{"api-chi-readme"}, c.project)
	fileGenerator([]string{"api-chi-makefile"}, c.project)
	fileGenerator([]string{"load-env"}, c.project)
	fileGenerator([]string{"api-server"}, c.project)
	fileGenerator([]string{"api-chi-routes"}, c.project)
	fileGenerator([]string{"api-chi-users"}, c.project)
	fileGenerator([]string{"api-errors"}, c.project)
	fileGenerator([]string{"api-response"}, c.project)

	fileGenerator([]string{"mock-repository"}, c.project)
	fileGenerator([]string{"user-test"}, c.project)
	fileGenerator([]string{"user"}, c.project)
	fileGenerator([]string{"user-repository-interface"}, c.project)
	fileGenerator([]string{"router-inferface"}, c.project)
	fileGenerator([]string{"user-service-interface"}, c.project)
	fileGenerator([]string{"user-service"}, c.project)
	fileGenerator([]string{"user-service-test"}, c.project)
	goGetPackages(c.project.Path, []string{"github.com/go-chi/chi/v5"})

	fmt.Printf("%v project created successfully", c.project.Name)
}

type ginBuilder struct {
	project *Project
}

func (g *ginBuilder) build() {
	fileGenerator([]string{"api-gin-main"}, g.project)
	fileGenerator([]string{"api-gin-env"}, g.project)
	fileGenerator([]string{"api-gin-readme"}, g.project)
	fileGenerator([]string{"api-gin-makefile"}, g.project)
	fileGenerator([]string{"load-env"}, g.project)
	fileGenerator([]string{"api-server"}, g.project)
	fileGenerator([]string{"api-gin-routes"}, g.project)
	fileGenerator([]string{"api-gin-users"}, g.project)
	fileGenerator([]string{"api-errors"}, g.project)
	fileGenerator([]string{"api-response"}, g.project)

	fileGenerator([]string{"mock-repository"}, g.project)
	fileGenerator([]string{"user-test"}, g.project)
	fileGenerator([]string{"user"}, g.project)
	fileGenerator([]string{"user-repository-interface"}, g.project)
	fileGenerator([]string{"router-inferface"}, g.project)
	fileGenerator([]string{"user-service-interface"}, g.project)
	fileGenerator([]string{"user-service"}, g.project)
	fileGenerator([]string{"user-service-test"}, g.project)
	goGetPackages(g.project.Path, []string{"github.com/gin-gonic/gin"})

	fmt.Printf("%v project created successfully", g.project.Name)
}

type httprouterBuilder struct {
	project *Project
}

func (h *httprouterBuilder) build() {
	fileGenerator([]string{"api-httprouter-main"}, h.project)
	fileGenerator([]string{"api-httprouter-env"}, h.project)
	fileGenerator([]string{"api-httprouter-readme"}, h.project)
	fileGenerator([]string{"api-httprouter-makefile"}, h.project)
	fileGenerator([]string{"load-env"}, h.project)
	fileGenerator([]string{"api-server"}, h.project)
	fileGenerator([]string{"api-httprouter-routes"}, h.project)
	fileGenerator([]string{"api-httprouter-users"}, h.project)
	fileGenerator([]string{"api-errors"}, h.project)
	fileGenerator([]string{"api-response"}, h.project)

	fileGenerator([]string{"mock-repository"}, h.project)
	fileGenerator([]string{"user-test"}, h.project)
	fileGenerator([]string{"user"}, h.project)
	fileGenerator([]string{"user-repository-interface"}, h.project)
	fileGenerator([]string{"router-inferface"}, h.project)
	fileGenerator([]string{"user-service-interface"}, h.project)
	fileGenerator([]string{"user-service"}, h.project)
	fileGenerator([]string{"user-service-test"}, h.project)
	goGetPackages(h.project.Path, []string{"github.com/julienschmidt/httprouter"})

	fmt.Printf("%v project created successfully", h.project.Name)
}

type gorillaBuilder struct {
	project *Project
}

func (g *gorillaBuilder) build() {
	fileGenerator([]string{"api-gorilla-main"}, g.project)
	fileGenerator([]string{"api-gorilla-env"}, g.project)
	fileGenerator([]string{"api-gorilla-readme"}, g.project)
	fileGenerator([]string{"api-gorilla-makefile"}, g.project)
	fileGenerator([]string{"load-env"}, g.project)
	fileGenerator([]string{"api-server"}, g.project)
	fileGenerator([]string{"api-gorilla-routes"}, g.project)
	fileGenerator([]string{"api-gorilla-users"}, g.project)
	fileGenerator([]string{"api-errors"}, g.project)
	fileGenerator([]string{"api-response"}, g.project)

	fileGenerator([]string{"mock-repository"}, g.project)
	fileGenerator([]string{"user-test"}, g.project)
	fileGenerator([]string{"user"}, g.project)
	fileGenerator([]string{"user-repository-interface"}, g.project)
	fileGenerator([]string{"router-inferface"}, g.project)
	fileGenerator([]string{"user-service-interface"}, g.project)
	fileGenerator([]string{"user-service"}, g.project)
	fileGenerator([]string{"user-service-test"}, g.project)
	goGetPackages(g.project.Path, []string{"github.com/gorilla/mux"})

	fmt.Printf("%v project created successfully", g.project.Name)
}

type httpBuilder struct {
	project *Project
}

func (h *httpBuilder) build() {
	fileGenerator([]string{"api-http-main"}, h.project)
	fileGenerator([]string{"api-http-env"}, h.project)
	fileGenerator([]string{"api-http-readme"}, h.project)
	fileGenerator([]string{"api-http-makefile"}, h.project)
	fileGenerator([]string{"load-env"}, h.project)
	fileGenerator([]string{"api-server"}, h.project)
	fileGenerator([]string{"api-http-routes"}, h.project)
	fileGenerator([]string{"api-http-users"}, h.project)
	fileGenerator([]string{"api-errors"}, h.project)
	fileGenerator([]string{"api-response"}, h.project)

	fileGenerator([]string{"mock-repository"}, h.project)
	fileGenerator([]string{"user-test"}, h.project)
	fileGenerator([]string{"user"}, h.project)
	fileGenerator([]string{"user-repository-interface"}, h.project)
	fileGenerator([]string{"router-inferface"}, h.project)
	fileGenerator([]string{"user-service-interface"}, h.project)
	fileGenerator([]string{"user-service"}, h.project)
	fileGenerator([]string{"user-service-test"}, h.project)

	fmt.Printf("%v project created successfully", h.project.Name)
}

type apiBuilderFactory func(p *Project) boilerplateBuilder

var apiBuilderMap = map[string]apiBuilderFactory{
	// "graphql": func(p *Project) boilerplateBuilder {
	// 	return &graphqlBuilder{p}
	// },
	"echo": func(p *Project) boilerplateBuilder {
		return &echoBuilder{p}
	},
	"chi": func(p *Project) boilerplateBuilder {
		return &chiBuilder{p}
	},
	"gin": func(p *Project) boilerplateBuilder {
		return &ginBuilder{p}
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
