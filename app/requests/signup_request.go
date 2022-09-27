package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

type SignupEmailExistsRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

func SignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {
	// Customer validation rule
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	// Customer validation message
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位数字",
		},
	}
	return validate(data, rules, messages)
}

func SignupEmailExist(data interface{}, c *gin.Context) map[string][]string {
	// Rules
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}

	// Error Message
	messages := govalidator.MapData{
		"email": []string{
			"required:Email  为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
	}

	// Start validation
	return validate(data, rules, messages)
}
