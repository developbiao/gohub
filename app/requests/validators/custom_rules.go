package validators

import (
	"errors"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"gohub/pkg/database"
	"strconv"
	"strings"
	"unicode/utf8"
)

func init() {
	// Custom not_exists, validation value must not exists database
	// not_exists:users,email
	// not_exists:users,email,32
	govalidator.AddCustomRule("not_exists", func(field string, rule string, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")
		// First argument is table name
		tableName := rng[0]
		// Second argument is field
		dbFiled := rng[1]

		// Third argument except id
		var exceptID string
		if len(rng) > 2 {
			exceptID = rng[2]
		}

		// User request value
		requestValue := value.(string)

		// Build sql query
		query := database.DB.Table(tableName).Where(dbFiled+" = ?", requestValue)
		if len(exceptID) > 0 {
			query.Where("id != ?", exceptID)
		}

		// Query database
		var count int64
		query.Count(&count)

		if count != 0 {
			// Custom message
			if message != "" {
				return errors.New(message)
			}
			// Default message
			return fmt.Errorf("%v 已被占用", requestValue)
		}
		// Validate pass
		return nil
	})

	// max_cn:8 chinese length maximum is 8
	govalidator.AddCustomRule("max_cn",
		func(field string, rule string, message string, value interface{}) error {
			valLength := utf8.RuneCountInString(value.(string))
			length, _ := strconv.Atoi(strings.TrimPrefix(rule, "max_cn:"))
			if valLength > length {
				if message != "" {
					return errors.New(message)
				}
				return fmt.Errorf("Length cannot exceed %d  characters", length)
			}
			return nil
		})

	// min_cn:2 chinese length cannot less than 2
	govalidator.AddCustomRule("min_cn",
		func(field string, rule string, message string, value interface{}) error {
			valLength := utf8.RuneCountInString(value.(string))
			length, _ := strconv.Atoi(strings.TrimPrefix(rule, "min_cn:"))
			if valLength < length {
				if message != "" {
					return errors.New(message)
				}
				return fmt.Errorf("Length cannot less than %d  characters", length)
			}
			return nil
		})

	// exists:categories,id validation field and value exists?
	govalidator.AddCustomRule("exists",
		func(field string, rule string, message string, value interface{}) error {
			rng := strings.Split(strings.Trim(rule, "exists:"), ",")
			// First argument is table name example: topics
			tableName := rng[0]

			// Second argument is field, example: id
			dbField := rng[1]

			// User request data
			requestValue := value.(string)

			// Query data
			var count int64
			database.DB.Table(tableName).Where(dbField+" =?", requestValue).Count(&count)
			if count == 0 {
				if message != "" {
					return errors.New(message)
				}
				return fmt.Errorf("%v Does not exists", requestValue)
			}
			return nil
		})
}
