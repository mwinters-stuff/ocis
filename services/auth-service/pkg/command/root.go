package command

import (
	"os"

	"github.com/owncloud/ocis/v2/ocis-pkg/clihelper"
	"github.com/owncloud/ocis/v2/services/auth-service/pkg/config"
	"github.com/urfave/cli/v2"
)

// GetCommands provides all commands for this service
func GetCommands(cfg *config.Config) cli.Commands {
	return []*cli.Command{
		// start this service
		Server(cfg),

		// interaction with this service

		// infos about this service
		Health(cfg),
		Version(cfg),
	}
}

// Execute is the entry point for the ocis-auth-service command.
func Execute(cfg *config.Config) error {
	app := clihelper.DefaultApp(&cli.App{
		Name:     "auth-service",
		Usage:    "Provide service authentication for oCIS",
		Commands: GetCommands(cfg),
	})

	return app.RunContext(cfg.Context, os.Args)
}
