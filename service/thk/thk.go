package thk

import (
	"log/slog"
	"sync"

	"github.com/zzzgydi/thanks/model"
)

func Thanks(repos []string, maxConcurrency uint) ([]*ThkContributor, error) {
	temp := make([]*tmpThk, 0)

	var mu sync.Mutex
	var wg sync.WaitGroup
	concurrency := make(chan struct{}, maxConcurrency)

	for _, repo := range repos {
		concurrency <- struct{}{}
		wg.Add(1)

		go func(repo string) {
			defer func() {
				wg.Done()
				<-concurrency
			}()

			score, contributors, err := ThkRepo(repo)
			if err != nil {
				slog.Error("failed to thk repo", "error", err)
				return
			}

			mu.Lock()
			temp = append(temp, &tmpThk{
				repo:         repo,
				score:        score,
				contributors: contributors,
			})
			mu.Unlock()
		}(repo)
	}

	wg.Wait()

	return tempToThkContributor(temp), nil
}

func ThkRepo(repo string) (float64, []*model.GitContributor, error) {
	git := NewGithubThk(repo)

	if err := git.Run(); err != nil {
		return 0, nil, err
	}

	return git.RepoScore(), git.GetContributors(), nil
}
