package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"gohub/pkg/console"
)

var CmdMakeCMD = &cobra.Command{
	Use:   "cmd",
	Short: "Create a command, should be snake_case, example: make cmd backup_database",
	Run:   runMakeCMD,
	Args:  cobra.ExactArgs(1), // 1 parameter Must be passed
}

func runMakeCMD(cmd *cobra.Command, args []string) {
	// Format the model name and return a Model object
	model := makeModelFromString(args[0])

	// Concatenate target file paths
	filePath := fmt.Sprintf("app/cmd/%s.go", model.PackageName)

	// Create file from template and replace variables
	createFileFromStub(filePath, "cmd", model)

	// Friendly reminder
	console.Success("command name:" + model.PackageName)
	console.Success("command variable name: cmd.Cmd" + model.StructName)
	console.Warning("please edit main.go's app.Commands slice to register command")
}
