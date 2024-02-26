package scaffolder

import "fmt"

type kafkaMicroserviceBuilder struct {
	project *Project
}

func (b *kafkaMicroserviceBuilder) build() {
	fileGenerator([]string{"kafka-microservice-readme"}, b.project)
	fileGenerator([]string{"kafka-microservice-env"}, b.project)
	fileGenerator([]string{"kafka-microservice-makefile"}, b.project)
	fileGenerator([]string{"kafka-microservice-docker-compose"}, b.project)
	fileGenerator([]string{"kafka-microservice-broker"}, b.project)
	fileGenerator([]string{"configs"}, b.project)
	fileGenerator([]string{"tools"}, b.project)
	fileGenerator([]string{"api-main"}, b.project)
	fileGenerator([]string{"message"}, b.project)
	fileGenerator([]string{"broker-service-interface"}, b.project)
	fileGenerator([]string{"broker-service"}, b.project)
	fileGenerator([]string{"message-broker-interface"}, b.project)
	fileGenerator([]string{"microservice-main"}, b.project)

	goGetPackages(b.project.Path, []string{"github.com/segmentio/kafka-go"})

	fmt.Printf("%v project created successfully", b.project.Name)
}

type rabbitmqMicroserviceBuilder struct {
	project *Project
}

func (b *rabbitmqMicroserviceBuilder) build() {
	fileGenerator([]string{"rabbitmq-microservice-readme"}, b.project)
	fileGenerator([]string{"rabbitmq-microservice-env"}, b.project)
	fileGenerator([]string{"rabbitmq-microservice-makefile"}, b.project)
	fileGenerator([]string{"rabbitmq-microservice-docker-compose"}, b.project)
	fileGenerator([]string{"rabbitmq-microservice-broker"}, b.project)
	fileGenerator([]string{"configs"}, b.project)
	fileGenerator([]string{"tools"}, b.project)
	fileGenerator([]string{"api-main"}, b.project)
	fileGenerator([]string{"message"}, b.project)
	fileGenerator([]string{"broker-service-interface"}, b.project)
	fileGenerator([]string{"broker-service"}, b.project)
	fileGenerator([]string{"message-broker-interface"}, b.project)
	fileGenerator([]string{"microservice-main"}, b.project)

	goGetPackages(b.project.Path, []string{"github.com/streadway/amqp"})

	fmt.Printf("%v project created successfully", b.project.Name)
}

type redisMicroserviceBuilder struct {
	project *Project
}

func (b *redisMicroserviceBuilder) build() {
	fileGenerator([]string{"redis-microservice-readme"}, b.project)
	fileGenerator([]string{"redis-microservice-env"}, b.project)
	fileGenerator([]string{"redis-microservice-makefile"}, b.project)
	fileGenerator([]string{"redis-microservice-docker-compose"}, b.project)
	fileGenerator([]string{"redis-microservice-broker"}, b.project)
	fileGenerator([]string{"configs"}, b.project)
	fileGenerator([]string{"tools"}, b.project)
	fileGenerator([]string{"api-main"}, b.project)
	fileGenerator([]string{"message"}, b.project)
	fileGenerator([]string{"broker-service-interface"}, b.project)
	fileGenerator([]string{"broker-service"}, b.project)
	fileGenerator([]string{"message-broker-interface"}, b.project)
	fileGenerator([]string{"microservice-main"}, b.project)

	goGetPackages(b.project.Path, []string{"github.com/redis/go-redis/v9"})

	fmt.Printf("%v project created successfully", b.project.Name)
}

type microserviceBuilderFactory func(p *Project) boilerplateBuilder

var microserviceBuilderMap = map[string]microserviceBuilderFactory{
	"kafka": func(p *Project) boilerplateBuilder {
		return &kafkaMicroserviceBuilder{p}
	},
	"rabbitmq": func(p *Project) boilerplateBuilder {
		return &rabbitmqMicroserviceBuilder{p}
	},
	"redis": func(p *Project) boilerplateBuilder {
		return &redisMicroserviceBuilder{p}
	},
}
