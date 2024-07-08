package model

import "github.com/zzzgydi/thanks/common"

type NodeRepo struct {
	Id   uint64 `json:"id" db:"id" gorm:"primary_key;autoIncrement"`
	Pkg  string `json:"pkg" db:"pkg" gorm:"uniqueIndex"`
	Repo string `json:"repo" db:"repo"`
}

func (n *NodeRepo) TableName() string {
	return "node_repo"
}

func CreateNodeRepo(nodeRepo *NodeRepo) error {
	return common.MDB.Create(nodeRepo).Error
}

func GetNodeRepoByPkg(pkg string) (*NodeRepo, error) {
	var nodeRepo NodeRepo
	if err := common.MDB.Where("pkg = ?", pkg).First(&nodeRepo).Error; err != nil {
		return nil, err
	}
	return &nodeRepo, nil
}
