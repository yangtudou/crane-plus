package image

import "testing"

func TestNormalizeRegistry(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty",
			input:    "",
			expected: "docker.io",
		},
		{
			name:     "docker",
			input:    "docker.io",
			expected: "docker.io",
		},
		{
			name:     "ghcr",
			input:    "ghcr.io",
			expected: "ghcr.io",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NormalizeRegistry(tt.input)

			if got != tt.expected {
				t.Fatalf(
					"NormalizeRegistry(%q) = %q, want %q",
					tt.input,
					got,
					tt.expected,
				)
			}
		})
	}
}

func TestIsDockerHub(t *testing.T) {
	if !IsDockerHub("docker.io") {
		t.Fatal("expected docker.io")
	}

	if IsDockerHub("ghcr.io") {
		t.Fatal("unexpected ghcr.io")
	}
}
