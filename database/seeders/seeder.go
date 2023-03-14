package seeders

import "gohub/pkg/seed"

func Initialize() {
	// Trigger current directory other file init method

	// Specified seeder priority
	seed.SetRunOrder([]string{
		"SeedUsersTable",
	})
}
