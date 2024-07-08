package model

import (
	"time"

	"github.com/zzzgydi/thanks/common"
)

type GitRepo struct {
	Id               uint64    `json:"id" db:"id" gorm:"primary_key;autoIncrement"`
	Repo             string    `json:"repo" db:"repo" gorm:"uniqueIndex"`
	StarCount        int       `json:"star_count" db:"star_count"`
	ForkCount        int       `json:"fork_count" db:"fork_count"`
	ContributorCount int       `json:"contributor_count" db:"contributor_count"`
	UpdatedAt        time.Time `json:"update_at" db:"update_at"`
}

func (g *GitRepo) TableName() string {
	return "git_repo"
}

func CreateGitRepo(gitRepo *GitRepo) error {
	return common.MDB.Create(gitRepo).Error
}

func GetGitRepoByRepoName(repo string) (*GitRepo, error) {
	var gitRepo GitRepo
	if err := common.MDB.Where("repo = ?", repo).First(&gitRepo).Error; err != nil {
		return nil, err
	}
	return &gitRepo, nil
}
