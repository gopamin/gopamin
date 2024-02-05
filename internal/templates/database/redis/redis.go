package templates

import (
	_ "embed"
)

//go:embed files/env.tmpl
var redisEnv []byte

func RedisEnvTemplate() ([]byte, string) {
	return redisEnv, ".env"
}

//go:embed files/makefile.tmpl
var redisMakefile []byte

func RedisMakefileTemplate() ([]byte, string) {
	return redisMakefile, "Makefile"
}

//go:embed files/readme.tmpl
var redisReadme []byte

func RedisReadmeTemplate() ([]byte, string) {
	return redisReadme, "README.md"
}

//go:embed files/docker-compose.tmpl
var redisDockerCompose []byte

func RedisDockerComposeTemplate() ([]byte, string) {
	return redisDockerCompose, "docker-compose.yml"
}

//go:embed files/cmd/main.tmpl
var redisMain []byte

func RedisMainTemplate() ([]byte, string) {
	return redisMain, "cmd/main.go"
}

//go:embed files/internal/adapters/redis/repository.tmpl
var redisRepository []byte

func RedisRepositoryTemplate() ([]byte, string) {
	return redisRepository, "internal/adapters/redis/repository.go"
}
