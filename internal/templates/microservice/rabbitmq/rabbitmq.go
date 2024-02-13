package templates

import (
	_ "embed"
)

//go:embed files/env.tmpl
var rabbitmqMicroserviceEnv []byte

func RabbitmqMicroserviceEnvTemplate() ([]byte, string) {
	return rabbitmqMicroserviceEnv, ".env"
}

//go:embed files/docker-compose.tmpl
var rabbitmqMicroserviceDockerCompose []byte

func RabbitmqMicroserviceDockerComposeTemplate() ([]byte, string) {
	return rabbitmqMicroserviceDockerCompose, "docker-compose.yml"
}

//go:embed files/makefile.tmpl
var rabbitmqMicroserviceMakefile []byte

func RabbitmqMicroserviceMakefileTemplate() ([]byte, string) {
	return rabbitmqMicroserviceMakefile, "Makefile"
}

//go:embed files/readme.tmpl
var rabbitmqMicroserviceReadme []byte

func RabbitmqMicroserviceReadmeTemplate() ([]byte, string) {
	return rabbitmqMicroserviceReadme, "README.md"
}

//go:embed files/internal/adapters/brokers/rabbitmq/rabbitmq.tmpl
var rabbitmqMicroserviceBroker []byte

func RabbitmqMicroserviceBrokerTemplate() ([]byte, string) {
	return rabbitmqMicroserviceBroker, "internal/adapters/brokers/rabbitmq/rabbitmq.go"
}
