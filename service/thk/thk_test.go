package thk_test

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"github.com/zzzgydi/thanks/service/parser"
	"github.com/zzzgydi/thanks/service/thk"
)

func TestThk(t *testing.T) {
	viper.AutomaticEnv()

	if err := thk.InitTest(); err != nil {
		t.Error("Failed to initialize thk:", err)
		return
	}

	pkJsonUrl := "https://github.com/zzzgydi/ailiuliu/raw/main/web/package.json"

	fmt.Println(pkJsonUrl)

	// http get
	resp, err := http.Get(pkJsonUrl)
	require.NoError(t, err)

	defer resp.Body.Close()

	// read body
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	nodeParser := parser.NewNodeParser(6)
	require.NotNil(t, nodeParser)

	repos, err := nodeParser.GetGitRepos(body)
	require.NoError(t, err)

	contributors, err := thk.Thanks(repos, 6)
	require.NoError(t, err)

	for i, c := range contributors {
		fmt.Printf("%-5d %-*s %-*s %.4f\n", i, 20, c.Login, 20, c.Repos[0].Repo, c.Total)
	}
}
