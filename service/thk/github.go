package thk

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/zzzgydi/thanks/model"
	"gorm.io/gorm"
)

type GithubThk struct {
	repo         string
	gitRepo      *model.GitRepo
	contributors []*model.GitContributor
}

func NewGithubThk(repo string) *GithubThk {
	return &GithubThk{repo: repo}
}

func (g *GithubThk) GetRepoInfo() *model.GitRepo {
	return g.gitRepo
}

func (g *GithubThk) GetContributors() []*model.GitContributor {
	return g.contributors
}

func (g *GithubThk) Run() error {
	// check cache first [from database]
	gr, err := model.GetGitRepoByRepoName(g.repo)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	g.gitRepo = gr

	if gr == nil {
		if err := g.fetchRepoInfo(); err != nil {
			return err
		}
		if err := g.fetchContributors(); err != nil {
			return err
		}
		// save to database
		err := model.CreateGitRepo(g.gitRepo)
		if err != nil {
			slog.Error("create git repo error", "repo", g.repo, "err", err)
			// return err // should not return here
		} else {
			for _, c := range g.contributors {
				c.RepoId = g.gitRepo.Id
			}

			if len(g.contributors) > 0 {
				err := model.CreateGitContributors(g.contributors)
				if err != nil {
					slog.Error("create git contributors error", "repo", g.repo, "err", err)
				}
			} else {
				slog.Warn("no contributors", "repo", g.repo)
			}
		}
	} else {
		// fetch contributors
		temp, err := model.GetGitContributorsByRepoId(gr.Id)
		if err != nil {
			return err
		}
		g.contributors = temp
	}

	// expired after 7 days
	if g.gitRepo.UpdatedAt.Before(time.Now().Add(-time.Hour * 24 * 7)) {
		// TODO
	}

	return nil
}

func (g *GithubThk) fetchRepoInfo() error {
	if g.gitRepo != nil {
		return nil
	}
	slog.Info("fetch repo info", "repo", g.repo)
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s", g.repo)

	var repoInfo RepoInfo
	if err := githubCurl(apiURL, "GET", nil, &repoInfo); err != nil {
		return err
	}

	g.gitRepo = &model.GitRepo{
		Repo:             g.repo,
		StarCount:        repoInfo.StargazersCount,
		ForkCount:        repoInfo.ForksCount,
		ContributorCount: 0,
		UpdatedAt:        time.Now(),
	}

	if g.contributors != nil {
		g.gitRepo.ContributorCount = len(g.contributors)
	}

	return nil
}

func (g *GithubThk) fetchContributors() error {
	if g.contributors != nil {
		return nil
	}
	slog.Info("fetch contributors", "repo", g.repo)
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/contributors?page=1&per_page=100", g.repo)

	var contributors []*Contributor
	if err := githubCurl(apiURL, "GET", nil, &contributors); err != nil {
		return err
	}

	// filter out non-user contributors, such as "bot"
	var filteredContributors []*model.GitContributor
	for _, c := range contributors {
		if c.Type == "User" {
			filteredContributors = append(filteredContributors, &model.GitContributor{
				Login:         c.Login,
				GitId:         c.Id,
				Contributions: c.Contributions,
			})
		}
	}

	// calculate total contributions and each contributor's score
	total := 0.0
	for _, c := range filteredContributors {
		total += float64(c.Contributions)
	}

	for _, c := range filteredContributors {
		c.Score = float64(c.Contributions) / total
	}

	g.contributors = filteredContributors
	if g.gitRepo != nil {
		g.gitRepo.ContributorCount = len(filteredContributors)
	}

	return nil
}

func (g *GithubThk) RepoScore() float64 {
	repoInfo := g.GetRepoInfo()

	contributorsCount := 0
	if g.contributors != nil {
		contributorsCount = len(g.contributors)
	}

	score := float64(repoInfo.StarCount)*weightStars +
		float64(repoInfo.ForkCount)*weightForks +
		float64(contributorsCount)*weightContributors

	if score < minRepoScore {
		score = minRepoScore
	}
	return score
}
