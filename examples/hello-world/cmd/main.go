package main

import (
	"fmt"
	"os"

	"hello-world/configs"
)

func init() {
	configs.LoadEnv()
}

func main() {
	projectName := os.Getenv("PROJECT_NAME")
	fmt.Println(projectName)
}
