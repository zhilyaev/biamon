package main

import (
	"fmt"
)

func (d *Docker) DockerTags() (a []string) {
	for i := 1; i < len(d.Destinations)-1; i++ {
		a = append(a, fmt.Sprintf("docker tag %q %q", d.Destinations[0], d.Destinations[i]))
	}

	return a
}
