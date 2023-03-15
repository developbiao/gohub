package database

import (
	"database/sql"
	"errors"
	"fmt"
	"gohub/pkg/config"
	"gohub/pkg/console"
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

// CurrentDatabase get database name
func CurrentDatabase() (dbname string) {
	dbname = DB.Migrator().CurrentDatabase()
	return
}

// DeleteAllTables  Delete all tables by connection kind
func DeleteAllTables() error {
	var err error
	switch config.Get("database.connection") {
	case "mysql":
		err = deleteMySQLTables()
	case "sqlite":
		err = deleteAllSqliteTables()
	default:
		panic(errors.New("database connection not supported"))
	}
	return err
}

func deleteMySQLTables() error {
	dbname := CurrentDatabase()
	var tables []string

	err := DB.Table("information_schema.tables").
		Where("table_schema = ?", dbname).
		Pluck("table_name", &tables).Error
	if err != nil {
		return err
	}

	//  Temporarily turn off checking for foreign keys
	DB.Exec("SET foreign_key_checks = 0;")

	// Delete all tables
	for _, table := range tables {
		err := DB.Migrator().DropTable(table)
		if err != nil {
			return err
		}
		console.Success("Deleted table:[" + table + "]")
	}

	// Turn On checking for foreign keys
	DB.Exec("SET foreign_key_checks = 1;")
	return nil
}

func deleteAllSqliteTables() error {
	var tables []string

	// Get All tables
	err := DB.Select(&tables, "SELECT name FROM sqlite_master WHERE type='table'").Error
	if err != nil {
		return err
	}

	// Delete all tables
	for _, table := range tables {
		err := DB.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}
	return nil
}

// TableName Get table name by model
func TableName(obj interface{}) string {
	stmt := &gorm.Statement{DB: DB}
	stmt.Parse(obj)
	return stmt.Schema.Table
}
