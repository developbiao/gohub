package requests

import "github.com/thedevsaddam/govalidator"

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
