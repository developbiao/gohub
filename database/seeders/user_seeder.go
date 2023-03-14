package seeders

import (
	"fmt"
	"gohub/database/factories"
	"gohub/pkg/console"
	"gohub/pkg/logger"
	"gohub/pkg/seed"
	"gorm.io/gorm"
)

func init() {
	// Add Seeder
	seed.Add("SeedUsersTable", func(db *gorm.DB) {
		users := factories.MakeUsers(10)

		// Batch create user (Note that batch creation does not call model hooks)
		result := db.Table("users").Create(&users)

		// Record error
		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		// Output running message
		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
