package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

var commands = []*cli.Command{
	{
		Name:  "config",
		Usage: "show render config",
		Action: func(c *cli.Context) error {
			b, err := templating(ConfigTplFile, nil)
			if err != nil {
				return err
			}

			fmt.Println(b.String())
			return nil
		},
	},
	{
		Name:  "login",
		Usage: "$ source <(biamon login)",
		Action: func(c *cli.Context) error {
			biamon, err := NewBiamonTpl(ConfigTplFile)
			if err != nil {
				return err
			}

			fmt.Println(strings.Join(biamon.Login(), "\n"))
			return nil
		},
	},
	{
		Name:  "build",
		Usage: "$ source <(biamon build)",
		Action: func(c *cli.Context) error {
			biamon, err := NewBiamonTpl(ConfigTplFile)
			if err != nil {
				return err
			}

			fmt.Println(strings.Join(biamon.Build(), "\n"))
			return nil
		},
	},
	{
		Name:  "push",
		Usage: "$ source <(biamon push)",
		Action: func(c *cli.Context) error {
			biamon, err := NewBiamonTpl(ConfigTplFile)
			if err != nil {
				return err
			}

			fmt.Println(strings.Join(biamon.Push(), "\n"))
			return nil
		},
	},
}

func command404(c *cli.Context, s string) {
	fmt.Printf("Command %q not found\n", s)
	os.Exit(127)
}
