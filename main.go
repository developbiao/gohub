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

	// Initialization logger
	bootstrap.SetupLogger()

	// Set gin running mode, support debug, release, test
	// Release mode block debug information officer recommend use on production
	// Default set release
	// In special cases you can manually change to debug mode
	gin.SetMode(gin.ReleaseMode)

	// Initialization database
	bootstrap.SetupDB()

	// Initialization redis
	bootstrap.SetupRedis()

	// Initialization gin instance
	router := gin.New()

	// Initialization and binding route
	bootstrap.SetupRoute(router)

	//logger.Dump(captcha.NewCaptcha().VerifyCaptcha("qaDAO2ccO0SSbXYXdu9G", "723469"), "Correct answer")
	//logger.Dump(captcha.NewCaptcha().VerifyCaptcha("qaDAO2ccO0SSbXYXdu9G", "0000"), "Error answer")

	// Running server on 3000 port
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		// Error exception
		fmt.Println(err.Error())
	}
}
