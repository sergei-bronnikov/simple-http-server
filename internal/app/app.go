package app

import (
	"github.com/sergei-bronnikov/simple-http-server/internal/httpserver"
	"github.com/urfave/cli/v2"
	"os"
)

const version = "dev"

func Run() error {
	app := &cli.App{
		Usage:       "",
		Description: "Simple http server",
		Version:     version,
		Args:        true,
		ArgsUsage:   "<target_directory>",
		Flags:       httpserver.CliFlags,
		Action: func(c *cli.Context) error {
			return httpserver.Run(c)
		},
	}

	if err := app.Run(os.Args); err != nil {
		return err
	}
	return nil
}
