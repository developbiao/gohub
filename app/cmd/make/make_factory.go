package make

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CmdMakeFactory = &cobra.Command{
	Use:   "factory",
	Short: "Create model's factory file, example: make factory user",
	Run:   runMakeFactory,
	Args:  cobra.ExactArgs(1),
}

func runMakeFactory(cmd *cobra.Command, args []string) {
	// Format user input to model
	model := makeModelFromString(args[0])

	// Contact target file path
	filePath := fmt.Sprintf("database/factories/%s_factory.go", model.PackageName)

	// Create file from factory template
	createFileFromStub(filePath, "factory", model)
}
