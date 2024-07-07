package thk_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zzzgydi/thanks/service/thk"
)

func TestParsePackageJson(t *testing.T) {
	data := []byte(`{
		"dependencies": {
			"express": "^4.17.1",
			"koa": "^2.11.0"
		},
		"devDependencies": {
			"mocha": "^8.2.1",
			"chai": "^4.3.4"
		}
	}`)
	deps, err := thk.ParsePackageJson(data)
	require.NoError(t, err)
	assert.ElementsMatch(t, []string{"express", "koa", "mocha", "chai"}, deps)
}

func TestGetGithubURLByNodeDep(t *testing.T) {

	url, err := thk.GetGithubURLByNodeDep("koa")
	require.NoError(t, err)

	fmt.Printf("koa: %s\n", url)

	url, err = thk.GetGithubURLByNodeDep("express")
	require.NoError(t, err)

	fmt.Printf("express: %s\n", url)

	url, err = thk.GetGithubURLByNodeDep("mocha")
	require.NoError(t, err)

	fmt.Printf("mocha: %s\n", url)

	url, err = thk.GetGithubURLByNodeDep("chai")
	require.NoError(t, err)

	fmt.Printf("chai: %s\n", url)
}

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
		repo, err := thk.ParseGithubURL(url)
		require.NoError(t, err)
		assert.Equal(t, outputs[i], repo)
		fmt.Printf("repo: %s\n", repo)
	}
}
