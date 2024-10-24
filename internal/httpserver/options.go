package httpserver

import (
	"fmt"
	"github.com/sergei-bronnikov/simple-http-server/internal/utils"
	"github.com/urfave/cli/v2"
	"os"
)

type options struct {
	targetPath string

	port int
}

const (
	defaultTarget = "./static"
	defaultPort   = 8080
)

var CliFlags = []cli.Flag{
	&cli.IntFlag{
		Name:    "port",
		Usage:   "",
		Aliases: []string{"p"},
	},
}

func parseOptions(ctx *cli.Context) (*options, error) {
	targetPath := ctx.Args().First()
	if targetPath == "" {
		targetPath = defaultTarget
	}
	stat, err := os.Stat(targetPath)
	if err != nil {
		return nil, err
	}
	if !stat.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", targetPath)
	}

	port := ctx.Int("port")
	if port == 0 {
		port = defaultPort
	}
	if !utils.PortIsFree(port) {
		return nil, fmt.Errorf("port %d is not available", port)
	}
	opts := &options{
		targetPath: targetPath,
		port:       port,
	}
	return opts, nil
}
