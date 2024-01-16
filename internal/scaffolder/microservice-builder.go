package scaffolder

import "fmt"

type kafkaBuilder struct {
	project *Project
}

func (k *kafkaBuilder) build() {
	fmt.Printf("%v -> %v \n", k.project.ProjectType, k.project.Platform)
}

type rabbitmqBuilder struct {
	project *Project
}

func (r *rabbitmqBuilder) build() {
	fmt.Printf("%v -> %v \n", r.project.ProjectType, r.project.Platform)
}

type grpcBuilder struct {
	project *Project
}

func (g *grpcBuilder) build() {
	fmt.Printf("%v -> %v \n", g.project.ProjectType, g.project.Platform)
}

type microserviceBuilderFactory func(p *Project) boilerplateBuilder

var microserviceBuilderMap = map[string]microserviceBuilderFactory{
	"kafka": func(p *Project) boilerplateBuilder {
		return &kafkaBuilder{p}
	},
	"rabbitmq": func(p *Project) boilerplateBuilder {
		return &rabbitmqBuilder{p}
	},
	"grpc": func(p *Project) boilerplateBuilder {
		return &grpcBuilder{p}
	},
}
