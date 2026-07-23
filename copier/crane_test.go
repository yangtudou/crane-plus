package copier

import (
	"context"
	"testing"

	"github.com/google/go-containerregistry/pkg/crane"
)

func TestCraneCopier_Copy(t *testing.T) {

	copier := New()

	copier.copyFunc = func(
		source string,
		target string,
		opts ...crane.Option,
	) error {

		if source != "docker.io/library/alpine:latest" {
			t.Fatalf(
				"unexpected source: %s",
				source,
			)
		}

		if target != "registry.example/alpine:latest" {
			t.Fatalf(
				"unexpected target: %s",
				target,
			)
		}

		if len(opts) != 2 {
			t.Fatalf(
				"expect 2 options, got %d",
				len(opts),
			)
		}

		return nil
	}

	err := copier.Copy(
		context.Background(),
		"docker.io/library/alpine:latest",
		"registry.example/alpine:latest",
		[]string{
			"linux/amd64",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}
