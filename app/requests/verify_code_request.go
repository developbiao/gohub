package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub/pkg/captcha"
)

type VerifyCodePhoneRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`
	Phone         string `json:"phone,omitempty" valid:"phone"`
}

// VerifyCodePhone verify phone captcha answer
func VerifyCodePhone(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone":          []string{"required", "digits:11"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:Phone number is required, parameter phone",
			"digits:Phone must is 11 digits",
		},
		"captcha_id": []string{
			"required:Captcha id is required, parameter captcha_id",
		},
		"captcha_answer": []string{
			"required:Captcha answer is required, parameter captcha_answer",
			"digits:Captcha answer must is 6 digits",
		},
	}
	errs := validate(data, rules, messages)
	_data := data.(*VerifyCodePhoneRequest)
	if ok := captcha.NewCaptcha().VerifyCaptcha(_data.CaptchaID, _data.CaptchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "Captcha answer is error")
	}
	return errs
}
