package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
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
		fmt.Println(err)
		os.Exit(1)
	}
}
