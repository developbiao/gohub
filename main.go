package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"gohub/app/cmd"
	cmdmake "gohub/app/cmd/make"
	"gohub/bootstrap"
	btsConfig "gohub/config"
	"gohub/pkg/config"
	"gohub/pkg/console"
	"os"
)

func init() {
	// Load config folder configs
	btsConfig.Initialize()
}

func main() {

	// Application Entrance
	var rootCmd = &cobra.Command{
		Use:   "Gohub",
		Short: "A simple web server project",
		Long:  `Default will run "server" command, you can use "-h" flag to see all subcommands`,
		PersistentPreRun: func(command *cobra.Command, args []string) {
			// Initialization config
			config.InitConfig(cmd.Env)

			// Initialization logger
			bootstrap.SetupLogger()

			// Initialization database
			bootstrap.SetupDB()

			// Initialization redis
			bootstrap.SetupRedis()
		},
	}

	// Register sub commands
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
		cmdmake.CmdMake,
	)

	// Configuration default run web server
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// Register global argument --env
	cmd.RegisterGlobalFlags(rootCmd)

	// Execute main command
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}

	// Initialization dependency command --env arguments from local
	//var env string
	//flag.StringVar(&env, "env", "", "load .env file, e.g: --env=testing")
	//flag.Parse()
	//config.InitConfig(env)
}
