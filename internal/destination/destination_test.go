package destination

import (
	"testing"

	"github.com/yyysay/registry-sync/internal/image"
	"github.com/yyysay/registry-sync/internal/mapper"
)

func TestDestinationBasename(t *testing.T) {
	src, err := image.Parse("cloudflare/cloudflared:latest")
	if err != nil {
		t.Fatal(err)
	}

	m := mapper.New(mapper.Basename)

	dst := New(
		"registry.cn-hangzhou.aliyuncs.com/myspace",
		m,
	)

	result := dst.Map(src)

	want := "registry.cn-hangzhou.aliyuncs.com/myspace/cloudflared:latest"

	if result.Reference != want {
		t.Fatalf("Reference = %q, want %q", result.Reference, want)
	}
}

func TestDestinationPreserve(t *testing.T) {
	src, err := image.Parse("cloudflare/cloudflared:latest")
	if err != nil {
		t.Fatal(err)
	}

	m := mapper.New(mapper.Preserve)

	dst := New(
		"harbor.example.com",
		m,
	)

	result := dst.Map(src)

	want := "harbor.example.com/cloudflare/cloudflared:latest"

	if result.Reference != want {
		t.Fatalf("Reference = %q, want %q", result.Reference, want)
	}
}

func TestDestinationGHCRBasename(t *testing.T) {
	src, err := image.Parse("ghcr.io/sagernet/sing-box:latest")
	if err != nil {
		t.Fatal(err)
	}

	m := mapper.New(mapper.Basename)

	dst := New(
		"registry.example.com/images",
		m,
	)

	result := dst.Map(src)

	want := "registry.example.com/images/sing-box:latest"

	if result.Reference != want {
		t.Fatalf("Reference = %q, want %q", result.Reference, want)
	}
}
