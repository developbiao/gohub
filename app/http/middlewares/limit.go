package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gohub/pkg/app"
	"gohub/pkg/limiter"
	"gohub/pkg/logger"
	"gohub/pkg/response"
	"net/http"
)

// LimitIP Global limiter middleware by ip
// limit format string  e.g:
//
// * 5 reqs/seconds: "5-S"
// * 10 reqs/minute: "10-M"
// * 1000 reqs/hour: "1000-H"
// * 2000 reqs/day: "2000-D"
func LimitIP(limit string) gin.HandlerFunc {
	if app.IsTesting() {
		limit = "1000000-H"
	}

	return func(c *gin.Context) {
		key := limiter.GetKeyIP(c)
		if ok := limitHandler(c, key, limit); !ok {
			return
		}
		c.Next()
	}
}

// LimitPerRoute limit by route
func LimitPerRoute(limit string) gin.HandlerFunc {
	if app.IsTesting() {
		limit = "1000000-H"
	}

	return func(c *gin.Context) {
		// Increment count by route
		c.Set("limiter-once", false)

		// Limit by ip
		key := limiter.GetKeyRouteWithIP(c)
		if ok := limitHandler(c, key, limit); !ok {
			return
		}
		c.Next()
	}

}

func limitHandler(c *gin.Context, key string, limit string) bool {
	// Get limit
	rate, err := limiter.CheckRate(c, key, limit)
	if err != nil {
		logger.LogIf(err)
		response.Abort500(c)
		return false
	}

	// ---- Set header information ----
	// X-RateLimit-Limit: 10000
	// x-RateLimit-Remaining: 9991
	// x-RateLimit-Reset: 1671450338
	c.Header("X-RateLimit-Limit", cast.ToString(rate.Limit))
	c.Header("X-RateLimit-Remaining", cast.ToString(rate.Remaining))
	c.Header("X-RateLimit-Reset", cast.ToString(rate.Reset))

	// Is Reached
	if rate.Reached {
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"message": "接口请求太频繁，请稍后再试~",
		})
		return false
	}
	return true
}
