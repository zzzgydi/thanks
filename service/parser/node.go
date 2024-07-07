package parser

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

var _ IParser = (*NodeParser)(nil)

type NodeParser struct {
	concurrency uint
}

func NewNodeParser(concurrency uint) *NodeParser {
	return &NodeParser{concurrency: concurrency}
}

func (l *NodeParser) GetGitRepos(data []byte) ([]string, error) {

	deps, err := parsePackageJson(data)
	if err != nil {
		return nil, err
	}

	repos := make(map[string]struct{})

	var mu sync.Mutex
	var wg sync.WaitGroup
	concurrency := make(chan struct{}, l.concurrency)

	for _, dep := range deps {
		concurrency <- struct{}{}
		wg.Add(1)

		go func(dep string) {
			defer func() {
				wg.Done()
				<-concurrency
			}()

			repo, err := getGithubURLByNodeDep(dep)
			if err != nil {
				return
			}

			mu.Lock()
			repos[repo] = struct{}{}
			mu.Unlock()
		}(dep)
	}

	wg.Wait()

	var result []string

	for repo := range repos {
		result = append(result, repo)
	}

	return result, nil
}

type packageJSON struct {
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

type nodeRepository struct {
	URL string `json:"url"`
}

type nodePackageInfo struct {
	Repository nodeRepository `json:"repository"`
}

func parsePackageJson(data []byte) ([]string, error) {
	var pj packageJSON
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

func getGithubURLByNodeDep(packageName string) (string, error) {
	url := fmt.Sprintf("https://registry.npmjs.org/%s", packageName)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get package info: %s", resp.Status)
	}

	var pkgInfo nodePackageInfo
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

	return ParseGithubRepo(repoUrl)
}
