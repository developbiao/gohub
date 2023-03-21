package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub/app/requests/validators"
	"gohub/pkg/auth"
)

type UserUpdateProfileRequest struct {
	Name         string `valid:"name" json:"name"`
	City         string `valid:"city" json:"city"`
	Introduction string `valid:"introduction" json:"introduction"`
}

type UserUpdateEmailRequest struct {
	Email      string `valid:"email" json:"email,omitempty"`
	VerifyCode string `valid:"verify_code" json:"verify_code,omitempty"`
}

func UserUpdateProfile(data interface{}, c *gin.Context) map[string][]string {
	// Query user UID, filter self duplicate name
	uid := auth.CurrentUID(c)
	rules := govalidator.MapData{
		"name":         []string{"required", "between:3,20", "alpha_num", "not_exists:users,name," + uid},
		"city":         []string{"min_cn:2", "max_cn:20"},
		"introduction": []string{"min_cn:2", "max_cn:250"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度在 3~20 之间",
			"not_exists:用户名已被占用",
		},
		"introduction": []string{
			"min_cn:自我介绍信息长度需至少 2 个字",
			"max_cn:自我介绍信息长度不能超过 240 个字",
		},
		"city": []string{
			"min_cn:城市至少 2 两个字",
			"max_cn:城市不能超过 20 个字",
		},
	}
	return validate(data, rules, messages)
}

func UserUpdateEmail(data interface{}, c *gin.Context) map[string][]string {
	currentUser := auth.CurrentUser(c)
	rules := govalidator.MapData{
		"email": []string{
			"required",
			"email",
			"min:4",
			"max:30",
			"not_exists:users,email," + currentUser.GetStringID(),
			"not_in:" + currentUser.Email,
		},
		"verify_code": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必须项",
			"min:Email 最小长度需要大于 4",
			"max:Email 长度需要小于 30",
			"email:Email 格式工不正确，请提供有效的邮箱地址",
			"not_exists:Email 已被占用",
			"not_in:Email 与旧 Email 一致",
		},
		"verify_code": []string{
			"required:验证码答案是必填",
			"digits:验证码长度必须为 6 位数字",
		},
	}

	errs := validate(data, rules, messages)
	_data := data.(*UserUpdateEmailRequest)
	errs = validators.ValidateVerifyCode(_data.Email, _data.VerifyCode, errs)
	return errs
}
