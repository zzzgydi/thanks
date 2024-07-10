package controller

import (
	"time"

	"github.com/zzzgydi/thanks/service/thk"
)

type CreateTaskRequest struct {
	Lang     string  `json:"lang"`
	Url      string  `json:"url"`
	MinScore float64 `json:"min_score"`
}

type TaskResponse struct {
	Id            string                `json:"id"`
	Lang          string                `json:"lang"`
	Contributions []*thk.ThkContributor `json:"contributions"`
	CreatedAt     time.Time             `json:"created_at"`
	UpdatedAt     time.Time             `json:"updated_at"`
}
