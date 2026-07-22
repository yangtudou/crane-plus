package mapper

import (
	"testing"

	"github.com/yyysay/registry-sync/internal/image"
)

func TestMapPreserve(t *testing.T) {
	src, err := image.Parse("cloudflare/cloudflared:latest")
	if err != nil {
		t.Fatal(err)
	}

	m := New(Preserve)

	dst := m.Map(src)

	if dst.Name != "cloudflare/cloudflared" {
		t.Fatalf("Name = %q", dst.Name)
	}

	if dst.Reference != "docker.io/cloudflare/cloudflared:latest" {
		t.Fatalf("Reference = %q", dst.Reference)
	}
}

func TestMapBasename(t *testing.T) {
	src, err := image.Parse("cloudflare/cloudflared:latest")
	if err != nil {
		t.Fatal(err)
	}

	m := New(Basename)

	dst := m.Map(src)

	if dst.Name != "cloudflared" {
		t.Fatalf("Name = %q", dst.Name)
	}

	if dst.Reference != "docker.io/cloudflared:latest" {
		t.Fatalf("Reference = %q", dst.Reference)
	}
}

func TestMapGHCRBasename(t *testing.T) {
	src, err := image.Parse("ghcr.io/sagernet/sing-box:latest")
	if err != nil {
		t.Fatal(err)
	}

	m := New(Basename)

	dst := m.Map(src)

	if dst.Name != "sing-box" {
		t.Fatalf("Name = %q", dst.Name)
	}

	if dst.Reference != "ghcr.io/sing-box:latest" {
		t.Fatalf("Reference = %q", dst.Reference)
	}
}
