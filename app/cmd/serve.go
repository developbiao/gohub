package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gohub/bootstrap"
	"gohub/pkg/config"
	"gohub/pkg/console"
	"gohub/pkg/logger"
)

// CmdServe represents the avaliable web sub-command
var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start web server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {
	// Set gin running mode, support debug, release, test
	// Release mode block debug information officer recommend use on production
	// Default set release
	// In special cases you can manually change to debug mode
	gin.SetMode(gin.ReleaseMode)

	// Initialization gin instance
	router := gin.New()

	// Initialization and binding route
	bootstrap.SetupRoute(router)

	// Running server
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		logger.ErrorString("CMD", "serve", err.Error())
		console.Exit("Unable to start server, error:" + err.Error())
	}

	// Test module verify is work
	testModule()
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
	console.Error("Console module is running ok!")
}
