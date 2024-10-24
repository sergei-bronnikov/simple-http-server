package app

import (
	"github.com/sergei-bronnikov/simple-http-server/internal/httpserver"
	"github.com/urfave/cli/v2"
	"os"
)

const version = "dev"

func Run() error {
	app := &cli.App{
		Name:        "Simple HTTP Server",
		Usage:       "",
		Description: "",
		Version:     version,
		Args:        true,
		ArgsUsage:   "",
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
