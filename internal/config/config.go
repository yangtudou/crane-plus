package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Rules []Rule `yaml:"rules"`
}

type Rule struct {
	Name        string      `yaml:"name"`
	Source      Source      `yaml:"source"`
	Destination Destination `yaml:"destination"`
}

type Source struct {
	Registry string `yaml:"registry"`
}

type Destination struct {
	Registry string `yaml:"registry"`
	Mode     string `yaml:"mode"`
}

func Load(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
