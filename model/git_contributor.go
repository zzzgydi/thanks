package model

import "github.com/zzzgydi/thanks/common"

type GitContributor struct {
	Id            uint64  `json:"id" db:"id" gorm:"primary_key;autoIncrement"`
	RepoId        uint64  `json:"repo_id" db:"repo_id" gorm:"index"`
	Login         string  `json:"login" db:"login"`
	GitId         int     `json:"git_id" db:"git_id"`
	Contributions int     `json:"contributions" db:"contributions"`
	Score         float64 `json:"score" db:"score"` // contributons / total_contributions
}

func (g *GitContributor) TableName() string {
	return "git_contributor"
}

func CreateGitContributor(gitContributor *GitContributor) error {
	return common.MDB.Create(gitContributor).Error
}

func CreateGitContributors(gitContributors []*GitContributor) error {
	return common.MDB.Create(gitContributors).Error
}

func GetGitContributorsByRepoId(repoId uint64) ([]*GitContributor, error) {
	var gitContributors []*GitContributor
	if err := common.MDB.Where("repo_id = ?", repoId).Find(&gitContributors).Error; err != nil {
		return nil, err
	}
	return gitContributors, nil
}

func GetAllGitContributors(repoIdList []uint64) ([]*GitContributor, error) {
	var gitContributors []*GitContributor
	if err := common.MDB.Where("repo_id IN ?", repoIdList).Find(&gitContributors).Error; err != nil {
		return nil, err
	}
	return gitContributors, nil
}
