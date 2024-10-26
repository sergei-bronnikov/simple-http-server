package httpserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"net/http"
)

func Run(ctx *cli.Context) error {
	opts, err := parseOptions(ctx)
	if err != nil {
		return err
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Static("/", opts.targetPath)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", opts.port),
		Handler: router,
	}
	return srv.ListenAndServe()
}
