package templates

import (
	_ "embed"
)

//go:embed files/env.tmpl
var env []byte

func EnvTemplate() ([]byte, string) {
	return env, ".env"
}

//go:embed files/makefile.tmpl
var makefile []byte

func MakefileTemplate() ([]byte, string) {
	return makefile, "Makefile"
}

//go:embed files/readme.tmpl
var readme []byte

func ReadmeTemplate() ([]byte, string) {
	return readme, "README.md"
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

//go:embed files/configs/configs.tmpl
var configs []byte

func ConfigsTemplate() ([]byte, string) {
	return configs, "configs/configs.go"
}

//go:embed files/configs/configs_test.tmpl
var configsTest []byte

func ConfigsTestTemplate() ([]byte, string) {
	return configsTest, "configs/configs_test.go"
}

//go:embed files/internal/core/domain/message.tmpl
var message []byte

func MessageTemplate() ([]byte, string) {
	return message, "internal/core/domain/message.go"
}

//go:embed files/internal/core/ports/broker-service.interface.tmpl
var brokerServiceInterface []byte

func BrokerServiceInterfaceTemplate() ([]byte, string) {
	return brokerServiceInterface, "internal/core/ports/broker-service.interface.go"
}

//go:embed files/internal/core/ports/message-broker.interface.tmpl
var messageBrokerInterface []byte

func MessageBrokerInterfaceTemplate() ([]byte, string) {
	return messageBrokerInterface, "internal/core/ports/message-broker.interface.go"
}

//go:embed files/internal/core/services/broker.service.tmpl
var brokerService []byte

func BrokerServiceTemplate() ([]byte, string) {
	return brokerService, "internal/core/services/broker.service.go"
}

//go:embed files/tools/tools.tmpl
var tools []byte

func ToolsTemplate() ([]byte, string) {
	return tools, "tools/tools.go"
}

//go:embed files/tools/tools_test.tmpl
var toolsTest []byte

func ToolsTestTemplate() ([]byte, string) {
	return toolsTest, "tools/tools_test.go"
}
