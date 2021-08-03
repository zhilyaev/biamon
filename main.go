package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var (
	Version = "dev"
	Ctl     *cli.App
)

func main() {
	Ctl = cli.NewApp()

	Ctl.Version = Version

	Ctl.Commands = commands
	Ctl.CommandNotFound = command404
	Ctl.Flags = flags

	err := Ctl.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
