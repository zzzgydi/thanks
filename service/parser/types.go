package parser

type IParser interface {
	GetGitRepos(data []byte) ([]string, error)
}
