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

//go:embed files/internal/adapters/mock/repository.tmpl
var mockRepository []byte

func MockRepositoryTemplate() ([]byte, string) {
	return mockRepository, "internal/adapters/mock/repository.go"
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
