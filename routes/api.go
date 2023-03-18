package routes

import (
	"github.com/gin-gonic/gin"
	controllers "gohub/app/http/controllers/api/v1"
	"gohub/app/http/controllers/api/v1/auth"
	"gohub/app/http/middlewares"
	"net/http"
)

func RegisterAPIRouters(r *gin.Engine) {
	// Test v1 group represent v1 version routes
	v1 := r.Group("/v1")

	// Global limiter middleware: every hour limit 60 reqs
	v1.Use(middlewares.LimitIP("200-H"))
	{
		// Registration a route
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"ping": "is working!",
			})
		})

		// Auth group
		authGroup := v1.Group("/auth")
		// Limit ip 1 hour / 1000 reqs
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			signup := new(auth.SignupController)
			// Check phone is registered
			authGroup.POST("/signup/phone/exist", middlewares.LimitIP("60-H"), signup.IsPhoneExist)
			// Check email is registered
			authGroup.POST("/signup/email/exist", middlewares.LimitIP("5-S"), signup.IsEmailExist)
			// Using phone registration
			authGroup.POST("/signup/using-phone", signup.SignupUsingPhone)
			// Using email registration
			authGroup.POST("/signup/using-email", signup.SignupUsingEmail)

			// Show captcha
			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", middlewares.LimitIP("50-H"), vcc.ShowCaptcha)
			// Using phone send sms code
			authGroup.POST("/verify-codes/phone", middlewares.LimitIP("20-H"), vcc.SendUsingPhone)
			// Using email send code
			authGroup.POST("/verify-codes/email", middlewares.LimitIP("20-H"), vcc.SendUsingEmail)

			// Using phone login
			lgc := new(auth.LoginController)
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), lgc.LoginByPhone)
			// Using password login
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)

			// Refresh token
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), lgc.RefreshToken)

			// Reset password
			pwc := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-phone", middlewares.GuestJWT(), pwc.ResetByPhone)
			authGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), pwc.ResetByEmail)

			// Get Current user
			uc := new(controllers.UsersController)
			v1.GET("/user", middlewares.AuthJWT(), uc.CurrentUser)
			usersGroup := v1.Group("/users")
			{
				usersGroup.GET("", uc.Index)
			}

			// Category
			cgc := new(controllers.CategoriesController)
			cgcGroup := v1.Group("/categories")
			{
				cgcGroup.GET("", cgc.Index)
				cgcGroup.POST("", middlewares.AuthJWT(), cgc.Store)
				cgcGroup.PUT("/:id", middlewares.AuthJWT(), cgc.Update)
				cgcGroup.DELETE("/:id", middlewares.AuthJWT(), cgc.Delete)
			}

			// Topic
			tpc := new(controllers.TopicsController)
			tpcGroup := v1.Group("/topics")
			{
				tpcGroup.POST("", middlewares.AuthJWT(), tpc.Store)
				tpcGroup.PUT("/:id", middlewares.AuthJWT(), tpc.Update)
				tpcGroup.DELETE("/:id", middlewares.AuthJWT(), tpc.Delete)
			}
		}
	}
}
