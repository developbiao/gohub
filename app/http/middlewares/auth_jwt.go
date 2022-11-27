package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/app/models/user"
	"gohub/pkg/config"
	"gohub/pkg/jwt"
	"gohub/pkg/response"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization:Bearer information from header
		claims, err := jwt.NewJWT().ParserToken(c)
		if err != nil {
			response.Unauthorized(c,
				fmt.Sprintf("Authorization filed, please check %v authorization document",
					config.GetString("app.name")))
			return
		}

		// Find user by user id
		userModel := user.Get(claims.UserID)
		if userModel.ID == 0 {
			response.Unauthorized(c, "Can not find user, many be user does not exists")
			return
		}

		// Save user information to context
		c.Set("current_user_id", userModel.GetStringID())
		c.Set("current_user_name", userModel.Name)
		c.Set("current_user", userModel)

		c.Next()
	}
}
