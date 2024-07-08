package thk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/exp/slog"
)

type RepoInfo struct {
	StargazersCount int `json:"stargazers_count"`
	ForksCount      int `json:"forks_count"`
}

type Contributor struct {
	Login         string  `json:"login"`
	Id            int     `json:"id"`
	Type          string  `json:"type"`
	Contributions int     `json:"contributions"`
	Score         float64 `json:"score,omitempty"`
}

func githubCurl(url, method string, body io.Reader, result any) error {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	if githubToken != "" {
		req.Header.Set("Authorization", "token "+githubToken)
	}
	resp, err := githubClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		slog.Error(err.Error(), "url", url, "respStatus", resp.Status)
		return err
	}
	return nil
}

func GetRepoInfo(repo string) (*RepoInfo, error) {
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s", repo)

	var repoInfo RepoInfo
	if err := githubCurl(apiURL, "GET", nil, &repoInfo); err != nil {
		return nil, err
	}

	return &repoInfo, nil
}

func GetContributors(repo string) ([]*Contributor, error) {
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/contributors?page=1&per_page=100", repo)

	var contributors []*Contributor
	if err := githubCurl(apiURL, "GET", nil, &contributors); err != nil {
		return nil, err
	}

	// filter out non-user contributors, such as "bot"
	var filteredContributors []*Contributor
	for _, c := range contributors {
		if c.Type == "User" {
			filteredContributors = append(filteredContributors, c)
		}
	}

	return filteredContributors, nil
}

func CalculateScore(repoInfo *RepoInfo, contributorsCount int) float64 {
	weightStargazers := 0.4
	weightForks := 0.2
	weightContributors := 0.2

	score := float64(repoInfo.StargazersCount)*weightStargazers +
		float64(repoInfo.ForksCount)*weightForks +
		float64(contributorsCount)*weightContributors

	if score < 1 {
		score = 1
	}

	return score
}

func CalculateContributorScore(contributors []*Contributor) {
	total := 0.0
	for _, c := range contributors {
		total += float64(c.Contributions)
	}

	for _, c := range contributors {
		c.Score = float64(c.Contributions) / total
	}
}
