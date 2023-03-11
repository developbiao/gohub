package migrate

import (
	"database/sql"
	"gorm.io/gorm"
)

// migrationFunc define up and down callback function type
type migrationFunc func(gorm.Migrator, *sql.DB)

// MigrationFile represent a migration file
type MigrationFile struct {
	Up       migrationFunc
	Down     migrationFunc
	FileName string
}

// migrationFiles all migration file
var migrationFiles []MigrationFile

// Add a migration file and registration
func Add(name string, up migrationFunc, down migrationFunc) {
	migrationFiles = append(migrationFiles, MigrationFile{
		FileName: name,
		Up:       up,
		Down:     down,
	})
}
