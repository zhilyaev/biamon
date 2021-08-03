package main

import "github.com/urfave/cli/v2"

var (
	ConfigTplFile string
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "file",
		Aliases:     []string{"f"},
		Value:       "biamon.yml",
		Usage:       "Path to config",
		EnvVars:     []string{"BIAMON_FILE"},
		Destination: &ConfigTplFile,
	},
}
