package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	stroageclient "taheri24.ir/blockchain/internal/cmd/storage_client"
	syncserver "taheri24.ir/blockchain/internal/cmd/sync_server"
	"taheri24.ir/blockchain/internal/cmd/wallet"
)

//go:embed banner.txt
var banner string

func main() {
	fmt.Println(banner)

	cliApp := cli.App{
		Commands: []*cli.Command{
			wallet.CliCommand,
			syncserver.CliCommand,
			stroageclient.CliCommand,
		},
	}
	if err := cliApp.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
