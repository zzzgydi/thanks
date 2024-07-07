package thk

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/viper"
	"github.com/zzzgydi/thanks/common/initializer"
)

var (
	githubToken  string
	githubClient *http.Client
)

func initThk() error {
	githubToken = viper.GetString("GITHUB_TOKEN")
	if githubToken == "" {
		return fmt.Errorf("GITHUB_TOKEN is not set")
	}

	githubClient = &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}

	return nil
}

func init() {
	initializer.Register("thk", initThk)
}

func InitTest() error {
	return initThk()
}
