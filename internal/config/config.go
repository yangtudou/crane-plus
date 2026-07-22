package config

import (
	"os"

	"github.com/yyysay/registry-sync/internal/mapper"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Destination Destination `yaml:"destination"`
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

func (d Destination) RepositoryMode() mapper.RepositoryMode {
	switch d.Mode {
	case "preserve":
		return mapper.Preserve
	default:
		return mapper.Basename
	}
}
