package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zzzgydi/thanks/router/controller"
	"github.com/zzzgydi/thanks/router/middleware"
)

func TaskRouter(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(middleware.LoggerMiddleware)

	api.POST("/task/create", controller.PostCreateTask)
	api.GET("/task/:id", controller.GetDetailTask)
}
