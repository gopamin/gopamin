package templates

import (
	_ "embed"
)

//go:embed files/env.tmpl
var kafkaMicroserviceEnv []byte

func KafkaMicroserviceEnvTemplate() ([]byte, string) {
	return kafkaMicroserviceEnv, ".env"
}

//go:embed files/docker-compose.tmpl
var kafkaMicroserviceDockerCompose []byte

func KafkaMicroserviceDockerComposeTemplate() ([]byte, string) {
	return kafkaMicroserviceDockerCompose, "docker-compose.yml"
}

//go:embed files/makefile.tmpl
var kafkaMicroserviceMakefile []byte

func KafkaMicroserviceMakefileTemplate() ([]byte, string) {
	return kafkaMicroserviceMakefile, "Makefile"
}

//go:embed files/readme.tmpl
var kafkaMicroserviceReadme []byte

func KafkaMicroserviceReadmeTemplate() ([]byte, string) {
	return kafkaMicroserviceReadme, "README.md"
}

//go:embed files/internal/adapters/brokers/kafka/kafka.tmpl
var kafkaMicroserviceBroker []byte

func KafkaMicroserviceBrokerTemplate() ([]byte, string) {
	return kafkaMicroserviceBroker, "internal/adapters/brokers/kafka/kafka.go"
}
