package main

import (
	"bytes"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"text/template"
)

type Config struct {
	Registries []*Registry
	Version    string

	Dockers []*Docker
}

type Docker struct {
	Dockerfile, Context string

	Destinations []string
	BuildFlags   []string `yaml:"build_flags"`
	BuildLabels  []string `yaml:"build_labels"`
	PushFlags    []string `yaml:"push_flags"`
}

type Registry struct {
	Username, Password, Server string
}

func NewBiamon(yml []byte) (*Config, error) { // nolint:revive
	b := &Config{
		Version: Version,
	}

	err := yaml.Unmarshal(yml, b)
	if err != nil {
		return b, err
	}

	return b, err
}

func NewBiamonTpl(file string) (*Config, error) {
	b, err := templating(file, nil)
	if err != nil {
		return nil, err
	}

	return NewBiamon(b.Bytes())
}

func templating(file string, data interface{}) (*bytes.Buffer, error) {
	if data == nil {
		data = map[string]interface{}{}
	}

	src, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	t := template.Must(template.New("tpl").Parse(string(src)))
	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		return nil, err
	}

	return &buf, nil
}
