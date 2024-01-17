package scaffolder

type boilerplateBuilder interface {
	build()
}

type director struct {
	builder boilerplateBuilder
}

func (w *director) construct() {
	w.builder.build()
}

type builderFactory func(p *Project) boilerplateBuilder

var buildersMap = map[string]builderFactory{
	"api": func(p *Project) boilerplateBuilder {
		apiBuilderFactory := apiBuilderMap[p.Platform]
		return apiBuilderFactory(p)
	},
	"hello-world": func(p *Project) boilerplateBuilder {
		helloWorldBuilderFactory := helloWorldBuilderMap[p.ProjectType]
		return helloWorldBuilderFactory(p)
	},
	"web-app": func(p *Project) boilerplateBuilder {
		webAppBuilderFactory := webAppBuilderMap[p.ProjectType]
		return webAppBuilderFactory(p)
	},
	"microservice": func(p *Project) boilerplateBuilder {
		microserviceBuilderFactory := microserviceBuilderMap[p.Platform]
		return microserviceBuilderFactory(p)
	},
}

func dbSelector(p *Project) {
	databaseBuilderFactory := databaseBuilderMap[p.Database]
	builder := databaseBuilderFactory(p)
	director := &director{builder: builder}
	director.construct()
}