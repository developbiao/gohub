package routes

import (
	"github.com/gin-gonic/gin"
	"gohub/app/http/controllers/api/v1/auth"
	"net/http"
)

func RegisterAPIRouters(r *gin.Engine) {
	// Test v1 group represent v1 version routes
	v1 := r.Group("/v1")
	{
		// Registration a route
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"ping": "is working!",
			})
		})

		// Auth group
		authGroup := v1.Group("auth")
		{
			signup := new(auth.SignupController)
			// Check phone is registered
			authGroup.POST("/signup/phone/exist", signup.IsPhoneExist)
			// Check email is registered
			authGroup.POST("/signup/email/exist", signup.IsEmailExist)

			// Show captcha
			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
		}
	}
}
