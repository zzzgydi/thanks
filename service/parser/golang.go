package parser

import (
	"regexp"

	"golang.org/x/mod/modfile"
)

var _ IParser = (*GoParser)(nil)

type GoParser struct{}

func NewGoParser() *GoParser {
	return &GoParser{}
}

func (g GoParser) GetGitRepos(data []byte) ([]string, error) {
	deps, err := parseGoMod(data)
	if err != nil {
		return nil, err
	}

	repos := make(map[string]struct{})
	for _, dep := range deps {
		if repo, err := ParseGithubRepo(dep); err == nil {
			repos[repo] = struct{}{}
		}
	}

	var result []string
	for repo := range repos {
		result = append(result, repo)
	}

	return result, nil
}

var (
	goModRe = regexp.MustCompile(`github\.com/[a-zA-Z0-9._-]+/[a-zA-Z0-9._-]+`)
)

func parseGoMod(data []byte) ([]string, error) {

	modFile, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		return nil, err
	}

	var deps []string

	for _, require := range modFile.Require {
		matches := goModRe.FindAllString(require.Mod.Path, -1)
		for _, match := range matches {
			deps = append(deps, "https://"+match)
		}
	}

	for _, replace := range modFile.Replace {
		matches := goModRe.FindAllString(replace.New.Path, -1)
		for _, match := range matches {
			deps = append(deps, "https://"+match)
		}
		matches = goModRe.FindAllString(replace.Old.Path, -1)
		for _, match := range matches {
			deps = append(deps, "https://"+match)
		}
	}

	return deps, nil
}
