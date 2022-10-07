package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/requests"
	"gohub/pkg/captcha"
	"gohub/pkg/logger"
	"gohub/pkg/response"
	"gohub/pkg/verifycode"
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

// SendUsingPhone using phone send sms code
func (vc *VerifyCodeController) SendUsingPhone(c *gin.Context) {
	// Validation phone captcha
	request := requests.VerifyCodePhoneRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCodePhone); !ok {
		return
	}

	// Send SMS
	if ok := verifycode.NewVerifyCode().SendSMS(request.Phone); !ok {
		response.Abort500(c, "Send sms failed")
	} else {
		response.Success(c)
	}
}

// SendUsingEmail using email send code
func (vc *VerifyCodeController) SendUsingEmail(c *gin.Context) {
	request := requests.VerifyCodeEmailRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCodeEmail); !ok {
		return
	}

	// Send email
	if err := verifycode.NewVerifyCode().SendEmail(request.Email); err != nil {
		response.Abort500(c, "Send email failed")
	} else {
		response.Success(c)
	}
}
