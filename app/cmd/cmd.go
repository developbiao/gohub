package cmd

import (
	"github.com/spf13/cobra"
	"gohub/pkg/helpers"
	"os"
)

// Env Evn save global flags --env value
var Env string

// RegisterGlobalFlags register global flags
func RegisterGlobalFlags(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().StringVarP(&Env, "env", "e", "", "load .env file, "+
		"example: --env=testing will use .env.testing file")
}

// RegisterDefaultCmd register default command
func RegisterDefaultCmd(rootCmd *cobra.Command, subCmd *cobra.Command) {
	cmd, _, err := rootCmd.Find(os.Args[1:])
	firstArg := helpers.FirstElement(os.Args[1:])
	if err == nil && cmd.Use == rootCmd.Use && firstArg != "-h" && firstArg != "--help" {
		args := append([]string{subCmd.Use}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}
}
