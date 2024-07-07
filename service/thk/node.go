package thk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type PackageJSON struct {
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

func ParsePackageJson(data []byte) ([]string, error) {
	var pj PackageJSON
	if err := json.Unmarshal(data, &pj); err != nil {
		return nil, err
	}

	var deps []string
	for dep := range pj.Dependencies {
		deps = append(deps, dep)
	}
	for dep := range pj.DevDependencies {
		deps = append(deps, dep)
	}

	return deps, nil
}

type Repository struct {
	URL string `json:"url"`
}

type PackageInfo struct {
	Repository Repository `json:"repository"`
}

func GetGithubURLByNodeDep(packageName string) (string, error) {
	url := fmt.Sprintf("https://registry.npmjs.org/%s", packageName)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get package info: %s", resp.Status)
	}

	var pkgInfo PackageInfo
	if err := json.NewDecoder(resp.Body).Decode(&pkgInfo); err != nil {
		return "", err
	}

	repoUrl := pkgInfo.Repository.URL

	if repoUrl == "" {
		return "", fmt.Errorf("no repository found")
	}

	if !strings.Contains(repoUrl, "github.com") {
		return "", fmt.Errorf("not a github repository")
	}

	repoUrl = strings.TrimSuffix(repoUrl, ".git")
	repoUrl = strings.TrimPrefix(repoUrl, "git+")

	return repoUrl, nil
}

func ParseGithubURL(input string) (string, error) {
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
