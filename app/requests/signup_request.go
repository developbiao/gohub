package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

func ValidateSignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {
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

	// Config init
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid", // Model struct tag identify
	}

	// Start validation
	return govalidator.New(opts).ValidateStruct()
}
