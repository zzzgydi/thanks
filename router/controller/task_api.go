package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zzzgydi/thanks/service/thk"
)

func PostCreateTask(c *gin.Context) {
	var req CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ReturnBadRequest(c, err)
		return
	}

	data, err := thk.LoadDataFromUrl(req.Url)
	if err != nil {
		ReturnServerError(c, err)
		return
	}

	minScore := req.MinScore
	if req.MinScore < 0.0001 {
		minScore = 0.0001
	}
	if req.MinScore > 0.1 {
		minScore = 0.1
	}

	task, err := thk.NewThankTask(req.Lang, minScore, data)
	if err != nil {
		ReturnServerError(c, err)
		return
	}

	contributions, err := task.Run()
	if err != nil {
		ReturnServerError(c, err)
		return
	}

	ReturnSuccess(c, &TaskResponse{
		Id:            task.Task().Id,
		Lang:          task.Task().Lang,
		CreatedAt:     task.Task().CreatedAt,
		UpdatedAt:     task.Task().UpdatedAt,
		Contributions: contributions,
	})
}
