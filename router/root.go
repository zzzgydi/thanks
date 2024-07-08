package router

import "github.com/gin-gonic/gin"

func RootRouter(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.Redirect(301, "https://thanks.com")
	})

	health := r.Group("/__internal__")

	health.GET("/health", func(c *gin.Context) {
		c.String(200, "ok")
	})
}
