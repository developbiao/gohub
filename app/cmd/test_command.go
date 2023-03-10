package cmd

import (
	"errors"
	"gohub/pkg/console"

	"github.com/spf13/cobra"
)

var CmdTestCommand = &cobra.Command{
	Use:   "test_command",
	Short: "HERE PUTS THE COMMAND DESCRIPTION",
	Run:   runTestCommand,
	Args:  cobra.ExactArgs(1), // 1 parameter must be passwd
}

func runTestCommand(cmd *cobra.Command, args []string) {

	console.Success("This is success message:" + args[0] + "\n")
	console.Warning("This is warning message\n")
	console.Error("This is error message\n")
	console.Warning("Terminal output please use english\n")
	console.Exit("exit method exist process and print\n")
	console.ExitIf(errors.New("when err != nil will be exit\n"))
}
