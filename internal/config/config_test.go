package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {

	dir := t.TempDir()

	file := filepath.Join(
		dir,
		"config.yaml",
	)

	content := `
runner:
  type: crane
  platform: linux/amd64
  workers: 3

rules:

  - name: ghcr
    source:
      registry: ghcr.io
    destination:
      registry: registry.local.y3-3am.top/github
      mode: preserve

  - name: mixed
    destination:
      registry: registry.local.y3-3am.top/aliyun
      mode: basename
`

	if err := os.WriteFile(
		file,
		[]byte(content),
		0644,
	); err != nil {

		t.Fatal(err)

	}

	cfg, err := Load(file)

	if err != nil {
		t.Fatalf(
			"Load() error = %v",
			err,
		)
	}

	// runner

	if cfg.Runner.Type != "crane" {

		t.Fatalf(
			"Runner.Type = %q",
			cfg.Runner.Type,
		)

	}

	if cfg.Runner.Platform != "linux/amd64" {

		t.Fatalf(
			"Runner.Platform = %q",
			cfg.Runner.Platform,
		)

	}

	if cfg.Runner.Workers != 3 {

		t.Fatalf(
			"Runner.Workers = %d",
			cfg.Runner.Workers,
		)

	}

	// rules count

	if len(cfg.Rules) != 2 {

		t.Fatalf(
			"Rules = %d, want 2",
			len(cfg.Rules),
		)

	}

	// normal rule

	ghcr := cfg.Rules[0]

	if ghcr.Name != "ghcr" {

		t.Fatalf(
			"Name = %q",
			ghcr.Name,
		)

	}

	if ghcr.Source.Registry != "ghcr.io" {

		t.Fatalf(
			"Source.Registry = %q",
			ghcr.Source.Registry,
		)

	}

	if ghcr.Destination.Registry != "registry.local.y3-3am.top/github" {

		t.Fatalf(
			"Destination.Registry = %q",
			ghcr.Destination.Registry,
		)

	}

	if ghcr.Destination.Mode != "preserve" {

		t.Fatalf(
			"Destination.Mode = %q",
			ghcr.Destination.Mode,
		)

	}

	// mixed rule

	mixed := cfg.Rules[1]

	if mixed.Name != "mixed" {

		t.Fatalf(
			"Name = %q",
			mixed.Name,
		)

	}

	if mixed.Source.Registry != "" {

		t.Fatalf(
			"mixed Source.Registry = %q, want empty",
			mixed.Source.Registry,
		)

	}

	if mixed.Destination.Registry != "registry.local.y3-3am.top/aliyun" {

		t.Fatalf(
			"mixed Destination.Registry = %q",
			mixed.Destination.Registry,
		)

	}

	if mixed.Destination.Mode != "basename" {

		t.Fatalf(
			"mixed Destination.Mode = %q",
			mixed.Destination.Mode,
		)

	}

}
