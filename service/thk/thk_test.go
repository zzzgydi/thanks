package thk_test

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/spf13/viper"
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
	if err != nil {
		t.Error("Failed to get package.json:", err)
		return
	}

	defer resp.Body.Close()

	// read body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error("Failed to read body:", err)
		return
	}

	contributors, err := thk.ThkNode(body)
	if err != nil {
		t.Error("Failed to thk node:", err)
		return
	}

	for i, c := range contributors {
		fmt.Printf("%-5d %-*s %-*s %.4f\n", i, 20, c.Login, 20, c.Repos[0].Repo, c.Total)
	}
}
