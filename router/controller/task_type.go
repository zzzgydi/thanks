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

type DetailTaskRequest struct {
	Id     string `json:"id"`
	Offset int    `json:"offset,omitempty"`
	Size   int    `json:"size,omitempty"`
}

type TaskResponse struct {
	Id        string                `json:"id"`
	Lang      string                `json:"lang"`
	Offset    int                   `json:"offset"`
	Total     int                   `json:"total"`
	List      []*thk.ThkContributor `json:"list"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}
