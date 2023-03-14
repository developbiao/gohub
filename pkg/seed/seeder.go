package seed

import (
	"gohub/pkg/console"
	"gohub/pkg/database"
	"gorm.io/gorm"
)

// Save all seeder
var seeders []Seeder

// Ordered seeder array
// Support must be executed in order, example topic create must dependency user
// So TopSeeder must run UserSeeder after executing
var orderedSeederNames []string

type SeederFunc func(*gorm.DB)

// Seeder corresponding to each database/seeders directory Seeder file
type Seeder struct {
	Func SeederFunc
	Name string
}

// Add register to seeders array
func Add(name string, fn SeederFunc) {
	seeders = append(seeders, Seeder{
		Name: name,
		Func: fn,
	})
}

// SetRunOrder set order seeder array
func SetRunOrder(names []string) {
	orderedSeederNames = names
}

// GetSeeder get seeder instance by name
func GetSeeder(name string) Seeder {
	for _, sdr := range seeders {
		if name == sdr.Name {
			return sdr
		}
	}
	return Seeder{}
}

// RunAll Run all seeder
func RunAll() {
	// First running ordered
	executed := make(map[string]string)
	for _, name := range orderedSeederNames {
		sdr := GetSeeder(name)
		if len(sdr.Name) > 0 {
			console.Warning("Running Ordered Seeder: " + sdr.Name)
			sdr.Func(database.DB)
			executed[name] = name
		}
	}

	// Continue running other seeder
	for _, sdr := range seeders {
		if _, ok := executed[sdr.Name]; !ok {
			console.Warning("Running Seeder: " + sdr.Name)
			sdr.Func(database.DB)
		}
	}
}

// RunSeeder run single seeder
func RunSeeder(name string) {
	for _, sdr := range seeders {
		if sdr.Name == name {
			sdr.Func(database.DB)
			break
		}
	}
}
