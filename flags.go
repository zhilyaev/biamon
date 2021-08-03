package main

import "github.com/urfave/cli/v2"

var (
	ConfigTplFile string
	Templater     string
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
	&cli.StringFlag{
		Name:        "templater",
		Aliases:     []string{"t"},
		Value:       "gomplate",
		Usage:       "You can use: gomplate or sprig",
		EnvVars:     []string{"BIAMON_TEMPLATER"},
		Destination: &Templater,
	},
}
