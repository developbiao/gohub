package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/pkg/captcha"
	"gohub/pkg/logger"
	"gohub/pkg/response"
)

// VerifyCodeController user controller
type VerifyCodeController struct {
	v1.BaseAPIController
}

// ShowCaptcha show captcha
func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	// Generate captcha
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	logger.LogIf(err)
	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}
