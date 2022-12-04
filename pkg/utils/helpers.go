package utils

import "strings"

func GetRepositoryName(repo string) string {
	// split repo string by "/"
	// and return last element
	return repo[strings.LastIndex(repo, "/")+1:]
}
