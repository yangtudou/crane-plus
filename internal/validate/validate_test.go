package validate

import (
	"testing"

	"github.com/yyysay/registry-sync/internal/config"
	"github.com/yyysay/registry-sync/internal/image"
)

func TestConfig(t *testing.T) {
	tests := []struct {
		name string
		cfg  config.Config
		ok   bool
	}{
		{
			name: "valid",
			cfg: config.Config{
				Rules: []config.Rule{
					{
						Name: "docker",
						Source: config.Source{
							Registry: "docker.io",
						},
						Destination: config.Destination{
							Registry: "registry.example.com/images",
							Mode:     "basename",
						},
					},
				},
			},
			ok: true,
		},
		{
			name: "empty rules",
			cfg: config.Config{
				Rules: nil,
			},
			ok: false,
		},
		{
			name: "empty registry",
			cfg: config.Config{
				Rules: []config.Rule{
					{
						Name: "docker",
						Source: config.Source{
							Registry: "docker.io",
						},
						Destination: config.Destination{
							Mode: "basename",
						},
					},
				},
			},
			ok: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Config(&tt.cfg)

			if tt.ok && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !tt.ok && err == nil {
				t.Fatal("expected error")
			}
		})
	}
}

func TestImages(t *testing.T) {
	img, err := image.Parse("nginx")
	if err != nil {
		t.Fatal(err)
	}

	if err := Images([]*image.Image{img}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := Images(nil); err == nil {
		t.Fatal("expected error")
	}
}
