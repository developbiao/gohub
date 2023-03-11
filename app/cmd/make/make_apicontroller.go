package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"gohub/pkg/console"
	"strings"
)

var CmdMakeAPIController = &cobra.Command{
	Use:   "apicontroller",
	Short: "Create api controller, example: make apicontroller v1/user",
	Run:   runMakeAPIController,
	Args:  cobra.ExactArgs(1),
}

func runMakeAPIController(cmd *cobra.Command, args []string) {
	// Process parameters
	array := strings.Split(args[0], "/")
	if len(array) != 2 {
		console.Exit("api controller name format: v1/user")
	}

	// apiVersion for contact target path
	apiVersion, name := array[0], array[1]
	model := makeModelFromString(name)

	// Build target directory
	filePath := fmt.Sprintf("app/http/controllers/api/%s/%s_controller.go", apiVersion, name)

	// Create controller file from template
	createFileFromStub(filePath, "apicontroller", model)
}
