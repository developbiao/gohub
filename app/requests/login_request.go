package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub/app/requests/validators"
)

type LoginByPhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

type LoginByPasswordRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`
	LoginID       string `valid:"login_id" json:"login_id"`
	Password      string `valid:"password" json:"password,omitempty"`
}

// LoginByPhone login user by phone
func LoginByPhone(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"verify_code": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"phone": []string{
			"required:Phone number is required, parameter is phone",
			"digits:Phone number length must is 11 digits",
		},
		"verify_code": []string{
			"required:Verify code is require, parameter is verify_code",
			"digits:Verify code number length must is 6 digits",
		},
	}

	errs := validate(data, rules, messages)

	// Phone sms code verify
	_data := data.(*LoginByPhoneRequest)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)
	return errs
}

// LoginByPassword login user by password
func LoginByPassword(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"login_id":       []string{"required", "min:3"},
		"password":       []string{"required", "min:6"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"login_id": []string{
			"required:Login ID is required, support phone,email username",
			"min:Login ID length must greater than 3",
		},
		"password": []string{
			"required:Password is required",
			"min:Password length must greater than 6",
		},
		"captcha_id": []string{
			"required:Captcha ID is required",
		},
		"captcha_answer": []string{
			"required:Captcha answer is required",
			"digits: Captcha answer length must greater than 6",
		},
	}

	errs := validate(data, rules, messages)

	// Captcha verify
	_data := data.(*LoginByPasswordRequest)
	errs = validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaAnswer, errs)
	return errs
}
