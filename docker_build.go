package main

import (
	"fmt"
	"strings"
)

func (d *Docker) Build() string {
	if d.Ctl == "" {
		d.Ctl = "docker"
	}

	tags := strings.Join(d.Destinations, " -t ")
	args := strings.Join(d.BuildFlags, " ")
	var labels string
	if len(d.BuildLabels) > 0 {
		labels = "--label " + strings.Join(d.BuildLabels, " --label ")
	}
	return fmt.Sprintf("%s build %s %s -t %s -f %s %s", d.Ctl, args, labels, tags, d.Dockerfile, d.Context)
}

func (c *Config) Build() (a []string) {
	for i := 0; i < len(c.Dockers); i++ {
		a = append(a, c.Dockers[i].Build())
	}
	return a
}
