package templates

import (
	_ "embed"
)

//go:embed files/dockerignore.tmpl
var dockerIgnore []byte

func DockerIgnoreTemplate() ([]byte, string) {
	return dockerIgnore, ".dockerignore"
}

//go:embed files/dockerfile.tmpl
var dockerFile []byte

func DockerFileTemplate() ([]byte, string) {
	return dockerFile, "Dockerfile"
}

//go:embed files/license.tmpl
var license []byte

func LicenseTemplate() ([]byte, string) {
	return license, "LICENSE"
}

//go:embed files/gitignore.tmpl
var gitIgnore []byte

func GitIgnoreTemplate() ([]byte, string) {
	return gitIgnore, ".gitignore"
}

//go:embed files/internal/adapters/repositories/mock/repository.tmpl
var mockRepository []byte

func MockRepositoryTemplate() ([]byte, string) {
	return mockRepository, "internal/adapters/repositories/mock/repository.go"
}

//go:embed files/internal/core/domain/user_test.tmpl
var userTest []byte

func UserTestTemplate() ([]byte, string) {
	return userTest, "internal/core/domain/user_test.go"
}

//go:embed files/internal/core/domain/user.tmpl
var user []byte

func UserTemplate() ([]byte, string) {
	return user, "internal/core/domain/user.go"
}

//go:embed files/internal/core/ports/user-repository.interface.tmpl
var userRepositoryInterface []byte

func UserRepositoryInterfaceTemplate() ([]byte, string) {
	return userRepositoryInterface, "internal/core/ports/user-repository.interface.go"
}

//go:embed files/internal/core/ports/user-service.interface.tmpl
var userServiceInterface []byte

func UserServiceInterfaceTemplate() ([]byte, string) {
	return userServiceInterface, "internal/core/ports/user-service.interface.go"
}

//go:embed files/internal/core/services/user.service.tmpl
var userService []byte

func UserServiceTemplate() ([]byte, string) {
	return userService, "internal/core/services/user.service.go"
}

//go:embed files/internal/core/services/user.service_test.tmpl
var userServiceTest []byte

func UserServiceTestTemplate() ([]byte, string) {
	return userServiceTest, "internal/core/services/user.service_test.go"
}

//go:embed files/internal/core/ports/router.interface.tmpl
var routerInferface []byte

func RouterInterfaceTemplate() ([]byte, string) {
	return routerInferface, "internal/core/ports/router.interface.go"
}

//go:embed files/internal/core/api/errors.tmpl
var apiErrors []byte

func ApiErrorsTemplate() ([]byte, string) {
	return apiErrors, "internal/core/api/errors.go"
}

//go:embed files/internal/core/api/response.tmpl
var apiResponse []byte

func ApiResponseTemplate() ([]byte, string) {
	return apiResponse, "internal/core/api/response.go"
}

//go:embed files/configs/load-env.tmpl
var loadEnv []byte

func LoadEnvTemplate() ([]byte, string) {
	return loadEnv, "configs/load-env.go"
}

//go:embed files/cmd/server/server.tmpl
var apiServer []byte

func ApiServerTemplate() ([]byte, string) {
	return apiServer, "cmd/server/server.go"
}

//go:embed files/cmd/main/api/main.tmpl
var apiMain []byte

func ApiMainTemplate() ([]byte, string) {
	return apiMain, "cmd/main.go"
}

//go:embed files/cmd/main/api/main-with-db.tmpl
var apiMainWithDb []byte

func ApiMainWithDbTemplate() ([]byte, string) {
	return apiMainWithDb, "cmd/main.go"
}
