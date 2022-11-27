package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
	btsConfig "gohub/config"
	"gohub/pkg/config"
	"gohub/pkg/logger"
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

	// Test get user form context
	//router.GET("/test_auth", middlewares.AuthJWT(), func(c *gin.Context) {
	//	userModel := auth.CurrentUser(c)
	//	response.Data(c, userModel)
	//})

	// Test module verify is work
	testModule()

	// Running server on 3000 port
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		// Error exception
		fmt.Println(err.Error())
	}
}

func testModule() {
	//  Verify captcha test
	//logger.Dump(captcha.NewCaptcha().VerifyCaptcha("qaDAO2ccO0SSbXYXdu9G", "723469"), "Correct answer")
	//logger.Dump(captcha.NewCaptcha().VerifyCaptcha("qaDAO2ccO0SSbXYXdu9G", "0000"), "Error answer")

	// Send sms test
	//sms.NewSMS().Send("1333000000", sms.Message{
	//	Template: config.GetString("sms.aliyun.template.code"),
	//	Data:     map[string]string{"code": "123456"},
	//})
	//
	//sms.NewSMS().Send("13330000000", sms.Message{
	//	Template: config.GetString("sms.test.template.code"),
	//	Data:     map[string]string{"code": "123456"},
	//})

	//verifycode.NewVerifyCode().SendSMS("13330000000")
	//if verifycode.NewVerifyCode().CheckAnswer("13330000000", "123456") {
	//	logger.DebugString("verifycode", "verify success", "123456")
	//}
	logger.DebugString("Test Module finished", "test module", "ok")
}
