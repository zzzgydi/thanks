package thk

import (
	"fmt"
	"strings"

	"github.com/zzzgydi/thanks/model"
	"github.com/zzzgydi/thanks/service/parser"
)

type ThankTask struct {
	lang   string
	data   []byte
	task   *model.Task
	parser parser.IParser
	repos  []string
}

func NewThankTask(lang string, data []byte) (*ThankTask, error) {
	tt := &ThankTask{
		lang: lang,
		data: data,
	}

	if err := tt.prepare(); err != nil {
		return nil, err
	}

	task, err := model.CreateTask(lang)
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
	}

	contributors, err := Thanks(tt.repos, 20)
	if err != nil {
		return nil, err
	}

	return contributors, nil
}
