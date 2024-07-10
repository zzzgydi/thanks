package router

import "github.com/gin-gonic/gin"

func RootRouter(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.Redirect(301, "https://github.com/zzzgydi/thanks")
	})

	r.GET("/health", func(c *gin.Context) {
		c.String(200, "ok")
	})
}
