package engine

import (
	"context"
	"testing"

	"registry-sync/model"
)

func TestExecuteMultipleTargets(t *testing.T) {

	mock := &mockCopier{}

	engine := New(mock)

	plan := model.Plan{

		Image: model.Image{
			Registry:   "docker.io",
			Repository: "cloudflare/cloudflared",
			Tag:        "latest",
			Platform: []string{
				"linux/amd64",
			},
		},

		Targets: []model.Target{
			{
				Name:     "local",
				Registry: "local.example",
			},
			{
				Name:      "aliyun",
				Registry:  "registry.example",
				Namespace: "test",
				Flatten:   true,
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

	if len(mock.calls) != 2 {
		t.Fatalf(
			"expect 2 copy calls, got %d",
			len(mock.calls),
		)
	}

	expected := []string{
		"docker.io/cloudflare/cloudflared:latest => local.example/cloudflare/cloudflared:latest ([linux/amd64])",
		"docker.io/cloudflare/cloudflared:latest => registry.example/test/cloudflared:latest ([linux/amd64])",
	}

	for i, want := range expected {

		if mock.calls[i] != want {

			t.Fatalf(
				"call[%d]\n got: %s\nwant: %s",
				i,
				mock.calls[i],
				want,
			)
		}
	}
}
