package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/hairyhenderson/gomplate/v3"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Registries []*Registry
	Version    string

	Dockers []*Docker
}

type Docker struct {
	Ctl string `default:"docker" yaml:"ctl"`
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

var ErrWrongTemplater = errors.New("wrong templater use sprig or gomplate")

func templating(file string, data interface{}) (*bytes.Buffer, error) {
	if data == nil {
		data = map[string]interface{}{}
	}

	src, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var fm template.FuncMap
	switch Templater {
	case "gomplate":
		fm = gomplate.Funcs(nil)
	case "sprig":
		fm = sprig.TxtFuncMap()
	default:
		return nil, ErrWrongTemplater
	}

	t := template.Must(template.New("tpl").Funcs(fm).Parse(string(src)))
	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		return nil, err
	}

	return &buf, nil
}
