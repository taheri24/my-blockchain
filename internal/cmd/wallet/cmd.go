package wallet

import "github.com/urfave/cli/v2"

func Run() {

}

var CliCommand *cli.Command = &cli.Command{
	Action: func(ctx *cli.Context) error {
		return nil
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "dataDir",
			Value: "./data",
		},
	},
}
