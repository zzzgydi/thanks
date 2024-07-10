package model

import (
	"time"

	"github.com/lithammer/shortuuid/v4"
	"github.com/zzzgydi/thanks/common"
)

type Task struct {
	Id        string    `json:"id" db:"id" gorm:"primary_key;type:char(22)"`
	Lang      string    `json:"lang" db:"lang"`
	MinScore  float64   `json:"min_score" db:"min_score"`
	Repos     string    `json:"repos" db:"repos"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (t *Task) TableName() string {
	return "task"
}

func CreateTask(lang string, minScore float64) (*Task, error) {
	task := &Task{
		Id:        shortuuid.New(),
		Lang:      lang,
		MinScore:  minScore,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := common.MDB.Create(task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func GetTaskById(id string) (*Task, error) {
	task := &Task{}
	if err := common.MDB.Where("id = ?", id).First(task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func SaveTask(task *Task) error {
	task.UpdatedAt = time.Now()
	return common.MDB.Save(task).Error
}
