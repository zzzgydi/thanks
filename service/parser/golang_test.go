package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGolangParser(t *testing.T) {
	data := []byte(`module github.com/zzzgydi/thanks

go 1.16

require (
	github.com/spf13/viper v1.7.1
	github.com/zzzgydi/slog-gorm v0.0.0-20240401110836-ab07ad30a474
	golang.org/x/mod v0.4.2
	github.com/bytedance/sonic v1.11.6 // indirect
	github.com/bytedance/sonic/loader v0.1.1 // indirect
)

replace github.com/spf13/viper => github.com/spf13/viper v1.7.1
replace golang.org/x/mod => golang.org/x/mod v0.4.2

`)

	parser := NewGoParser()
	require.NotNil(t, parser)

	repos, err := parser.GetGitRepos(data)
	require.NoError(t, err)

	for _, repo := range repos {
		fmt.Println(repo)
	}

	require.Equal(t, 3, len(repos))
}
