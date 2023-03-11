package make

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CmdMakeRequest = &cobra.Command{
	Use:   "request",
	Short: "Create request file, example make request user",
	Run:   runMakeRequest,
	Args:  cobra.ExactArgs(1),
}

func runMakeRequest(cmd *cobra.Command, args []string) {
	// Format user input
	model := makeModelFromString(args[0])
	// request file path
	filePath := fmt.Sprintf("app/requests/%s_request.go", model.PackageName)
	// Create file from request template
	createFileFromStub(filePath, "request", model)
}
