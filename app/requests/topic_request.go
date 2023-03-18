package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type TopicRequest struct {
	Title      string `valid:"title" json:"title,omitempty"`
	Body       string `valid:"body" json:"body,omitempty"`
	CategoryID string `valid:"category_id" json:"category_id,omitempty"`
}

func TopicSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"title":       []string{"required", "min_cn:3", "max_cn:40"},
		"body":        []string{"required", "min_cn:10", "max_cn:50000"},
		"category_id": []string{"required", "exists:categories,id"},
	}
	messages := govalidator.MapData{
		"title": []string{
			"required:标题名称为必填项",
			"min_cn:标题长度需至少 3 个字",
			"max_cn:标题长度不能超过 40 个字",
		},
		"body": []string{
			"required:帖子内容为必填项",
			"min_cn:帖子长度需至少 3 个字",
			"max_cn:帖子长度不能超过 40 个字",
		},
		"category_id": []string{
			"required:帖子分类为必填项",
			"exists:帖子分类未找到",
		},
	}
	return validate(data, rules, messages)
}
