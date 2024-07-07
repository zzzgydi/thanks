package thk_test

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/zzzgydi/thanks/service/thk"
)

func TestCalculateScore(t *testing.T) {
	viper.AutomaticEnv()

	if err := thk.InitTest(); err != nil {
		t.Error("Failed to initialize thk:", err)
		return
	}

	repo := "expressjs/express"
	// repo := "zzzgydi/thanks"
	repoInfo, err := thk.GetRepoInfo(repo)
	if err != nil {
		t.Error("Failed to get repository information:", err)
		return
	}

	contributors, err := thk.GetContributors(repo)
	if err != nil {
		t.Error("Failed to get contributors:", err)
		return
	}

	score := thk.CalculateScore(repoInfo, len(contributors))

	fmt.Printf("The repository score is %.2f\n", score)

	thk.CalculateContributorScore(contributors)

	for i, c := range contributors {
		fmt.Printf("%-5d %-*s %.2f\n", i, 20, c.Login, c.Score*score)
	}
}
