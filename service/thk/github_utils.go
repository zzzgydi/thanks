package thk

import (
	"encoding/json"
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
