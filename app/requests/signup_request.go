package requests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

type SignupEmailExistsRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

// ValidatorFunc  validation function type
type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

// Validate validate function handler user customer rule
func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	// 1. Paring request support json, from data, URL QUERY
	if err := c.ShouldBind(obj); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Request parameters paring err, " +
				"please ensure parameters is json, upload file is multipart header",
		})
		fmt.Println(err.Error())
		return false
	}

	// 2. Form data validation
	errs := handler(obj, c)
	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求验证不能过，具体请查看 errors",
			"error":   errs,
		})
		return false
	}
	return true
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
