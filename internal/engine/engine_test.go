package engine

import (
	"testing"

	"github.com/yyysay/registry-sync/internal/destination"
	"github.com/yyysay/registry-sync/internal/image"
	"github.com/yyysay/registry-sync/internal/mapper"
	"github.com/yyysay/registry-sync/internal/rule"
	"github.com/yyysay/registry-sync/internal/source"
)

func TestGenerate(t *testing.T) {

	r := rule.New(
		"ghcr-to-aliyun",
		source.New(
			"ghcr.io",
			source.Default,
		),
		destination.New(
			"registry.example.com/images",
			mapper.New(mapper.Basename),
		),
	)

	engine := New(
		[]*rule.Rule{r},
	)

	img, err := image.Parse(
		"ghcr.io/sagernet/sing-box:latest",
	)

	if err != nil {
		t.Fatal(err)
	}

	tasks, err := engine.Generate(
		[]*image.Image{img},
	)

	if err != nil {
		t.Fatalf(
			"Generate() error = %v",
			err,
		)
	}

	if len(tasks) != 1 {

		t.Fatalf(
			"tasks = %d",
			len(tasks),
		)

	}

	want := "registry.example.com/images/sing-box:latest"

	if tasks[0].Target.Reference != want {

		t.Fatalf(
			"Target = %q, want %q",
			tasks[0].Target.Reference,
			want,
		)

	}

}

func TestGenerateNoMatch(t *testing.T) {

	r := rule.New(
		"docker-only",
		source.New(
			"docker.io",
			source.Default,
		),
		destination.New(
			"registry.example.com/images",
			mapper.New(mapper.Basename),
		),
	)

	engine := New(
		[]*rule.Rule{r},
	)

	img, err := image.Parse(
		"ghcr.io/sagernet/sing-box:latest",
	)

	if err != nil {
		t.Fatal(err)
	}

	tasks, err := engine.Generate(
		[]*image.Image{img},
	)

	if err != nil {
		t.Fatalf(
			"Generate() error = %v",
			err,
		)
	}

	if len(tasks) != 0 {

		t.Fatalf(
			"tasks = %d, want 0",
			len(tasks),
		)

	}

}

func TestGenerateMixedFallback(t *testing.T) {

	mixed := rule.New(
		"mixed",
		nil,
		destination.New(
			"registry.example.com/images",
			mapper.New(mapper.Basename),
		),
	)

	engine := New(
		[]*rule.Rule{mixed},
	)

	img, err := image.Parse(
		"nginx:latest",
	)

	if err != nil {
		t.Fatal(err)
	}

	tasks, err := engine.Generate(
		[]*image.Image{img},
	)

	if err != nil {
		t.Fatalf(
			"Generate() error = %v",
			err,
		)
	}

	if len(tasks) != 1 {

		t.Fatalf(
			"tasks = %d, want 1",
			len(tasks),
		)

	}

	want := "registry.example.com/images/nginx:latest"

	if tasks[0].Target.Reference != want {

		t.Fatalf(
			"Target = %q, want %q",
			tasks[0].Target.Reference,
			want,
		)

	}

}

func TestGenerateRulePriorityOverMixed(t *testing.T) {

	ghcr := rule.New(
		"ghcr",
		source.New(
			"ghcr.io",
			source.Default,
		),
		destination.New(
			"registry.example.com/github",
			mapper.New(mapper.Preserve),
		),
	)

	mixed := rule.New(
		"mixed",
		nil,
		destination.New(
			"registry.example.com/default",
			mapper.New(mapper.Basename),
		),
	)

	engine := New(
		[]*rule.Rule{
			ghcr,
			mixed,
		},
	)

	img, err := image.Parse(
		"ghcr.io/sagernet/sing-box:latest",
	)

	if err != nil {
		t.Fatal(err)
	}

	tasks, err := engine.Generate(
		[]*image.Image{img},
	)

	if err != nil {
		t.Fatalf(
			"Generate() error = %v",
			err,
		)
	}

	if len(tasks) != 1 {

		t.Fatalf(
			"tasks = %d, want 1",
			len(tasks),
		)

	}

	want := "registry.example.com/github/sagernet/sing-box:latest"

	if tasks[0].Target.Reference != want {

		t.Fatalf(
			"Target = %q, want %q",
			tasks[0].Target.Reference,
			want,
		)

	}

}
