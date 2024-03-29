package cmd

import (
	"github.com/spf13/cobra"
	"gohub/database/seeders"
	"gohub/pkg/console"
	"gohub/pkg/seed"
)

var CmdDBSeed = &cobra.Command{
	Use:   "seed",
	Short: "Insert fake data to the database",
	Run:   runSeeders,
	Args:  cobra.MaximumNArgs(1), // Maximum 1 arguments
}

func runSeeders(command *cobra.Command, args []string) {
	seeders.Initialize()
	if len(args) > 0 {
		// There are cases where parameters are passwd
		name := args[0]
		seeder := seed.GetSeeder(name)
		if len(seeder.Name) > 0 {
			seed.RunSeeder(name)
		} else {
			console.Error("Seeder does not found: " + name)
		}
	} else {
		seed.RunAll()
	}
	console.Success("Done seeding.")
}
