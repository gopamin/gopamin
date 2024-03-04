package configs

import (
	"os"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	t.Run("read environment variables", func(t *testing.T) {
		LoadEnv()

		projectName := os.Getenv("PROJECT_NAME")
		if projectName != "allah" {
			t.Fatalf("expected allah but got %v", projectName)
		}
	})
}


