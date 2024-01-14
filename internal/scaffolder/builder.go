package scaffolder

type tmpBuilder interface {
	build()
}

type tmpDirector struct {
	builder tmpBuilder
}

func (w *tmpDirector) construct() {
	w.builder.build()
}

type tmpBuilderFactory func(p Project) tmpBuilder

var tmpBuilderMap = map[string]tmpBuilderFactory{
	"api": func(p Project) tmpBuilder {
		apiBuilderFactory := apiBuilderMap[p.Platform]
		return apiBuilderFactory(p)
	},
	"hello-world": func(p Project) tmpBuilder {
		helloWorldBuilderFactory := helloWorldBuilderMap[p.ProjectType]
		return helloWorldBuilderFactory(p)
	},
	"web-app": func(p Project) tmpBuilder {
		webAppBuilderFactory := webAppBuilderMap[p.ProjectType]
		return webAppBuilderFactory(p)
	},
	"microservice": func(p Project) tmpBuilder {
		webAppBuilderFactory := microserviceBuilderMap[p.Platform]
		return webAppBuilderFactory(p)
	},
}
