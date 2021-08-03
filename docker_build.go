package main

import (
	"fmt"
	"strings"
)

func (d *Docker) Build() string {
	tags := strings.Join(d.Destinations, " -t ")
	args := strings.Join(d.BuildFlags, " ")
	labels := strings.Join(d.BuildLabels, " --label ")
	return fmt.Sprintf("docker build %s --label %s -t %s -f %s %s", args, labels, tags, d.Dockerfile, d.Context)
}

func (c *Config) Build() (a []string) {
	for i := 0; i < len(c.Dockers); i++ {
		a = append(a, c.Dockers[i].Build())
	}
	return a
}
