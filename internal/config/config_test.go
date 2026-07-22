package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	dir := t.TempDir()

	file := filepath.Join(dir, "config.yaml")

	content := `
rules:
  - name: docker-to-aliyun
    source:
      registry: docker.io
    destination:
      registry: registry.cn-hangzhou.aliyuncs.com/myspace
      mode: basename

  - name: ghcr-to-aliyun
    source:
      registry: ghcr.io
    destination:
      registry: registry.cn-hangzhou.aliyuncs.com/myspace
      mode: preserve
`

	if err := os.WriteFile(file, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(file)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if len(cfg.Rules) != 2 {
		t.Fatalf("Rules = %d, want 2", len(cfg.Rules))
	}

	if cfg.Rules[0].Name != "docker-to-aliyun" {
		t.Fatalf("Name = %q", cfg.Rules[0].Name)
	}

	if cfg.Rules[0].Source.Registry != "docker.io" {
		t.Fatalf("Source.Registry = %q", cfg.Rules[0].Source.Registry)
	}

	if cfg.Rules[0].Destination.Registry != "registry.cn-hangzhou.aliyuncs.com/myspace" {
		t.Fatalf(
			"Destination.Registry = %q",
			cfg.Rules[0].Destination.Registry,
		)
	}

	if cfg.Rules[1].Destination.Mode != "preserve" {
		t.Fatalf(
			"Mode = %q",
			cfg.Rules[1].Destination.Mode,
		)
	}
}
