package controller

import (
	"fmt"
	"strconv"

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

	total := len(contributions)
	size := 20
	if total < size {
		size = total
	}

	ReturnSuccess(c, &TaskResponse{
		Id:        task.Task().Id,
		Lang:      task.Task().Lang,
		Offset:    0,
		Total:     total,
		List:      contributions[:size],
		CreatedAt: task.Task().CreatedAt,
		UpdatedAt: task.Task().UpdatedAt,
	})
}

func GetDetailTask(c *gin.Context) {
	taskId := c.Param("id")
	if taskId == "" {
		ReturnBadRequest(c, fmt.Errorf("task id is empty"))
		return
	}

	offsetStr := c.DefaultQuery("offset", "0")
	sizeStr := c.DefaultQuery("size", "100")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		size = 100
	}

	task, err := thk.NewThankTaskFromId(taskId)
	if err != nil {
		ReturnServerError(c, err)
		return
	}

	contributions, err := task.Run()
	if err != nil {
		ReturnServerError(c, err)
		return
	}

	list := contributions

	if offset >= len(contributions) {
		list = []*thk.ThkContributor{}
	} else {
		if offset+size > len(contributions) {
			size = len(contributions)
		} else {
			size = offset + size
		}
		list = contributions[offset:size]
	}

	ReturnSuccess(c, &TaskResponse{
		Id:        task.Task().Id,
		Lang:      task.Task().Lang,
		Offset:    offset,
		Total:     len(contributions),
		List:      list,
		CreatedAt: task.Task().CreatedAt,
		UpdatedAt: task.Task().UpdatedAt,
	})
}
