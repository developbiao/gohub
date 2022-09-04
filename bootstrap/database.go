package bootstrap

import (
	"errors"
	"fmt"
	"gohub/pkg/config"
	"gohub/pkg/database"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func SetupDB() {
	var dbConfig gorm.Dialector
	switch config.Get("database.connection") {
	case "mysql":
		// Construct DSN
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?"+
			"charset=%v&parseTime=True&multiStatements=true&loc=Local",
			config.Get("database.mysql.username"),
			config.Get("database.mysql.password"),
			config.Get("database.mysql.host"),
			config.Get("database.mysql.port"),
			config.Get("database.mysql.database"),
			config.Get("database.mysql.charset"),
		)
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	case "sqlite":
		// Initialize sqlite
		database := config.Get("database.sqlite.database")
		dbConfig = sqlite.Open(database)
	default:
		panic(errors.New("database connection not supported"))
	}

	// Connection database, and set GORM logger mode
	database.Connect(dbConfig, logger.Default.LogMode(logger.Info))

	// Set max connections
	database.SQLDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))

	// Set max idle
	database.SQLDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))

	// Set each connection expire time
	database.SQLDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)

}
