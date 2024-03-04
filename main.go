package main

import (
	"github.com/gopamin/gopamin/configs"
	"github.com/gopamin/gopamin/internal/commands"
)

func init() {
	configs.LoadEnv()
}

func main() {
	commands.Execute()
}
