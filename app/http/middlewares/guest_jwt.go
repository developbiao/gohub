package middlewares

import (
	"github.com/gin-gonic/gin"
	"gohub/pkg/jwt"
	"gohub/pkg/response"
)

// GuestJWT guest middleware check
func GuestJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) > 0 {
			_, err := jwt.NewJWT().ParserToken(c)
			if err == nil {
				response.Unauthorized(c, "Please use guest visit")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
