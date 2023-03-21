package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub/pkg/auth"
)

type UserUpdateProfileRequest struct {
	Name         string `valid:"name" json:"name"`
	City         string `valid:"city" json:"city"`
	Introduction string `valid:"introduction" json:"introduction"`
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
