package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
	btsConfig "gohub/config"
	"gohub/pkg/config"
)

func init() {
	// Load config folder configs
	btsConfig.Initialize()
}

func main() {

	// Initialization dependency command --env arguments
	var env string
	flag.StringVar(&env, "env", "", "load .env file, e.g: --env=testing")
	flag.Parse()
	config.InitConfig(env)

	// Initialization database
	bootstrap.SetupDB()

	// Initialization gin instance
	router := gin.New()

	// Initialization and binding route
	bootstrap.SetupRoute(router)

	// Running server on 3000 port
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		// Error exception
		fmt.Println(err.Error())
	}
}
