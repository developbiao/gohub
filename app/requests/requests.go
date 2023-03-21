package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub/pkg/response"
)

// ValidatorFunc  validation function type
type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

// Validate validate function handler user customer rule
func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	// 1. Paring request support json, from data, URL QUERY
	if err := c.ShouldBind(obj); err != nil {
		response.BadRequest(c, err, "Request parameters paring err,"+
			"please ensure parameters is json, upload file is multipart header")
		return false
	}

	// 2. Form data validation
	errs := handler(obj, c)
	if len(errs) > 0 {
		response.ValidationError(c, errs)
		return false
	}
	return true
}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	// Init config
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid", // In model struct using identify
	}

	// Start validation
	return govalidator.New(opts).ValidateStruct()
}

// validateFile validate file
func validateFile(c *gin.Context, data interface{}, rules govalidator.MapData,
	messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Request:       c.Request,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}
	return govalidator.New(opts).Validate()
}
