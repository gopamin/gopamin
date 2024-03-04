package configs

import (
	"github.com/gopamin/gopamin/tools"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	rootDir, err := tools.FindRootDir()
	if err != nil {
		panic(err)
	}
	envFilePath := filepath.Join(rootDir, ".env")
	godotenv.Load(envFilePath)
}
