package utils

import (
	"os"
	"strings"
)

func GetRepositoryName(repo string) string {
	// split repo string by "/"
	// and return last element
	return repo[strings.LastIndex(repo, "/")+1:]
}

func GetEnvVar(key string, def string) string {
	// get environment variable
	// and return value
	ev := os.Getenv(key)
	if ev == "" {
		return def
	}
	return ev
}
