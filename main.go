package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
)

func main() {
	// Initialization gin instance
	router := gin.New()

	// Initialization and binding route
	bootstrap.SetupRoute(router)

	// Running server on 3000 port
	err := router.Run(":3000")
	if err != nil {
		// Error exception
		fmt.Println(err.Error())
	}
}
