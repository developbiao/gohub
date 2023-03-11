package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var CmdMakeModel = &cobra.Command{
	Use:   "model",
	Short: "Create model file, example: make model user",
	Run:   runMakeModel,
	Args:  cobra.ExactArgs(1), // Must pass 1 parameter
}

func runMakeModel(cmd *cobra.Command, args []string) {
	// Format user input
	model := makeModelFromString(args[0])

	// Ensure model folder exists, example: `app/models/user`
	dir := fmt.Sprintf("app/models/%s/", model.PackageName)

	// osMkdirAll Creates directory use 0777
	os.MkdirAll(dir, os.ModePerm)

	// Replace variables
	createFileFromStub(dir+model.PackageName+"_model.go", "model/model", model)
	createFileFromStub(dir+model.PackageName+"_util.go", "model/model_util", model)
	createFileFromStub(dir+model.PackageName+"_hooks.go", "model/model_hooks", model)
}