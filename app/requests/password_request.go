package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub/app/requests/validators"
)

type RestByPhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Password   string `valid:"password" json:"password,omitempty"`
}

// RestByPhone reset password by phone
func RestByPhone(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"verify_code": []string{"required", "digits:6"},
		"password":    []string{"required", "min:6"},
	}
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填荐，参数名称 phone",
			"digits:手机号长度必须为 11 位数字",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度为 6 位数字",
		},
		"password": []string{
			"required: 密码是必填项",
			"min:密码长度需大于 6 ",
		},
	}

	errs := validate(data, rules, messages)
	_data := data.(*RestByPhoneRequest)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)
	return errs
}
