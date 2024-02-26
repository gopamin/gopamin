package scaffolder

import "fmt"

type echoBuilder struct {
	project *Project
}

func (b *echoBuilder) build() {
	env := []string{"api-echo-env"}
	readme := []string{"api-echo-readme"}
	makefile := []string{"api-echo-makefile"}

	if b.project.Database == "" {
		b.project.Database = "mock"
		dbSelector(b.project)
		fileGenerator([]string{"api-main"}, b.project)
	} else {
		fileGenerator([]string{"api-main-with-db"}, b.project)
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
	fileGenerator([]string{"server"}, b.project)
	fileGenerator([]string{"api-echo-routes"}, b.project)
	fileGenerator([]string{"api-echo-users"}, b.project)
	fileGenerator([]string{"api-errors"}, b.project)
	fileGenerator([]string{"api-response"}, b.project)
	fileGenerator([]string{"router-inferface"}, b.project)

	goGetPackages(b.project.Path, []string{"github.com/labstack/echo/v4"})

	fmt.Printf("%v project created successfully", b.project.Name)
}

type chiBuilder struct {
	project *Project
}

func (b *chiBuilder) build() {
	env := []string{"api-chi-env"}
	readme := []string{"api-chi-readme"}
	makefile := []string{"api-chi-makefile"}

	if b.project.Database == "" {
		b.project.Database = "mock"
		dbSelector(b.project)
		fileGenerator([]string{"api-main"}, b.project)
	} else {
		fileGenerator([]string{"api-main-with-db"}, b.project)
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
	fileGenerator([]string{"server"}, b.project)
	fileGenerator([]string{"api-chi-routes"}, b.project)
	fileGenerator([]string{"api-chi-users"}, b.project)
	fileGenerator([]string{"api-errors"}, b.project)
	fileGenerator([]string{"api-response"}, b.project)
	fileGenerator([]string{"router-inferface"}, b.project)

	goGetPackages(b.project.Path, []string{"github.com/go-chi/chi/v5"})

	fmt.Printf("%v project created successfully", b.project.Name)
}

type ginBuilder struct {
	project *Project
}

func (b *ginBuilder) build() {
	env := []string{"api-gin-env"}
	readme := []string{"api-gin-readme"}
	makefile := []string{"api-gin-makefile"}

	if b.project.Database == "" {
		b.project.Database = "mock"
		dbSelector(b.project)
		fileGenerator([]string{"api-main"}, b.project)
	} else {
		fileGenerator([]string{"api-main-with-db"}, b.project)
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
	fileGenerator([]string{"server"}, b.project)
	fileGenerator([]string{"api-gin-routes"}, b.project)
	fileGenerator([]string{"api-gin-users"}, b.project)
	fileGenerator([]string{"api-errors"}, b.project)
	fileGenerator([]string{"api-response"}, b.project)
	fileGenerator([]string{"router-inferface"}, b.project)

	goGetPackages(b.project.Path, []string{"github.com/gin-gonic/gin"})

	fmt.Printf("%v project created successfully", b.project.Name)
}

type httprouterBuilder struct {
	project *Project
}

func (b *httprouterBuilder) build() {
	env := []string{"api-httprouter-env"}
	readme := []string{"api-httprouter-readme"}
	makefile := []string{"api-httprouter-makefile"}

	if b.project.Database == "" {
		b.project.Database = "mock"
		dbSelector(b.project)
		fileGenerator([]string{"api-main"}, b.project)
	} else {
		fileGenerator([]string{"api-main-with-db"}, b.project)
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
	fileGenerator([]string{"server"}, b.project)
	fileGenerator([]string{"api-httprouter-routes"}, b.project)
	fileGenerator([]string{"api-httprouter-users"}, b.project)
	fileGenerator([]string{"api-errors"}, b.project)
	fileGenerator([]string{"api-response"}, b.project)
	fileGenerator([]string{"router-inferface"}, b.project)

	goGetPackages(b.project.Path, []string{"github.com/julienschmidt/httprouter"})

	fmt.Printf("%v project created successfully", b.project.Name)
}

type gorillaBuilder struct {
	project *Project
}

func (b *gorillaBuilder) build() {
	env := []string{"api-gorilla-env"}
	readme := []string{"api-gorilla-readme"}
	makefile := []string{"api-gorilla-makefile"}

	if b.project.Database == "" {
		b.project.Database = "mock"
		dbSelector(b.project)
		fileGenerator([]string{"api-main"}, b.project)
	} else {
		fileGenerator([]string{"api-main-with-db"}, b.project)
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
	fileGenerator([]string{"server"}, b.project)
	fileGenerator([]string{"api-gorilla-routes"}, b.project)
	fileGenerator([]string{"api-gorilla-users"}, b.project)
	fileGenerator([]string{"api-errors"}, b.project)
	fileGenerator([]string{"api-response"}, b.project)
	fileGenerator([]string{"router-inferface"}, b.project)

	goGetPackages(b.project.Path, []string{"github.com/gorilla/mux"})

	fmt.Printf("%v project created successfully", b.project.Name)
}

type httpBuilder struct {
	project *Project
}

func (b *httpBuilder) build() {
	env := []string{"api-http-env"}
	readme := []string{"api-http-readme"}
	makefile := []string{"api-http-makefile"}

	if b.project.Database == "" {
		b.project.Database = "mock"
		dbSelector(b.project)
		fileGenerator([]string{"api-main"}, b.project)
	} else {
		fileGenerator([]string{"api-main-with-db"}, b.project)
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
	fileGenerator([]string{"server"}, b.project)
	fileGenerator([]string{"api-http-routes"}, b.project)
	fileGenerator([]string{"api-http-users"}, b.project)
	fileGenerator([]string{"api-errors"}, b.project)
	fileGenerator([]string{"api-response"}, b.project)
	fileGenerator([]string{"router-inferface"}, b.project)

	fmt.Printf("%v project created successfully", b.project.Name)
}

type apiBuilderFactory func(p *Project) boilerplateBuilder

var apiBuilderMap = map[string]apiBuilderFactory{
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
