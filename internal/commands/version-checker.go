package commands

import (
	"io"
	"net/http"
	"strings"
)

const ENV_URL = "https://raw.githubusercontent.com/gopamin/gopamin/master/.env"

func versionChecker() (bool, string) {
	res, err := http.Get(ENV_URL)
	if err != nil {
		return true, ""
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return true, ""
	}
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return true, ""
	}

	newVersion := extractVersion(string(bytes))
	if VERSION != newVersion {
		return false, newVersion
	}

	return true, ""
}

func extractVersion(input string) string {
	parts := strings.Split(input, "=")
	if len(parts) >= 2 {
		return strings.TrimSpace(parts[1])
	}

	return ""
}
