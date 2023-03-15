package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type PaginationRequest struct {
	Sort    string `valid:"sort" form:"sort"`
	Order   string `valid:"order" form:"order"`
	PerPage string `valid:"per_page" form:"per_page"`
	Page    string `valid:"page" from:"page"`
}

func Pagination(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"sort":     []string{"in:id,created_at,updated_at"},
		"order":    []string{"in:asc,desc"},
		"per_page": []string{"numeric_between:2,100"},
		"page":     []string{"numeric"},
	}
	messages := govalidator.MapData{
		"sort": []string{
			"in:sort fields only supported: id,created_at,updated_at",
		},
		"order": []string{
			"in:order rule only supported asc（ASC）,desc（DESC）",
		},
		"per_page": []string{
			"numeric_between:per_page number between: 2~100",
		},
		"page": []string{
			"numeric:page must be number",
		},
	}
	return validate(data, rules, messages)
}
