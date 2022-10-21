package validators

import (
	"errors"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"gohub/pkg/database"
	"strings"
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
}
