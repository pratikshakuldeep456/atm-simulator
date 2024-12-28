package main

import (
	"log"
	"os"
	"pratikshakuldeep456/atm-stimulator/pkg"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:  "atm-simulator",
		Usage: "simulates atm machine",
		Commands: []*cli.Command{
			{
				Name:   "start",
				Usage:  "start the session",
				Action: pkg.StartCLI,
			},
			{
				Name:   "setting",
				Usage:  "allow user to set currency",
				Action: pkg.SetCurrency,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
