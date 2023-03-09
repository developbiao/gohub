package cmd

import (
	"github.com/spf13/cobra"
	"gohub/pkg/console"
	"gohub/pkg/helpers"
)

var CmdKey = &cobra.Command{
	Use:   "key",
	Short: "Generate App Key, will print the generate key",
	Run:   runKeyGenerate,
	Args:  cobra.NoArgs, // no argument
}

func runKeyGenerate(cmd *cobra.Command, args []string) {
	console.Success("--------\n")
	console.Success("App Key:" + helpers.RandomString(32) + "\n")
	console.Success("--------\n")
	console.Warning("please go to .env file to change the APP_KEY option\n")
}
