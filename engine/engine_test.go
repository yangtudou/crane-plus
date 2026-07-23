package engine

import (
	"context"
	"fmt"
	"testing"

	"registry-sync/model"
)

type mockCopier struct {
	calls []string
}

func (m *mockCopier) Copy(
	ctx context.Context,
	source string,
	target string,
	platform []string,
) error {

	call := fmt.Sprintf(
		"%s => %s (%v)",
		source,
		target,
		platform,
	)

	m.calls = append(m.calls, call)

	return nil
}

func TestBuildImageName(t *testing.T) {

	got := BuildImageName(
		"docker.io",
		"cloudflare/cloudflared",
		"latest",
	)

	want := "docker.io/cloudflare/cloudflared:latest"

	if got != want {
		t.Fatalf(
			"got %s want %s",
			got,
			want,
		)
	}
}

func TestBuildMirrorImage(t *testing.T) {

	image := model.Image{
		Registry:   "ghcr.io",
		Repository: "sagernet/sing-box",
		Tag:        "latest",
	}

	mirror := model.Mirror{
		URL:  "aliyun.example",
		Type: "flatten",
	}

	got := BuildMirrorImage(
		image,
		mirror,
	)

	want := "aliyun.example/sing-box:latest"

	if got != want {
		t.Fatalf(
			"got %s want %s",
			got,
			want,
		)
	}
}

func TestBuildTargetImage(t *testing.T) {

	image := model.Image{
		Registry:   "docker.io",
		Repository: "cloudflare/cloudflared",
		Tag:        "latest",
	}

	target := model.Target{
		Registry:  "registry.example",
		Namespace: "test",
		Flatten:   true,
	}

	got := BuildTargetImage(
		image,
		target,
	)

	want := "registry.example/test/cloudflared:latest"

	if got != want {
		t.Fatalf(
			"got %s want %s",
			got,
			want,
		)
	}
}

func TestEngineExecute(t *testing.T) {

	mock := &mockCopier{}

	engine := New(mock)

	plan := model.Plan{
		Image: model.Image{
			Registry:   "docker.io",
			Repository: "cloudflare/cloudflared",
			Tag:        "latest",
			Platform: []string{
				"linux/arm64",
			},
		},
		Targets: []model.Target{
			{
				Registry: "local.example",
			},
		},
	}

	err := engine.Execute(
		context.Background(),
		plan,
	)

	if err != nil {
		t.Fatal(err)
	}

	if len(mock.calls) == 0 {
		t.Fatal("copier was not called")
	}
}
