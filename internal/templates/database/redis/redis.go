package templates

import (
	_ "embed"
)

//go:embed files/internal/databases/redis.tmpl
var redisDatabase []byte

//go:embed files/env.tmpl
var redisEnv []byte

//go:embed files/makefile.tmpl
var redisMakefile []byte

//go:embed files/readme.tmpl
var redisReadme []byte

//go:embed files/docker-compose.tmpl
var redisDockerCompose []byte

//go:embed files/cmd/main.tmpl
var redisMain []byte

func RedisMainTemplate() ([]byte, string) {
	return redisMain, "cmd/main.go"
}

func RedisDockerComposeTemplate() ([]byte, string) {
	return redisDockerCompose, "docker-compose.yml"
}

func RedisDatabaseTemplate() ([]byte, string) {
	return redisDatabase, "internal/databases/redis.go"
}

func RedisEnvTemplate() ([]byte, string) {
	return redisEnv, ".env"
}

func RedisMakefileTemplate() ([]byte, string) {
	return redisMakefile, "Makefile"
}

func RedisReadmeTemplate() ([]byte, string) {
	return redisReadme, "README.md"
}
