package source

import "testing"

func TestNew(t *testing.T) {
	src := New(
		"docker.io",
		Default,
	)

	if src.Registry != "docker.io" {
		t.Fatalf(
			"Registry = %q",
			src.Registry,
		)
	}

	if src.Mode != Default {
		t.Fatalf(
			"Mode = %v",
			src.Mode,
		)
	}
}
