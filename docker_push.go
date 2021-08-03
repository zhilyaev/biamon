package main

import (
	"fmt"
	"strings"
)

func (d *Docker) Push() (a []string) {
	args := strings.Join(d.PushFlags, " ")

	for i := 0; i < len(d.Destinations)-1; i++ {
		a = append(a, fmt.Sprintf("docker push %s %s", args, d.Destinations[i]))
	}
	return a
}

func (c *Config) Push() (a []string) {
	for i := 0; i < len(c.Dockers); i++ {
		a = append(a, c.Dockers[i].Push()...)
	}
	return a
}
