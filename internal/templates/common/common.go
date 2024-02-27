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

//go:embed files/cmd/server/server.tmpl
var server []byte

func ServerTemplate() ([]byte, string) {
	return server, "cmd/server/server.go"
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

//go:embed files/cmd/main/web-app/main-with-db.tmpl
var webAppMainWithDb []byte

func WebAppMainWithDbTemplate() ([]byte, string) {
	return webAppMainWithDb, "cmd/main.go"
}

//go:embed files/cmd/main/hello-world/main.tmpl
var helloWorldMain []byte

func HelloWorldMainTemplate() ([]byte, string) {
	return helloWorldMain, "cmd/main.go"
}

//go:embed files/cmd/main/hello-world/main_test.tmpl
var helloWorldMainTest []byte

func HelloWorldMainTestTemplate() ([]byte, string) {
	return helloWorldMainTest, "cmd/main_test.go"
}

//go:embed files/cmd/main/hello-world/main-with-db.tmpl
var helloWorldMainWithDb []byte

func HelloWorldMainWithDbTemplate() ([]byte, string) {
	return helloWorldMainWithDb, "cmd/main.go"
}

//go:embed files/assets/users.tmpl
var webAppUsersHtmlTemplate []byte

func WebAppUsersHtmlTemplateTemplate() ([]byte, string) {
	return webAppUsersHtmlTemplate, "assets/users.html"
}

//go:embed files/assets/user.tmpl
var webAppUserHtmlTemplate []byte

func WebAppUserHtmlTemplateTemplate() ([]byte, string) {
	return webAppUserHtmlTemplate, "assets/user.html"
}

//go:embed files/assets/styles.tmpl
var webAppStyles []byte

func WebAppStylesTemplate() ([]byte, string) {
	return webAppStyles, "assets/styles.css"
}

//go:embed files/cmd/main/web-app/main.tmpl
var webAppMain []byte

func WebAppMainTemplate() ([]byte, string) {
	return webAppMain, "cmd/main.go"
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

//go:embed files/cmd/main/microservice/main.tmpl
var microserviceMain []byte

func MicroserviceMainTemplate() ([]byte, string) {
	return microserviceMain, "cmd/main.go"
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
