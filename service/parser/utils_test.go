package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseGithubUrl(t *testing.T) {
	urls := []string{
		"https://github.com/remarkjs/remark-math.git#main",
		"https://github.com/expressjs/express.git",
		"https://github.com/expressjs/express",
		"http://github.com/expressjs/express",
	}

	outputs := []string{
		"remarkjs/remark-math",
		"expressjs/express",
		"expressjs/express",
		"expressjs/express",
	}

	for i, url := range urls {
		repo, err := ParseGithubRepo(url)
		require.NoError(t, err)
		assert.Equal(t, outputs[i], repo)
		fmt.Printf("repo: %s\n", repo)
	}
}
