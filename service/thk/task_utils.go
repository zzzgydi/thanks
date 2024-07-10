package thk

import "io"

func LoadDataFromUrl(url string) ([]byte, error) {
	resp, err := githubClient.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
