package parser

import (
	"fmt"
	"net/url"
	"strings"
)

// url to github repo
func ParseGithubRepo(input string) (string, error) {
	parsedURL, err := url.Parse(input)
	if err != nil {
		return "", err
	}

	path := parsedURL.Path
	path = strings.TrimSuffix(path, ".git")

	if idx := strings.Index(path, "#"); idx != -1 {
		path = path[:idx]
	}

	if !strings.HasPrefix(path, "/") {
		return "", fmt.Errorf("invalid URL format")
	}

	path = path[1:]

	return path, nil
}
