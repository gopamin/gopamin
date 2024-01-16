package templates

import (
	_ "embed"
)

//go:embed files/dockerignore.template
var dockerIgnore []byte

//go:embed files/dockerfile.template
var dockerFile []byte

//go:embed files/license.template
var license []byte

//go:embed files/gitignore.template
var gitIgnore []byte

func DockerFileTemplate() ([]byte, string) {
	return dockerFile, "Dockerfile"
}

func LicenseTemplate() ([]byte, string) {
	return license, "LICENSE"
}

func GitIgnoreTemplate() ([]byte, string) {
	return gitIgnore, ".gitignore"
}

func DockerIgnoreTemplate() ([]byte, string) {
	return dockerIgnore, ".dockerignore"
}
