// Package migrate process database migrate
package migrate

import (
	"gohub/pkg/database"
	"gorm.io/gorm"
)

// Migrator database migration operator
type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
}

// Migration represent migrations table
type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	Batch     int
}

// NewMigrator create migrator instance
func NewMigrator() *Migrator {
	migrator := &Migrator{
		Folder:   "database/migrations/",
		DB:       database.DB,
		Migrator: database.DB.Migrator(),
	}
	migrator.createMigrationsTable()
	return migrator
}

// CreateMigrationsTable check and create migrations table
func (migrator *Migrator) createMigrationsTable() {
	migration := Migrator{}

	// If not exist create migrations table
	if !migrator.Migrator.HasTable(&migration) {
		migrator.Migrator.CreateTable(&migration)
	}
}
