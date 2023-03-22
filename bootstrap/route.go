package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gohub/app/http/middlewares"
	"gohub/routes"
	"net/http"
	"strings"
)

func SetupRoute(router *gin.Engine) {
	// Registration global middleware
	registerGlobalMiddleWare(router)

	// Register api routers
	routes.RegisterAPIRouters(router)

	// Registration API route
	setup404Handler(router)

	// Step public resource
	router.Static("/upload", "./public/upload")
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
	)
}

func setup404Handler(router *gin.Engine) {
	// Processing 404 request
	router.NoRoute(func(c *gin.Context) {
		// Get Accept value from header
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "Page not found 404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "Page not found 404",
			})
		}
	})
}
