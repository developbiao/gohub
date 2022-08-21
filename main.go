package main

import "C"
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	// Initialization gin instance
	r := gin.Default()

	// Registration middleware
	r.Use(gin.Logger(), gin.Recovery())

	// Registration route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "Hello Gin is working!",
		})
	})

	// Processing 404 request
	r.NoRoute(func(c *gin.Context) {
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

	// Running server on 8000 port
	r.Run(":8000")
}
