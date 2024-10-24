package httpserver

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"net/http"
)

func Run(ctx *cli.Context) error {
	opts, err := parseOptions(ctx)
	if err != nil {
		return err
	}

	http.Handle("/", http.FileServer(http.Dir(opts.targetPath)))
	err = http.ListenAndServe(fmt.Sprintf(":%d", opts.port), nil)
	if err != nil {
		return err
	}
	return nil
}
