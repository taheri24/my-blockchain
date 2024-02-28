package syncserver

import (
	"fmt"
	"net/http"

	"github.com/urfave/cli/v2"
	"taheri24.ir/blockchain/internal/mux"
)

var startServer cli.ActionFunc = func(ctx *cli.Context) error {
	listenAddr := fmt.Sprintf(":%d", ctx.Int("port"))
	serverMux := http.NewServeMux()
	add := mux.GroupFn(serverMux)
	add("/api", mux.Api())

	return http.ListenAndServe(listenAddr, serverMux)
}

var CliCommand *cli.Command = &cli.Command{

	Action: startServer,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:  "port",
			Value: 2090,
		},
		&cli.StringFlag{
			Name:  "dataDir",
			Value: "./data",
		},
	},
}
