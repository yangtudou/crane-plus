package rule

import (
	"testing"

	"github.com/yyysay/registry-sync/internal/destination"
	"github.com/yyysay/registry-sync/internal/mapper"
	"github.com/yyysay/registry-sync/internal/source"
)

func TestNew(t *testing.T) {

	src := source.New(
		"docker.io",
		source.Default,
	)

	dst := destination.New(
		"registry.example.com/images",
		mapper.New(mapper.Basename),
	)

	r := New(
		"docker-to-example",
		src,
		dst,
	)

	if r.Name != "docker-to-example" {

		t.Fatalf(
			"Name = %q",
			r.Name,
		)

	}

	if r.Source.Registry != "docker.io" {

		t.Fatalf(
			"Source.Registry = %q",
			r.Source.Registry,
		)

	}

	if r.Destination.Registry != "registry.example.com/images" {

		t.Fatalf(
			"Destination.Registry = %q",
			r.Destination.Registry,
		)

	}

}

func TestIsMixed(t *testing.T) {

	dst := destination.New(
		"registry.example.com/images",
		mapper.New(mapper.Basename),
	)

	r := New(
		"mixed",
		nil,
		dst,
	)

	if !r.IsMixed() {

		t.Fatal(
			"expected mixed rule",
		)

	}

	if r.Name != "mixed" {

		t.Fatalf(
			"Name = %q",
			r.Name,
		)

	}

	if r.Destination.Registry != "registry.example.com/images" {

		t.Fatalf(
			"Destination.Registry = %q",
			r.Destination.Registry,
		)

	}

}
