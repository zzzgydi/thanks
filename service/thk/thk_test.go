package thk_test

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"github.com/zzzgydi/thanks/common/config"
	"github.com/zzzgydi/thanks/common/initializer"
	"github.com/zzzgydi/thanks/common/logger"
	"github.com/zzzgydi/thanks/service/parser"
	"github.com/zzzgydi/thanks/service/thk"
)

func TestThkNode(t *testing.T) {
	startTime := time.Now().Unix()
	defer func() {
		endTime := time.Now().Unix()
		fmt.Printf("----------------\nCost Time: %d\n", endTime-startTime)
	}()

	viper.AutomaticEnv()
	logger.InitLogger(viper.GetString("TEST_ROOT_DIR"))
	config.InitConfig()
	initializer.InitInitializer()

	pkJsonUrl := "https://github.com/zzzgydi/ailiuliu/raw/main/web/package.json"

	// http get
	resp, err := http.Get(pkJsonUrl)
	require.NoError(t, err)

	defer resp.Body.Close()

	// read body
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	nodeParser := parser.NewNodeParser(20)
	require.NotNil(t, nodeParser)

	repos, err := nodeParser.GetGitRepos(body)
	require.NoError(t, err)

	contributors, err := thk.Thanks(repos, 20)
	require.NoError(t, err)

	for i, c := range contributors {
		repos := make([]string, 0, len(c.Repos))
		for _, r := range c.Repos {
			repos = append(repos, r.Repo)
		}

		repo := strings.Join(repos, ", ")

		fmt.Printf("%-5d %-*s %.4f%%\t%s\n", i, 30, c.Login, c.Total*100, repo)
	}
}

func TestThkGolang(t *testing.T) {
	startTime := time.Now().Unix()
	defer func() {
		endTime := time.Now().Unix()
		fmt.Printf("----------------\nCost Time: %d\n", endTime-startTime)
	}()

	viper.AutomaticEnv()
	logger.InitLogger(viper.GetString("TEST_ROOT_DIR"))
	config.InitConfig()
	initializer.InitInitializer()

	modUrl := "https://github.com/zzzgydi/thanks/raw/main/go.mod"

	// http get
	resp, err := http.Get(modUrl)
	require.NoError(t, err)

	defer resp.Body.Close()

	// read body
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	nodeParser := parser.NewGoParser()
	require.NotNil(t, nodeParser)

	repos, err := nodeParser.GetGitRepos(body)
	require.NoError(t, err)

	contributors, err := thk.Thanks(repos, 20)
	require.NoError(t, err)

	for i, c := range contributors {
		repos := make([]string, 0, len(c.Repos))
		for _, r := range c.Repos {
			repos = append(repos, r.Repo)
		}

		repo := strings.Join(repos, ", ")

		fmt.Printf("%-5d %-*s %.4f%%\t%s\n", i, 30, c.Login, c.Total*100, repo)
	}
}
