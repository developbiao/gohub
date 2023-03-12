// Package migrate process database migrate
package migrate

import (
	"fmt"
	"gohub/pkg/console"
	"gohub/pkg/database"
	"gohub/pkg/file"
	"gorm.io/gorm"
	"os"
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
	migration := Migration{}

	// If not exist create migrations table
	if !migrator.Migrator.HasTable(&migration) {
		migrator.Migrator.CreateTable(&migration)
	}
}

// Up run execute migration
func (migrator *Migrator) Up() {
	// Read all migration files, ensure sort by date
	migrateFiles := migrator.readAllMigrationFiles()

	// Get current batch value
	batch := migrator.getBatch()

	// Get all migration data
	var migrations []Migration
	migrator.DB.Find(&migrations)

	// flag database is new
	runed := false
	for _, mfile := range migrateFiles {
		if mfile.isNotMigrated(migrations) {
			migrator.runUpMigration(mfile, batch)
			runed = true
		}
	}

	if !runed {
		console.Success("database is up to date.")
	}
}

// getBatch get batch
func (migrator *Migrator) getBatch() int {
	// default value is 1
	batch := 1

	// Get last execute migration data
	lastMigration := Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)

	// If exist value increment
	if lastMigration.ID > 0 {
		batch = lastMigration.Batch + 1
	}
	return batch
}

// readAllMigrationFiles Read file from directory, ensure files is sort by date
func (migrator *Migrator) readAllMigrationFiles() []MigrationFile {
	// Read in database/migrations/ directory all files
	files, err := os.ReadDir(migrator.Folder)
	console.ExitIf(err)

	console.Warning("Debug read all migration files")
	fmt.Println("Files", files)

	var migrateFiles []MigrationFile
	for _, f := range files {
		// Remove file extension ".go"
		fileName := file.FileNameWithoutExtension(f.Name())

		// Get MigrationFile object by migrate file
		mfile := getMigrationFile(fileName)

		// Check file is available append to array
		if len(mfile.FileName) > 0 {
			migrateFiles = append(migrateFiles, mfile)
		}
	}
	fmt.Println("Match Files", migrateFiles)
	// Return sorted array
	return migrateFiles
}

// runUpMigration run up to migration
func (migrator *Migrator) runUpMigration(mfile MigrationFile, batch int) {
	// execute up block SQL
	if mfile.Up != nil {
		console.Warning("migration " + mfile.FileName)
		mfile.Up(database.DB.Migrator(), database.SQLDB)
		console.Success("migrated " + mfile.FileName)
	}

	// Sync to database
	err := migrator.DB.Create(&Migration{Migration: mfile.FileName, Batch: batch}).Error
	console.ExitIf(err)
}
