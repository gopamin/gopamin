package templates

import (
	_ "embed"
)

//go:embed files/env.tmpl
var redisMicroserviceEnv []byte

func RedisMicroserviceEnvTemplate() ([]byte, string) {
	return redisMicroserviceEnv, ".env"
}

//go:embed files/docker-compose.tmpl
var redisMicroserviceDockerCompose []byte

func RedisMicroserviceDockerComposeTemplate() ([]byte, string) {
	return redisMicroserviceDockerCompose, "docker-compose.yml"
}

//go:embed files/makefile.tmpl
var redisMicroserviceMakefile []byte

func RedisMicroserviceMakefileTemplate() ([]byte, string) {
	return redisMicroserviceMakefile, "Makefile"
}

//go:embed files/readme.tmpl
var redisMicroserviceReadme []byte

func RedisMicroserviceReadmeTemplate() ([]byte, string) {
	return redisMicroserviceReadme, "README.md"
}

//go:embed files/internal/adapters/brokers/redis/redis.tmpl
var redisMicroserviceBroker []byte

func RedisMicroserviceBrokerTemplate() ([]byte, string) {
	return redisMicroserviceBroker, "internal/adapters/brokers/redis/redis.go"
}
