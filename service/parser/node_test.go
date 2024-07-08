package parser

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zzzgydi/thanks/common/config"
	"github.com/zzzgydi/thanks/common/initializer"
	"github.com/zzzgydi/thanks/common/logger"
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
	deps, err := parsePackageJson(data)
	require.NoError(t, err)
	assert.ElementsMatch(t, []string{"express", "koa", "mocha", "chai"}, deps)
}

func TestGetGithubURLByNodeDep(t *testing.T) {

	deps := []string{"koa", "express", "mocha", "chai"}

	for _, dep := range deps {
		url, err := getGithubURLByNodeDep(dep)
		require.NoError(t, err)
		fmt.Printf("%s: %s\n", dep, url)
	}
}

func TestNodeParser(t *testing.T) {
	viper.AutomaticEnv()
	logger.InitLogger(viper.GetString("TEST_ROOT_DIR"))
	config.InitConfig()
	initializer.InitInitializer()

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

	parser := NewNodeParser(6)
	require.NotNil(t, parser)

	repos, err := parser.GetGitRepos(data)
	require.NoError(t, err)

	for _, repo := range repos {
		fmt.Println(repo)
	}

	assert.Equal(t, 4, len(repos))
}
