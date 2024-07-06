package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zzzgydi/thanks/router/controller"
	"github.com/zzzgydi/thanks/router/utils"
)

func RateLimiterMiddleware() gin.HandlerFunc {
	ipLimiter := utils.NewRateLimiter("z:rl:ip", 20, 60*60)     // 1小时 最多15次
	userLimiter := utils.NewRateLimiter("z:rl:user", 60, 60*60) // 1小时 最多60次

	return func(c *gin.Context) {
		// 先判断是否为登录用户
		// 如果是登录用户，则使用用户ID作为限流key
		userId := utils.GetUserID(c)
		ip := c.ClientIP()

		var allow bool
		var err error

		// 仅正式登录的用户用用户ID作为限流key
		// 游客用ip限流
		if userId != "" {
			allow, err = userLimiter.Allow(c, userId)
		} else {
			allow, err = ipLimiter.Allow(c, ip)
		}

		if err != nil {
			controller.ReturnServerError(c, err)
			return
		}

		if trace := utils.GetTraceLogger(c); trace != nil {
			trace.Trace("rl_allow", allow)
		}

		if !allow {
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}

		c.Next()
	}
}
