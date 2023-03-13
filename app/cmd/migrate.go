package cmd

import (
	"github.com/spf13/cobra"
	"gohub/database/migrations"
	"gohub/pkg/migrate"
)

var CmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration",
}

var CmdMigrateUp = &cobra.Command{
	Use:   "up",
	Short: "Run un migrated migrations",
	Run:   runUp,
}

var CmdMigrateRollback = &cobra.Command{
	Use: "down",
	// Alias name migrate down == migrate rollback
	Aliases: []string{"rollback"},
	Short:   "Reverse the up command",
	Run:     runDown,
}

var CmdMigrateReset = &cobra.Command{
	Use:   "reset",
	Short: "Rollback all database migrations",
	Run:   runReset,
}

var CmdMigrateReFresh = &cobra.Command{
	Use:   "refresh",
	Short: "Reset and re-run all migrations",
	Run:   runRefresh,
}

var CmdMigrateFresh = &cobra.Command{
	Use:   "fresh",
	Short: "Drop all tables and re-run all migrations",
	Run:   runFresh,
}

func init() {
	CmdMigrate.AddCommand(
		CmdMigrateUp,
		CmdMigrateRollback,
		CmdMigrateReset,
		CmdMigrateReFresh,
		CmdMigrateFresh,
	)
}

func migrator() *migrate.Migrator {
	// Register database/migrations allow migration files
	migrations.Initialize()
	// Initialize migrator
	return migrate.NewMigrator()
}

func runUp(cmd *cobra.Command, args []string) {
	migrator().Up()
}

func runDown(cmd *cobra.Command, args []string) {
	migrator().Rollback()
}

func runReset(cmd *cobra.Command, args []string) {
	migrator().Reset()
}

func runRefresh(cmd *cobra.Command, args []string) {
	migrator().Refresh()
}

func runFresh(cmd *cobra.Command, args []string) {
	migrator().Fresh()
}
