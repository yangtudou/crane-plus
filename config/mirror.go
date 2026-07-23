package config

type MirrorConfig struct {
	Name string      `yaml:"name"`
	URL  string      `yaml:"url"`
	Type MirrorType  `yaml:"type,omitempty"`
	Auth *AuthConfig `yaml:"auth,omitempty"`
}

type MirrorType string

const (
	MirrorTypeExact   MirrorType = ""
	MirrorTypeFlatten MirrorType = "flatten"
)
