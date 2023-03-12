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

// getMigrationFile get MigrationFile Object by migration file name
func getMigrationFile(name string) MigrationFile {
	for _, mfile := range migrationFiles {
		if name == mfile.FileName {
			return mfile
		}
	}
	return MigrationFile{}
}

// isNotMigrated check is migrated
func (mfile MigrationFile) isNotMigrated(migrations []Migration) bool {
	for _, migration := range migrations {
		if migration.Migration == mfile.FileName {
			return false
		}
	}
	return true
}
