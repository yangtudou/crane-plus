package config

type Config struct {
	Dest     map[string]DestinationConfig `yaml:"dest"`
	Mirror   []MirrorConfig               `yaml:"mirror"`
	Sync     []string                     `yaml:"sync"`
	Platform []string                     `yaml:"platform,omitempty"`
	Sources  map[string]SourceConfig      `yaml:"sources"`
}
