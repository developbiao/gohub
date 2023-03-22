package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gohub/pkg/response"
)

func ForceUA() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Request.Header["User-Agent"]) == 0 {
			response.BadRequest(c, errors.New("User-Agent header not detected"),
				"Request User-Agent is required")
			return
		}
		c.Next()
	}
}
