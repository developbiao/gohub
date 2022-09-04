package database

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// DB objects
var DB *gorm.DB
var SQLDB *sql.DB

// Connect to database
func Connect(dbConfig gorm.Dialector, _logger gormlogger.Interface) {
	// Using gorm.Open connection to database
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})

	// Error happen
	if err != nil {
		fmt.Println(err.Error())
	}

	// Get base sqlDB
	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}
