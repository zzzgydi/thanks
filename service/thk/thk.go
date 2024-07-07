package thk

import (
	"log/slog"
	"sync"
)

func ThkNode(pkgJson []byte) ([]*ThkContributor, error) {
	// parse package.json
	deps, err := ParsePackageJson(pkgJson)
	if err != nil {
		return nil, err
	}

	temp := make([]*tmpThk, 0)

	var mu sync.Mutex
	var wg sync.WaitGroup
	concurrency := make(chan struct{}, 6)

	for _, dep := range deps {
		concurrency <- struct{}{}
		wg.Add(1)

		go func(dep string) {
			defer func() {
				wg.Done()
				<-concurrency
			}()

			githubUrl, err := GetGithubURLByNodeDep(dep)
			if err != nil {
				slog.Error("failed to get github url by node dependency", "dep", dep, "error", err)
				return
			}

			githubRepo, err := ParseGithubURL(githubUrl)
			if err != nil {
				slog.Error("failed to parse github url", "url", githubUrl, "error", err)
				return
			}

			score, contributors, err := ThkRepo(githubRepo)
			if err != nil {
				slog.Error("failed to thk repo", "error", err)
				return
			}

			mu.Lock()
			temp = append(temp, &tmpThk{
				repo:         githubRepo,
				score:        score,
				contributors: contributors,
			})
			mu.Unlock()
		}(dep)
	}

	wg.Wait()

	return tempToThkContributor(temp), nil
}

func ThkRepo(repo string) (float64, []*Contributor, error) {
	repoInfo, err := GetRepoInfo(repo)
	if err != nil {
		return 0, nil, err
	}

	contributors, err := GetContributors(repo)
	if err != nil {
		return 0, nil, err
	}

	score := CalculateScore(repoInfo, len(contributors))
	CalculateContributorScore(contributors)

	return score, contributors, nil
}
