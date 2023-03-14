package seed

import "gorm.io/gorm"

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
