package main

import "fmt"

func (r *Registry) Login() string {
	return fmt.Sprintf("docker login -u %s -p %s %s", r.Username, r.Password, r.Server)
}

func (r *Registry) LoginStdin() string {
	return fmt.Sprintf("echo %s | docker login -u %s %s --password-stdin", r.Password, r.Username, r.Server)
}

func (c *Config) Login() (a []string) {
	for _, r := range c.Registries {
		a = append(a, r.Login())
	}

	return a
}

func (c *Config) LoginStdin() (a []string) {
	for _, r := range c.Registries {
		a = append(a, r.LoginStdin())
	}

	return a
}
