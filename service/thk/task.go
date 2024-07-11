package thk

import (
	"fmt"
	"strings"

	"github.com/zzzgydi/thanks/model"
	"github.com/zzzgydi/thanks/service/parser"
)

type ThankTask struct {
	lang     string
	minScore float64
	data     []byte
	task     *model.Task
	parser   parser.IParser
	repos    []string
}

func NewThankTask(lang string, minScore float64, data []byte) (*ThankTask, error) {
	tt := &ThankTask{
		lang:     lang,
		minScore: minScore,
		data:     data,
	}

	if err := tt.prepare(); err != nil {
		return nil, err
	}

	task, err := model.CreateTask(lang, minScore)
	if err != nil {
		return nil, err
	}

	tt.task = task
	return tt, nil
}

func NewThankTaskFromId(id string) (*ThankTask, error) {
	task, err := model.GetTaskById(id)
	if err != nil {
		return nil, err
	}

	repos := strings.Split(task.Repos, ";")

	tt := &ThankTask{
		lang:  task.Lang,
		task:  task,
		repos: repos,
	}

	return tt, nil
}

func (tt *ThankTask) prepare() error {
	lang := tt.lang
	if lang == "golang" {
		tt.parser = parser.NewGoParser()
	} else if lang == "node" {
		tt.parser = parser.NewNodeParser(20)
	} else {
		return fmt.Errorf("unsupported language: %s", lang)
	}
	return nil
}

func (tt *ThankTask) Task() *model.Task {
	return tt.task
}

func (tt *ThankTask) Run() ([]*ThkContributor, error) {
	if tt.repos == nil && tt.parser != nil && tt.data != nil {
		repos, err := tt.parser.GetGitRepos(tt.data)
		if err != nil {
			return nil, err
		}

		tt.repos = repos
		tt.task.Repos = strings.Join(repos, ";")

		err = model.SaveTask(tt.task)
		if err != nil {
			return nil, err
		}

		return Thanks(tt.repos, 20)
	} else if tt.repos != nil {
		return ThanksReadOnly(tt.repos)
	}

	return nil, fmt.Errorf("unhandled case")

	// minScore := tt.minScore
	// if minScore < 0.0001 {
	// 	minScore = 0.0001
	// }
	// if minScore > 0.1 {
	// 	minScore = 0.1
	// }

	// ret := make([]*ThkContributor, 0, len(contributors))
	// for _, c := range contributors {
	// 	if c.Total >= minScore {
	// 		ret = append(ret, c)
	// 	}
	// }

	// return ret, nil
}
